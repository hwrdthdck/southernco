// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	ddconfig "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"
)

const (
	// typeStr is the type of the exporter
	typeStr = "datadog"
)

type factory struct {
	onceMetadata sync.Once
}

// NewFactory creates a Datadog exporter factory
func NewFactory() component.ExporterFactory {
	f := &factory{}
	return component.NewExporterFactory(
		typeStr,
		f.createDefaultConfig,
		component.WithMetricsExporter(f.createMetricsExporter),
		component.WithTracesExporter(f.createTracesExporter),
	)
}

func defaulttimeoutSettings() exporterhelper.TimeoutSettings {
	return exporterhelper.TimeoutSettings{
		Timeout: 15 * time.Second,
	}
}

// createDefaultConfig creates the default exporter configuration
func (*factory) createDefaultConfig() config.Exporter {
	return &ddconfig.Config{
		ExporterSettings: config.NewExporterSettings(config.NewComponentID(typeStr)),
		TimeoutSettings:  defaulttimeoutSettings(),
		RetrySettings:    exporterhelper.NewDefaultRetrySettings(),
		QueueSettings:    exporterhelper.NewDefaultQueueSettings(),

		Metrics: ddconfig.MetricsConfig{
			SendMonotonic: true,
			DeltaTTL:      3600,
			Quantiles:     true,
			ExporterConfig: ddconfig.MetricsExporterConfig{
				ResourceAttributesAsTags:             false,
				InstrumentationLibraryMetadataAsTags: false,
			},
			HistConfig: ddconfig.HistogramConfig{
				Mode:         "distributions",
				SendCountSum: false,
			},
			SumConfig: ddconfig.SumConfig{
				CumulativeMonotonicMode: ddconfig.CumulativeMonotonicSumModeToDelta,
			},
		},

		Traces: ddconfig.TracesConfig{
			SampleRate:      1,
			IgnoreResources: []string{},
		},

		HostMetadata: ddconfig.HostMetadataConfig{
			Enabled:        true,
			HostnameSource: ddconfig.HostnameSourceFirstResource,
		},

		SendMetadata:        true,
		UseResourceMetadata: true,
	}
}

// createMetricsExporter creates a metrics exporter based on this config.
func (f *factory) createMetricsExporter(
	ctx context.Context,
	set component.ExporterCreateSettings,
	c config.Exporter,
) (component.MetricsExporter, error) {

	cfg := c.(*ddconfig.Config)

	set.Logger.Info("sanitizing Datadog metrics exporter configuration")
	if err := cfg.Sanitize(set.Logger); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	var pushMetricsFn consumer.ConsumeMetricsFunc

	if cfg.OnlyMetadata {
		pushMetricsFn = func(_ context.Context, md pmetric.Metrics) error {
			// only sending metadata use only metrics
			f.onceMetadata.Do(func() {
				attrs := pcommon.NewMap()
				if md.ResourceMetrics().Len() > 0 {
					attrs = md.ResourceMetrics().At(0).Resource().Attributes()
				}
				go metadata.Pusher(ctx, set, newMetadataConfigfromConfig(cfg), attrs)
			})

			return nil
		}
	} else {
		exp, err := newMetricsExporter(ctx, set, cfg, &f.onceMetadata)
		if err != nil {
			cancel()
			return nil, err
		}
		pushMetricsFn = exp.PushMetricsDataScrubbed
	}

	exporter, err := exporterhelper.NewMetricsExporter(
		cfg,
		set,
		pushMetricsFn,
		// explicitly disable since we rely on http.Client timeout logic.
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0 * time.Second}),
		// We use our own custom mechanism for retries, since we hit several endpoints.
		exporterhelper.WithRetry(exporterhelper.RetrySettings{Enabled: false}),
		exporterhelper.WithQueue(cfg.QueueSettings),
		exporterhelper.WithShutdown(func(context.Context) error {
			cancel()
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}
	return resourcetotelemetry.WrapMetricsExporter(
		resourcetotelemetry.Settings{Enabled: cfg.Metrics.ExporterConfig.ResourceAttributesAsTags}, exporter), nil
}

// createTracesExporter creates a trace exporter based on this config.
func (f *factory) createTracesExporter(
	ctx context.Context,
	set component.ExporterCreateSettings,
	c config.Exporter,
) (component.TracesExporter, error) {

	cfg := c.(*ddconfig.Config)

	set.Logger.Info("sanitizing Datadog traces exporter configuration")
	if err := cfg.Sanitize(set.Logger); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	var pushTracesFn consumer.ConsumeTracesFunc

	if cfg.OnlyMetadata {
		pushTracesFn = func(_ context.Context, td ptrace.Traces) error {
			// only sending metadata, use only attributes
			f.onceMetadata.Do(func() {
				attrs := pcommon.NewMap()
				if td.ResourceSpans().Len() > 0 {
					attrs = td.ResourceSpans().At(0).Resource().Attributes()
				}
				go metadata.Pusher(ctx, set, newMetadataConfigfromConfig(cfg), attrs)
			})

			return nil
		}
	} else {
		pushTracesFn = newTracesExporter(ctx, set, cfg, &f.onceMetadata).pushTraceDataScrubbed
	}

	return exporterhelper.NewTracesExporter(
		cfg,
		set,
		pushTracesFn,
		// explicitly disable since we rely on http.Client timeout logic.
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0 * time.Second}),
		// We don't do retries on traces because of deduping concerns on APM Events.
		exporterhelper.WithRetry(exporterhelper.RetrySettings{Enabled: false}),
		exporterhelper.WithQueue(cfg.QueueSettings),
		exporterhelper.WithShutdown(func(context.Context) error {
			cancel()
			return nil
		}),
	)
}
