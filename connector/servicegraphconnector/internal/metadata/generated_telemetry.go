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

const ScopeName = "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector"

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter(ScopeName)
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer(ScopeName)
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                             metric.Meter
	ConnectorServicegraphDroppedSpans metric.Int64Counter
	ConnectorServicegraphExpiredEdges metric.Int64Counter
	ConnectorServicegraphTotalEdges   metric.Int64Counter
	level                             configtelemetry.Level
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
	builder.ConnectorServicegraphDroppedSpans, err = builder.meter.Int64Counter(
		"otelcol_connector_servicegraph_dropped_spans",
		metric.WithDescription("Number of spans dropped when trying to add edges"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ConnectorServicegraphExpiredEdges, err = builder.meter.Int64Counter(
		"otelcol_connector_servicegraph_expired_edges",
		metric.WithDescription("Number of edges that expired before finding its matching span"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ConnectorServicegraphTotalEdges, err = builder.meter.Int64Counter(
		"otelcol_connector_servicegraph_total_edges",
		metric.WithDescription("Total number of unique edges"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
