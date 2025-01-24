// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                         metric.Meter
	LoadbalancerBackendLatency    metric.Int64Histogram
	LoadbalancerBackendOutcome    metric.Int64Counter
	LoadbalancerNumBackendUpdates metric.Int64Counter
	LoadbalancerNumBackends       metric.Int64Gauge
	LoadbalancerNumResolutions    metric.Int64Counter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.LoadbalancerBackendLatency, err = builder.meter.Int64Histogram(
		"otelcol_loadbalancer_backend_latency",
		metric.WithDescription("Response latency in ms for the backends."),
		metric.WithUnit("ms"),
		metric.WithExplicitBucketBoundaries([]float64{5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000}...),
	)
	errs = errors.Join(errs, err)
	builder.LoadbalancerBackendOutcome, err = builder.meter.Int64Counter(
		"otelcol_loadbalancer_backend_outcome",
		metric.WithDescription("Number of successes and failures for each endpoint."),
		metric.WithUnit("{outcomes}"),
	)
	errs = errors.Join(errs, err)
	builder.LoadbalancerNumBackendUpdates, err = builder.meter.Int64Counter(
		"otelcol_loadbalancer_num_backend_updates",
		metric.WithDescription("Number of times the list of backends was updated."),
		metric.WithUnit("{updates}"),
	)
	errs = errors.Join(errs, err)
	builder.LoadbalancerNumBackends, err = builder.meter.Int64Gauge(
		"otelcol_loadbalancer_num_backends",
		metric.WithDescription("Current number of backends in use."),
		metric.WithUnit("{backends}"),
	)
	errs = errors.Join(errs, err)
	builder.LoadbalancerNumResolutions, err = builder.meter.Int64Counter(
		"otelcol_loadbalancer_num_resolutions",
		metric.WithDescription("Number of times the resolver has triggered new resolutions."),
		metric.WithUnit("{resolutions}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
