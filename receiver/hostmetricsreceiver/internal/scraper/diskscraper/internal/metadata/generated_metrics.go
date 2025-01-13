// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// AttributeDirection specifies the value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionRead
	AttributeDirectionWrite
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionRead:
		return "read"
	case AttributeDirectionWrite:
		return "write"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"read":  AttributeDirectionRead,
	"write": AttributeDirectionWrite,
}

type metricSystemDiskIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.io metric with initial data.
func (m *metricSystemDiskIo) init() {
	m.data.SetName("system.disk.io")
	m.data.SetDescription("Disk bytes transferred.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskIo) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskIo(cfg MetricConfig) metricSystemDiskIo {
	m := metricSystemDiskIo{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskIoTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.io_time metric with initial data.
func (m *metricSystemDiskIoTime) init() {
	m.data.SetName("system.disk.io_time")
	m.data.SetDescription("Time disk spent activated. On Windows, this is calculated as the inverse of disk idle time.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskIoTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, deviceAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskIoTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskIoTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskIoTime(cfg MetricConfig) metricSystemDiskIoTime {
	m := metricSystemDiskIoTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskMerged struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.merged metric with initial data.
func (m *metricSystemDiskMerged) init() {
	m.data.SetName("system.disk.merged")
	m.data.SetDescription("The number of disk reads/writes merged into single physical disk access operations.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskMerged) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskMerged) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskMerged) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskMerged(cfg MetricConfig) metricSystemDiskMerged {
	m := metricSystemDiskMerged{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskOperationTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.operation_time metric with initial data.
func (m *metricSystemDiskOperationTime) init() {
	m.data.SetName("system.disk.operation_time")
	m.data.SetDescription("Time spent in disk operations.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskOperationTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskOperationTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskOperationTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskOperationTime(cfg MetricConfig) metricSystemDiskOperationTime {
	m := metricSystemDiskOperationTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.operations metric with initial data.
func (m *metricSystemDiskOperations) init() {
	m.data.SetName("system.disk.operations")
	m.data.SetDescription("Disk operations count.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskOperations) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskOperations(cfg MetricConfig) metricSystemDiskOperations {
	m := metricSystemDiskOperations{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskPendingOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.pending_operations metric with initial data.
func (m *metricSystemDiskPendingOperations) init() {
	m.data.SetName("system.disk.pending_operations")
	m.data.SetDescription("The queue size of pending I/O operations.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskPendingOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskPendingOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskPendingOperations) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskPendingOperations(cfg MetricConfig) metricSystemDiskPendingOperations {
	m := metricSystemDiskPendingOperations{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskWeightedIoTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.weighted_io_time metric with initial data.
func (m *metricSystemDiskWeightedIoTime) init() {
	m.data.SetName("system.disk.weighted_io_time")
	m.data.SetDescription("Time disk spent activated multiplied by the queue length.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskWeightedIoTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, deviceAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("device", deviceAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskWeightedIoTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskWeightedIoTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskWeightedIoTime(cfg MetricConfig) metricSystemDiskWeightedIoTime {
	m := metricSystemDiskWeightedIoTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                            MetricsBuilderConfig // config of the metrics builder.
	startTime                         pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                   int                  // maximum observed number of metrics per resource.
	metricsBuffer                     pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                         component.BuildInfo  // contains version information.
	metricSystemDiskIo                metricSystemDiskIo
	metricSystemDiskIoTime            metricSystemDiskIoTime
	metricSystemDiskMerged            metricSystemDiskMerged
	metricSystemDiskOperationTime     metricSystemDiskOperationTime
	metricSystemDiskOperations        metricSystemDiskOperations
	metricSystemDiskPendingOperations metricSystemDiskPendingOperations
	metricSystemDiskWeightedIoTime    metricSystemDiskWeightedIoTime
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                            mbc,
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		buildInfo:                         settings.BuildInfo,
		metricSystemDiskIo:                newMetricSystemDiskIo(mbc.Metrics.SystemDiskIo),
		metricSystemDiskIoTime:            newMetricSystemDiskIoTime(mbc.Metrics.SystemDiskIoTime),
		metricSystemDiskMerged:            newMetricSystemDiskMerged(mbc.Metrics.SystemDiskMerged),
		metricSystemDiskOperationTime:     newMetricSystemDiskOperationTime(mbc.Metrics.SystemDiskOperationTime),
		metricSystemDiskOperations:        newMetricSystemDiskOperations(mbc.Metrics.SystemDiskOperations),
		metricSystemDiskPendingOperations: newMetricSystemDiskPendingOperations(mbc.Metrics.SystemDiskPendingOperations),
		metricSystemDiskWeightedIoTime:    newMetricSystemDiskWeightedIoTime(mbc.Metrics.SystemDiskWeightedIoTime),
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemDiskIo.emit(ils.Metrics())
	mb.metricSystemDiskIoTime.emit(ils.Metrics())
	mb.metricSystemDiskMerged.emit(ils.Metrics())
	mb.metricSystemDiskOperationTime.emit(ils.Metrics())
	mb.metricSystemDiskOperations.emit(ils.Metrics())
	mb.metricSystemDiskPendingOperations.emit(ils.Metrics())
	mb.metricSystemDiskWeightedIoTime.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSystemDiskIoDataPoint adds a data point to system.disk.io metric.
func (mb *MetricsBuilder) RecordSystemDiskIoDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemDiskIo.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemDiskIoTimeDataPoint adds a data point to system.disk.io_time metric.
func (mb *MetricsBuilder) RecordSystemDiskIoTimeDataPoint(ts pcommon.Timestamp, val float64, deviceAttributeValue string) {
	mb.metricSystemDiskIoTime.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue)
}

// RecordSystemDiskMergedDataPoint adds a data point to system.disk.merged metric.
func (mb *MetricsBuilder) RecordSystemDiskMergedDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemDiskMerged.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemDiskOperationTimeDataPoint adds a data point to system.disk.operation_time metric.
func (mb *MetricsBuilder) RecordSystemDiskOperationTimeDataPoint(ts pcommon.Timestamp, val float64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemDiskOperationTime.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemDiskOperationsDataPoint adds a data point to system.disk.operations metric.
func (mb *MetricsBuilder) RecordSystemDiskOperationsDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemDiskOperations.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemDiskPendingOperationsDataPoint adds a data point to system.disk.pending_operations metric.
func (mb *MetricsBuilder) RecordSystemDiskPendingOperationsDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string) {
	mb.metricSystemDiskPendingOperations.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue)
}

// RecordSystemDiskWeightedIoTimeDataPoint adds a data point to system.disk.weighted_io_time metric.
func (mb *MetricsBuilder) RecordSystemDiskWeightedIoTimeDataPoint(ts pcommon.Timestamp, val float64, deviceAttributeValue string) {
	mb.metricSystemDiskWeightedIoTime.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
