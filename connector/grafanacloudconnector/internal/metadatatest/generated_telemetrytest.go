// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

// Deprecated: [v0.119.0] Use componenttest.Telemetry
type Telemetry struct {
	*componenttest.Telemetry
}

// Deprecated: [v0.119.0] Use componenttest.NewTelemetry
func SetupTelemetry(opts ...componenttest.TelemetryOption) Telemetry {
	return Telemetry{Telemetry: componenttest.NewTelemetry(opts...)}
}

// Deprecated: [v0.119.0] Use metadatatest.NewSettings
func (tt *Telemetry) NewSettings() connector.Settings {
	return NewSettings(tt.Telemetry)
}

func NewSettings(tt *componenttest.Telemetry) connector.Settings {
	set := connectortest.NewNopSettings()
	set.ID = component.NewID(component.MustNewType("grafanacloud"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

// Deprecated: [v0.119.0] Use metadatatest.AssertEqual*
func (tt *Telemetry) AssertMetrics(t *testing.T, expected []metricdata.Metrics, opts ...metricdatatest.Option) {
	var md metricdata.ResourceMetrics
	require.NoError(t, tt.Reader.Collect(context.Background(), &md))
	// ensure all required metrics are present
	for _, want := range expected {
		got := getMetricFromResource(want.Name, md)
		metricdatatest.AssertEqual(t, want, got, opts...)
	}

	// ensure no additional metrics are emitted
	require.Equal(t, len(expected), lenMetrics(md))
}

func AssertEqualGrafanacloudDatapointCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_grafanacloud_datapoint_count",
		Description: "Number of datapoints sent to Grafana Cloud",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_grafanacloud_datapoint_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualGrafanacloudFlushCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_grafanacloud_flush_count",
		Description: "Number of metrics flushes",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_grafanacloud_flush_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualGrafanacloudHostCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_grafanacloud_host_count",
		Description: "Number of unique hosts",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_grafanacloud_host_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func getMetricFromResource(name string, got metricdata.ResourceMetrics) metricdata.Metrics {
	for _, sm := range got.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				return m
			}
		}
	}

	return metricdata.Metrics{}
}

func lenMetrics(got metricdata.ResourceMetrics) int {
	metricsCount := 0
	for _, sm := range got.ScopeMetrics {
		metricsCount += len(sm.Metrics)
	}

	return metricsCount
}
