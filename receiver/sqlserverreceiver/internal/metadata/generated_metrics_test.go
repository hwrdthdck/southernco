// Code generated by mdatagen. DO NOT EDIT.
package metadata

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testMetricsSet int

const (
	testMetricsSetDefault testMetricsSet = iota
	testMetricsSetAll
	testMetricsSetNo
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name       string
		metricsSet testMetricsSet
	}{
		{
			name:       "default",
			metricsSet: testMetricsSetDefault,
		},
		{
			name:       "all_metrics",
			metricsSet: testMetricsSetAll,
		},
		{
			name:       "no_metrics",
			metricsSet: testMetricsSetNo,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadConfig(t, test.name), settings, WithStartTime(start))

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
			mb.RecordSqlserverPageOperationRateDataPoint(ts, 1, AttributePageOperations(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSqlserverPageSplitRateDataPoint(ts, 1)

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

			metrics := mb.Emit(WithSqlserverDatabaseName("attr-val"))

			if test.metricsSet == testMetricsSetNo {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("sqlserver.database.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.SqlserverDatabaseName.Enabled, ok)
			if mb.resourceAttributesSettings.SqlserverDatabaseName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 1)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testMetricsSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testMetricsSetAll {
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "read", attrVal.Str())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
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
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "sqlserver.transaction_log.growth.count":
					assert.False(t, validatedMetrics["sqlserver.transaction_log.growth.count"], "Found a duplicate in the metrics slice: sqlserver.transaction_log.growth.count")
					validatedMetrics["sqlserver.transaction_log.growth.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of transaction log expansions for a database.", ms.At(i).Description())
					assert.Equal(t, "{growths}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
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
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
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

func loadConfig(t *testing.T, name string) MetricsSettings {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsSettings()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
