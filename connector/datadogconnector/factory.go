// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate mdatagen metadata.yaml

package datadogconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector"

import (
	"context"

	"github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/metricsclient"
	"github.com/DataDog/datadog-agent/pkg/trace/timing"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/featuregate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector/internal/metadata"
)

// NativeIngestFeatureGate is the feature gate that controls native OTel spans ingestion in Datadog APM stats
var NativeIngestFeatureGate = featuregate.GlobalRegistry().MustRegister(
	"connector.datadogconnector.NativeIngest",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, datadogconnector uses the native OTel API to ingest OTel spans and produce APM stats."),
	featuregate.WithRegisterFromVersion("v0.102.0"),
)

// NewFactory creates a factory for tailtracer connector.
func NewFactory() connector.Factory {
	//  OTel connector factory to make a factory for connectors
	return connector.NewFactory(
		metadata.Type,
		createDefaultConfig,
		connector.WithTracesToMetrics(createTracesToMetricsConnector, metadata.TracesToMetricsStability),
		connector.WithTracesToTraces(createTracesToTracesConnector, metadata.TracesToTracesStability))
}

func createDefaultConfig() component.Config {
	return &Config{
		Traces: TracesConfig{
			IgnoreResources: []string{},
			TraceBuffer:     1000,
		},
	}
}

// defines the consumer type of the connector
// we want to consume traces and export metrics therefore define nextConsumer as metrics, consumer is the next component in the pipeline
func createTracesToMetricsConnector(_ context.Context, params connector.CreateSettings, cfg component.Config, nextConsumer consumer.Metrics) (c connector.Traces, err error) {
	metricsClient := metricsclient.InitializeMetricClient(params.MeterProvider, metricsclient.ConnectorSourceTag)
	timingReporter := timing.New(metricsClient)
	if NativeIngestFeatureGate.IsEnabled() {
		params.Logger.Info("Datadog connector using the native OTel API to ingest OTel spans and produce APM stats")
		c, err = newTraceToMetricConnectorNative(params.TelemetrySettings, cfg, nextConsumer, metricsClient, timingReporter)
	} else {
		params.Logger.Info("Datadog connector using the legacy processing pipelines to ingest OTel spans and produce APM stats")
		c, err = newTraceToMetricConnector(params.TelemetrySettings, cfg, nextConsumer, metricsClient, timingReporter)
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func createTracesToTracesConnector(_ context.Context, params connector.CreateSettings, _ component.Config, nextConsumer consumer.Traces) (connector.Traces, error) {
	return newTraceToTraceConnector(params.Logger, nextConsumer), nil
}
