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

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/kafkareceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/kafkareceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                    metric.Meter
	KafkaReceiverCurrentOffset               metric.Int64Gauge
	KafkaReceiverMessages                    metric.Int64Counter
	KafkaReceiverOffsetLag                   metric.Int64Gauge
	KafkaReceiverPartitionClose              metric.Int64Counter
	KafkaReceiverPartitionStart              metric.Int64Counter
	KafkaReceiverUnmarshalFailedLogRecords   metric.Int64Counter
	KafkaReceiverUnmarshalFailedMetricPoints metric.Int64Counter
	KafkaReceiverUnmarshalFailedSpans        metric.Int64Counter
	level                                    configtelemetry.Level
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
	builder.KafkaReceiverCurrentOffset, err = builder.meter.Int64Gauge(
		"otelcol_kafka_receiver_current_offset",
		metric.WithDescription("Current message offset"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverMessages, err = builder.meter.Int64Counter(
		"otelcol_kafka_receiver_messages",
		metric.WithDescription("Number of received messages"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverOffsetLag, err = builder.meter.Int64Gauge(
		"otelcol_kafka_receiver_offset_lag",
		metric.WithDescription("Current offset lag"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverPartitionClose, err = builder.meter.Int64Counter(
		"otelcol_kafka_receiver_partition_close",
		metric.WithDescription("Number of finished partitions"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverPartitionStart, err = builder.meter.Int64Counter(
		"otelcol_kafka_receiver_partition_start",
		metric.WithDescription("Number of started partitions"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverUnmarshalFailedLogRecords, err = builder.meter.Int64Counter(
		"otelcol_kafka_receiver_unmarshal_failed_log_records",
		metric.WithDescription("Number of log records failed to be unmarshaled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverUnmarshalFailedMetricPoints, err = builder.meter.Int64Counter(
		"otelcol_kafka_receiver_unmarshal_failed_metric_points",
		metric.WithDescription("Number of metric points failed to be unmarshaled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverUnmarshalFailedSpans, err = builder.meter.Int64Counter(
		"otelcol_kafka_receiver_unmarshal_failed_spans",
		metric.WithDescription("Number of spans failed to be unmarshaled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
