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
			if test.metricsSet == testMetricsSetDefault || test.metricsSet == testMetricsSetAll {
				assert.Equal(t, "[WARNING] `kafka.brokers` should not be enabled: The metric is deprecated and will be removed. Use `kafka.brokers.count`", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersConsumerFetchRateDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersIncomingByteRateDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersOutgoingByteRateDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersRequestLatencyDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersRequestRateDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersRequestSizeDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersRequestsInFlightDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersResponseRateDataPoint(ts, 1, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaBrokersResponseSizeDataPoint(ts, 1, 1)

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
					assert.Equal(t, "[DEPRECATED] Number of brokers in the cluster.", ms.At(i).Description())
					assert.Equal(t, "{brokers}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "kafka.brokers.consumer_fetch_rate":
					assert.False(t, validatedMetrics["kafka.brokers.consumer_fetch_rate"], "Found a duplicate in the metrics slice: kafka.brokers.consumer_fetch_rate")
					validatedMetrics["kafka.brokers.consumer_fetch_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average consumer fetch Rate", ms.At(i).Description())
					assert.Equal(t, "{fetches}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.count":
					assert.False(t, validatedMetrics["kafka.brokers.count"], "Found a duplicate in the metrics slice: kafka.brokers.count")
					validatedMetrics["kafka.brokers.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of brokers in the cluster.", ms.At(i).Description())
					assert.Equal(t, "{brokers}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "kafka.brokers.incoming_byte_rate":
					assert.False(t, validatedMetrics["kafka.brokers.incoming_byte_rate"], "Found a duplicate in the metrics slice: kafka.brokers.incoming_byte_rate")
					validatedMetrics["kafka.brokers.incoming_byte_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average tncoming Byte Rate in bytes/second", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.outgoing_byte_rate":
					assert.False(t, validatedMetrics["kafka.brokers.outgoing_byte_rate"], "Found a duplicate in the metrics slice: kafka.brokers.outgoing_byte_rate")
					validatedMetrics["kafka.brokers.outgoing_byte_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average outgoing Byte Rate in bytes/second.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.request_latency":
					assert.False(t, validatedMetrics["kafka.brokers.request_latency"], "Found a duplicate in the metrics slice: kafka.brokers.request_latency")
					validatedMetrics["kafka.brokers.request_latency"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Request latency Average in ms", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.request_rate":
					assert.False(t, validatedMetrics["kafka.brokers.request_rate"], "Found a duplicate in the metrics slice: kafka.brokers.request_rate")
					validatedMetrics["kafka.brokers.request_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average request rate per second.", ms.At(i).Description())
					assert.Equal(t, "{requests}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.request_size":
					assert.False(t, validatedMetrics["kafka.brokers.request_size"], "Found a duplicate in the metrics slice: kafka.brokers.request_size")
					validatedMetrics["kafka.brokers.request_size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average request size in bytes", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.requests_in_flight":
					assert.False(t, validatedMetrics["kafka.brokers.requests_in_flight"], "Found a duplicate in the metrics slice: kafka.brokers.requests_in_flight")
					validatedMetrics["kafka.brokers.requests_in_flight"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Requests in flight", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.response_rate":
					assert.False(t, validatedMetrics["kafka.brokers.response_rate"], "Found a duplicate in the metrics slice: kafka.brokers.response_rate")
					validatedMetrics["kafka.brokers.response_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average response rate per second", ms.At(i).Description())
					assert.Equal(t, "{response}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
				case "kafka.brokers.response_size":
					assert.False(t, validatedMetrics["kafka.brokers.response_size"], "Found a duplicate in the metrics slice: kafka.brokers.response_size")
					validatedMetrics["kafka.brokers.response_size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average response size in bytes", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 1, attrVal.Int())
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
