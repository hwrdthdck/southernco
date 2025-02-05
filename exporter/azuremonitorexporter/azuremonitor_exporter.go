// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"context"
	"sync"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type azureMonitorExporter struct {
	config           *Config
	transportChannel appinsights.TelemetryChannel
	logger           *zap.Logger
	packer           *metricPacker
	startOnce        sync.Once
	endOnce          sync.Once
}

func (exporter *azureMonitorExporter) Start(_ context.Context, _ component.Host) (err error) {
	exporter.startOnce.Do(func() {
		connectionVars, err := parseConnectionString(exporter.config)
		if err != nil {
			return
		}

		exporter.config.InstrumentationKey = configopaque.String(connectionVars.InstrumentationKey)
		exporter.config.Endpoint = connectionVars.IngestionURL
		telemetryConfiguration := appinsights.NewTelemetryConfiguration(connectionVars.InstrumentationKey)
		telemetryConfiguration.EndpointUrl = connectionVars.IngestionURL
		telemetryConfiguration.MaxBatchSize = exporter.config.MaxBatchSize
		telemetryConfiguration.MaxBatchInterval = exporter.config.MaxBatchInterval

		telemetryClient := appinsights.NewTelemetryClientFromConfig(telemetryConfiguration)
		exporter.transportChannel = telemetryClient.Channel()
	})

	return nil
}

func (exporter *azureMonitorExporter) Shutdown(_ context.Context) (err error) {
	exporter.startOnce.Do(func() {
		if exporter.transportChannel != nil {
			select {
			case <-exporter.transportChannel.Close(exporter.config.ShutdownTimeout):
			case <-time.After(exporter.config.ShutdownTimeout):
				exporter.transportChannel.Stop()
			}
		}
	})

	return nil
}

func (exporter *azureMonitorExporter) consumeLogs(_ context.Context, logData plog.Logs) error {
	resourceLogs := logData.ResourceLogs()
	logPacker := newLogPacker(exporter.logger)

	for i := 0; i < resourceLogs.Len(); i++ {
		scopeLogs := resourceLogs.At(i).ScopeLogs()
		resource := resourceLogs.At(i).Resource()
		for j := 0; j < scopeLogs.Len(); j++ {
			logs := scopeLogs.At(j).LogRecords()
			scope := scopeLogs.At(j).Scope()
			for k := 0; k < logs.Len(); k++ {
				envelope := logPacker.LogRecordToEnvelope(logs.At(k), resource, scope)
				envelope.IKey = string(exporter.config.InstrumentationKey)
				exporter.transportChannel.Send(envelope)
			}
		}
	}
	// Flush the transport channel to force the telemetry to be sent
	exporter.transportChannel.Flush()
	return nil
}

func (exporter *azureMonitorExporter) consumeMetrics(_ context.Context, metricData pmetric.Metrics) error {
	resourceMetrics := metricData.ResourceMetrics()

	for i := 0; i < resourceMetrics.Len(); i++ {
		scopeMetrics := resourceMetrics.At(i).ScopeMetrics()
		resource := resourceMetrics.At(i).Resource()
		for j := 0; j < scopeMetrics.Len(); j++ {
			metrics := scopeMetrics.At(j).Metrics()
			scope := scopeMetrics.At(j).Scope()
			for k := 0; k < metrics.Len(); k++ {
				for _, envelope := range exporter.packer.MetricToEnvelopes(metrics.At(k), resource, scope) {
					envelope.IKey = string(exporter.config.InstrumentationKey)
					exporter.transportChannel.Send(envelope)
				}
			}
		}
	}

	// Flush the transport channel to force the telemetry to be sent
	exporter.transportChannel.Flush()
	return nil
}

type traceVisitor struct {
	processed int
	err       error
	exporter  *azureMonitorExporter
}

// Called for each tuple of Resource, InstrumentationScope, and Span
func (v *traceVisitor) visit(
	resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	span ptrace.Span,
) (ok bool) {
	envelopes, err := spanToEnvelopes(resource, scope, span, v.exporter.config.SpanEventsEnabled, v.exporter.logger)
	if err != nil {
		// record the error and short-circuit
		v.err = consumererror.NewPermanent(err)
		return false
	}

	for _, envelope := range envelopes {
		envelope.IKey = string(v.exporter.config.InstrumentationKey)

		// This is a fire and forget operation
		v.exporter.transportChannel.Send(envelope)
	}

	// Flush the transport channel to force the telemetry to be sent
	v.exporter.transportChannel.Flush()
	v.processed++

	return true
}

func (exporter *azureMonitorExporter) consumeTraces(_ context.Context, traceData ptrace.Traces) error {
	spanCount := traceData.SpanCount()
	if spanCount == 0 {
		return nil
	}

	visitor := &traceVisitor{exporter: exporter}
	accept(traceData, visitor)
	return visitor.err
}

// Returns a new instance of the log exporter
func newAzureMonitorExporter(config *Config, set exporter.Settings) AzureMonitorExporter {
	return &azureMonitorExporter{
		config:    config,
		logger:    set.Logger,
		packer:    newMetricPacker(set.Logger),
		startOnce: sync.Once{},
		endOnce:   sync.Once{},
	}
}
