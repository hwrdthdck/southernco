// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

// Deprecated: [v0.108.0] use LeveledMeter instead.
func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                metric.Meter
	DeltatocumulativeDatapointsDropped   metric.Int64Counter
	DeltatocumulativeDatapointsProcessed metric.Int64Counter
	DeltatocumulativeGapsLength          metric.Int64Counter
	DeltatocumulativeStreamsEvicted      metric.Int64Counter
	DeltatocumulativeStreamsLimit        metric.Int64Gauge
	DeltatocumulativeStreamsMaxStale     metric.Int64Gauge
	DeltatocumulativeStreamsTracked      metric.Int64UpDownCounter
	meters                               map[configtelemetry.Level]metric.Meter
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
	builder := TelemetryBuilder{meters: map[configtelemetry.Level]metric.Meter{}}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meters[configtelemetry.LevelBasic] = LeveledMeter(settings, configtelemetry.LevelBasic)
	var err, errs error
	builder.DeltatocumulativeDatapointsDropped, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_deltatocumulative.datapoints.dropped",
		metric.WithDescription("number of datapoints dropped due to given 'reason'"),
		metric.WithUnit("{datapoint}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeDatapointsProcessed, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_deltatocumulative.datapoints.processed",
		metric.WithDescription("number of datapoints processed"),
		metric.WithUnit("{datapoint}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeGapsLength, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_deltatocumulative.gaps.length",
		metric.WithDescription("total duration where data was expected but not received"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsEvicted, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_deltatocumulative.streams.evicted",
		metric.WithDescription("number of streams evicted"),
		metric.WithUnit("{stream}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsLimit, err = builder.meters[configtelemetry.LevelBasic].Int64Gauge(
		"otelcol_deltatocumulative.streams.limit",
		metric.WithDescription("upper limit of tracked streams"),
		metric.WithUnit("{stream}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsMaxStale, err = builder.meters[configtelemetry.LevelBasic].Int64Gauge(
		"otelcol_deltatocumulative.streams.max_stale",
		metric.WithDescription("duration after which streams inactive streams are dropped"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsTracked, err = builder.meters[configtelemetry.LevelBasic].Int64UpDownCounter(
		"otelcol_deltatocumulative.streams.tracked",
		metric.WithDescription("number of streams tracked"),
		metric.WithUnit("{dps}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
