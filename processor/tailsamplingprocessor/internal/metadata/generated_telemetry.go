// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                               metric.Meter
	ProcessorTailSamplingCountSpansSampled              metric.Int64Counter
	ProcessorTailSamplingCountTracesSampled             metric.Int64Counter
	ProcessorTailSamplingEarlyReleasesFromCacheDecision metric.Int64Counter
	ProcessorTailSamplingGlobalCountTracesSampled       metric.Int64Counter
	ProcessorTailSamplingNewTraceIDReceived             metric.Int64Counter
	ProcessorTailSamplingSamplingDecisionLatency        metric.Int64Histogram
	ProcessorTailSamplingSamplingDecisionTimerLatency   metric.Int64Histogram
	ProcessorTailSamplingSamplingLateSpanAge            metric.Int64Histogram
	ProcessorTailSamplingSamplingPolicyEvaluationError  metric.Int64Counter
	ProcessorTailSamplingSamplingTraceDroppedTooEarly   metric.Int64Counter
	ProcessorTailSamplingSamplingTraceRemovalAge        metric.Int64Histogram
	ProcessorTailSamplingSamplingTracesOnMemory         metric.Int64Gauge
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
	builder.ProcessorTailSamplingCountSpansSampled, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_count_spans_sampled",
		metric.WithDescription("Count of spans that were sampled or not per sampling policy"),
		metric.WithUnit("{spans}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingCountTracesSampled, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_count_traces_sampled",
		metric.WithDescription("Count of traces that were sampled or not per sampling policy"),
		metric.WithUnit("{traces}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingEarlyReleasesFromCacheDecision, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_early_releases_from_cache_decision",
		metric.WithDescription("Number of spans that were able to be immediately released due to a decision cache hit."),
		metric.WithUnit("{spans}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingGlobalCountTracesSampled, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_global_count_traces_sampled",
		metric.WithDescription("Global count of traces that were sampled or not by at least one policy"),
		metric.WithUnit("{traces}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingNewTraceIDReceived, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_new_trace_id_received",
		metric.WithDescription("Counts the arrival of new traces"),
		metric.WithUnit("{traces}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingDecisionLatency, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Histogram(
		"otelcol_processor_tail_sampling_sampling_decision_latency",
		metric.WithDescription("Latency (in microseconds) of a given sampling policy"),
		metric.WithUnit("µs"),
		metric.WithExplicitBucketBoundaries([]float64{1, 2, 5, 10, 25, 50, 75, 100, 150, 200, 300, 400, 500, 750, 1000, 2000, 3000, 4000, 5000, 10000, 20000, 30000, 50000}...),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingDecisionTimerLatency, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Histogram(
		"otelcol_processor_tail_sampling_sampling_decision_timer_latency",
		metric.WithDescription("Latency (in microseconds) of each run of the sampling decision timer"),
		metric.WithUnit("µs"),
		metric.WithExplicitBucketBoundaries([]float64{1, 2, 5, 10, 25, 50, 75, 100, 150, 200, 300, 400, 500, 750, 1000, 2000, 3000, 4000, 5000, 10000, 20000, 30000, 50000}...),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingLateSpanAge, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Histogram(
		"otelcol_processor_tail_sampling_sampling_late_span_age",
		metric.WithDescription("Time (in seconds) from the sampling decision was taken and the arrival of a late span"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingPolicyEvaluationError, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_sampling_policy_evaluation_error",
		metric.WithDescription("Count of sampling policy evaluation errors"),
		metric.WithUnit("{errors}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingTraceDroppedTooEarly, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_processor_tail_sampling_sampling_trace_dropped_too_early",
		metric.WithDescription("Count of traces that needed to be dropped before the configured wait time"),
		metric.WithUnit("{traces}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingTraceRemovalAge, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Histogram(
		"otelcol_processor_tail_sampling_sampling_trace_removal_age",
		metric.WithDescription("Time (in seconds) from arrival of a new trace until its removal from memory"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorTailSamplingSamplingTracesOnMemory, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Gauge(
		"otelcol_processor_tail_sampling_sampling_traces_on_memory",
		metric.WithDescription("Tracks the number of traces current on memory"),
		metric.WithUnit("{traces}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}

func getLeveledMeter(meter metric.Meter, cfgLevel, srvLevel configtelemetry.Level) metric.Meter {
	if cfgLevel <= srvLevel {
		return meter
	}
	return noopmetric.Meter{}
}
