// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                       metric.Meter
	OtelsvcK8sIPLookupMiss      metric.Int64Counter
	OtelsvcK8sNamespaceAdded    metric.Int64Counter
	OtelsvcK8sNamespaceDeleted  metric.Int64Counter
	OtelsvcK8sNamespaceUpdated  metric.Int64Counter
	OtelsvcK8sNodeAdded         metric.Int64Counter
	OtelsvcK8sNodeDeleted       metric.Int64Counter
	OtelsvcK8sNodeUpdated       metric.Int64Counter
	OtelsvcK8sPodAdded          metric.Int64Counter
	OtelsvcK8sPodDeleted        metric.Int64Counter
	OtelsvcK8sPodTableSize      metric.Int64Gauge
	OtelsvcK8sPodUpdated        metric.Int64Counter
	OtelsvcK8sReplicasetAdded   metric.Int64Counter
	OtelsvcK8sReplicasetDeleted metric.Int64Counter
	OtelsvcK8sReplicasetUpdated metric.Int64Counter
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
	builder.OtelsvcK8sIPLookupMiss, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_ip_lookup_miss",
		metric.WithDescription("Number of times pod by IP lookup failed."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sNamespaceAdded, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_namespace_added",
		metric.WithDescription("Number of namespace add events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sNamespaceDeleted, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_namespace_deleted",
		metric.WithDescription("Number of namespace delete events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sNamespaceUpdated, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_namespace_updated",
		metric.WithDescription("Number of namespace update events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sNodeAdded, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_node_added",
		metric.WithDescription("Number of node add events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sNodeDeleted, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_node_deleted",
		metric.WithDescription("Number of node delete events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sNodeUpdated, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_node_updated",
		metric.WithDescription("Number of node update events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sPodAdded, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_pod_added",
		metric.WithDescription("Number of pod add events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sPodDeleted, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_pod_deleted",
		metric.WithDescription("Number of pod delete events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sPodTableSize, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Gauge(
		"otelcol_otelsvc_k8s_pod_table_size",
		metric.WithDescription("Size of table containing pod info"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sPodUpdated, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_pod_updated",
		metric.WithDescription("Number of pod update events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sReplicasetAdded, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_replicaset_added",
		metric.WithDescription("Number of ReplicaSet add events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sReplicasetDeleted, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_replicaset_deleted",
		metric.WithDescription("Number of ReplicaSet delete events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelsvcK8sReplicasetUpdated, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_otelsvc_k8s_replicaset_updated",
		metric.WithDescription("Number of ReplicaSet update events received"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}

func getLeveledMeter(meter metric.Meter, cfgLevel, srvLevel configtelemetry.Level) metric.Meter {
	if cfgLevel <= srvLevel {
		return meter
	}
	return noop.Meter{}
}
