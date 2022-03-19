// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/disk metrics.
type MetricsSettings struct {
	SystemDiskIo                MetricSettings `mapstructure:"system.disk.io"`
	SystemDiskIoTime            MetricSettings `mapstructure:"system.disk.io_time"`
	SystemDiskMerged            MetricSettings `mapstructure:"system.disk.merged"`
	SystemDiskOperationTime     MetricSettings `mapstructure:"system.disk.operation_time"`
	SystemDiskOperations        MetricSettings `mapstructure:"system.disk.operations"`
	SystemDiskPendingOperations MetricSettings `mapstructure:"system.disk.pending_operations"`
	SystemDiskWeightedIoTime    MetricSettings `mapstructure:"system.disk.weighted_io_time"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemDiskIo: MetricSettings{
			Enabled: true,
		},
		SystemDiskIoTime: MetricSettings{
			Enabled: true,
		},
		SystemDiskMerged: MetricSettings{
			Enabled: true,
		},
		SystemDiskOperationTime: MetricSettings{
			Enabled: true,
		},
		SystemDiskOperations: MetricSettings{
			Enabled: true,
		},
		SystemDiskPendingOperations: MetricSettings{
			Enabled: true,
		},
		SystemDiskWeightedIoTime: MetricSettings{
			Enabled: true,
		},
	}
}

type metricSystemDiskIo struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.io metric with initial data.
func (m *metricSystemDiskIo) init() {
	m.data.SetName("system.disk.io")
	m.data.SetDescription("Disk bytes transferred.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskIo) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskIo) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskIo(settings MetricSettings) metricSystemDiskIo {
	m := metricSystemDiskIo{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskIoTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.io_time metric with initial data.
func (m *metricSystemDiskIoTime) init() {
	m.data.SetName("system.disk.io_time")
	m.data.SetDescription("Time disk spent activated. On Windows, this is calculated as the inverse of disk idle time.")
	m.data.SetUnit("s")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskIoTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, deviceAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskIoTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskIoTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskIoTime(settings MetricSettings) metricSystemDiskIoTime {
	m := metricSystemDiskIoTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskMerged struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.merged metric with initial data.
func (m *metricSystemDiskMerged) init() {
	m.data.SetName("system.disk.merged")
	m.data.SetDescription("The number of disk reads merged into single physical disk access operations.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskMerged) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskMerged) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskMerged) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskMerged(settings MetricSettings) metricSystemDiskMerged {
	m := metricSystemDiskMerged{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskOperationTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.operation_time metric with initial data.
func (m *metricSystemDiskOperationTime) init() {
	m.data.SetName("system.disk.operation_time")
	m.data.SetDescription("Time spent in disk operations.")
	m.data.SetUnit("s")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskOperationTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskOperationTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskOperationTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskOperationTime(settings MetricSettings) metricSystemDiskOperationTime {
	m := metricSystemDiskOperationTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.operations metric with initial data.
func (m *metricSystemDiskOperations) init() {
	m.data.SetName("system.disk.operations")
	m.data.SetDescription("Disk operations count.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskOperations(settings MetricSettings) metricSystemDiskOperations {
	m := metricSystemDiskOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskPendingOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.pending_operations metric with initial data.
func (m *metricSystemDiskPendingOperations) init() {
	m.data.SetName("system.disk.pending_operations")
	m.data.SetDescription("The queue size of pending I/O operations.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskPendingOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskPendingOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskPendingOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskPendingOperations(settings MetricSettings) metricSystemDiskPendingOperations {
	m := metricSystemDiskPendingOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemDiskWeightedIoTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.disk.weighted_io_time metric with initial data.
func (m *metricSystemDiskWeightedIoTime) init() {
	m.data.SetName("system.disk.weighted_io_time")
	m.data.SetDescription("Time disk spent activated multiplied by the queue length.")
	m.data.SetUnit("s")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemDiskWeightedIoTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, deviceAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemDiskWeightedIoTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemDiskWeightedIoTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemDiskWeightedIoTime(settings MetricSettings) metricSystemDiskWeightedIoTime {
	m := metricSystemDiskWeightedIoTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                         pdata.Timestamp
	metricSystemDiskIo                metricSystemDiskIo
	metricSystemDiskIoTime            metricSystemDiskIoTime
	metricSystemDiskMerged            metricSystemDiskMerged
	metricSystemDiskOperationTime     metricSystemDiskOperationTime
	metricSystemDiskOperations        metricSystemDiskOperations
	metricSystemDiskPendingOperations metricSystemDiskPendingOperations
	metricSystemDiskWeightedIoTime    metricSystemDiskWeightedIoTime
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                         pdata.NewTimestampFromTime(time.Now()),
		metricSystemDiskIo:                newMetricSystemDiskIo(settings.SystemDiskIo),
		metricSystemDiskIoTime:            newMetricSystemDiskIoTime(settings.SystemDiskIoTime),
		metricSystemDiskMerged:            newMetricSystemDiskMerged(settings.SystemDiskMerged),
		metricSystemDiskOperationTime:     newMetricSystemDiskOperationTime(settings.SystemDiskOperationTime),
		metricSystemDiskOperations:        newMetricSystemDiskOperations(settings.SystemDiskOperations),
		metricSystemDiskPendingOperations: newMetricSystemDiskPendingOperations(settings.SystemDiskPendingOperations),
		metricSystemDiskWeightedIoTime:    newMetricSystemDiskWeightedIoTime(settings.SystemDiskWeightedIoTime),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	mb.metricSystemDiskIo.emit(metrics)
	mb.metricSystemDiskIoTime.emit(metrics)
	mb.metricSystemDiskMerged.emit(metrics)
	mb.metricSystemDiskOperationTime.emit(metrics)
	mb.metricSystemDiskOperations.emit(metrics)
	mb.metricSystemDiskPendingOperations.emit(metrics)
	mb.metricSystemDiskWeightedIoTime.emit(metrics)
}

// RecordSystemDiskIoDataPoint adds a data point to system.disk.io metric.
func (mb *MetricsBuilder) RecordSystemDiskIoDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemDiskIo.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemDiskIoTimeDataPoint adds a data point to system.disk.io_time metric.
func (mb *MetricsBuilder) RecordSystemDiskIoTimeDataPoint(ts pdata.Timestamp, val float64, deviceAttributeValue string) {
	mb.metricSystemDiskIoTime.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue)
}

// RecordSystemDiskMergedDataPoint adds a data point to system.disk.merged metric.
func (mb *MetricsBuilder) RecordSystemDiskMergedDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemDiskMerged.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemDiskOperationTimeDataPoint adds a data point to system.disk.operation_time metric.
func (mb *MetricsBuilder) RecordSystemDiskOperationTimeDataPoint(ts pdata.Timestamp, val float64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemDiskOperationTime.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemDiskOperationsDataPoint adds a data point to system.disk.operations metric.
func (mb *MetricsBuilder) RecordSystemDiskOperationsDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemDiskOperations.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemDiskPendingOperationsDataPoint adds a data point to system.disk.pending_operations metric.
func (mb *MetricsBuilder) RecordSystemDiskPendingOperationsDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string) {
	mb.metricSystemDiskPendingOperations.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue)
}

// RecordSystemDiskWeightedIoTimeDataPoint adds a data point to system.disk.weighted_io_time metric.
func (mb *MetricsBuilder) RecordSystemDiskWeightedIoTimeDataPoint(ts pdata.Timestamp, val float64, deviceAttributeValue string) {
	mb.metricSystemDiskWeightedIoTime.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// NewMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) NewMetricData() pdata.Metrics {
	md := pdata.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/hostmetricsreceiver/disk")
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Device (Name of the disk.)
	Device string
	// Direction (Direction of flow of bytes/operations (read or write).)
	Direction string
}{
	"device",
	"direction",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	Read  string
	Write string
}{
	"read",
	"write",
}
