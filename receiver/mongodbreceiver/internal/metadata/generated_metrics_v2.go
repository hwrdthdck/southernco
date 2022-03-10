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

// MetricsSettings provides settings for mongodbreceiver metrics.
type MetricsSettings struct {
	MongodbCacheOperations MetricSettings `mapstructure:"mongodb.cache.operations"`
	MongodbCollectionCount MetricSettings `mapstructure:"mongodb.collection.count"`
	MongodbConnectionCount MetricSettings `mapstructure:"mongodb.connection.count"`
	MongodbDataSize        MetricSettings `mapstructure:"mongodb.data.size"`
	MongodbExtentCount     MetricSettings `mapstructure:"mongodb.extent.count"`
	MongodbGlobalLockTime  MetricSettings `mapstructure:"mongodb.global_lock.time"`
	MongodbIndexCount      MetricSettings `mapstructure:"mongodb.index.count"`
	MongodbIndexSize       MetricSettings `mapstructure:"mongodb.index.size"`
	MongodbMemoryUsage     MetricSettings `mapstructure:"mongodb.memory.usage"`
	MongodbObjectCount     MetricSettings `mapstructure:"mongodb.object.count"`
	MongodbOperationCount  MetricSettings `mapstructure:"mongodb.operation.count"`
	MongodbStorageSize     MetricSettings `mapstructure:"mongodb.storage.size"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		MongodbCacheOperations: MetricSettings{
			Enabled: true,
		},
		MongodbCollectionCount: MetricSettings{
			Enabled: true,
		},
		MongodbConnectionCount: MetricSettings{
			Enabled: true,
		},
		MongodbDataSize: MetricSettings{
			Enabled: true,
		},
		MongodbExtentCount: MetricSettings{
			Enabled: true,
		},
		MongodbGlobalLockTime: MetricSettings{
			Enabled: true,
		},
		MongodbIndexCount: MetricSettings{
			Enabled: true,
		},
		MongodbIndexSize: MetricSettings{
			Enabled: true,
		},
		MongodbMemoryUsage: MetricSettings{
			Enabled: true,
		},
		MongodbObjectCount: MetricSettings{
			Enabled: true,
		},
		MongodbOperationCount: MetricSettings{
			Enabled: true,
		},
		MongodbStorageSize: MetricSettings{
			Enabled: true,
		},
	}
}

type metricMongodbCacheOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.cache.operations metric with initial data.
func (m *metricMongodbCacheOperations) init() {
	m.data.SetName("mongodb.cache.operations")
	m.data.SetDescription("The number of cache operations of the instance.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbCacheOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, typeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Type, pdata.NewAttributeValueString(typeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbCacheOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbCacheOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbCacheOperations(settings MetricSettings) metricMongodbCacheOperations {
	m := metricMongodbCacheOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbCollectionCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.collection.count metric with initial data.
func (m *metricMongodbCollectionCount) init() {
	m.data.SetName("mongodb.collection.count")
	m.data.SetDescription("The number of collections.")
	m.data.SetUnit("{collections}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbCollectionCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbCollectionCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbCollectionCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbCollectionCount(settings MetricSettings) metricMongodbCollectionCount {
	m := metricMongodbCollectionCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbConnectionCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.connection.count metric with initial data.
func (m *metricMongodbConnectionCount) init() {
	m.data.SetName("mongodb.connection.count")
	m.data.SetDescription("The number of connections.")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbConnectionCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, connectionTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.ConnectionType, pdata.NewAttributeValueString(connectionTypeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbConnectionCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbConnectionCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbConnectionCount(settings MetricSettings) metricMongodbConnectionCount {
	m := metricMongodbConnectionCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbDataSize struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.data.size metric with initial data.
func (m *metricMongodbDataSize) init() {
	m.data.SetName("mongodb.data.size")
	m.data.SetDescription("The size of the collection. Data compression does not affect this value.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbDataSize) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbDataSize) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbDataSize) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbDataSize(settings MetricSettings) metricMongodbDataSize {
	m := metricMongodbDataSize{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbExtentCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.extent.count metric with initial data.
func (m *metricMongodbExtentCount) init() {
	m.data.SetName("mongodb.extent.count")
	m.data.SetDescription("The number of extents.")
	m.data.SetUnit("{extents}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbExtentCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbExtentCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbExtentCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbExtentCount(settings MetricSettings) metricMongodbExtentCount {
	m := metricMongodbExtentCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbGlobalLockTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.global_lock.time metric with initial data.
func (m *metricMongodbGlobalLockTime) init() {
	m.data.SetName("mongodb.global_lock.time")
	m.data.SetDescription("The time the global lock has been held.")
	m.data.SetUnit("ms")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricMongodbGlobalLockTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbGlobalLockTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbGlobalLockTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbGlobalLockTime(settings MetricSettings) metricMongodbGlobalLockTime {
	m := metricMongodbGlobalLockTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbIndexCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.index.count metric with initial data.
func (m *metricMongodbIndexCount) init() {
	m.data.SetName("mongodb.index.count")
	m.data.SetDescription("The number of indexes.")
	m.data.SetUnit("{indexes}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbIndexCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbIndexCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbIndexCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbIndexCount(settings MetricSettings) metricMongodbIndexCount {
	m := metricMongodbIndexCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbIndexSize struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.index.size metric with initial data.
func (m *metricMongodbIndexSize) init() {
	m.data.SetName("mongodb.index.size")
	m.data.SetDescription("Sum of the space allocated to all indexes in the database, including free index space.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbIndexSize) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbIndexSize) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbIndexSize) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbIndexSize(settings MetricSettings) metricMongodbIndexSize {
	m := metricMongodbIndexSize{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbMemoryUsage struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.memory.usage metric with initial data.
func (m *metricMongodbMemoryUsage) init() {
	m.data.SetName("mongodb.memory.usage")
	m.data.SetDescription("The amount of memory used.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbMemoryUsage) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, memoryTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.MemoryType, pdata.NewAttributeValueString(memoryTypeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbMemoryUsage) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbMemoryUsage(settings MetricSettings) metricMongodbMemoryUsage {
	m := metricMongodbMemoryUsage{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbObjectCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.object.count metric with initial data.
func (m *metricMongodbObjectCount) init() {
	m.data.SetName("mongodb.object.count")
	m.data.SetDescription("The number of objects.")
	m.data.SetUnit("{objects}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbObjectCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbObjectCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbObjectCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbObjectCount(settings MetricSettings) metricMongodbObjectCount {
	m := metricMongodbObjectCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbOperationCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.operation.count metric with initial data.
func (m *metricMongodbOperationCount) init() {
	m.data.SetName("mongodb.operation.count")
	m.data.SetDescription("The number of operations executed.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbOperationCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Operation, pdata.NewAttributeValueString(operationAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbOperationCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbOperationCount(settings MetricSettings) metricMongodbOperationCount {
	m := metricMongodbOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricMongodbStorageSize struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mongodb.storage.size metric with initial data.
func (m *metricMongodbStorageSize) init() {
	m.data.SetName("mongodb.storage.size")
	m.data.SetDescription("The total amount of storage allocated to this collection.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMongodbStorageSize) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMongodbStorageSize) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMongodbStorageSize) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMongodbStorageSize(settings MetricSettings) metricMongodbStorageSize {
	m := metricMongodbStorageSize{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	data                         *pdata.Metrics // data buffer for generated metric.
	startTime                    pdata.Timestamp
	metricMongodbCacheOperations metricMongodbCacheOperations
	metricMongodbCollectionCount metricMongodbCollectionCount
	metricMongodbConnectionCount metricMongodbConnectionCount
	metricMongodbDataSize        metricMongodbDataSize
	metricMongodbExtentCount     metricMongodbExtentCount
	metricMongodbGlobalLockTime  metricMongodbGlobalLockTime
	metricMongodbIndexCount      metricMongodbIndexCount
	metricMongodbIndexSize       metricMongodbIndexSize
	metricMongodbMemoryUsage     metricMongodbMemoryUsage
	metricMongodbObjectCount     metricMongodbObjectCount
	metricMongodbOperationCount  metricMongodbOperationCount
	metricMongodbStorageSize     metricMongodbStorageSize
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
		startTime:                    pdata.NewTimestampFromTime(time.Now()),
		metricMongodbCacheOperations: newMetricMongodbCacheOperations(settings.MongodbCacheOperations),
		metricMongodbCollectionCount: newMetricMongodbCollectionCount(settings.MongodbCollectionCount),
		metricMongodbConnectionCount: newMetricMongodbConnectionCount(settings.MongodbConnectionCount),
		metricMongodbDataSize:        newMetricMongodbDataSize(settings.MongodbDataSize),
		metricMongodbExtentCount:     newMetricMongodbExtentCount(settings.MongodbExtentCount),
		metricMongodbGlobalLockTime:  newMetricMongodbGlobalLockTime(settings.MongodbGlobalLockTime),
		metricMongodbIndexCount:      newMetricMongodbIndexCount(settings.MongodbIndexCount),
		metricMongodbIndexSize:       newMetricMongodbIndexSize(settings.MongodbIndexSize),
		metricMongodbMemoryUsage:     newMetricMongodbMemoryUsage(settings.MongodbMemoryUsage),
		metricMongodbObjectCount:     newMetricMongodbObjectCount(settings.MongodbObjectCount),
		metricMongodbOperationCount:  newMetricMongodbOperationCount(settings.MongodbOperationCount),
		metricMongodbStorageSize:     newMetricMongodbStorageSize(settings.MongodbStorageSize),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit() pdata.Metrics {
	if mb.data == nil {
		return pdata.NewMetrics()
	}
	mb.metricMongodbCacheOperations.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbCollectionCount.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbConnectionCount.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbDataSize.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbExtentCount.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbGlobalLockTime.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbIndexCount.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbIndexSize.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbMemoryUsage.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbObjectCount.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbOperationCount.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	mb.metricMongodbStorageSize.emit(mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics())
	return *mb.data
}

// EnsureCapacity TODO
func (mb *MetricsBuilder) EnsureCapacity(length int) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().EnsureCapacity(length)
}

// IncreaseCapacity TODO
func (mb *MetricsBuilder) IncreaseCapacity(length int) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().EnsureCapacity(length + mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().Len())
}

// RecordMongodbCacheOperationsDataPoint adds a data point to mongodb.cache.operations metric.
func (mb *MetricsBuilder) RecordMongodbCacheOperationsDataPoint(ts pdata.Timestamp, val int64, typeAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbCacheOperations.recordDataPoint(mb.startTime, ts, val, typeAttributeValue)
}

// RecordMongodbCollectionCountDataPoint adds a data point to mongodb.collection.count metric.
func (mb *MetricsBuilder) RecordMongodbCollectionCountDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbCollectionCount.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordMongodbConnectionCountDataPoint adds a data point to mongodb.connection.count metric.
func (mb *MetricsBuilder) RecordMongodbConnectionCountDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, connectionTypeAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbConnectionCount.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, connectionTypeAttributeValue)
}

// RecordMongodbDataSizeDataPoint adds a data point to mongodb.data.size metric.
func (mb *MetricsBuilder) RecordMongodbDataSizeDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbDataSize.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordMongodbExtentCountDataPoint adds a data point to mongodb.extent.count metric.
func (mb *MetricsBuilder) RecordMongodbExtentCountDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbExtentCount.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordMongodbGlobalLockTimeDataPoint adds a data point to mongodb.global_lock.time metric.
func (mb *MetricsBuilder) RecordMongodbGlobalLockTimeDataPoint(ts pdata.Timestamp, val int64) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbGlobalLockTime.recordDataPoint(mb.startTime, ts, val)
}

// RecordMongodbIndexCountDataPoint adds a data point to mongodb.index.count metric.
func (mb *MetricsBuilder) RecordMongodbIndexCountDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbIndexCount.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordMongodbIndexSizeDataPoint adds a data point to mongodb.index.size metric.
func (mb *MetricsBuilder) RecordMongodbIndexSizeDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbIndexSize.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordMongodbMemoryUsageDataPoint adds a data point to mongodb.memory.usage metric.
func (mb *MetricsBuilder) RecordMongodbMemoryUsageDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, memoryTypeAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbMemoryUsage.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, memoryTypeAttributeValue)
}

// RecordMongodbObjectCountDataPoint adds a data point to mongodb.object.count metric.
func (mb *MetricsBuilder) RecordMongodbObjectCountDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbObjectCount.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordMongodbOperationCountDataPoint adds a data point to mongodb.operation.count metric.
func (mb *MetricsBuilder) RecordMongodbOperationCountDataPoint(ts pdata.Timestamp, val int64, operationAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbOperationCount.recordDataPoint(mb.startTime, ts, val, operationAttributeValue)
}

// RecordMongodbStorageSizeDataPoint adds a data point to mongodb.storage.size metric.
func (mb *MetricsBuilder) RecordMongodbStorageSizeDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.metricMongodbStorageSize.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// newMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) newMetricData() *pdata.Metrics {
	md := pdata.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/mongodbreceiver")
	return &md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// ConnectionType (The status of the connection.)
	ConnectionType string
	// Database (The name of a database.)
	Database string
	// MemoryType (The type of memory used.)
	MemoryType string
	// Operation (The MongoDB operation being counted.)
	Operation string
	// Type (The result of a cache request.)
	Type string
}{
	"type",
	"database",
	"type",
	"operation",
	"type",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeConnectionType are the possible values that the attribute "connection_type" can have.
var AttributeConnectionType = struct {
	Active    string
	Available string
	Current   string
}{
	"active",
	"available",
	"current",
}

// AttributeMemoryType are the possible values that the attribute "memory_type" can have.
var AttributeMemoryType = struct {
	Resident string
	Virtual  string
}{
	"resident",
	"virtual",
}

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Insert  string
	Query   string
	Update  string
	Delete  string
	Getmore string
	Command string
}{
	"insert",
	"query",
	"update",
	"delete",
	"getmore",
	"command",
}

// AttributeType are the possible values that the attribute "type" can have.
var AttributeType = struct {
	Hit  string
	Miss string
}{
	"hit",
	"miss",
}
