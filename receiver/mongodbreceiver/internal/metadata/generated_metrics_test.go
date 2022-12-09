// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), receivertest.NewNopCreateSettings(), WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["mongodb.cache.operations"] = true
	mb.RecordMongodbCacheOperationsDataPoint(ts, 1, AttributeType(1))

	enabledMetrics["mongodb.collection.count"] = true
	mb.RecordMongodbCollectionCountDataPoint(ts, 1, "attr-val")

	enabledMetrics["mongodb.connection.count"] = true
	mb.RecordMongodbConnectionCountDataPoint(ts, 1, "attr-val", AttributeConnectionType(1))

	enabledMetrics["mongodb.cursor.count"] = true
	mb.RecordMongodbCursorCountDataPoint(ts, 1)

	enabledMetrics["mongodb.cursor.timeout.count"] = true
	mb.RecordMongodbCursorTimeoutCountDataPoint(ts, 1)

	enabledMetrics["mongodb.data.size"] = true
	mb.RecordMongodbDataSizeDataPoint(ts, 1, "attr-val")

	enabledMetrics["mongodb.database.count"] = true
	mb.RecordMongodbDatabaseCountDataPoint(ts, 1)

	enabledMetrics["mongodb.document.operation.count"] = true
	mb.RecordMongodbDocumentOperationCountDataPoint(ts, 1, "attr-val", AttributeOperation(1))

	enabledMetrics["mongodb.extent.count"] = true
	mb.RecordMongodbExtentCountDataPoint(ts, 1, "attr-val")

	enabledMetrics["mongodb.global_lock.time"] = true
	mb.RecordMongodbGlobalLockTimeDataPoint(ts, 1)

	enabledMetrics["mongodb.index.access.count"] = true
	mb.RecordMongodbIndexAccessCountDataPoint(ts, 1, "attr-val", "attr-val")

	enabledMetrics["mongodb.index.count"] = true
	mb.RecordMongodbIndexCountDataPoint(ts, 1, "attr-val")

	enabledMetrics["mongodb.index.size"] = true
	mb.RecordMongodbIndexSizeDataPoint(ts, 1, "attr-val")

	mb.RecordMongodbLockAcquireCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))

	mb.RecordMongodbLockAcquireTimeDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))

	mb.RecordMongodbLockAcquireWaitCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))

	mb.RecordMongodbLockDeadlockCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))

	enabledMetrics["mongodb.memory.usage"] = true
	mb.RecordMongodbMemoryUsageDataPoint(ts, 1, "attr-val", AttributeMemoryType(1))

	enabledMetrics["mongodb.network.io.receive"] = true
	mb.RecordMongodbNetworkIoReceiveDataPoint(ts, 1)

	enabledMetrics["mongodb.network.io.transmit"] = true
	mb.RecordMongodbNetworkIoTransmitDataPoint(ts, 1)

	enabledMetrics["mongodb.network.request.count"] = true
	mb.RecordMongodbNetworkRequestCountDataPoint(ts, 1)

	enabledMetrics["mongodb.object.count"] = true
	mb.RecordMongodbObjectCountDataPoint(ts, 1, "attr-val")

	enabledMetrics["mongodb.operation.count"] = true
	mb.RecordMongodbOperationCountDataPoint(ts, 1, AttributeOperation(1))

	enabledMetrics["mongodb.operation.time"] = true
	mb.RecordMongodbOperationTimeDataPoint(ts, 1, AttributeOperation(1))

	enabledMetrics["mongodb.session.count"] = true
	mb.RecordMongodbSessionCountDataPoint(ts, 1)

	enabledMetrics["mongodb.storage.size"] = true
	mb.RecordMongodbStorageSizeDataPoint(ts, 1, "attr-val")

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		MongodbCacheOperations:        MetricSettings{Enabled: true},
		MongodbCollectionCount:        MetricSettings{Enabled: true},
		MongodbConnectionCount:        MetricSettings{Enabled: true},
		MongodbCursorCount:            MetricSettings{Enabled: true},
		MongodbCursorTimeoutCount:     MetricSettings{Enabled: true},
		MongodbDataSize:               MetricSettings{Enabled: true},
		MongodbDatabaseCount:          MetricSettings{Enabled: true},
		MongodbDocumentOperationCount: MetricSettings{Enabled: true},
		MongodbExtentCount:            MetricSettings{Enabled: true},
		MongodbGlobalLockTime:         MetricSettings{Enabled: true},
		MongodbIndexAccessCount:       MetricSettings{Enabled: true},
		MongodbIndexCount:             MetricSettings{Enabled: true},
		MongodbIndexSize:              MetricSettings{Enabled: true},
		MongodbLockAcquireCount:       MetricSettings{Enabled: true},
		MongodbLockAcquireTime:        MetricSettings{Enabled: true},
		MongodbLockAcquireWaitCount:   MetricSettings{Enabled: true},
		MongodbLockDeadlockCount:      MetricSettings{Enabled: true},
		MongodbMemoryUsage:            MetricSettings{Enabled: true},
		MongodbNetworkIoReceive:       MetricSettings{Enabled: true},
		MongodbNetworkIoTransmit:      MetricSettings{Enabled: true},
		MongodbNetworkRequestCount:    MetricSettings{Enabled: true},
		MongodbObjectCount:            MetricSettings{Enabled: true},
		MongodbOperationCount:         MetricSettings{Enabled: true},
		MongodbOperationTime:          MetricSettings{Enabled: true},
		MongodbSessionCount:           MetricSettings{Enabled: true},
		MongodbStorageSize:            MetricSettings{Enabled: true},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordMongodbCacheOperationsDataPoint(ts, 1, AttributeType(1))
	mb.RecordMongodbCollectionCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbConnectionCountDataPoint(ts, 1, "attr-val", AttributeConnectionType(1))
	mb.RecordMongodbCursorCountDataPoint(ts, 1)
	mb.RecordMongodbCursorTimeoutCountDataPoint(ts, 1)
	mb.RecordMongodbDataSizeDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbDatabaseCountDataPoint(ts, 1)
	mb.RecordMongodbDocumentOperationCountDataPoint(ts, 1, "attr-val", AttributeOperation(1))
	mb.RecordMongodbExtentCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbGlobalLockTimeDataPoint(ts, 1)
	mb.RecordMongodbIndexAccessCountDataPoint(ts, 1, "attr-val", "attr-val")
	mb.RecordMongodbIndexCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbIndexSizeDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbLockAcquireCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbLockAcquireTimeDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbLockAcquireWaitCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbLockDeadlockCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbMemoryUsageDataPoint(ts, 1, "attr-val", AttributeMemoryType(1))
	mb.RecordMongodbNetworkIoReceiveDataPoint(ts, 1)
	mb.RecordMongodbNetworkIoTransmitDataPoint(ts, 1)
	mb.RecordMongodbNetworkRequestCountDataPoint(ts, 1)
	mb.RecordMongodbObjectCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbOperationCountDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordMongodbOperationTimeDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordMongodbSessionCountDataPoint(ts, 1)
	mb.RecordMongodbStorageSizeDataPoint(ts, 1, "attr-val")

	metrics := mb.Emit(WithDatabase("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("database")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "mongodb.cache.operations":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of cache operations of the instance.", ms.At(i).Description())
			assert.Equal(t, "{operations}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "hit", attrVal.Str())
			validatedMetrics["mongodb.cache.operations"] = struct{}{}
		case "mongodb.collection.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of collections.", ms.At(i).Description())
			assert.Equal(t, "{collections}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.collection.count"] = struct{}{}
		case "mongodb.connection.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of connections.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "active", attrVal.Str())
			validatedMetrics["mongodb.connection.count"] = struct{}{}
		case "mongodb.cursor.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of open cursors maintained for clients.", ms.At(i).Description())
			assert.Equal(t, "{cursors}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.cursor.count"] = struct{}{}
		case "mongodb.cursor.timeout.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of cursors that have timed out.", ms.At(i).Description())
			assert.Equal(t, "{cursors}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.cursor.timeout.count"] = struct{}{}
		case "mongodb.data.size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The size of the collection. Data compression does not affect this value.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.data.size"] = struct{}{}
		case "mongodb.database.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of existing databases.", ms.At(i).Description())
			assert.Equal(t, "{databases}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.database.count"] = struct{}{}
		case "mongodb.document.operation.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of document operations executed.", ms.At(i).Description())
			assert.Equal(t, "{documents}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "insert", attrVal.Str())
			validatedMetrics["mongodb.document.operation.count"] = struct{}{}
		case "mongodb.extent.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of extents.", ms.At(i).Description())
			assert.Equal(t, "{extents}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.extent.count"] = struct{}{}
		case "mongodb.global_lock.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The time the global lock has been held.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.global_lock.time"] = struct{}{}
		case "mongodb.index.access.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of times an index has been accessed.", ms.At(i).Description())
			assert.Equal(t, "{accesses}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("collection")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.index.access.count"] = struct{}{}
		case "mongodb.index.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of indexes.", ms.At(i).Description())
			assert.Equal(t, "{indexes}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.index.count"] = struct{}{}
		case "mongodb.index.size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Sum of the space allocated to all indexes in the database, including free index space.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.index.size"] = struct{}{}
		case "mongodb.lock.acquire.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times the lock was acquired in the specified mode.", ms.At(i).Description())
			assert.Equal(t, "{count}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_type")
			assert.True(t, ok)
			assert.Equal(t, "parallel_batch_write_mode", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_mode")
			assert.True(t, ok)
			assert.Equal(t, "shared", attrVal.Str())
			validatedMetrics["mongodb.lock.acquire.count"] = struct{}{}
		case "mongodb.lock.acquire.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Cumulative wait time for the lock acquisitions.", ms.At(i).Description())
			assert.Equal(t, "microseconds", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_type")
			assert.True(t, ok)
			assert.Equal(t, "parallel_batch_write_mode", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_mode")
			assert.True(t, ok)
			assert.Equal(t, "shared", attrVal.Str())
			validatedMetrics["mongodb.lock.acquire.time"] = struct{}{}
		case "mongodb.lock.acquire.wait_count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times the lock acquisitions encountered waits because the locks were held in a conflicting mode.", ms.At(i).Description())
			assert.Equal(t, "{count}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_type")
			assert.True(t, ok)
			assert.Equal(t, "parallel_batch_write_mode", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_mode")
			assert.True(t, ok)
			assert.Equal(t, "shared", attrVal.Str())
			validatedMetrics["mongodb.lock.acquire.wait_count"] = struct{}{}
		case "mongodb.lock.deadlock.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times the lock acquisitions encountered deadlocks.", ms.At(i).Description())
			assert.Equal(t, "{count}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_type")
			assert.True(t, ok)
			assert.Equal(t, "parallel_batch_write_mode", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("lock_mode")
			assert.True(t, ok)
			assert.Equal(t, "shared", attrVal.Str())
			validatedMetrics["mongodb.lock.deadlock.count"] = struct{}{}
		case "mongodb.memory.usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of memory used.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "resident", attrVal.Str())
			validatedMetrics["mongodb.memory.usage"] = struct{}{}
		case "mongodb.network.io.receive":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of bytes received.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.network.io.receive"] = struct{}{}
		case "mongodb.network.io.transmit":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of by transmitted.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.network.io.transmit"] = struct{}{}
		case "mongodb.network.request.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of requests received by the server.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.network.request.count"] = struct{}{}
		case "mongodb.object.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of objects.", ms.At(i).Description())
			assert.Equal(t, "{objects}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.object.count"] = struct{}{}
		case "mongodb.operation.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of operations executed.", ms.At(i).Description())
			assert.Equal(t, "{operations}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "insert", attrVal.Str())
			validatedMetrics["mongodb.operation.count"] = struct{}{}
		case "mongodb.operation.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total time spent performing operations.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "insert", attrVal.Str())
			validatedMetrics["mongodb.operation.time"] = struct{}{}
		case "mongodb.session.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of active sessions.", ms.At(i).Description())
			assert.Equal(t, "{sessions}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["mongodb.session.count"] = struct{}{}
		case "mongodb.storage.size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total amount of storage allocated to this collection.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["mongodb.storage.size"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		MongodbCacheOperations:        MetricSettings{Enabled: false},
		MongodbCollectionCount:        MetricSettings{Enabled: false},
		MongodbConnectionCount:        MetricSettings{Enabled: false},
		MongodbCursorCount:            MetricSettings{Enabled: false},
		MongodbCursorTimeoutCount:     MetricSettings{Enabled: false},
		MongodbDataSize:               MetricSettings{Enabled: false},
		MongodbDatabaseCount:          MetricSettings{Enabled: false},
		MongodbDocumentOperationCount: MetricSettings{Enabled: false},
		MongodbExtentCount:            MetricSettings{Enabled: false},
		MongodbGlobalLockTime:         MetricSettings{Enabled: false},
		MongodbIndexAccessCount:       MetricSettings{Enabled: false},
		MongodbIndexCount:             MetricSettings{Enabled: false},
		MongodbIndexSize:              MetricSettings{Enabled: false},
		MongodbLockAcquireCount:       MetricSettings{Enabled: false},
		MongodbLockAcquireTime:        MetricSettings{Enabled: false},
		MongodbLockAcquireWaitCount:   MetricSettings{Enabled: false},
		MongodbLockDeadlockCount:      MetricSettings{Enabled: false},
		MongodbMemoryUsage:            MetricSettings{Enabled: false},
		MongodbNetworkIoReceive:       MetricSettings{Enabled: false},
		MongodbNetworkIoTransmit:      MetricSettings{Enabled: false},
		MongodbNetworkRequestCount:    MetricSettings{Enabled: false},
		MongodbObjectCount:            MetricSettings{Enabled: false},
		MongodbOperationCount:         MetricSettings{Enabled: false},
		MongodbOperationTime:          MetricSettings{Enabled: false},
		MongodbSessionCount:           MetricSettings{Enabled: false},
		MongodbStorageSize:            MetricSettings{Enabled: false},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())
	mb.RecordMongodbCacheOperationsDataPoint(ts, 1, AttributeType(1))
	mb.RecordMongodbCollectionCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbConnectionCountDataPoint(ts, 1, "attr-val", AttributeConnectionType(1))
	mb.RecordMongodbCursorCountDataPoint(ts, 1)
	mb.RecordMongodbCursorTimeoutCountDataPoint(ts, 1)
	mb.RecordMongodbDataSizeDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbDatabaseCountDataPoint(ts, 1)
	mb.RecordMongodbDocumentOperationCountDataPoint(ts, 1, "attr-val", AttributeOperation(1))
	mb.RecordMongodbExtentCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbGlobalLockTimeDataPoint(ts, 1)
	mb.RecordMongodbIndexAccessCountDataPoint(ts, 1, "attr-val", "attr-val")
	mb.RecordMongodbIndexCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbIndexSizeDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbLockAcquireCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbLockAcquireTimeDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbLockAcquireWaitCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbLockDeadlockCountDataPoint(ts, 1, "attr-val", AttributeLockType(1), AttributeLockMode(1))
	mb.RecordMongodbMemoryUsageDataPoint(ts, 1, "attr-val", AttributeMemoryType(1))
	mb.RecordMongodbNetworkIoReceiveDataPoint(ts, 1)
	mb.RecordMongodbNetworkIoTransmitDataPoint(ts, 1)
	mb.RecordMongodbNetworkRequestCountDataPoint(ts, 1)
	mb.RecordMongodbObjectCountDataPoint(ts, 1, "attr-val")
	mb.RecordMongodbOperationCountDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordMongodbOperationTimeDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordMongodbSessionCountDataPoint(ts, 1)
	mb.RecordMongodbStorageSizeDataPoint(ts, 1, "attr-val")

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
