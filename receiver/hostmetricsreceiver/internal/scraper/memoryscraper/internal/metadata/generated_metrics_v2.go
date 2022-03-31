// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"strconv"
	"time"

	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/receiver/scrapererror"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/memory metrics.
type MetricsSettings struct {
	SystemMemoryUsage       MetricSettings `mapstructure:"system.memory.usage"`
	SystemMemoryUtilization MetricSettings `mapstructure:"system.memory.utilization"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemMemoryUsage: MetricSettings{
			Enabled: true,
		},
		SystemMemoryUtilization: MetricSettings{
			Enabled: false,
		},
	}
}

type metricSystemMemoryUsage struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.memory.usage metric with initial data.
func (m *metricSystemMemoryUsage) init() {
	m.data.SetName("system.memory.usage")
	m.data.SetDescription("Bytes of memory in use.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemMemoryUsage) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemMemoryUsage) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemMemoryUsage(settings MetricSettings) metricSystemMemoryUsage {
	m := metricSystemMemoryUsage{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemMemoryUtilization struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.memory.utilization metric with initial data.
func (m *metricSystemMemoryUtilization) init() {
	m.data.SetName("system.memory.utilization")
	m.data.SetDescription("Percentage of memory bytes in use.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemMemoryUtilization) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemMemoryUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemMemoryUtilization) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemMemoryUtilization(settings MetricSettings) metricSystemMemoryUtilization {
	m := metricSystemMemoryUtilization{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                     pdata.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity               int             // maximum observed number of metrics per resource.
	resourceCapacity              int             // maximum observed number of resource attributes.
	metricsBuffer                 pdata.Metrics   // accumulates metrics data before emitting.
	metricSystemMemoryUsage       metricSystemMemoryUsage
	metricSystemMemoryUtilization metricSystemMemoryUtilization
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
		startTime:                     pdata.NewTimestampFromTime(time.Now()),
		metricsBuffer:                 pdata.NewMetrics(),
		metricSystemMemoryUsage:       newMetricSystemMemoryUsage(settings.SystemMemoryUsage),
		metricSystemMemoryUtilization: newMetricSystemMemoryUtilization(settings.SystemMemoryUtilization),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pdata.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceOption applies changes to provided resource.
type ResourceOption func(pdata.Resource)

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pdata.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/memory")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemMemoryUsage.emit(ils.Metrics())
	mb.metricSystemMemoryUtilization.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pdata.Metrics {
	mb.EmitForResource(ro...)
	metrics := pdata.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSystemMemoryUsageDataPoint adds a data point to system.memory.usage metric.
func (mb *MetricsBuilder) RecordSystemMemoryUsageDataPoint(ts pdata.Timestamp, val int64, stateAttributeValue string) {
	mb.metricSystemMemoryUsage.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// ParseSystemMemoryUsageDataPoint attempts to parse and add a data point to system.memory.usage metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseSystemMemoryUsageDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, stateAttributeValue string) bool {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
		return false
	} else {
		mb.metricSystemMemoryUsage.recordDataPoint(mb.startTime, ts, i, stateAttributeValue)
		return true
	}
}

// RecordSystemMemoryUtilizationDataPoint adds a data point to system.memory.utilization metric.
func (mb *MetricsBuilder) RecordSystemMemoryUtilizationDataPoint(ts pdata.Timestamp, val float64, stateAttributeValue string) {
	mb.metricSystemMemoryUtilization.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// ParseSystemMemoryUtilizationDataPoint attempts to parse and add a data point to system.memory.utilization metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseSystemMemoryUtilizationDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, stateAttributeValue string) bool {
	if f, err := strconv.ParseFloat(val, 64); err != nil {
		errors.AddPartial(1, err)
		return false
	} else {
		mb.metricSystemMemoryUtilization.recordDataPoint(mb.startTime, ts, f, stateAttributeValue)
		return true
	}
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// State (Breakdown of memory usage by type.)
	State string
}{
	"state",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Buffered          string
	Cached            string
	Inactive          string
	Free              string
	SlabReclaimable   string
	SlabUnreclaimable string
	Used              string
}{
	"buffered",
	"cached",
	"inactive",
	"free",
	"slab_reclaimable",
	"slab_unreclaimable",
	"used",
}
