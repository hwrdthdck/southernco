// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type metricContainerBlockioIoServiceBytesRecursiveRead struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.blockio.io_service_bytes_recursive.read metric with initial data.
func (m *metricContainerBlockioIoServiceBytesRecursiveRead) init() {
	m.data.SetName("container.blockio.io_service_bytes_recursive.read")
	m.data.SetDescription("Number of bytes transferred from the disk by the container")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricContainerBlockioIoServiceBytesRecursiveRead) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerBlockioIoServiceBytesRecursiveRead) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerBlockioIoServiceBytesRecursiveRead) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerBlockioIoServiceBytesRecursiveRead(cfg MetricConfig) metricContainerBlockioIoServiceBytesRecursiveRead {
	m := metricContainerBlockioIoServiceBytesRecursiveRead{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerBlockioIoServiceBytesRecursiveWrite struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.blockio.io_service_bytes_recursive.write metric with initial data.
func (m *metricContainerBlockioIoServiceBytesRecursiveWrite) init() {
	m.data.SetName("container.blockio.io_service_bytes_recursive.write")
	m.data.SetDescription("Number of bytes transferred to the disk by the container")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricContainerBlockioIoServiceBytesRecursiveWrite) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerBlockioIoServiceBytesRecursiveWrite) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerBlockioIoServiceBytesRecursiveWrite) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerBlockioIoServiceBytesRecursiveWrite(cfg MetricConfig) metricContainerBlockioIoServiceBytesRecursiveWrite {
	m := metricContainerBlockioIoServiceBytesRecursiveWrite{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerCPUPercent struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.cpu.percent metric with initial data.
func (m *metricContainerCPUPercent) init() {
	m.data.SetName("container.cpu.percent")
	m.data.SetDescription("Percent of CPU used by the container.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricContainerCPUPercent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerCPUPercent) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerCPUPercent) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerCPUPercent(cfg MetricConfig) metricContainerCPUPercent {
	m := metricContainerCPUPercent{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerCPUUsagePercpu struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.cpu.usage.percpu metric with initial data.
func (m *metricContainerCPUUsagePercpu) init() {
	m.data.SetName("container.cpu.usage.percpu")
	m.data.SetDescription("Total CPU time consumed per CPU-core.")
	m.data.SetUnit("ns")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricContainerCPUUsagePercpu) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, coreAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("core", coreAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerCPUUsagePercpu) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerCPUUsagePercpu) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerCPUUsagePercpu(cfg MetricConfig) metricContainerCPUUsagePercpu {
	m := metricContainerCPUUsagePercpu{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerCPUUsageSystem struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.cpu.usage.system metric with initial data.
func (m *metricContainerCPUUsageSystem) init() {
	m.data.SetName("container.cpu.usage.system")
	m.data.SetDescription("System CPU usage.")
	m.data.SetUnit("ns")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricContainerCPUUsageSystem) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerCPUUsageSystem) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerCPUUsageSystem) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerCPUUsageSystem(cfg MetricConfig) metricContainerCPUUsageSystem {
	m := metricContainerCPUUsageSystem{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerCPUUsageTotal struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.cpu.usage.total metric with initial data.
func (m *metricContainerCPUUsageTotal) init() {
	m.data.SetName("container.cpu.usage.total")
	m.data.SetDescription("Total CPU time consumed.")
	m.data.SetUnit("ns")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricContainerCPUUsageTotal) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerCPUUsageTotal) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerCPUUsageTotal) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerCPUUsageTotal(cfg MetricConfig) metricContainerCPUUsageTotal {
	m := metricContainerCPUUsageTotal{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerMemoryPercent struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.memory.percent metric with initial data.
func (m *metricContainerMemoryPercent) init() {
	m.data.SetName("container.memory.percent")
	m.data.SetDescription("Percentage of memory used.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricContainerMemoryPercent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerMemoryPercent) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerMemoryPercent) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerMemoryPercent(cfg MetricConfig) metricContainerMemoryPercent {
	m := metricContainerMemoryPercent{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerMemoryUsageLimit struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.memory.usage.limit metric with initial data.
func (m *metricContainerMemoryUsageLimit) init() {
	m.data.SetName("container.memory.usage.limit")
	m.data.SetDescription("Memory limit of the container.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricContainerMemoryUsageLimit) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerMemoryUsageLimit) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerMemoryUsageLimit) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerMemoryUsageLimit(cfg MetricConfig) metricContainerMemoryUsageLimit {
	m := metricContainerMemoryUsageLimit{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerMemoryUsageTotal struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.memory.usage.total metric with initial data.
func (m *metricContainerMemoryUsageTotal) init() {
	m.data.SetName("container.memory.usage.total")
	m.data.SetDescription("Memory usage of the container.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricContainerMemoryUsageTotal) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerMemoryUsageTotal) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerMemoryUsageTotal) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerMemoryUsageTotal(cfg MetricConfig) metricContainerMemoryUsageTotal {
	m := metricContainerMemoryUsageTotal{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerNetworkIoUsageRxBytes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.network.io.usage.rx_bytes metric with initial data.
func (m *metricContainerNetworkIoUsageRxBytes) init() {
	m.data.SetName("container.network.io.usage.rx_bytes")
	m.data.SetDescription("Bytes received by the container.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricContainerNetworkIoUsageRxBytes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerNetworkIoUsageRxBytes) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerNetworkIoUsageRxBytes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerNetworkIoUsageRxBytes(cfg MetricConfig) metricContainerNetworkIoUsageRxBytes {
	m := metricContainerNetworkIoUsageRxBytes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricContainerNetworkIoUsageTxBytes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills container.network.io.usage.tx_bytes metric with initial data.
func (m *metricContainerNetworkIoUsageTxBytes) init() {
	m.data.SetName("container.network.io.usage.tx_bytes")
	m.data.SetDescription("Bytes sent by the container.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricContainerNetworkIoUsageTxBytes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricContainerNetworkIoUsageTxBytes) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricContainerNetworkIoUsageTxBytes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricContainerNetworkIoUsageTxBytes(cfg MetricConfig) metricContainerNetworkIoUsageTxBytes {
	m := metricContainerNetworkIoUsageTxBytes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                                             MetricsBuilderConfig // config of the metrics builder.
	startTime                                          pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                                    int                  // maximum observed number of metrics per resource.
	metricsBuffer                                      pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                                          component.BuildInfo  // contains version information.
	metricContainerBlockioIoServiceBytesRecursiveRead  metricContainerBlockioIoServiceBytesRecursiveRead
	metricContainerBlockioIoServiceBytesRecursiveWrite metricContainerBlockioIoServiceBytesRecursiveWrite
	metricContainerCPUPercent                          metricContainerCPUPercent
	metricContainerCPUUsagePercpu                      metricContainerCPUUsagePercpu
	metricContainerCPUUsageSystem                      metricContainerCPUUsageSystem
	metricContainerCPUUsageTotal                       metricContainerCPUUsageTotal
	metricContainerMemoryPercent                       metricContainerMemoryPercent
	metricContainerMemoryUsageLimit                    metricContainerMemoryUsageLimit
	metricContainerMemoryUsageTotal                    metricContainerMemoryUsageTotal
	metricContainerNetworkIoUsageRxBytes               metricContainerNetworkIoUsageRxBytes
	metricContainerNetworkIoUsageTxBytes               metricContainerNetworkIoUsageTxBytes
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:        mbc,
		startTime:     pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer: pmetric.NewMetrics(),
		buildInfo:     settings.BuildInfo,
		metricContainerBlockioIoServiceBytesRecursiveRead:  newMetricContainerBlockioIoServiceBytesRecursiveRead(mbc.Metrics.ContainerBlockioIoServiceBytesRecursiveRead),
		metricContainerBlockioIoServiceBytesRecursiveWrite: newMetricContainerBlockioIoServiceBytesRecursiveWrite(mbc.Metrics.ContainerBlockioIoServiceBytesRecursiveWrite),
		metricContainerCPUPercent:                          newMetricContainerCPUPercent(mbc.Metrics.ContainerCPUPercent),
		metricContainerCPUUsagePercpu:                      newMetricContainerCPUUsagePercpu(mbc.Metrics.ContainerCPUUsagePercpu),
		metricContainerCPUUsageSystem:                      newMetricContainerCPUUsageSystem(mbc.Metrics.ContainerCPUUsageSystem),
		metricContainerCPUUsageTotal:                       newMetricContainerCPUUsageTotal(mbc.Metrics.ContainerCPUUsageTotal),
		metricContainerMemoryPercent:                       newMetricContainerMemoryPercent(mbc.Metrics.ContainerMemoryPercent),
		metricContainerMemoryUsageLimit:                    newMetricContainerMemoryUsageLimit(mbc.Metrics.ContainerMemoryUsageLimit),
		metricContainerMemoryUsageTotal:                    newMetricContainerMemoryUsageTotal(mbc.Metrics.ContainerMemoryUsageTotal),
		metricContainerNetworkIoUsageRxBytes:               newMetricContainerNetworkIoUsageRxBytes(mbc.Metrics.ContainerNetworkIoUsageRxBytes),
		metricContainerNetworkIoUsageTxBytes:               newMetricContainerNetworkIoUsageTxBytes(mbc.Metrics.ContainerNetworkIoUsageTxBytes),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
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
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/podmanreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricContainerBlockioIoServiceBytesRecursiveRead.emit(ils.Metrics())
	mb.metricContainerBlockioIoServiceBytesRecursiveWrite.emit(ils.Metrics())
	mb.metricContainerCPUPercent.emit(ils.Metrics())
	mb.metricContainerCPUUsagePercpu.emit(ils.Metrics())
	mb.metricContainerCPUUsageSystem.emit(ils.Metrics())
	mb.metricContainerCPUUsageTotal.emit(ils.Metrics())
	mb.metricContainerMemoryPercent.emit(ils.Metrics())
	mb.metricContainerMemoryUsageLimit.emit(ils.Metrics())
	mb.metricContainerMemoryUsageTotal.emit(ils.Metrics())
	mb.metricContainerNetworkIoUsageRxBytes.emit(ils.Metrics())
	mb.metricContainerNetworkIoUsageTxBytes.emit(ils.Metrics())

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
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordContainerBlockioIoServiceBytesRecursiveReadDataPoint adds a data point to container.blockio.io_service_bytes_recursive.read metric.
func (mb *MetricsBuilder) RecordContainerBlockioIoServiceBytesRecursiveReadDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerBlockioIoServiceBytesRecursiveRead.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerBlockioIoServiceBytesRecursiveWriteDataPoint adds a data point to container.blockio.io_service_bytes_recursive.write metric.
func (mb *MetricsBuilder) RecordContainerBlockioIoServiceBytesRecursiveWriteDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerBlockioIoServiceBytesRecursiveWrite.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerCPUPercentDataPoint adds a data point to container.cpu.percent metric.
func (mb *MetricsBuilder) RecordContainerCPUPercentDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricContainerCPUPercent.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerCPUUsagePercpuDataPoint adds a data point to container.cpu.usage.percpu metric.
func (mb *MetricsBuilder) RecordContainerCPUUsagePercpuDataPoint(ts pcommon.Timestamp, val int64, coreAttributeValue string) {
	mb.metricContainerCPUUsagePercpu.recordDataPoint(mb.startTime, ts, val, coreAttributeValue)
}

// RecordContainerCPUUsageSystemDataPoint adds a data point to container.cpu.usage.system metric.
func (mb *MetricsBuilder) RecordContainerCPUUsageSystemDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerCPUUsageSystem.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerCPUUsageTotalDataPoint adds a data point to container.cpu.usage.total metric.
func (mb *MetricsBuilder) RecordContainerCPUUsageTotalDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerCPUUsageTotal.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerMemoryPercentDataPoint adds a data point to container.memory.percent metric.
func (mb *MetricsBuilder) RecordContainerMemoryPercentDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricContainerMemoryPercent.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerMemoryUsageLimitDataPoint adds a data point to container.memory.usage.limit metric.
func (mb *MetricsBuilder) RecordContainerMemoryUsageLimitDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerMemoryUsageLimit.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerMemoryUsageTotalDataPoint adds a data point to container.memory.usage.total metric.
func (mb *MetricsBuilder) RecordContainerMemoryUsageTotalDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerMemoryUsageTotal.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerNetworkIoUsageRxBytesDataPoint adds a data point to container.network.io.usage.rx_bytes metric.
func (mb *MetricsBuilder) RecordContainerNetworkIoUsageRxBytesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerNetworkIoUsageRxBytes.recordDataPoint(mb.startTime, ts, val)
}

// RecordContainerNetworkIoUsageTxBytesDataPoint adds a data point to container.network.io.usage.tx_bytes metric.
func (mb *MetricsBuilder) RecordContainerNetworkIoUsageTxBytesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricContainerNetworkIoUsageTxBytes.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
