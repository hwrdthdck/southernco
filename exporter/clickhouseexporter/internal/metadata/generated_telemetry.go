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
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                       metric.Meter
	ExporterClickhouseSentLogs                  metric.Int64Counter
	ExporterClickhouseSentLogsBatchSize         metric.Int64Histogram
	ExporterClickhouseSentLogsLatency           metric.Int64Histogram
	ExporterClickhouseSentMetricPoints          metric.Int64Counter
	ExporterClickhouseSentMetricPointsBatchSize metric.Int64Histogram
	ExporterClickhouseSentMetricPointsLatency   metric.Int64Histogram
	ExporterClickhouseSentSpans                 metric.Int64Counter
	ExporterClickhouseSentSpansBatchSize        metric.Int64Histogram
	ExporterClickhouseSentSpansLatency          metric.Int64Histogram
	meters                                      map[configtelemetry.Level]metric.Meter
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
	builder.ExporterClickhouseSentLogs, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_clickhouse_sent_logs",
		metric.WithDescription("Number of logs the exporter sent."),
		metric.WithUnit("{logs}"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentLogsBatchSize, err = builder.meters[configtelemetry.LevelBasic].Int64Histogram(
		"otelcol_exporter_clickhouse_sent_logs_batch_size",
		metric.WithDescription("How many counts of a batch are sent to Clickhouse. Clickhouse recommend inserting metric in packets of at least 1000 rows, or no more than a single request per second."),
		metric.WithUnit("int"),
		metric.WithExplicitBucketBoundaries([]float64{100, 250, 500, 750, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 50000, 100000}...),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentLogsLatency, err = builder.meters[configtelemetry.LevelBasic].Int64Histogram(
		"otelcol_exporter_clickhouse_sent_logs_latency",
		metric.WithDescription("Latency (in milliseconds) of each sent log batch."),
		metric.WithUnit("ms"),
		metric.WithExplicitBucketBoundaries([]float64{5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000}...),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentMetricPoints, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_clickhouse_sent_metric_points",
		metric.WithDescription("Number of metric points the exporter sent."),
		metric.WithUnit("{metric_points}"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentMetricPointsBatchSize, err = builder.meters[configtelemetry.LevelBasic].Int64Histogram(
		"otelcol_exporter_clickhouse_sent_metric_points_batch_size",
		metric.WithDescription("How many counts of a batch are sent to Clickhouse. Clickhouse recommend inserting metric in packets of at least 1000 rows, or no more than a single request per second."),
		metric.WithUnit("int"),
		metric.WithExplicitBucketBoundaries([]float64{100, 250, 500, 750, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 50000, 100000}...),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentMetricPointsLatency, err = builder.meters[configtelemetry.LevelBasic].Int64Histogram(
		"otelcol_exporter_clickhouse_sent_metric_points_latency",
		metric.WithDescription("Metric."),
		metric.WithUnit("1"),
		metric.WithExplicitBucketBoundaries([]float64{5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000}...),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentSpans, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_clickhouse_sent_spans",
		metric.WithDescription("Number of spans the exporter sent."),
		metric.WithUnit("{spans}"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentSpansBatchSize, err = builder.meters[configtelemetry.LevelBasic].Int64Histogram(
		"otelcol_exporter_clickhouse_sent_spans_batch_size",
		metric.WithDescription("How many counts of a batch are sent to Clickhouse. Clickhouse recommend inserting metric in packets of at least 1000 rows, or no more than a single request per second."),
		metric.WithUnit("int"),
		metric.WithExplicitBucketBoundaries([]float64{100, 250, 500, 750, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 50000, 100000}...),
	)
	errs = errors.Join(errs, err)
	builder.ExporterClickhouseSentSpansLatency, err = builder.meters[configtelemetry.LevelBasic].Int64Histogram(
		"otelcol_exporter_clickhouse_sent_spans_latency",
		metric.WithDescription("Latency (in milliseconds) of each sent span batch."),
		metric.WithUnit("ms"),
		metric.WithExplicitBucketBoundaries([]float64{5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000}...),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
