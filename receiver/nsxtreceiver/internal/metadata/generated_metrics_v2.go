// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for nsxtreceiver metrics.
type MetricsSettings struct {
	NsxInterfacePacketCount        MetricSettings `mapstructure:"nsx.interface.packet.count"`
	NsxInterfaceThroughput         MetricSettings `mapstructure:"nsx.interface.throughput"`
	NsxNodeCacheMemoryUsage        MetricSettings `mapstructure:"nsx.node.cache.memory.usage"`
	NsxNodeCPUUtilization          MetricSettings `mapstructure:"nsx.node.cpu.utilization"`
	NsxNodeDiskUsage               MetricSettings `mapstructure:"nsx.node.disk.usage"`
	NsxNodeDiskUtilization         MetricSettings `mapstructure:"nsx.node.disk.utilization"`
	NsxNodeLoadBalancerUtilization MetricSettings `mapstructure:"nsx.node.load_balancer.utilization"`
	NsxNodeMemoryUsage             MetricSettings `mapstructure:"nsx.node.memory.usage"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		NsxInterfacePacketCount: MetricSettings{
			Enabled: true,
		},
		NsxInterfaceThroughput: MetricSettings{
			Enabled: true,
		},
		NsxNodeCacheMemoryUsage: MetricSettings{
			Enabled: true,
		},
		NsxNodeCPUUtilization: MetricSettings{
			Enabled: true,
		},
		NsxNodeDiskUsage: MetricSettings{
			Enabled: true,
		},
		NsxNodeDiskUtilization: MetricSettings{
			Enabled: true,
		},
		NsxNodeLoadBalancerUtilization: MetricSettings{
			Enabled: true,
		},
		NsxNodeMemoryUsage: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeCPUProcessClass specifies the a value cpu.process.class attribute.
type AttributeCPUProcessClass int

const (
	_ AttributeCPUProcessClass = iota
	AttributeCPUProcessClassDatapath
	AttributeCPUProcessClassServices
)

// String returns the string representation of the AttributeCPUProcessClass.
func (av AttributeCPUProcessClass) String() string {
	switch av {
	case AttributeCPUProcessClassDatapath:
		return "datapath"
	case AttributeCPUProcessClassServices:
		return "services"
	}
	return ""
}

// MapAttributeCPUProcessClass is a helper map of string to AttributeCPUProcessClass attribute value.
var MapAttributeCPUProcessClass = map[string]AttributeCPUProcessClass{
	"datapath": AttributeCPUProcessClassDatapath,
	"services": AttributeCPUProcessClassServices,
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionReceived
	AttributeDirectionTransmitted
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionReceived:
		return "received"
	case AttributeDirectionTransmitted:
		return "transmitted"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"received":    AttributeDirectionReceived,
	"transmitted": AttributeDirectionTransmitted,
}

// AttributePacketType specifies the a value packet.type attribute.
type AttributePacketType int

const (
	_ AttributePacketType = iota
	AttributePacketTypeDropped
	AttributePacketTypeErrored
	AttributePacketTypeSuccess
)

// String returns the string representation of the AttributePacketType.
func (av AttributePacketType) String() string {
	switch av {
	case AttributePacketTypeDropped:
		return "dropped"
	case AttributePacketTypeErrored:
		return "errored"
	case AttributePacketTypeSuccess:
		return "success"
	}
	return ""
}

// MapAttributePacketType is a helper map of string to AttributePacketType attribute value.
var MapAttributePacketType = map[string]AttributePacketType{
	"dropped": AttributePacketTypeDropped,
	"errored": AttributePacketTypeErrored,
	"success": AttributePacketTypeSuccess,
}

type metricNsxInterfacePacketCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.interface.packet.count metric with initial data.
func (m *metricNsxInterfacePacketCount) init() {
	m.data.SetName("nsx.interface.packet.count")
	m.data.SetDescription("The number of packets flowing through the network interface on the node.")
	m.data.SetUnit("{packets}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxInterfacePacketCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string, packetTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Direction, pcommon.NewValueString(directionAttributeValue))
	dp.Attributes().Insert(A.PacketType, pcommon.NewValueString(packetTypeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxInterfacePacketCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxInterfacePacketCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxInterfacePacketCount(settings MetricSettings) metricNsxInterfacePacketCount {
	m := metricNsxInterfacePacketCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxInterfaceThroughput struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.interface.throughput metric with initial data.
func (m *metricNsxInterfaceThroughput) init() {
	m.data.SetName("nsx.interface.throughput")
	m.data.SetDescription("The number of Bytes flowing through the network interface.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxInterfaceThroughput) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Direction, pcommon.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxInterfaceThroughput) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxInterfaceThroughput) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxInterfaceThroughput(settings MetricSettings) metricNsxInterfaceThroughput {
	m := metricNsxInterfaceThroughput{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxNodeCacheMemoryUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.node.cache.memory.usage metric with initial data.
func (m *metricNsxNodeCacheMemoryUsage) init() {
	m.data.SetName("nsx.node.cache.memory.usage")
	m.data.SetDescription("The memory usage of the node's cache")
	m.data.SetUnit("KBy")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricNsxNodeCacheMemoryUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxNodeCacheMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxNodeCacheMemoryUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxNodeCacheMemoryUsage(settings MetricSettings) metricNsxNodeCacheMemoryUsage {
	m := metricNsxNodeCacheMemoryUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxNodeCPUUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.node.cpu.utilization metric with initial data.
func (m *metricNsxNodeCPUUtilization) init() {
	m.data.SetName("nsx.node.cpu.utilization")
	m.data.SetDescription("The average amount of CPU being used by the node.")
	m.data.SetUnit("%")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxNodeCPUUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, cpuProcessClassAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.CPUProcessClass, pcommon.NewValueString(cpuProcessClassAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxNodeCPUUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxNodeCPUUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxNodeCPUUtilization(settings MetricSettings) metricNsxNodeCPUUtilization {
	m := metricNsxNodeCPUUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxNodeDiskUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.node.disk.usage metric with initial data.
func (m *metricNsxNodeDiskUsage) init() {
	m.data.SetName("nsx.node.disk.usage")
	m.data.SetDescription("The amount of storage space used by the node.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxNodeDiskUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, diskAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Disk, pcommon.NewValueString(diskAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxNodeDiskUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxNodeDiskUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxNodeDiskUsage(settings MetricSettings) metricNsxNodeDiskUsage {
	m := metricNsxNodeDiskUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxNodeDiskUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.node.disk.utilization metric with initial data.
func (m *metricNsxNodeDiskUtilization) init() {
	m.data.SetName("nsx.node.disk.utilization")
	m.data.SetDescription("The percentage of storage space utilized.")
	m.data.SetUnit("%")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxNodeDiskUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, diskAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Disk, pcommon.NewValueString(diskAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxNodeDiskUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxNodeDiskUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxNodeDiskUtilization(settings MetricSettings) metricNsxNodeDiskUtilization {
	m := metricNsxNodeDiskUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxNodeLoadBalancerUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.node.load_balancer.utilization metric with initial data.
func (m *metricNsxNodeLoadBalancerUtilization) init() {
	m.data.SetName("nsx.node.load_balancer.utilization")
	m.data.SetDescription("The utilization of load balancers by the node")
	m.data.SetUnit("%")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxNodeLoadBalancerUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, loadBalancerAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.LoadBalancer, pcommon.NewValueString(loadBalancerAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxNodeLoadBalancerUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxNodeLoadBalancerUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxNodeLoadBalancerUtilization(settings MetricSettings) metricNsxNodeLoadBalancerUtilization {
	m := metricNsxNodeLoadBalancerUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxNodeMemoryUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsx.node.memory.usage metric with initial data.
func (m *metricNsxNodeMemoryUsage) init() {
	m.data.SetName("nsx.node.memory.usage")
	m.data.SetDescription("The memory usage of the node")
	m.data.SetUnit("KBy")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricNsxNodeMemoryUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxNodeMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxNodeMemoryUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxNodeMemoryUsage(settings MetricSettings) metricNsxNodeMemoryUsage {
	m := metricNsxNodeMemoryUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                            pcommon.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity                      int               // maximum observed number of metrics per resource.
	resourceCapacity                     int               // maximum observed number of resource attributes.
	metricsBuffer                        pmetric.Metrics   // accumulates metrics data before emitting.
	metricNsxInterfacePacketCount        metricNsxInterfacePacketCount
	metricNsxInterfaceThroughput         metricNsxInterfaceThroughput
	metricNsxNodeCacheMemoryUsage        metricNsxNodeCacheMemoryUsage
	metricNsxNodeCPUUtilization          metricNsxNodeCPUUtilization
	metricNsxNodeDiskUsage               metricNsxNodeDiskUsage
	metricNsxNodeDiskUtilization         metricNsxNodeDiskUtilization
	metricNsxNodeLoadBalancerUtilization metricNsxNodeLoadBalancerUtilization
	metricNsxNodeMemoryUsage             metricNsxNodeMemoryUsage
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
		startTime:                            pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                        pmetric.NewMetrics(),
		metricNsxInterfacePacketCount:        newMetricNsxInterfacePacketCount(settings.NsxInterfacePacketCount),
		metricNsxInterfaceThroughput:         newMetricNsxInterfaceThroughput(settings.NsxInterfaceThroughput),
		metricNsxNodeCacheMemoryUsage:        newMetricNsxNodeCacheMemoryUsage(settings.NsxNodeCacheMemoryUsage),
		metricNsxNodeCPUUtilization:          newMetricNsxNodeCPUUtilization(settings.NsxNodeCPUUtilization),
		metricNsxNodeDiskUsage:               newMetricNsxNodeDiskUsage(settings.NsxNodeDiskUsage),
		metricNsxNodeDiskUtilization:         newMetricNsxNodeDiskUtilization(settings.NsxNodeDiskUtilization),
		metricNsxNodeLoadBalancerUtilization: newMetricNsxNodeLoadBalancerUtilization(settings.NsxNodeLoadBalancerUtilization),
		metricNsxNodeMemoryUsage:             newMetricNsxNodeMemoryUsage(settings.NsxNodeMemoryUsage),
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

// ResourceOption applies changes to provided resource.
type ResourceOption func(pcommon.Resource)

// WithNsxInterfaceID sets provided value as "nsx.interface.id" attribute for current resource.
func WithNsxInterfaceID(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("nsx.interface.id", val)
	}
}

// WithNsxNodeID sets provided value as "nsx.node.id" attribute for current resource.
func WithNsxNodeID(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("nsx.node.id", val)
	}
}

// WithNsxNodeName sets provided value as "nsx.node.name" attribute for current resource.
func WithNsxNodeName(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("nsx.node.name", val)
	}
}

// WithNsxNodeType sets provided value as "nsx.node.type" attribute for current resource.
func WithNsxNodeType(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("nsx.node.type", val)
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/nsxtreceiver")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricNsxInterfacePacketCount.emit(ils.Metrics())
	mb.metricNsxInterfaceThroughput.emit(ils.Metrics())
	mb.metricNsxNodeCacheMemoryUsage.emit(ils.Metrics())
	mb.metricNsxNodeCPUUtilization.emit(ils.Metrics())
	mb.metricNsxNodeDiskUsage.emit(ils.Metrics())
	mb.metricNsxNodeDiskUtilization.emit(ils.Metrics())
	mb.metricNsxNodeLoadBalancerUtilization.emit(ils.Metrics())
	mb.metricNsxNodeMemoryUsage.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pmetric.Metrics {
	mb.EmitForResource(ro...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordNsxInterfacePacketCountDataPoint adds a data point to nsx.interface.packet.count metric.
func (mb *MetricsBuilder) RecordNsxInterfacePacketCountDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection, packetTypeAttributeValue AttributePacketType) {
	mb.metricNsxInterfacePacketCount.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String(), packetTypeAttributeValue.String())
}

// RecordNsxInterfaceThroughputDataPoint adds a data point to nsx.interface.throughput metric.
func (mb *MetricsBuilder) RecordNsxInterfaceThroughputDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricNsxInterfaceThroughput.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordNsxNodeCacheMemoryUsageDataPoint adds a data point to nsx.node.cache.memory.usage metric.
func (mb *MetricsBuilder) RecordNsxNodeCacheMemoryUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNsxNodeCacheMemoryUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordNsxNodeCPUUtilizationDataPoint adds a data point to nsx.node.cpu.utilization metric.
func (mb *MetricsBuilder) RecordNsxNodeCPUUtilizationDataPoint(ts pcommon.Timestamp, val float64, cpuProcessClassAttributeValue AttributeCPUProcessClass) {
	mb.metricNsxNodeCPUUtilization.recordDataPoint(mb.startTime, ts, val, cpuProcessClassAttributeValue.String())
}

// RecordNsxNodeDiskUsageDataPoint adds a data point to nsx.node.disk.usage metric.
func (mb *MetricsBuilder) RecordNsxNodeDiskUsageDataPoint(ts pcommon.Timestamp, val int64, diskAttributeValue string) {
	mb.metricNsxNodeDiskUsage.recordDataPoint(mb.startTime, ts, val, diskAttributeValue)
}

// RecordNsxNodeDiskUtilizationDataPoint adds a data point to nsx.node.disk.utilization metric.
func (mb *MetricsBuilder) RecordNsxNodeDiskUtilizationDataPoint(ts pcommon.Timestamp, val float64, diskAttributeValue string) {
	mb.metricNsxNodeDiskUtilization.recordDataPoint(mb.startTime, ts, val, diskAttributeValue)
}

// RecordNsxNodeLoadBalancerUtilizationDataPoint adds a data point to nsx.node.load_balancer.utilization metric.
func (mb *MetricsBuilder) RecordNsxNodeLoadBalancerUtilizationDataPoint(ts pcommon.Timestamp, val float64, loadBalancerAttributeValue string) {
	mb.metricNsxNodeLoadBalancerUtilization.recordDataPoint(mb.startTime, ts, val, loadBalancerAttributeValue)
}

// RecordNsxNodeMemoryUsageDataPoint adds a data point to nsx.node.memory.usage metric.
func (mb *MetricsBuilder) RecordNsxNodeMemoryUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNsxNodeMemoryUsage.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// CPUProcessClass (The CPU usage of the architecture allocated for either DPDK (datapath) or non-DPDK (services) processes)
	CPUProcessClass string
	// Direction (The direction of network flow)
	Direction string
	// Disk (The name of the mounted storage)
	Disk string
	// LoadBalancer (The name of the load balancer being utilized)
	LoadBalancer string
	// PacketType (The type of packet counter)
	PacketType string
}{
	"cpu.process.class",
	"direction",
	"disk",
	"load_balancer",
	"type",
}

// A is an alias for Attributes.
var A = Attributes
