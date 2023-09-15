// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type worker struct {
	running        *atomic.Bool    // pointer to shared flag that indicates it's time to stop the test
	metricType     metricType      // type of metric to generate
	numMetrics     int             // how many metrics the worker has to generate (only when duration==0)
	totalDuration  time.Duration   // how long to run the test for (overrides `numMetrics`)
	limitPerSecond rate.Limit      // how many metrics per second to generate
	wg             *sync.WaitGroup // notify when done
	logger         *zap.Logger     // logger
	index          int             // worker index
}

func (w worker) simulateMetrics(res *resource.Resource, exporter sdkmetric.Exporter) {
	limiter := rate.NewLimiter(w.limitPerSecond, 1)

	var i int64
	for w.running.Load() {
		var metrics []metricdata.Metrics
		if w.metricType == metricTypeGauge || w.metricType == metricTypeAll {
			metrics = append(metrics, metricdata.Metrics{
				Name: "gen.metric.gauge",
				Data: metricdata.Gauge[int64]{
					DataPoints: []metricdata.DataPoint[int64]{
						{
							StartTime: time.Now(),
							Time:      time.Now().Add(1 * time.Second),
							Value:     i,
						},
					},
				},
			})
		} else if w.metricType == metricTypeSum || w.metricType == metricTypeAll {
			metrics = append(metrics, metricdata.Metrics{
				Name: "gen.metric.sum",
				Data: metricdata.Sum[int64]{
					IsMonotonic: true,
					Temporality: metricdata.DeltaTemporality,
					DataPoints: []metricdata.DataPoint[int64]{
						{
							StartTime: time.Now().Add(-1 * time.Second),
							Time:      time.Now(),
							Value:     i,
						},
					},
				},
			})
		} else if w.metricType == metricTypeHistogram || w.metricType == metricTypeAll {
			value := 24.42
			metrics = append(metrics, metricdata.Metrics{
				Name: "gen.metric.histogram",
				Data: metricdata.Histogram[float64]{
					Temporality: metricdata.CumulativeTemporality,
					DataPoints: []metricdata.HistogramDataPoint[float64]{
						{
							Sum:   float64(float32(value)),
							Max:   metricdata.NewExtrema(float64(float32(value))),
							Min:   metricdata.NewExtrema(float64(float32(value))),
							Count: 1,
						},
					},
				},
			})
		} else {
			w.logger.Fatal("unknown metric type")
		}

		rm := metricdata.ResourceMetrics{
			Resource:     res,
			ScopeMetrics: []metricdata.ScopeMetrics{{Metrics: metrics}},
		}

		if err := exporter.Export(context.Background(), &rm); err != nil {
			w.logger.Fatal("exporter failed", zap.Error(err))
		}
		if err := limiter.Wait(context.Background()); err != nil {
			w.logger.Fatal("limiter wait failed, retry", zap.Error(err))
		}

		i++
		if w.numMetrics != 0 && i >= int64(w.numMetrics) {
			break
		}
	}

	w.logger.Info("metrics generated", zap.Int64("metrics", i))
	w.wg.Done()
}
