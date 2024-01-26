// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricExporter struct {
	config           *Config
	transportChannel transportChannel
	logger           *zap.Logger
	packer           *metricPacker
}

func (exporter *metricExporter) onMetricData(_ context.Context, metricData pmetric.Metrics) error {
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

	return nil
}

func (exporter *metricExporter) Shutdown(_ context.Context) error {
	shutdownTimeout := 1 * time.Second

	select {
	case <-exporter.transportChannel.Close():
		return nil
	case <-time.After(shutdownTimeout):
		// Currently, due to a dependency's bug (https://github.com/microsoft/ApplicationInsights-Go/issues/70),
		// the timeout will always be hit before Close is complete. This is not an error, but it will leak a goroutine.
		// There's nothing that we can do about this for now, so there's no reason to log or return an error
		// here.
		return nil
	}
}

// Returns a new instance of the metric exporter
func newMetricsExporter(config *Config, transportChannel transportChannel, set exporter.CreateSettings) (exporter.Metrics, error) {
	exporter := &metricExporter{
		config:           config,
		transportChannel: transportChannel,
		logger:           set.Logger,
		packer:           newMetricPacker(set.Logger),
	}

	return exporterhelper.NewMetricsExporter(
		context.TODO(),
		set,
		config,
		exporter.onMetricData,
		exporterhelper.WithQueue(config.QueueSettings),
		exporterhelper.WithShutdown(exporter.Shutdown))
}
