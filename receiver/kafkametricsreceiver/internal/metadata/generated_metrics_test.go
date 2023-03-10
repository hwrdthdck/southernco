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
			mb.RecordKafkaBrokersDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupLagDataPoint(ts, 1, "attr-val", "attr-val", 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupLagSumDataPoint(ts, 1, "attr-val", "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupMembersDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupOffsetDataPoint(ts, 1, "attr-val", "attr-val", 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupOffsetSumDataPoint(ts, 1, "attr-val", "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionCurrentOffsetDataPoint(ts, 1, "attr-val", 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionOldestOffsetDataPoint(ts, 1, "attr-val", 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionReplicasDataPoint(ts, 1, "attr-val", 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionReplicasInSyncDataPoint(ts, 1, "attr-val", 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaTopicPartitionsDataPoint(ts, 1, "attr-val")

			metrics := mb.Emit()

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 0)

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
				case "kafka.brokers":
					assert.False(t, validatedMetrics["kafka.brokers"], "Found a duplicate in the metrics slice: kafka.brokers")
					validatedMetrics["kafka.brokers"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of brokers in the cluster.", ms.At(i).Description())
					assert.Equal(t, "{brokers}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "kafka.consumer_group.lag":
					assert.False(t, validatedMetrics["kafka.consumer_group.lag"], "Found a duplicate in the metrics slice: kafka.consumer_group.lag")
					validatedMetrics["kafka.consumer_group.lag"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current approximate lag of consumer group at partition of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.consumer_group.lag_sum":
					assert.False(t, validatedMetrics["kafka.consumer_group.lag_sum"], "Found a duplicate in the metrics slice: kafka.consumer_group.lag_sum")
					validatedMetrics["kafka.consumer_group.lag_sum"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current approximate sum of consumer group lag across all partitions of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "kafka.consumer_group.members":
					assert.False(t, validatedMetrics["kafka.consumer_group.members"], "Found a duplicate in the metrics slice: kafka.consumer_group.members")
					validatedMetrics["kafka.consumer_group.members"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of members in the consumer group", ms.At(i).Description())
					assert.Equal(t, "{members}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "kafka.consumer_group.offset":
					assert.False(t, validatedMetrics["kafka.consumer_group.offset"], "Found a duplicate in the metrics slice: kafka.consumer_group.offset")
					validatedMetrics["kafka.consumer_group.offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current offset of the consumer group at partition of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.consumer_group.offset_sum":
					assert.False(t, validatedMetrics["kafka.consumer_group.offset_sum"], "Found a duplicate in the metrics slice: kafka.consumer_group.offset_sum")
					validatedMetrics["kafka.consumer_group.offset_sum"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Sum of consumer group offset across partitions of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "kafka.partition.current_offset":
					assert.False(t, validatedMetrics["kafka.partition.current_offset"], "Found a duplicate in the metrics slice: kafka.partition.current_offset")
					validatedMetrics["kafka.partition.current_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current offset of partition of topic.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.partition.oldest_offset":
					assert.False(t, validatedMetrics["kafka.partition.oldest_offset"], "Found a duplicate in the metrics slice: kafka.partition.oldest_offset")
					validatedMetrics["kafka.partition.oldest_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Oldest offset of partition of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.partition.replicas":
					assert.False(t, validatedMetrics["kafka.partition.replicas"], "Found a duplicate in the metrics slice: kafka.partition.replicas")
					validatedMetrics["kafka.partition.replicas"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of replicas for partition of topic", ms.At(i).Description())
					assert.Equal(t, "{replicas}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.partition.replicas_in_sync":
					assert.False(t, validatedMetrics["kafka.partition.replicas_in_sync"], "Found a duplicate in the metrics slice: kafka.partition.replicas_in_sync")
					validatedMetrics["kafka.partition.replicas_in_sync"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of synchronized replicas of partition", ms.At(i).Description())
					assert.Equal(t, "{replicas}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.topic.partitions":
					assert.False(t, validatedMetrics["kafka.topic.partitions"], "Found a duplicate in the metrics slice: kafka.topic.partitions")
					validatedMetrics["kafka.topic.partitions"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of partitions in topic.", ms.At(i).Description())
					assert.Equal(t, "{partitions}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
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
