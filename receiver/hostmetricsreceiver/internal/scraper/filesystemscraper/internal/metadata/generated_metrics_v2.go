// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/filesystem metrics.
type MetricsSettings struct {
	SystemFilesystemInodesUsage MetricSettings `mapstructure:"system.filesystem.inodes.usage"`
	SystemFilesystemUsage       MetricSettings `mapstructure:"system.filesystem.usage"`
	SystemFilesystemUtilization MetricSettings `mapstructure:"system.filesystem.utilization"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemFilesystemInodesUsage: MetricSettings{
			Enabled: true,
		},
		SystemFilesystemUsage: MetricSettings{
			Enabled: true,
		},
		SystemFilesystemUtilization: MetricSettings{
			Enabled: false,
		},
	}
}

// AttributeState specifies the a value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateFree
	AttributeStateReserved
	AttributeStateUsed
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateFree:
		return "free"
	case AttributeStateReserved:
		return "reserved"
	case AttributeStateUsed:
		return "used"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"free":     AttributeStateFree,
	"reserved": AttributeStateReserved,
	"used":     AttributeStateUsed,
}

type metricSystemFilesystemInodesUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.filesystem.inodes.usage metric with initial data.
func (m *metricSystemFilesystemInodesUsage) init() {
	m.data.SetName("system.filesystem.inodes.usage")
	m.data.SetDescription("FileSystem inodes used.")
	m.data.SetUnit("{inodes}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemFilesystemInodesUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, modeAttributeValue string, mountpointAttributeValue string, typeAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("mode", pcommon.NewValueString(modeAttributeValue))
	dp.Attributes().Insert("mountpoint", pcommon.NewValueString(mountpointAttributeValue))
	dp.Attributes().Insert("type", pcommon.NewValueString(typeAttributeValue))
	dp.Attributes().Insert("state", pcommon.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemFilesystemInodesUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemFilesystemInodesUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemFilesystemInodesUsage(settings MetricSettings) metricSystemFilesystemInodesUsage {
	m := metricSystemFilesystemInodesUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemFilesystemUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.filesystem.usage metric with initial data.
func (m *metricSystemFilesystemUsage) init() {
	m.data.SetName("system.filesystem.usage")
	m.data.SetDescription("Filesystem bytes used.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemFilesystemUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, modeAttributeValue string, mountpointAttributeValue string, typeAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("mode", pcommon.NewValueString(modeAttributeValue))
	dp.Attributes().Insert("mountpoint", pcommon.NewValueString(mountpointAttributeValue))
	dp.Attributes().Insert("type", pcommon.NewValueString(typeAttributeValue))
	dp.Attributes().Insert("state", pcommon.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemFilesystemUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemFilesystemUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemFilesystemUsage(settings MetricSettings) metricSystemFilesystemUsage {
	m := metricSystemFilesystemUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemFilesystemUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.filesystem.utilization metric with initial data.
func (m *metricSystemFilesystemUtilization) init() {
	m.data.SetName("system.filesystem.utilization")
	m.data.SetDescription("Fraction of filesystem bytes used.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemFilesystemUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, deviceAttributeValue string, modeAttributeValue string, mountpointAttributeValue string, typeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("mode", pcommon.NewValueString(modeAttributeValue))
	dp.Attributes().Insert("mountpoint", pcommon.NewValueString(mountpointAttributeValue))
	dp.Attributes().Insert("type", pcommon.NewValueString(typeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemFilesystemUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemFilesystemUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemFilesystemUtilization(settings MetricSettings) metricSystemFilesystemUtilization {
	m := metricSystemFilesystemUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                         pcommon.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity                   int               // maximum observed number of metrics per resource.
	resourceCapacity                  int               // maximum observed number of resource attributes.
	metricsBuffer                     pmetric.Metrics   // accumulates metrics data before emitting.
	metricSystemFilesystemInodesUsage metricSystemFilesystemInodesUsage
	metricSystemFilesystemUsage       metricSystemFilesystemUsage
	metricSystemFilesystemUtilization metricSystemFilesystemUtilization
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		metricSystemFilesystemInodesUsage: newMetricSystemFilesystemInodesUsage(settings.SystemFilesystemInodesUsage),
		metricSystemFilesystemUsage:       newMetricSystemFilesystemUsage(settings.SystemFilesystemUsage),
		metricSystemFilesystemUtilization: newMetricSystemFilesystemUtilization(settings.SystemFilesystemUtilization),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).DataType() {
			case pmetric.MetricDataTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricDataTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/filesystem")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemFilesystemInodesUsage.emit(ils.Metrics())
	mb.metricSystemFilesystemUsage.emit(ils.Metrics())
	mb.metricSystemFilesystemUtilization.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSystemFilesystemInodesUsageDataPoint adds a data point to system.filesystem.inodes.usage metric.
func (mb *MetricsBuilder) RecordSystemFilesystemInodesUsageDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, modeAttributeValue string, mountpointAttributeValue string, typeAttributeValue string, stateAttributeValue AttributeState) {
	mb.metricSystemFilesystemInodesUsage.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, modeAttributeValue, mountpointAttributeValue, typeAttributeValue, stateAttributeValue.String())
}

// RecordSystemFilesystemUsageDataPoint adds a data point to system.filesystem.usage metric.
func (mb *MetricsBuilder) RecordSystemFilesystemUsageDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, modeAttributeValue string, mountpointAttributeValue string, typeAttributeValue string, stateAttributeValue AttributeState) {
	mb.metricSystemFilesystemUsage.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, modeAttributeValue, mountpointAttributeValue, typeAttributeValue, stateAttributeValue.String())
}

// RecordSystemFilesystemUtilizationDataPoint adds a data point to system.filesystem.utilization metric.
func (mb *MetricsBuilder) RecordSystemFilesystemUtilizationDataPoint(ts pcommon.Timestamp, val float64, deviceAttributeValue string, modeAttributeValue string, mountpointAttributeValue string, typeAttributeValue string) {
	mb.metricSystemFilesystemUtilization.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, modeAttributeValue, mountpointAttributeValue, typeAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
