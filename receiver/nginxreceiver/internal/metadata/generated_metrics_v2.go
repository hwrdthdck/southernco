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
	startTime                      pdata.Timestamp
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

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	mb.metricNginxConnectionsAccepted.emit(metrics)
	mb.metricNginxConnectionsCurrent.emit(metrics)
	mb.metricNginxConnectionsHandled.emit(metrics)
	mb.metricNginxRequests.emit(metrics)
}

// RecordNginxConnectionsAcceptedDataPoint adds a data point to nginx.connections_accepted metric.
func (mb *MetricsBuilder) RecordNginxConnectionsAcceptedDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricNginxConnectionsAccepted.recordDataPoint(mb.startTime, ts, val)
}

// RecordNginxConnectionsCurrentDataPoint adds a data point to nginx.connections_current metric.
func (mb *MetricsBuilder) RecordNginxConnectionsCurrentDataPoint(ts pdata.Timestamp, val int64, stateAttributeValue string) {
	mb.metricNginxConnectionsCurrent.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// RecordNginxConnectionsHandledDataPoint adds a data point to nginx.connections_handled metric.
func (mb *MetricsBuilder) RecordNginxConnectionsHandledDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricNginxConnectionsHandled.recordDataPoint(mb.startTime, ts, val)
}

// RecordNginxRequestsDataPoint adds a data point to nginx.requests metric.
func (mb *MetricsBuilder) RecordNginxRequestsDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricNginxRequests.recordDataPoint(mb.startTime, ts, val)
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
	ilm.InstrumentationLibrary().SetName("otelcol/nginxreceiver")
	return md
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
