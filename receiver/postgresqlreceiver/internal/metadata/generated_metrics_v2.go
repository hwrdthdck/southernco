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

// MetricsSettings provides settings for postgresqlreceiver metrics.
type MetricsSettings struct {
	PostgresqlBackends   MetricSettings `mapstructure:"postgresql.backends"`
	PostgresqlBlocksRead MetricSettings `mapstructure:"postgresql.blocks_read"`
	PostgresqlCommits    MetricSettings `mapstructure:"postgresql.commits"`
	PostgresqlDbSize     MetricSettings `mapstructure:"postgresql.db_size"`
	PostgresqlOperations MetricSettings `mapstructure:"postgresql.operations"`
	PostgresqlRollbacks  MetricSettings `mapstructure:"postgresql.rollbacks"`
	PostgresqlRows       MetricSettings `mapstructure:"postgresql.rows"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		PostgresqlBackends: MetricSettings{
			Enabled: true,
		},
		PostgresqlBlocksRead: MetricSettings{
			Enabled: true,
		},
		PostgresqlCommits: MetricSettings{
			Enabled: true,
		},
		PostgresqlDbSize: MetricSettings{
			Enabled: true,
		},
		PostgresqlOperations: MetricSettings{
			Enabled: true,
		},
		PostgresqlRollbacks: MetricSettings{
			Enabled: true,
		},
		PostgresqlRows: MetricSettings{
			Enabled: true,
		},
	}
}

type metricPostgresqlBackends struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.backends metric with initial data.
func (m *metricPostgresqlBackends) init() {
	m.data.SetName("postgresql.backends")
	m.data.SetDescription("The number of backends.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlBackends) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlBackends) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlBackends) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlBackends(settings MetricSettings) metricPostgresqlBackends {
	m := metricPostgresqlBackends{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlBlocksRead struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.blocks_read metric with initial data.
func (m *metricPostgresqlBlocksRead) init() {
	m.data.SetName("postgresql.blocks_read")
	m.data.SetDescription("The number of blocks read.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlBlocksRead) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.Table, pdata.NewValueString(tableAttributeValue))
	dp.Attributes().Insert(A.Source, pdata.NewValueString(sourceAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlBlocksRead) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlBlocksRead) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlBlocksRead(settings MetricSettings) metricPostgresqlBlocksRead {
	m := metricPostgresqlBlocksRead{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlCommits struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.commits metric with initial data.
func (m *metricPostgresqlCommits) init() {
	m.data.SetName("postgresql.commits")
	m.data.SetDescription("The number of commits.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlCommits) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlCommits) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlCommits) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlCommits(settings MetricSettings) metricPostgresqlCommits {
	m := metricPostgresqlCommits{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlDbSize struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.db_size metric with initial data.
func (m *metricPostgresqlDbSize) init() {
	m.data.SetName("postgresql.db_size")
	m.data.SetDescription("The database disk usage.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlDbSize) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlDbSize) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlDbSize) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlDbSize(settings MetricSettings) metricPostgresqlDbSize {
	m := metricPostgresqlDbSize{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.operations metric with initial data.
func (m *metricPostgresqlOperations) init() {
	m.data.SetName("postgresql.operations")
	m.data.SetDescription("The number of db row operations.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.Table, pdata.NewValueString(tableAttributeValue))
	dp.Attributes().Insert(A.Operation, pdata.NewValueString(operationAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlOperations(settings MetricSettings) metricPostgresqlOperations {
	m := metricPostgresqlOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlRollbacks struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.rollbacks metric with initial data.
func (m *metricPostgresqlRollbacks) init() {
	m.data.SetName("postgresql.rollbacks")
	m.data.SetDescription("The number of rollbacks.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlRollbacks) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlRollbacks) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlRollbacks) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlRollbacks(settings MetricSettings) metricPostgresqlRollbacks {
	m := metricPostgresqlRollbacks{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlRows struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.rows metric with initial data.
func (m *metricPostgresqlRows) init() {
	m.data.SetName("postgresql.rows")
	m.data.SetDescription("The number of rows in the database.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlRows) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.Table, pdata.NewValueString(tableAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlRows) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlRows) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlRows(settings MetricSettings) metricPostgresqlRows {
	m := metricPostgresqlRows{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                  pdata.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity            int             // maximum observed number of metrics per resource.
	resourceCapacity           int             // maximum observed number of resource attributes.
	metricsBuffer              pdata.Metrics   // accumulates metrics data before emitting.
	metricPostgresqlBackends   metricPostgresqlBackends
	metricPostgresqlBlocksRead metricPostgresqlBlocksRead
	metricPostgresqlCommits    metricPostgresqlCommits
	metricPostgresqlDbSize     metricPostgresqlDbSize
	metricPostgresqlOperations metricPostgresqlOperations
	metricPostgresqlRollbacks  metricPostgresqlRollbacks
	metricPostgresqlRows       metricPostgresqlRows
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
		startTime:                  pdata.NewTimestampFromTime(time.Now()),
		metricsBuffer:              pdata.NewMetrics(),
		metricPostgresqlBackends:   newMetricPostgresqlBackends(settings.PostgresqlBackends),
		metricPostgresqlBlocksRead: newMetricPostgresqlBlocksRead(settings.PostgresqlBlocksRead),
		metricPostgresqlCommits:    newMetricPostgresqlCommits(settings.PostgresqlCommits),
		metricPostgresqlDbSize:     newMetricPostgresqlDbSize(settings.PostgresqlDbSize),
		metricPostgresqlOperations: newMetricPostgresqlOperations(settings.PostgresqlOperations),
		metricPostgresqlRollbacks:  newMetricPostgresqlRollbacks(settings.PostgresqlRollbacks),
		metricPostgresqlRows:       newMetricPostgresqlRows(settings.PostgresqlRows),
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
	ils.Scope().SetName("otelcol/postgresqlreceiver")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricPostgresqlBackends.emit(ils.Metrics())
	mb.metricPostgresqlBlocksRead.emit(ils.Metrics())
	mb.metricPostgresqlCommits.emit(ils.Metrics())
	mb.metricPostgresqlDbSize.emit(ils.Metrics())
	mb.metricPostgresqlOperations.emit(ils.Metrics())
	mb.metricPostgresqlRollbacks.emit(ils.Metrics())
	mb.metricPostgresqlRows.emit(ils.Metrics())
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

// RecordPostgresqlBackendsDataPoint adds a data point to postgresql.backends metric.
func (mb *MetricsBuilder) RecordPostgresqlBackendsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlBackends.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// ParsePostgresqlBackendsDataPoint attempts to parse and add a data point to postgresql.backends metric.
func (mb *MetricsBuilder) ParsePostgresqlBackendsDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlBackends.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue)
	}
}

// RecordPostgresqlBlocksReadDataPoint adds a data point to postgresql.blocks_read metric.
func (mb *MetricsBuilder) RecordPostgresqlBlocksReadDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue string) {
	mb.metricPostgresqlBlocksRead.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, sourceAttributeValue)
}

// ParsePostgresqlBlocksReadDataPoint attempts to parse and add a data point to postgresql.blocks_read metric.
func (mb *MetricsBuilder) ParsePostgresqlBlocksReadDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlBlocksRead.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue, tableAttributeValue, sourceAttributeValue)
	}
}

// RecordPostgresqlCommitsDataPoint adds a data point to postgresql.commits metric.
func (mb *MetricsBuilder) RecordPostgresqlCommitsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlCommits.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// ParsePostgresqlCommitsDataPoint attempts to parse and add a data point to postgresql.commits metric.
func (mb *MetricsBuilder) ParsePostgresqlCommitsDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlCommits.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue)
	}
}

// RecordPostgresqlDbSizeDataPoint adds a data point to postgresql.db_size metric.
func (mb *MetricsBuilder) RecordPostgresqlDbSizeDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlDbSize.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// ParsePostgresqlDbSizeDataPoint attempts to parse and add a data point to postgresql.db_size metric.
func (mb *MetricsBuilder) ParsePostgresqlDbSizeDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlDbSize.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue)
	}
}

// RecordPostgresqlOperationsDataPoint adds a data point to postgresql.operations metric.
func (mb *MetricsBuilder) RecordPostgresqlOperationsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue string) {
	mb.metricPostgresqlOperations.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, operationAttributeValue)
}

// ParsePostgresqlOperationsDataPoint attempts to parse and add a data point to postgresql.operations metric.
func (mb *MetricsBuilder) ParsePostgresqlOperationsDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlOperations.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue, tableAttributeValue, operationAttributeValue)
	}
}

// RecordPostgresqlRollbacksDataPoint adds a data point to postgresql.rollbacks metric.
func (mb *MetricsBuilder) RecordPostgresqlRollbacksDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlRollbacks.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// ParsePostgresqlRollbacksDataPoint attempts to parse and add a data point to postgresql.rollbacks metric.
func (mb *MetricsBuilder) ParsePostgresqlRollbacksDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlRollbacks.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue)
	}
}

// RecordPostgresqlRowsDataPoint adds a data point to postgresql.rows metric.
func (mb *MetricsBuilder) RecordPostgresqlRowsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue string) {
	mb.metricPostgresqlRows.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, stateAttributeValue)
}

// ParsePostgresqlRowsDataPoint attempts to parse and add a data point to postgresql.rows metric.
func (mb *MetricsBuilder) ParsePostgresqlRowsDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue string) {
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		errors.AddPartial(1, err)
	} else {
		mb.metricPostgresqlRows.recordDataPoint(mb.startTime, ts, i, databaseAttributeValue, tableAttributeValue, stateAttributeValue)
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
	// Database (The name of the database.)
	Database string
	// Operation (The database operation.)
	Operation string
	// Source (The block read source type.)
	Source string
	// State (The tuple (row) state.)
	State string
	// Table (The schema name followed by the table name.)
	Table string
}{
	"database",
	"operation",
	"source",
	"state",
	"table",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Ins    string
	Upd    string
	Del    string
	HotUpd string
}{
	"ins",
	"upd",
	"del",
	"hot_upd",
}

// AttributeSource are the possible values that the attribute "source" can have.
var AttributeSource = struct {
	HeapRead  string
	HeapHit   string
	IdxRead   string
	IdxHit    string
	ToastRead string
	ToastHit  string
	TidxRead  string
	TidxHit   string
}{
	"heap_read",
	"heap_hit",
	"idx_read",
	"idx_hit",
	"toast_read",
	"toast_hit",
	"tidx_read",
	"tidx_hit",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Dead string
	Live string
}{
	"dead",
	"live",
}
