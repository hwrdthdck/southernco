// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen/internal/common"
)

func Start(cfg *Config) error {
	logger, err := common.CreateLogger(cfg.SkipSettingGRPCLogger)
	if err != nil {
		return err
	}

	var exp *otlptrace.Exporter
	if cfg.UseHTTP {
		logger.Info("starting HTTP exporter")
		exp, err = otlptracehttp.New(context.Background(), httpExporterOptions(cfg)...)
	} else {
		logger.Info("starting gRPC exporter")
		exp, err = otlptracegrpc.New(context.Background(), grpcExporterOptions(cfg)...)
	}

	if err != nil {
		return fmt.Errorf("failed to obtain OTLP exporter: %w", err)
	}
	defer func() {
		logger.Info("stopping the exporter")
		if tempError := exp.Shutdown(context.Background()); tempError != nil {
			logger.Error("failed to stop the exporter", zap.Error(tempError))
		}
	}()

	var ssp sdktrace.SpanProcessor
	if cfg.Batch {
		ssp = sdktrace.NewBatchSpanProcessor(exp, sdktrace.WithBatchTimeout(time.Second))
		defer func() {
			logger.Info("stop the batch span processor")
			if tempError := ssp.Shutdown(context.Background()); tempError != nil {
				logger.Error("failed to stop the batch span processor", zap.Error(err))
			}
		}()
	}

	var attributes []attribute.KeyValue
	// may be overridden by `-otlp-attributes service.name="foo"`
	attributes = append(attributes, semconv.ServiceNameKey.String(cfg.ServiceName))
	attributes = append(attributes, cfg.GetAttributes()...)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, attributes...)),
	)

	if cfg.Batch {
		tracerProvider.RegisterSpanProcessor(ssp)
	}
	otel.SetTracerProvider(tracerProvider)

	if err = Run(cfg, logger); err != nil {
		logger.Error("failed to execute the test scenario.", zap.Error(err))
		return err
	}

	return nil
}

// Run executes the test scenario.
func Run(c *Config, logger *zap.Logger) error {
	if c.TotalDuration > 0 {
		c.NumTraces = 0
	} else if c.NumTraces <= 0 {
		return fmt.Errorf("either `traces` or `duration` must be greater than 0")
	}

	limit := rate.Limit(c.Rate)
	if c.Rate == 0 {
		limit = rate.Inf
		logger.Info("generation of traces isn't being throttled")
	} else {
		logger.Info("generation of traces is limited", zap.Float64("per-second", float64(limit)))
	}

	var statusCode codes.Code

	switch strings.ToLower(c.StatusCode) {
	case "0", "unset", "":
		statusCode = codes.Unset
	case "1", "error":
		statusCode = codes.Error
	case "2", "ok":
		statusCode = codes.Ok
	default:
		return fmt.Errorf("expected `status-code` to be one of (Unset, Error, Ok) or (0, 1, 2), got %q instead", c.StatusCode)
	}

	wg := sync.WaitGroup{}

	running := &atomic.Bool{}
	running.Store(true)

	telemetryAttributes := c.GetTelemetryAttributes()

	for i := 0; i < c.WorkerCount; i++ {
		wg.Add(1)
		w := worker{
			numTraces:        c.NumTraces,
			propagateContext: c.PropagateContext,
			statusCode:       statusCode,
			limitPerSecond:   limit,
			totalDuration:    c.TotalDuration,
			running:          running,
			wg:               &wg,
			logger:           logger.With(zap.Int("worker", i)),
			loadSize:         c.LoadSize,
		}

		go w.simulateTraces(telemetryAttributes)
	}
	if c.TotalDuration > 0 {
		time.Sleep(c.TotalDuration)
		running.Store(false)
	}
	wg.Wait()
	return nil
}
