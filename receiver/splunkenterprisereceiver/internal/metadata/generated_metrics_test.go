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

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketCountDataPoint(ts, 1, "index.name-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketDirsCountDataPoint(ts, 1, "index.name-val", "bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketDirsSizeDataPoint(ts, 1, "index.name-val", "bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketEventCountDataPoint(ts, 1, "index.name-val", "bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketHotCountDataPoint(ts, 1, "index.name-val", "bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketWarmCountDataPoint(ts, 1, "index.name-val", "bucket.dir-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedEventCountDataPoint(ts, 1, "index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedRawSizeDataPoint(ts, 1, "index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedTotalSizeDataPoint(ts, 1, "index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkLicenseIndexUsageDataPoint(ts, 1, "index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkServerIntrospectionIndexerThroughputDataPoint(ts, 1, "status-val")

			res := pcommon.NewResource()
			res.Attributes().PutStr("k1", "v1")
			metrics := mb.Emit(WithResource(res))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "splunk.data.indexes.extended.bucket.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of buckets per index", ms.At(i).Description())
					assert.Equal(t, "{Buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.dirs.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.dirs.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.dirs.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.dirs.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of buckets", ms.At(i).Description())
					assert.Equal(t, "{Buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.dirs.size":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.dirs.size"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.dirs.size")
					validatedMetrics["splunk.data.indexes.extended.bucket.dirs.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Size (fractional MB) on disk of this bucket super-directory", ms.At(i).Description())
					assert.Equal(t, "MBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.event.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.event.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.event.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.event.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of events in this bucket super-directory", ms.At(i).Description())
					assert.Equal(t, "{Events}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.hot.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.hot.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.hot.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.hot.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "(If size > 0) Number of hot buckets", ms.At(i).Description())
					assert.Equal(t, "{Buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.warm.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.warm.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.warm.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.warm.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "(If size > 0) Number of warm buckets", ms.At(i).Description())
					assert.Equal(t, "{Buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.event.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.event.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.event.count")
					validatedMetrics["splunk.data.indexes.extended.event.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of events for index, excluding frozen events. Approximately equal to the event_count sum of all buckets.", ms.At(i).Description())
					assert.Equal(t, "{Events}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
				case "splunk.data.indexes.extended.raw.size":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.raw.size"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.raw.size")
					validatedMetrics["splunk.data.indexes.extended.raw.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Cumulative size (fractional MB) on disk of the <bucket>/rawdata/ directories of all buckets in this index, excluding frozen", ms.At(i).Description())
					assert.Equal(t, "MBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
				case "splunk.data.indexes.extended.total.size":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.total.size"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.total.size")
					validatedMetrics["splunk.data.indexes.extended.total.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Size (fractional MB) on disk of this index", ms.At(i).Description())
					assert.Equal(t, "MBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
				case "splunk.license.index.usage":
					assert.False(t, validatedMetrics["splunk.license.index.usage"], "Found a duplicate in the metrics slice: splunk.license.index.usage")
					validatedMetrics["splunk.license.index.usage"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexed license usage per index", ms.At(i).Description())
					assert.Equal(t, "GBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "index.name-val", attrVal.Str())
				case "splunk.server.introspection.indexer.throughput":
					assert.False(t, validatedMetrics["splunk.server.introspection.indexer.throughput"], "Found a duplicate in the metrics slice: splunk.server.introspection.indexer.throughput")
					validatedMetrics["splunk.server.introspection.indexer.throughput"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking average KBps throughput of indexer", ms.At(i).Description())
					assert.Equal(t, "KBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "status-val", attrVal.Str())
				}
			}
		})
	}
}
