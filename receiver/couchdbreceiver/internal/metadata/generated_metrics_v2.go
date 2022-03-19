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

// MetricsSettings provides settings for couchdbreceiver metrics.
type MetricsSettings struct {
	CouchdbAverageRequestTime MetricSettings `mapstructure:"couchdb.average_request_time"`
	CouchdbDatabaseOpen       MetricSettings `mapstructure:"couchdb.database.open"`
	CouchdbDatabaseOperations MetricSettings `mapstructure:"couchdb.database.operations"`
	CouchdbFileDescriptorOpen MetricSettings `mapstructure:"couchdb.file_descriptor.open"`
	CouchdbHttpdBulkRequests  MetricSettings `mapstructure:"couchdb.httpd.bulk_requests"`
	CouchdbHttpdRequests      MetricSettings `mapstructure:"couchdb.httpd.requests"`
	CouchdbHttpdResponses     MetricSettings `mapstructure:"couchdb.httpd.responses"`
	CouchdbHttpdViews         MetricSettings `mapstructure:"couchdb.httpd.views"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		CouchdbAverageRequestTime: MetricSettings{
			Enabled: true,
		},
		CouchdbDatabaseOpen: MetricSettings{
			Enabled: true,
		},
		CouchdbDatabaseOperations: MetricSettings{
			Enabled: true,
		},
		CouchdbFileDescriptorOpen: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdBulkRequests: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdRequests: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdResponses: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdViews: MetricSettings{
			Enabled: true,
		},
	}
}

type metricCouchdbAverageRequestTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.average_request_time metric with initial data.
func (m *metricCouchdbAverageRequestTime) init() {
	m.data.SetName("couchdb.average_request_time")
	m.data.SetDescription("The average duration of a served request.")
	m.data.SetUnit("ms")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
}

func (m *metricCouchdbAverageRequestTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbAverageRequestTime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbAverageRequestTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbAverageRequestTime(settings MetricSettings) metricCouchdbAverageRequestTime {
	m := metricCouchdbAverageRequestTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbDatabaseOpen struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.database.open metric with initial data.
func (m *metricCouchdbDatabaseOpen) init() {
	m.data.SetName("couchdb.database.open")
	m.data.SetDescription("The number of open databases.")
	m.data.SetUnit("{databases}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricCouchdbDatabaseOpen) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbDatabaseOpen) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbDatabaseOpen) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbDatabaseOpen(settings MetricSettings) metricCouchdbDatabaseOpen {
	m := metricCouchdbDatabaseOpen{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbDatabaseOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.database.operations metric with initial data.
func (m *metricCouchdbDatabaseOperations) init() {
	m.data.SetName("couchdb.database.operations")
	m.data.SetDescription("The number of database operations.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbDatabaseOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Operation, pdata.NewValueString(operationAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbDatabaseOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbDatabaseOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbDatabaseOperations(settings MetricSettings) metricCouchdbDatabaseOperations {
	m := metricCouchdbDatabaseOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbFileDescriptorOpen struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.file_descriptor.open metric with initial data.
func (m *metricCouchdbFileDescriptorOpen) init() {
	m.data.SetName("couchdb.file_descriptor.open")
	m.data.SetDescription("The number of open file descriptors.")
	m.data.SetUnit("{files}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricCouchdbFileDescriptorOpen) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbFileDescriptorOpen) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbFileDescriptorOpen) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbFileDescriptorOpen(settings MetricSettings) metricCouchdbFileDescriptorOpen {
	m := metricCouchdbFileDescriptorOpen{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdBulkRequests struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.bulk_requests metric with initial data.
func (m *metricCouchdbHttpdBulkRequests) init() {
	m.data.SetName("couchdb.httpd.bulk_requests")
	m.data.SetDescription("The number of bulk requests.")
	m.data.SetUnit("{requests}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricCouchdbHttpdBulkRequests) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdBulkRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdBulkRequests) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdBulkRequests(settings MetricSettings) metricCouchdbHttpdBulkRequests {
	m := metricCouchdbHttpdBulkRequests{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdRequests struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.requests metric with initial data.
func (m *metricCouchdbHttpdRequests) init() {
	m.data.SetName("couchdb.httpd.requests")
	m.data.SetDescription("The number of HTTP requests by method.")
	m.data.SetUnit("{requests}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbHttpdRequests) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, httpMethodAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.HTTPMethod, pdata.NewValueString(httpMethodAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdRequests) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdRequests(settings MetricSettings) metricCouchdbHttpdRequests {
	m := metricCouchdbHttpdRequests{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdResponses struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.responses metric with initial data.
func (m *metricCouchdbHttpdResponses) init() {
	m.data.SetName("couchdb.httpd.responses")
	m.data.SetDescription("The number of each HTTP status code.")
	m.data.SetUnit("{responses}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbHttpdResponses) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, httpStatusCodeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.HTTPStatusCode, pdata.NewValueString(httpStatusCodeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdResponses) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdResponses) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdResponses(settings MetricSettings) metricCouchdbHttpdResponses {
	m := metricCouchdbHttpdResponses{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdViews struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.views metric with initial data.
func (m *metricCouchdbHttpdViews) init() {
	m.data.SetName("couchdb.httpd.views")
	m.data.SetDescription("The number of views read.")
	m.data.SetUnit("{views}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbHttpdViews) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, viewAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.View, pdata.NewValueString(viewAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdViews) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdViews) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdViews(settings MetricSettings) metricCouchdbHttpdViews {
	m := metricCouchdbHttpdViews{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                       pdata.Timestamp
	metricCouchdbAverageRequestTime metricCouchdbAverageRequestTime
	metricCouchdbDatabaseOpen       metricCouchdbDatabaseOpen
	metricCouchdbDatabaseOperations metricCouchdbDatabaseOperations
	metricCouchdbFileDescriptorOpen metricCouchdbFileDescriptorOpen
	metricCouchdbHttpdBulkRequests  metricCouchdbHttpdBulkRequests
	metricCouchdbHttpdRequests      metricCouchdbHttpdRequests
	metricCouchdbHttpdResponses     metricCouchdbHttpdResponses
	metricCouchdbHttpdViews         metricCouchdbHttpdViews
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
		startTime:                       pdata.NewTimestampFromTime(time.Now()),
		metricCouchdbAverageRequestTime: newMetricCouchdbAverageRequestTime(settings.CouchdbAverageRequestTime),
		metricCouchdbDatabaseOpen:       newMetricCouchdbDatabaseOpen(settings.CouchdbDatabaseOpen),
		metricCouchdbDatabaseOperations: newMetricCouchdbDatabaseOperations(settings.CouchdbDatabaseOperations),
		metricCouchdbFileDescriptorOpen: newMetricCouchdbFileDescriptorOpen(settings.CouchdbFileDescriptorOpen),
		metricCouchdbHttpdBulkRequests:  newMetricCouchdbHttpdBulkRequests(settings.CouchdbHttpdBulkRequests),
		metricCouchdbHttpdRequests:      newMetricCouchdbHttpdRequests(settings.CouchdbHttpdRequests),
		metricCouchdbHttpdResponses:     newMetricCouchdbHttpdResponses(settings.CouchdbHttpdResponses),
		metricCouchdbHttpdViews:         newMetricCouchdbHttpdViews(settings.CouchdbHttpdViews),
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
	mb.metricCouchdbAverageRequestTime.emit(metrics)
	mb.metricCouchdbDatabaseOpen.emit(metrics)
	mb.metricCouchdbDatabaseOperations.emit(metrics)
	mb.metricCouchdbFileDescriptorOpen.emit(metrics)
	mb.metricCouchdbHttpdBulkRequests.emit(metrics)
	mb.metricCouchdbHttpdRequests.emit(metrics)
	mb.metricCouchdbHttpdResponses.emit(metrics)
	mb.metricCouchdbHttpdViews.emit(metrics)
}

// RecordCouchdbAverageRequestTimeDataPoint adds a data point to couchdb.average_request_time metric.
func (mb *MetricsBuilder) RecordCouchdbAverageRequestTimeDataPoint(ts pdata.Timestamp, val float64) {
	mb.metricCouchdbAverageRequestTime.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbDatabaseOpenDataPoint adds a data point to couchdb.database.open metric.
func (mb *MetricsBuilder) RecordCouchdbDatabaseOpenDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricCouchdbDatabaseOpen.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbDatabaseOperationsDataPoint adds a data point to couchdb.database.operations metric.
func (mb *MetricsBuilder) RecordCouchdbDatabaseOperationsDataPoint(ts pdata.Timestamp, val int64, operationAttributeValue string) {
	mb.metricCouchdbDatabaseOperations.recordDataPoint(mb.startTime, ts, val, operationAttributeValue)
}

// RecordCouchdbFileDescriptorOpenDataPoint adds a data point to couchdb.file_descriptor.open metric.
func (mb *MetricsBuilder) RecordCouchdbFileDescriptorOpenDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricCouchdbFileDescriptorOpen.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbHttpdBulkRequestsDataPoint adds a data point to couchdb.httpd.bulk_requests metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdBulkRequestsDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricCouchdbHttpdBulkRequests.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbHttpdRequestsDataPoint adds a data point to couchdb.httpd.requests metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdRequestsDataPoint(ts pdata.Timestamp, val int64, httpMethodAttributeValue string) {
	mb.metricCouchdbHttpdRequests.recordDataPoint(mb.startTime, ts, val, httpMethodAttributeValue)
}

// RecordCouchdbHttpdResponsesDataPoint adds a data point to couchdb.httpd.responses metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdResponsesDataPoint(ts pdata.Timestamp, val int64, httpStatusCodeAttributeValue string) {
	mb.metricCouchdbHttpdResponses.recordDataPoint(mb.startTime, ts, val, httpStatusCodeAttributeValue)
}

// RecordCouchdbHttpdViewsDataPoint adds a data point to couchdb.httpd.views metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdViewsDataPoint(ts pdata.Timestamp, val int64, viewAttributeValue string) {
	mb.metricCouchdbHttpdViews.recordDataPoint(mb.startTime, ts, val, viewAttributeValue)
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
	ilm.InstrumentationLibrary().SetName("otelcol/couchdbreceiver")
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// CouchdbNodeName (The name of the node.)
	CouchdbNodeName string
	// HTTPMethod (An HTTP request method.)
	HTTPMethod string
	// HTTPStatusCode (An HTTP status code.)
	HTTPStatusCode string
	// Operation (The operation type.)
	Operation string
	// View (The view type.)
	View string
}{
	"couchdb.node.name",
	"http.method",
	"http.status_code",
	"operation",
	"view",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeHTTPMethod are the possible values that the attribute "http.method" can have.
var AttributeHTTPMethod = struct {
	COPY    string
	DELETE  string
	GET     string
	HEAD    string
	OPTIONS string
	POST    string
	PUT     string
}{
	"COPY",
	"DELETE",
	"GET",
	"HEAD",
	"OPTIONS",
	"POST",
	"PUT",
}

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Writes string
	Reads  string
}{
	"writes",
	"reads",
}

// AttributeView are the possible values that the attribute "view" can have.
var AttributeView = struct {
	TemporaryViewReads string
	ViewReads          string
}{
	"temporary_view_reads",
	"view_reads",
}
