// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
		{
			name:        "filter_set_include",
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "filter_set_exclude",
			resAttrsSet: testDataSetAll,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverBatchRequestRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverBatchSQLCompilationRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverBatchSQLRecompilationRateDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordSqlserverDatabaseCountDataPoint(ts, "1", AttributeDatabaseStatusOnline)

			allMetricsCount++
			mb.RecordSqlserverDatabaseIoDataPoint(ts, "1", "physical_filename-val", "logical_filename-val", "file_type-val", AttributeDirectionRead)

			allMetricsCount++
			mb.RecordSqlserverDatabaseLatencyDataPoint(ts, 1, "physical_filename-val", "logical_filename-val", "file_type-val", AttributeDirectionRead)

			allMetricsCount++
			mb.RecordSqlserverDatabaseOperationsDataPoint(ts, "1", "physical_filename-val", "logical_filename-val", "file_type-val", AttributeDirectionRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverLockWaitRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverLockWaitTimeAvgDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageBufferCacheHitRatioDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageCheckpointFlushRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageLazyWriteRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageLifeExpectancyDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageOperationRateDataPoint(ts, 1, AttributePageOperationsRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageSplitRateDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordSqlserverProcessesBlockedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryExecutionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalElapsedTimeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalGrantKbDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalLogicalReadsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalLogicalWritesDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalPhysicalReadsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalRowsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverQueryTotalWorkerTimeDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordSqlserverResourcePoolDiskThrottledReadRateDataPoint(ts, "1")

			allMetricsCount++
			mb.RecordSqlserverResourcePoolDiskThrottledWriteRateDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionWriteRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionLogFlushDataRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionLogFlushRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionLogFlushWaitRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionLogGrowthCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionLogShrinkCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverTransactionLogUsageDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverUserConnectionCountDataPoint(ts, 1)

			rb := mb.NewResourceBuilder()
			rb.SetSqlserverComputerName("sqlserver.computer.name-val")
			rb.SetSqlserverDatabaseName("sqlserver.database.name-val")
			rb.SetSqlserverInstanceName("sqlserver.instance.name-val")
			rb.SetSqlserverQueryHash("sqlserver.query.hash-val")
			rb.SetSqlserverQueryPlanHandle("sqlserver.query_plan.handle-val")
			rb.SetSqlserverQueryPlanHash("sqlserver.query_plan.hash-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "sqlserver.batch.request.rate":
					assert.False(t, validatedMetrics["sqlserver.batch.request.rate"], "Found a duplicate in the metrics slice: sqlserver.batch.request.rate")
					validatedMetrics["sqlserver.batch.request.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of batch requests received by SQL Server.", ms.At(i).Description())
					assert.Equal(t, "{requests}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.batch.sql_compilation.rate":
					assert.False(t, validatedMetrics["sqlserver.batch.sql_compilation.rate"], "Found a duplicate in the metrics slice: sqlserver.batch.sql_compilation.rate")
					validatedMetrics["sqlserver.batch.sql_compilation.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of SQL compilations needed.", ms.At(i).Description())
					assert.Equal(t, "{compilations}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.batch.sql_recompilation.rate":
					assert.False(t, validatedMetrics["sqlserver.batch.sql_recompilation.rate"], "Found a duplicate in the metrics slice: sqlserver.batch.sql_recompilation.rate")
					validatedMetrics["sqlserver.batch.sql_recompilation.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of SQL recompilations needed.", ms.At(i).Description())
					assert.Equal(t, "{compilations}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.database.count":
					assert.False(t, validatedMetrics["sqlserver.database.count"], "Found a duplicate in the metrics slice: sqlserver.database.count")
					validatedMetrics["sqlserver.database.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of databases", ms.At(i).Description())
					assert.Equal(t, "{databases}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database.status")
					assert.True(t, ok)
					assert.EqualValues(t, "online", attrVal.Str())
				case "sqlserver.database.io":
					assert.False(t, validatedMetrics["sqlserver.database.io"], "Found a duplicate in the metrics slice: sqlserver.database.io")
					validatedMetrics["sqlserver.database.io"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of bytes of I/O on this file.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("physical_filename")
					assert.True(t, ok)
					assert.EqualValues(t, "physical_filename-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("logical_filename")
					assert.True(t, ok)
					assert.EqualValues(t, "logical_filename-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("file_type")
					assert.True(t, ok)
					assert.EqualValues(t, "file_type-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "sqlserver.database.latency":
					assert.False(t, validatedMetrics["sqlserver.database.latency"], "Found a duplicate in the metrics slice: sqlserver.database.latency")
					validatedMetrics["sqlserver.database.latency"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total time that the users waited for I/O issued on this file.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("physical_filename")
					assert.True(t, ok)
					assert.EqualValues(t, "physical_filename-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("logical_filename")
					assert.True(t, ok)
					assert.EqualValues(t, "logical_filename-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("file_type")
					assert.True(t, ok)
					assert.EqualValues(t, "file_type-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "sqlserver.database.operations":
					assert.False(t, validatedMetrics["sqlserver.database.operations"], "Found a duplicate in the metrics slice: sqlserver.database.operations")
					validatedMetrics["sqlserver.database.operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of operations issued on the file.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("physical_filename")
					assert.True(t, ok)
					assert.EqualValues(t, "physical_filename-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("logical_filename")
					assert.True(t, ok)
					assert.EqualValues(t, "logical_filename-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("file_type")
					assert.True(t, ok)
					assert.EqualValues(t, "file_type-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "sqlserver.lock.wait.rate":
					assert.False(t, validatedMetrics["sqlserver.lock.wait.rate"], "Found a duplicate in the metrics slice: sqlserver.lock.wait.rate")
					validatedMetrics["sqlserver.lock.wait.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of lock requests resulting in a wait.", ms.At(i).Description())
					assert.Equal(t, "{requests}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.lock.wait_time.avg":
					assert.False(t, validatedMetrics["sqlserver.lock.wait_time.avg"], "Found a duplicate in the metrics slice: sqlserver.lock.wait_time.avg")
					validatedMetrics["sqlserver.lock.wait_time.avg"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average wait time for all lock requests that had to wait.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.page.buffer_cache.hit_ratio":
					assert.False(t, validatedMetrics["sqlserver.page.buffer_cache.hit_ratio"], "Found a duplicate in the metrics slice: sqlserver.page.buffer_cache.hit_ratio")
					validatedMetrics["sqlserver.page.buffer_cache.hit_ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Pages found in the buffer pool without having to read from disk.", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.page.checkpoint.flush.rate":
					assert.False(t, validatedMetrics["sqlserver.page.checkpoint.flush.rate"], "Found a duplicate in the metrics slice: sqlserver.page.checkpoint.flush.rate")
					validatedMetrics["sqlserver.page.checkpoint.flush.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of pages flushed by operations requiring dirty pages to be flushed.", ms.At(i).Description())
					assert.Equal(t, "{pages}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.page.lazy_write.rate":
					assert.False(t, validatedMetrics["sqlserver.page.lazy_write.rate"], "Found a duplicate in the metrics slice: sqlserver.page.lazy_write.rate")
					validatedMetrics["sqlserver.page.lazy_write.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of lazy writes moving dirty pages to disk.", ms.At(i).Description())
					assert.Equal(t, "{writes}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.page.life_expectancy":
					assert.False(t, validatedMetrics["sqlserver.page.life_expectancy"], "Found a duplicate in the metrics slice: sqlserver.page.life_expectancy")
					validatedMetrics["sqlserver.page.life_expectancy"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time a page will stay in the buffer pool.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.page.operation.rate":
					assert.False(t, validatedMetrics["sqlserver.page.operation.rate"], "Found a duplicate in the metrics slice: sqlserver.page.operation.rate")
					validatedMetrics["sqlserver.page.operation.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of physical database page operations issued.", ms.At(i).Description())
					assert.Equal(t, "{operations}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "sqlserver.page.split.rate":
					assert.False(t, validatedMetrics["sqlserver.page.split.rate"], "Found a duplicate in the metrics slice: sqlserver.page.split.rate")
					validatedMetrics["sqlserver.page.split.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of pages split as a result of overflowing index pages.", ms.At(i).Description())
					assert.Equal(t, "{pages}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.processes.blocked":
					assert.False(t, validatedMetrics["sqlserver.processes.blocked"], "Found a duplicate in the metrics slice: sqlserver.processes.blocked")
					validatedMetrics["sqlserver.processes.blocked"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of processes that are currently blocked", ms.At(i).Description())
					assert.Equal(t, "{processes}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.query.execution_count":
					assert.False(t, validatedMetrics["sqlserver.query.execution_count"], "Found a duplicate in the metrics slice: sqlserver.query.execution_count")
					validatedMetrics["sqlserver.query.execution_count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of executions of the query", ms.At(i).Description())
					assert.Equal(t, "{executions}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.query.total_elapsed_time":
					assert.False(t, validatedMetrics["sqlserver.query.total_elapsed_time"], "Found a duplicate in the metrics slice: sqlserver.query.total_elapsed_time")
					validatedMetrics["sqlserver.query.total_elapsed_time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total time taken by the query", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.query.total_grant_kb":
					assert.False(t, validatedMetrics["sqlserver.query.total_grant_kb"], "Found a duplicate in the metrics slice: sqlserver.query.total_grant_kb")
					validatedMetrics["sqlserver.query.total_grant_kb"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total memory granted to the query", ms.At(i).Description())
					assert.Equal(t, "kb", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.query.total_logical_reads":
					assert.False(t, validatedMetrics["sqlserver.query.total_logical_reads"], "Found a duplicate in the metrics slice: sqlserver.query.total_logical_reads")
					validatedMetrics["sqlserver.query.total_logical_reads"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The total logical reads performed by the query", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.query.total_logical_writes":
					assert.False(t, validatedMetrics["sqlserver.query.total_logical_writes"], "Found a duplicate in the metrics slice: sqlserver.query.total_logical_writes")
					validatedMetrics["sqlserver.query.total_logical_writes"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The total logical writes performed by the query", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.query.total_physical_reads":
					assert.False(t, validatedMetrics["sqlserver.query.total_physical_reads"], "Found a duplicate in the metrics slice: sqlserver.query.total_physical_reads")
					validatedMetrics["sqlserver.query.total_physical_reads"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The total physical reads performed by the query", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.query.total_rows":
					assert.False(t, validatedMetrics["sqlserver.query.total_rows"], "Found a duplicate in the metrics slice: sqlserver.query.total_rows")
					validatedMetrics["sqlserver.query.total_rows"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The total rows returned by the query", ms.At(i).Description())
					assert.Equal(t, "{rows}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.query.total_worker_time":
					assert.False(t, validatedMetrics["sqlserver.query.total_worker_time"], "Found a duplicate in the metrics slice: sqlserver.query.total_worker_time")
					validatedMetrics["sqlserver.query.total_worker_time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total CPU time taken by the query", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.resource_pool.disk.throttled.read.rate":
					assert.False(t, validatedMetrics["sqlserver.resource_pool.disk.throttled.read.rate"], "Found a duplicate in the metrics slice: sqlserver.resource_pool.disk.throttled.read.rate")
					validatedMetrics["sqlserver.resource_pool.disk.throttled.read.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of read operations that were throttled in the last second", ms.At(i).Description())
					assert.Equal(t, "{reads}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.resource_pool.disk.throttled.write.rate":
					assert.False(t, validatedMetrics["sqlserver.resource_pool.disk.throttled.write.rate"], "Found a duplicate in the metrics slice: sqlserver.resource_pool.disk.throttled.write.rate")
					validatedMetrics["sqlserver.resource_pool.disk.throttled.write.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of write operations that were throttled in the last second", ms.At(i).Description())
					assert.Equal(t, "{writes}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.transaction.rate":
					assert.False(t, validatedMetrics["sqlserver.transaction.rate"], "Found a duplicate in the metrics slice: sqlserver.transaction.rate")
					validatedMetrics["sqlserver.transaction.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of transactions started for the database (not including XTP-only transactions).", ms.At(i).Description())
					assert.Equal(t, "{transactions}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.transaction.write.rate":
					assert.False(t, validatedMetrics["sqlserver.transaction.write.rate"], "Found a duplicate in the metrics slice: sqlserver.transaction.write.rate")
					validatedMetrics["sqlserver.transaction.write.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of transactions that wrote to the database and committed.", ms.At(i).Description())
					assert.Equal(t, "{transactions}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.transaction_log.flush.data.rate":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.flush.data.rate"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.flush.data.rate")
					validatedMetrics["sqlserver.transaction_log.flush.data.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Total number of log bytes flushed.", ms.At(i).Description())
					assert.Equal(t, "By/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.transaction_log.flush.rate":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.flush.rate"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.flush.rate")
					validatedMetrics["sqlserver.transaction_log.flush.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of log flushes.", ms.At(i).Description())
					assert.Equal(t, "{flushes}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.transaction_log.flush.wait.rate":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.flush.wait.rate"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.flush.wait.rate")
					validatedMetrics["sqlserver.transaction_log.flush.wait.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of commits waiting for a transaction log flush.", ms.At(i).Description())
					assert.Equal(t, "{commits}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "sqlserver.transaction_log.growth.count":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.growth.count"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.growth.count")
					validatedMetrics["sqlserver.transaction_log.growth.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of transaction log expansions for a database.", ms.At(i).Description())
					assert.Equal(t, "{growths}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.transaction_log.shrink.count":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.shrink.count"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.shrink.count")
					validatedMetrics["sqlserver.transaction_log.shrink.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of transaction log shrinks for a database.", ms.At(i).Description())
					assert.Equal(t, "{shrinks}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.transaction_log.usage":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.usage"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.usage")
					validatedMetrics["sqlserver.transaction_log.usage"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percent of transaction log space used.", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "sqlserver.user.connection.count":
					assert.False(t, validatedMetrics["sqlserver.user.connection.count"], "Found a duplicate in the metrics slice: sqlserver.user.connection.count")
					validatedMetrics["sqlserver.user.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of users connected to the SQL Server.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
