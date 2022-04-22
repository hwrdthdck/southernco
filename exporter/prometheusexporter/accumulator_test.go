// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheusexporter

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

func TestInvalidDataType(t *testing.T) {
	a := newAccumulator(zap.NewNop(), 1*time.Hour).(*lastValueAccumulator)
	metric := pmetric.NewMetric()
	metric.SetDataType(-100)
	n := a.addMetric(metric, pcommon.NewInstrumentationScope(), pcommon.NewMap(), time.Now())
	require.Zero(t, n)
}

func TestAccumulateDeltaAggregation(t *testing.T) {
	tests := []struct {
		name       string
		fillMetric func(time.Time, pmetric.Metric)
	}{
		{
			name: "IntSum",
			fillMetric: func(ts time.Time, metric pmetric.Metric) {
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityDelta)
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetIntVal(42)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "Sum",
			fillMetric: func(ts time.Time, metric pmetric.Metric) {
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityDelta)
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetDoubleVal(42.42)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "Histogram",
			fillMetric: func(ts time.Time, metric pmetric.Metric) {
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeHistogram)
				metric.Histogram().SetAggregationTemporality(pmetric.MetricAggregationTemporalityDelta)
				metric.SetDescription("test description")
				dp := metric.Histogram().DataPoints().AppendEmpty()
				dp.SetBucketCounts([]uint64{5, 2})
				dp.SetCount(7)
				dp.SetExplicitBounds([]float64{3.5, 10.0})
				dp.SetSum(42.42)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resourceMetrics := pmetric.NewResourceMetrics()
			ilm := resourceMetrics.ScopeMetrics().AppendEmpty()
			ilm.Scope().SetName("test")
			tt.fillMetric(time.Now(), ilm.Metrics().AppendEmpty())

			a := newAccumulator(zap.NewNop(), 1*time.Hour).(*lastValueAccumulator)
			n := a.Accumulate(resourceMetrics)
			require.Equal(t, 0, n)

			signature := timeseriesSignature(ilm.Scope().Name(), ilm.Metrics().At(0), pcommon.NewMap())
			v, ok := a.registeredMetrics.Load(signature)
			require.False(t, ok)
			require.Nil(t, v)
		})
	}
}

