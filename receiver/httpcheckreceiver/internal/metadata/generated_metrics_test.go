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
			mb.RecordHttpcheckDurationDataPoint(ts, 1, "http.url-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordHttpcheckErrorDataPoint(ts, 1, "http.url-val", "error.message-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordHttpcheckStatusDataPoint(ts, 1, "http.url-val", 16, "http.method-val", "http.status_class-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if test.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "httpcheck.duration":
					assert.False(t, validatedMetrics["httpcheck.duration"], "Found a duplicate in the metrics slice: httpcheck.duration")
					validatedMetrics["httpcheck.duration"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Measures the duration of the HTTP check.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("http.url")
					assert.True(t, ok)
					assert.EqualValues(t, "http.url-val", attrVal.Str())
				case "httpcheck.error":
					assert.False(t, validatedMetrics["httpcheck.error"], "Found a duplicate in the metrics slice: httpcheck.error")
					validatedMetrics["httpcheck.error"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Records errors occurring during HTTP check.", ms.At(i).Description())
					assert.Equal(t, "{error}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("http.url")
					assert.True(t, ok)
					assert.EqualValues(t, "http.url-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("error.message")
					assert.True(t, ok)
					assert.EqualValues(t, "error.message-val", attrVal.Str())
				case "httpcheck.status":
					assert.False(t, validatedMetrics["httpcheck.status"], "Found a duplicate in the metrics slice: httpcheck.status")
					validatedMetrics["httpcheck.status"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "1 if the check resulted in status_code matching the status_class, otherwise 0.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("http.url")
					assert.True(t, ok)
					assert.EqualValues(t, "http.url-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("http.status_code")
					assert.True(t, ok)
					assert.EqualValues(t, 16, attrVal.Int())
					attrVal, ok = dp.Attributes().Get("http.method")
					assert.True(t, ok)
					assert.EqualValues(t, "http.method-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("http.status_class")
					assert.True(t, ok)
					assert.EqualValues(t, "http.status_class-val", attrVal.Str())
				}
			}
		})
	}
}
