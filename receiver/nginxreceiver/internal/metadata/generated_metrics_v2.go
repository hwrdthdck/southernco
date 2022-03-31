// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"strconv"
	"time"

	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/receiver/scrapererror"

	"go.uber.org/zap"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for nginxreceiver metrics.
type MetricsSettings struct {
	NginxConnectionsAccepted MetricSettings `mapstructure:"nginx.connections_accepted"`
	NginxConnectionsCurrent  MetricSettings `mapstructure:"nginx.connections_current"`
	NginxConnectionsHandled  MetricSettings `mapstructure:"nginx.connections_handled"`
	NginxRequests            MetricSettings `mapstructure:"nginx.requests"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		NginxConnectionsAccepted: MetricSettings{
			Enabled: true,
		},
		NginxConnectionsCurrent: MetricSettings{
			Enabled: true,
		},
		NginxConnectionsHandled: MetricSettings{
			Enabled: true,
		},
		NginxRequests: MetricSettings{
			Enabled: true,
		},
	}
}

type metricNginxConnectionsAccepted struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.connections_accepted metric with initial data.
func (m *metricNginxConnectionsAccepted) init() {
	m.data.SetName("nginx.connections_accepted")
	m.data.SetDescription("The total number of accepted client connections")
	m.data.SetUnit("connections")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricNginxConnectionsAccepted) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxConnectionsAccepted) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxConnectionsAccepted) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxConnectionsAccepted(settings MetricSettings) metricNginxConnectionsAccepted {
	m := metricNginxConnectionsAccepted{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricNginxConnectionsCurrent struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.connections_current metric with initial data.
func (m *metricNginxConnectionsCurrent) init() {
	m.data.SetName("nginx.connections_current")
	m.data.SetDescription("The current number of nginx connections by state")
	m.data.SetUnit("connections")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNginxConnectionsCurrent) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxConnectionsCurrent) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxConnectionsCurrent) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxConnectionsCurrent(settings MetricSettings) metricNginxConnectionsCurrent {
	m := metricNginxConnectionsCurrent{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricNginxConnectionsHandled struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.connections_handled metric with initial data.
func (m *metricNginxConnectionsHandled) init() {
	m.data.SetName("nginx.connections_handled")
	m.data.SetDescription("The total number of handled connections. Generally, the parameter value is the same as nginx.connections_accepted unless some resource limits have been reached (for example, the worker_connections limit).")
	m.data.SetUnit("connections")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricNginxConnectionsHandled) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxConnectionsHandled) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxConnectionsHandled) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxConnectionsHandled(settings MetricSettings) metricNginxConnectionsHandled {
	m := metricNginxConnectionsHandled{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricNginxRequests struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.requests metric with initial data.
func (m *metricNginxRequests) init() {
	m.data.SetName("nginx.requests")
	m.data.SetDescription("Total number of requests made to the server since it started")
	m.data.SetUnit("requests")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricNginxRequests) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxRequests) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxRequests(settings MetricSettings) metricNginxRequests {
	m := metricNginxRequests{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                      pdata.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity                int             // maximum observed number of metrics per resource.
	resourceCapacity               int             // maximum observed number of resource attributes.
	metricsBuffer                  pdata.Metrics   // accumulates metrics data before emitting.
	metricNginxConnectionsAccepted metricNginxConnectionsAccepted
	metricNginxConnectionsCurrent  metricNginxConnectionsCurrent
	metricNginxConnectionsHandled  metricNginxConnectionsHandled
	metricNginxRequests            metricNginxRequests
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
		startTime:                      pdata.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pdata.NewMetrics(),
		metricNginxConnectionsAccepted: newMetricNginxConnectionsAccepted(settings.NginxConnectionsAccepted),
		metricNginxConnectionsCurrent:  newMetricNginxConnectionsCurrent(settings.NginxConnectionsCurrent),
		metricNginxConnectionsHandled:  newMetricNginxConnectionsHandled(settings.NginxConnectionsHandled),
		metricNginxRequests:            newMetricNginxRequests(settings.NginxRequests),
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
	ils.Scope().SetName("otelcol/nginxreceiver")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricNginxConnectionsAccepted.emit(ils.Metrics())
	mb.metricNginxConnectionsCurrent.emit(ils.Metrics())
	mb.metricNginxConnectionsHandled.emit(ils.Metrics())
	mb.metricNginxRequests.emit(ils.Metrics())
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

func logFailedParse(logger *zap.Logger, expectedType, metric, value string) {
	logger.Info(
		"failed to parse value",
		zap.String("expectedType", expectedType),
		zap.String("metric", metric),
		zap.String("value", value),
	)
}

// RecordNginxConnectionsAcceptedDataPoint adds a data point to nginx.connections_accepted metric.
func (mb *MetricsBuilder) RecordNginxConnectionsAcceptedDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricNginxConnectionsAccepted.recordDataPoint(mb.startTime, ts, val)
}

// ParseNginxConnectionsAcceptedDataPoint attempts to parse and add a data point to nginx.connections_accepted metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseNginxConnectionsAcceptedDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, logger *zap.Logger) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
		logFailedParse(logger, "int", "NginxConnectionsAccepted", val)
	} else {
		mb.metricNginxConnectionsAccepted.recordDataPoint(mb.startTime, ts, i)
	}
}

// RecordNginxConnectionsCurrentDataPoint adds a data point to nginx.connections_current metric.
func (mb *MetricsBuilder) RecordNginxConnectionsCurrentDataPoint(ts pdata.Timestamp, val int64, stateAttributeValue string) {
	mb.metricNginxConnectionsCurrent.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// ParseNginxConnectionsCurrentDataPoint attempts to parse and add a data point to nginx.connections_current metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseNginxConnectionsCurrentDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, logger *zap.Logger, stateAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
		logFailedParse(logger, "int", "NginxConnectionsCurrent", val)
	} else {
		mb.metricNginxConnectionsCurrent.recordDataPoint(mb.startTime, ts, i, stateAttributeValue)
	}
}

// RecordNginxConnectionsHandledDataPoint adds a data point to nginx.connections_handled metric.
func (mb *MetricsBuilder) RecordNginxConnectionsHandledDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricNginxConnectionsHandled.recordDataPoint(mb.startTime, ts, val)
}

// ParseNginxConnectionsHandledDataPoint attempts to parse and add a data point to nginx.connections_handled metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseNginxConnectionsHandledDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, logger *zap.Logger) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
		logFailedParse(logger, "int", "NginxConnectionsHandled", val)
	} else {
		mb.metricNginxConnectionsHandled.recordDataPoint(mb.startTime, ts, i)
	}
}

// RecordNginxRequestsDataPoint adds a data point to nginx.requests metric.
func (mb *MetricsBuilder) RecordNginxRequestsDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricNginxRequests.recordDataPoint(mb.startTime, ts, val)
}

// ParseNginxRequestsDataPoint attempts to parse and add a data point to nginx.requests metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseNginxRequestsDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, logger *zap.Logger) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
		logFailedParse(logger, "int", "NginxRequests", val)
	} else {
		mb.metricNginxRequests.recordDataPoint(mb.startTime, ts, i)
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
	// State (The state of a connection)
	State string
}{
	"state",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Active  string
	Reading string
	Writing string
	Waiting string
}{
	"active",
	"reading",
	"writing",
	"waiting",
}
