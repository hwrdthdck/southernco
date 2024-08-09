// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

const ScopeName = "github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter(ScopeName)
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer(ScopeName)
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                 metric.Meter
	RoutingProcessorNonRoutedLogRecords   metric.Int64Counter
	RoutingProcessorNonRoutedMetricPoints metric.Int64Counter
	RoutingProcessorNonRoutedSpans        metric.Int64Counter
	level                                 configtelemetry.Level
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// WithLevel sets the current telemetry level for the component.
func WithLevel(lvl configtelemetry.Level) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.level = lvl
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{level: configtelemetry.LevelBasic}
	for _, op := range options {
		op(&builder)
	}
	var err, errs error
	if builder.level >= configtelemetry.LevelBasic {
		builder.meter = Meter(settings)
	} else {
		builder.meter = noop.Meter{}
	}
	builder.RoutingProcessorNonRoutedLogRecords, err = builder.meter.Int64Counter(
		"otelcol_routing_processor_non_routed_log_records",
		metric.WithDescription("Number of log records that were not routed to some or all exporters."),
		metric.WithUnit("{records}"),
	)
	errs = errors.Join(errs, err)
	builder.RoutingProcessorNonRoutedMetricPoints, err = builder.meter.Int64Counter(
		"otelcol_routing_processor_non_routed_metric_points",
		metric.WithDescription("Number of metric points that were not routed to some or all exporters."),
		metric.WithUnit("{datapoints}"),
	)
	errs = errors.Join(errs, err)
	builder.RoutingProcessorNonRoutedSpans, err = builder.meter.Int64Counter(
		"otelcol_routing_processor_non_routed_spans",
		metric.WithDescription("Number of spans that were not routed to some or all exporters."),
		metric.WithUnit("{spans}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