func TestAccumulateMetrics(t *testing.T) {
	tests := []struct {
		name   string
		metric func(time.Time, float64, pmetric.MetricSlice)
	}{
		{
			name: "IntGauge",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeGauge)
				metric.SetDescription("test description")
				dp := metric.Gauge().DataPoints().AppendEmpty()
				dp.SetIntVal(int64(v))
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "Gauge",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeGauge)
				metric.SetDescription("test description")
				dp := metric.Gauge().DataPoints().AppendEmpty()
				dp.SetDoubleVal(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "IntSum",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.SetDescription("test description")
				metric.Sum().SetIsMonotonic(false)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetIntVal(int64(v))
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "Sum",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.SetDescription("test description")
				metric.Sum().SetIsMonotonic(false)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetDoubleVal(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "MonotonicIntSum",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.SetDescription("test description")
				metric.Sum().SetIsMonotonic(true)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetIntVal(int64(v))
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "MonotonicSum",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.Sum().SetIsMonotonic(true)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				metric.SetDescription("test description")
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetDoubleVal(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "Histogram",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeHistogram)
				metric.Histogram().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				metric.SetDescription("test description")
				dp := metric.Histogram().DataPoints().AppendEmpty()
				dp.SetBucketCounts([]uint64{5, 2})
				dp.SetCount(7)
				dp.SetExplicitBounds([]float64{3.5, 10.0})
				dp.SetSum(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
			},
		},
		{
			name: "Summary",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSummary)
				metric.SetDescription("test description")
				dp := metric.Summary().DataPoints().AppendEmpty()
				dp.SetCount(10)
				dp.SetSum(0.012)
				dp.SetCount(10)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
				fillQuantileValue := func(pN, value float64, dest pmetric.ValueAtQuantile) {
					dest.SetQuantile(pN)
					dest.SetValue(value)
				}
				fillQuantileValue(0.50, 190, dp.QuantileValues().AppendEmpty())
				fillQuantileValue(0.99, 817, dp.QuantileValues().AppendEmpty())
			},
		},
		{
			name: "StalenessMarkerGauge",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeGauge)
				metric.SetDescription("test description")
				dp := metric.Gauge().DataPoints().AppendEmpty()
				dp.SetDoubleVal(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
				dp.SetFlags(pmetric.MetricDataPointFlags(pmetric.MetricDataPointFlagNoRecordedValue))
			},
		},
		{
			name: "StalenessMarkerSum",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSum)
				metric.SetDescription("test description")
				metric.Sum().SetIsMonotonic(false)
				metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				dp := metric.Sum().DataPoints().AppendEmpty()
				dp.SetDoubleVal(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
				dp.SetFlags(pmetric.MetricDataPointFlags(pmetric.MetricDataPointFlagNoRecordedValue))
			},
		},
		{
			name: "StalenessMarkerHistogram",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeHistogram)
				metric.Histogram().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
				metric.SetDescription("test description")
				dp := metric.Histogram().DataPoints().AppendEmpty()
				dp.SetBucketCounts([]uint64{5, 2})
				dp.SetCount(7)
				dp.SetExplicitBounds([]float64{3.5, 10.0})
				dp.SetSum(v)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
				dp.SetFlags(pmetric.MetricDataPointFlags(pmetric.MetricDataPointFlagNoRecordedValue))
			},
		},
		{
			name: "StalenessMarkerSummary",
			metric: func(ts time.Time, v float64, metrics pmetric.MetricSlice) {
				metric := metrics.AppendEmpty()
				metric.SetName("test_metric")
				metric.SetDataType(pmetric.MetricDataTypeSummary)
				metric.SetDescription("test description")
				dp := metric.Summary().DataPoints().AppendEmpty()
				dp.SetCount(10)
				dp.SetSum(0.012)
				dp.SetCount(10)
				dp.Attributes().InsertString("label_1", "1")
				dp.Attributes().InsertString("label_2", "2")
				dp.SetTimestamp(pcommon.NewTimestampFromTime(ts))
				dp.SetFlags(pmetric.MetricDataPointFlags(pmetric.MetricDataPointFlagNoRecordedValue))
				fillQuantileValue := func(pN, value float64, dest pmetric.ValueAtQuantile) {
					dest.SetQuantile(pN)
					dest.SetValue(value)
				}
				fillQuantileValue(0.50, 190, dp.QuantileValues().AppendEmpty())
				fillQuantileValue(0.99, 817, dp.QuantileValues().AppendEmpty())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ts1 := time.Now().Add(-3 * time.Second)
			ts2 := time.Now().Add(-2 * time.Second)
			ts3 := time.Now().Add(-1 * time.Second)

			resourceMetrics2 := pmetric.NewResourceMetrics()
			ilm2 := resourceMetrics2.ScopeMetrics().AppendEmpty()
			ilm2.Scope().SetName("test")
			tt.metric(ts2, 21, ilm2.Metrics())
			tt.metric(ts1, 13, ilm2.Metrics())

			a := newAccumulator(zap.NewNop(), 1*time.Hour).(*lastValueAccumulator)

			// 2 metric arrived
			n := a.Accumulate(resourceMetrics2)
			if strings.HasPrefix(tt.name, "StalenessMarker") {
				require.Equal(t, 0, n)
				return
			}
			require.Equal(t, 1, n)

			m2Labels, _, m2Value, m2Temporality, m2IsMonotonic := getMetricProperties(ilm2.Metrics().At(0))

			signature := timeseriesSignature(ilm2.Scope().Name(), ilm2.Metrics().At(0), m2Labels)
			m, ok := a.registeredMetrics.Load(signature)
			require.True(t, ok)

			v := m.(*accumulatedValue)
			vLabels, vTS, vValue, vTemporality, vIsMonotonic := getMetricProperties(ilm2.Metrics().At(0))

			require.Equal(t, v.instrumentationLibrary.Name(), "test")
			require.Equal(t, v.value.DataType(), ilm2.Metrics().At(0).DataType())
			vLabels.Range(func(k string, v pcommon.Value) bool {
				r, _ := m2Labels.Get(k)
				require.Equal(t, r, v)
				return true
			})
			require.Equal(t, m2Labels.Len(), vLabels.Len())
			require.Equal(t, m2Value, vValue)
			require.Equal(t, ts2.Unix(), vTS.Unix())
			require.Greater(t, v.updated.Unix(), vTS.Unix())
			require.Equal(t, m2Temporality, vTemporality)
			require.Equal(t, m2IsMonotonic, vIsMonotonic)

			// 3 metrics arrived
			resourceMetrics3 := pmetric.NewResourceMetrics()
			ilm3 := resourceMetrics3.ScopeMetrics().AppendEmpty()
			ilm3.Scope().SetName("test")
			tt.metric(ts2, 21, ilm3.Metrics())
			tt.metric(ts3, 34, ilm3.Metrics())
			tt.metric(ts1, 13, ilm3.Metrics())

			_, _, m3Value, _, _ := getMetricProperties(ilm3.Metrics().At(1))

			n = a.Accumulate(resourceMetrics3)
			require.Equal(t, 2, n)

			m, ok = a.registeredMetrics.Load(signature)
			require.True(t, ok)
			v = m.(*accumulatedValue)
			_, vTS, vValue, _, _ = getMetricProperties(v.value)

			require.Equal(t, m3Value, vValue)
			require.Equal(t, ts3.Unix(), vTS.Unix())
		})
	}
}

