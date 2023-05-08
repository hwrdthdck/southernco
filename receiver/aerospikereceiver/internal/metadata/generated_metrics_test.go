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
			mb := NewMetricsBuilder(loadConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceDiskAvailableDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceGeojsonRegionQueryCellsDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceGeojsonRegionQueryFalsePositiveDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceGeojsonRegionQueryPointsDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceGeojsonRegionQueryRequestsDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceMemoryFreeDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceMemoryUsageDataPoint(ts, "1", AttributeNamespaceComponent(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceQueryCountDataPoint(ts, "1", AttributeQueryType(1), AttributeIndexType(1), AttributeQueryResult(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceScanCountDataPoint(ts, "1", AttributeScanType(1), AttributeScanResult(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNamespaceTransactionCountDataPoint(ts, "1", AttributeTransactionType(1), AttributeTransactionResult(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNodeConnectionCountDataPoint(ts, "1", AttributeConnectionType(1), AttributeConnectionOp(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNodeConnectionOpenDataPoint(ts, "1", AttributeConnectionType(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNodeMemoryFreeDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordAerospikeNodeQueryTrackedDataPoint(ts, "1")

			metrics := mb.Emit(WithAerospikeNamespace("attr-val"), WithAerospikeNodeName("attr-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("aerospike.namespace")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.AerospikeNamespace.Enabled, ok)
			if mb.resourceAttributesConfig.AerospikeNamespace.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("aerospike.node.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.AerospikeNodeName.Enabled, ok)
			if mb.resourceAttributesConfig.AerospikeNodeName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 2)

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
				case "aerospike.namespace.disk.available":
					assert.False(t, validatedMetrics["aerospike.namespace.disk.available"], "Found a duplicate in the metrics slice: aerospike.namespace.disk.available")
					validatedMetrics["aerospike.namespace.disk.available"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Minimum percentage of contiguous disk space free to the namespace across all devices", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.namespace.geojson.region_query_cells":
					assert.False(t, validatedMetrics["aerospike.namespace.geojson.region_query_cells"], "Found a duplicate in the metrics slice: aerospike.namespace.geojson.region_query_cells")
					validatedMetrics["aerospike.namespace.geojson.region_query_cells"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of cell coverings for query region queried", ms.At(i).Description())
					assert.Equal(t, "{cells}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.namespace.geojson.region_query_false_positive":
					assert.False(t, validatedMetrics["aerospike.namespace.geojson.region_query_false_positive"], "Found a duplicate in the metrics slice: aerospike.namespace.geojson.region_query_false_positive")
					validatedMetrics["aerospike.namespace.geojson.region_query_false_positive"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of points outside the region.", ms.At(i).Description())
					assert.Equal(t, "{points}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.namespace.geojson.region_query_points":
					assert.False(t, validatedMetrics["aerospike.namespace.geojson.region_query_points"], "Found a duplicate in the metrics slice: aerospike.namespace.geojson.region_query_points")
					validatedMetrics["aerospike.namespace.geojson.region_query_points"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of points within the region.", ms.At(i).Description())
					assert.Equal(t, "{points}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.namespace.geojson.region_query_requests":
					assert.False(t, validatedMetrics["aerospike.namespace.geojson.region_query_requests"], "Found a duplicate in the metrics slice: aerospike.namespace.geojson.region_query_requests")
					validatedMetrics["aerospike.namespace.geojson.region_query_requests"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of geojson queries on the system since the uptime of the node.", ms.At(i).Description())
					assert.Equal(t, "{queries}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.namespace.memory.free":
					assert.False(t, validatedMetrics["aerospike.namespace.memory.free"], "Found a duplicate in the metrics slice: aerospike.namespace.memory.free")
					validatedMetrics["aerospike.namespace.memory.free"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of the namespace's memory which is still free", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.namespace.memory.usage":
					assert.False(t, validatedMetrics["aerospike.namespace.memory.usage"], "Found a duplicate in the metrics slice: aerospike.namespace.memory.usage")
					validatedMetrics["aerospike.namespace.memory.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Memory currently used by each component of the namespace", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("component")
					assert.True(t, ok)
					assert.Equal(t, "data", attrVal.Str())
				case "aerospike.namespace.query.count":
					assert.False(t, validatedMetrics["aerospike.namespace.query.count"], "Found a duplicate in the metrics slice: aerospike.namespace.query.count")
					validatedMetrics["aerospike.namespace.query.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of query operations performed on the namespace", ms.At(i).Description())
					assert.Equal(t, "{queries}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "aggregation", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("index")
					assert.True(t, ok)
					assert.Equal(t, "primary", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("result")
					assert.True(t, ok)
					assert.Equal(t, "abort", attrVal.Str())
				case "aerospike.namespace.scan.count":
					assert.False(t, validatedMetrics["aerospike.namespace.scan.count"], "Found a duplicate in the metrics slice: aerospike.namespace.scan.count")
					validatedMetrics["aerospike.namespace.scan.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of scan operations performed on the namespace", ms.At(i).Description())
					assert.Equal(t, "{scans}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "aggregation", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("result")
					assert.True(t, ok)
					assert.Equal(t, "abort", attrVal.Str())
				case "aerospike.namespace.transaction.count":
					assert.False(t, validatedMetrics["aerospike.namespace.transaction.count"], "Found a duplicate in the metrics slice: aerospike.namespace.transaction.count")
					validatedMetrics["aerospike.namespace.transaction.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of transactions performed on the namespace", ms.At(i).Description())
					assert.Equal(t, "{transactions}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "delete", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("result")
					assert.True(t, ok)
					assert.Equal(t, "error", attrVal.Str())
				case "aerospike.node.connection.count":
					assert.False(t, validatedMetrics["aerospike.node.connection.count"], "Found a duplicate in the metrics slice: aerospike.node.connection.count")
					validatedMetrics["aerospike.node.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of connections opened and closed to the node", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "client", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.Equal(t, "close", attrVal.Str())
				case "aerospike.node.connection.open":
					assert.False(t, validatedMetrics["aerospike.node.connection.open"], "Found a duplicate in the metrics slice: aerospike.node.connection.open")
					validatedMetrics["aerospike.node.connection.open"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of open connections to the node", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "client", attrVal.Str())
				case "aerospike.node.memory.free":
					assert.False(t, validatedMetrics["aerospike.node.memory.free"], "Found a duplicate in the metrics slice: aerospike.node.memory.free")
					validatedMetrics["aerospike.node.memory.free"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of the node's memory which is still free", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "aerospike.node.query.tracked":
					assert.False(t, validatedMetrics["aerospike.node.query.tracked"], "Found a duplicate in the metrics slice: aerospike.node.query.tracked")
					validatedMetrics["aerospike.node.query.tracked"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of queries tracked by the system.", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}

func loadConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
