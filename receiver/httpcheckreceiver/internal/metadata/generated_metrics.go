// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for httpcheckreceiver metrics.
type MetricsSettings struct {
	HttpcheckDuration MetricSettings `mapstructure:"httpcheck.duration"`
	HttpcheckStatus   MetricSettings `mapstructure:"httpcheck.status"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		HttpcheckDuration: MetricSettings{
			Enabled: true,
		},
		HttpcheckStatus: MetricSettings{
			Enabled: true,
		},
	}
}

type metricHttpcheckDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills httpcheck.duration metric with initial data.
func (m *metricHttpcheckDuration) init() {
	m.data.SetName("httpcheck.duration")
	m.data.SetDescription("Measures the duration of the HTTP check.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricHttpcheckDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, httpURLAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("http.url", httpURLAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricHttpcheckDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricHttpcheckDuration) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricHttpcheckDuration(settings MetricSettings) metricHttpcheckDuration {
	m := metricHttpcheckDuration{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricHttpcheckStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills httpcheck.status metric with initial data.
func (m *metricHttpcheckStatus) init() {
	m.data.SetName("httpcheck.status")
	m.data.SetDescription("1 if the endpoint resulted in status http.status_code, otherwise 0.")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricHttpcheckStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, httpURLAttributeValue string, httpStatusCodeAttributeValue int64, httpMethodAttributeValue string, httpStatusClassAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("http.url", httpURLAttributeValue)
	dp.Attributes().PutInt("http.status_code", httpStatusCodeAttributeValue)
	dp.Attributes().PutStr("http.method", httpMethodAttributeValue)
	dp.Attributes().PutStr("http.status_class", httpStatusClassAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricHttpcheckStatus) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricHttpcheckStatus) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricHttpcheckStatus(settings MetricSettings) metricHttpcheckStatus {
	m := metricHttpcheckStatus{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime               pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity         int                 // maximum observed number of metrics per resource.
	resourceCapacity        int                 // maximum observed number of resource attributes.
	metricsBuffer           pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo               component.BuildInfo // contains version information
	metricHttpcheckDuration metricHttpcheckDuration
	metricHttpcheckStatus   metricHttpcheckStatus
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, buildInfo component.BuildInfo, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:               pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:           pmetric.NewMetrics(),
		buildInfo:               buildInfo,
		metricHttpcheckDuration: newMetricHttpcheckDuration(settings.HttpcheckDuration),
		metricHttpcheckStatus:   newMetricHttpcheckStatus(settings.HttpcheckStatus),
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
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/httpcheckreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricHttpcheckDuration.emit(ils.Metrics())
	mb.metricHttpcheckStatus.emit(ils.Metrics())
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

// RecordHttpcheckDurationDataPoint adds a data point to httpcheck.duration metric.
func (mb *MetricsBuilder) RecordHttpcheckDurationDataPoint(ts pcommon.Timestamp, val int64, httpURLAttributeValue string) {
	mb.metricHttpcheckDuration.recordDataPoint(mb.startTime, ts, val, httpURLAttributeValue)
}

// RecordHttpcheckStatusDataPoint adds a data point to httpcheck.status metric.
func (mb *MetricsBuilder) RecordHttpcheckStatusDataPoint(ts pcommon.Timestamp, val int64, httpURLAttributeValue string, httpStatusCodeAttributeValue int64, httpMethodAttributeValue string, httpStatusClassAttributeValue string) {
	mb.metricHttpcheckStatus.recordDataPoint(mb.startTime, ts, val, httpURLAttributeValue, httpStatusCodeAttributeValue, httpMethodAttributeValue, httpStatusClassAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