func getMetricProperties(metric pmetric.Metric) (
	attributes pcommon.Map,
	ts time.Time,
	value float64,
	temporality pmetric.MetricAggregationTemporality,
	isMonotonic bool,
) {
	switch metric.DataType() {
	case pmetric.MetricDataTypeGauge:
		attributes = metric.Gauge().DataPoints().At(0).Attributes()
		ts = metric.Gauge().DataPoints().At(0).Timestamp().AsTime()
		dp := metric.Gauge().DataPoints().At(0)
		switch dp.ValueType() {
		case pmetric.NumberDataPointValueTypeInt:
			value = float64(dp.IntVal())
		case pmetric.NumberDataPointValueTypeDouble:
			value = dp.DoubleVal()
		}
		temporality = pmetric.MetricAggregationTemporalityUnspecified
		isMonotonic = false
	case pmetric.MetricDataTypeSum:
		attributes = metric.Sum().DataPoints().At(0).Attributes()
		ts = metric.Sum().DataPoints().At(0).Timestamp().AsTime()
		dp := metric.Sum().DataPoints().At(0)
		switch dp.ValueType() {
		case pmetric.NumberDataPointValueTypeInt:
			value = float64(dp.IntVal())
		case pmetric.NumberDataPointValueTypeDouble:
			value = dp.DoubleVal()
		}
		temporality = metric.Sum().AggregationTemporality()
		isMonotonic = metric.Sum().IsMonotonic()
	case pmetric.MetricDataTypeHistogram:
		attributes = metric.Histogram().DataPoints().At(0).Attributes()
		ts = metric.Histogram().DataPoints().At(0).Timestamp().AsTime()
		value = metric.Histogram().DataPoints().At(0).Sum()
		temporality = metric.Histogram().AggregationTemporality()
		isMonotonic = true
	case pmetric.MetricDataTypeSummary:
		attributes = metric.Summary().DataPoints().At(0).Attributes()
		ts = metric.Summary().DataPoints().At(0).Timestamp().AsTime()
		value = metric.Summary().DataPoints().At(0).Sum()
	default:
		log.Panicf("Invalid data type %s", metric.DataType().String())
	}

	return
}
