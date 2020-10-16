// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package splunkhecexporter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/splunk"
)

func Test_metricDataToSplunk(t *testing.T) {
	logger := zap.NewNop()

	unixSecs := int64(1574092046)
	unixNSecs := int64(11 * time.Millisecond)
	tsUnix := time.Unix(unixSecs, unixNSecs)
	tsMSecs := timestampToEpochMilliseconds(pdata.TimestampUnixNano(tsUnix.UnixNano()))

	doubleVal := 1234.5678
	int64Val := int64(123)

	distributionBounds := []float64{1, 2, 4}
	distributionCounts := []uint64{4, 2, 3}

	tests := []struct {
		name                     string
		metricsDataFn            func() pdata.Metrics
		wantSplunkMetrics        []splunk.Event
		wantNumDroppedTimeseries int
	}{
		{
			name: "gauges",
			metricsDataFn: func() pdata.Metrics {

				metrics := pdata.NewMetrics()
				rm := pdata.NewResourceMetrics()
				rm.InitEmpty()
				metrics.ResourceMetrics().Append(rm)
				rm.Resource().InitEmpty()
				rm.Resource().Attributes().InsertString("k0", "v0")
				rm.Resource().Attributes().InsertString("k1", "v1")
				ilm := pdata.NewInstrumentationLibraryMetrics()
				ilm.InitEmpty()

				doubleGauge := pdata.NewMetric()
				doubleGauge.InitEmpty()
				doubleGauge.SetName("gauge_double_with_dims")
				doubleGauge.SetDataType(pdata.MetricDataTypeDoubleGauge)
				doubleGauge.DoubleGauge().InitEmpty()
				doubleDataPt := pdata.NewDoubleDataPoint()
				doubleDataPt.InitEmpty()
				doubleDataPt.SetValue(doubleVal)
				doubleDataPt.SetTimestamp(pdata.TimestampUnixNano(tsUnix.UnixNano()))
				doubleGauge.DoubleGauge().DataPoints().Append(doubleDataPt)
				ilm.Metrics().Append(doubleGauge)

				intGauge := pdata.NewMetric()
				intGauge.InitEmpty()
				intGauge.SetName("gauge_int_with_dims")
				intGauge.SetDataType(pdata.MetricDataTypeIntGauge)
				intGauge.IntGauge().InitEmpty()
				intDataPt := pdata.NewIntDataPoint()
				intDataPt.InitEmpty()
				intDataPt.SetValue(int64Val)
				intDataPt.SetTimestamp(pdata.TimestampUnixNano(tsUnix.UnixNano()))
				intDataPt.SetTimestamp(pdata.TimestampUnixNano(tsUnix.UnixNano()))
				intGauge.IntGauge().DataPoints().Append(intDataPt)
				ilm.Metrics().Append(intGauge)

				rm.InstrumentationLibraryMetrics().Append(ilm)
				return metrics
			},
			wantSplunkMetrics: []splunk.Event{
				commonSplunkMetric("gauge_double_with_dims", tsMSecs, []string{"k0", "k1"}, []interface{}{"v0", "v1"}, doubleVal),
				commonSplunkMetric("gauge_int_with_dims", tsMSecs, []string{"k0", "k1"}, []interface{}{"v0", "v1"}, int64Val),
			},
		},
		{
			name: "double_histogram",
			metricsDataFn: func() pdata.Metrics {

				metrics := pdata.NewMetrics()
				rm := pdata.NewResourceMetrics()
				rm.InitEmpty()
				metrics.ResourceMetrics().Append(rm)
				rm.Resource().InitEmpty()
				rm.Resource().Attributes().InsertString("k0", "v0")
				rm.Resource().Attributes().InsertString("k1", "v1")
				ilm := pdata.NewInstrumentationLibraryMetrics()
				ilm.InitEmpty()

				doubleHistogram := pdata.NewMetric()
				doubleHistogram.InitEmpty()
				doubleHistogram.SetName("double_histogram_with_dims")
				doubleHistogram.SetDataType(pdata.MetricDataTypeDoubleHistogram)
				doubleHistogram.DoubleHistogram().InitEmpty()
				doubleHistogramPt := pdata.NewDoubleHistogramDataPoint()
				doubleHistogramPt.InitEmpty()
				doubleHistogramPt.SetExplicitBounds(distributionBounds)
				doubleHistogramPt.SetBucketCounts(distributionCounts)
				doubleHistogramPt.SetSum(23)
				doubleHistogramPt.SetCount(7)
				doubleHistogramPt.SetTimestamp(pdata.TimestampUnixNano(tsUnix.UnixNano()))
				doubleHistogram.DoubleHistogram().DataPoints().Append(doubleHistogramPt)
				ilm.Metrics().Append(doubleHistogram)

				rm.InstrumentationLibraryMetrics().Append(ilm)
				return metrics
			},
			wantSplunkMetrics: []splunk.Event{
				{
					Host:       "unknown",
					Source:     "",
					SourceType: "",
					Event:      "metric",
					Time:       tsMSecs,
					Fields: map[string]interface{}{
						"k0":                                     "v0",
						"k1":                                     "v1",
						"metric_name:double_histogram_with_dims": float64(23),
						"metric_name:double_histogram_with_dims_1.000000": uint64(4),
						"metric_name:double_histogram_with_dims_2.000000": uint64(2),
						"metric_name:double_histogram_with_dims_4.000000": uint64(3),
						"metric_name:double_histogram_with_dims_count":    uint64(7),
					},
				},
			},
		},
		{
			name: "int_histogram",
			metricsDataFn: func() pdata.Metrics {

				metrics := pdata.NewMetrics()
				rm := pdata.NewResourceMetrics()
				rm.InitEmpty()
				metrics.ResourceMetrics().Append(rm)
				rm.Resource().InitEmpty()
				rm.Resource().Attributes().InsertString("k0", "v0")
				rm.Resource().Attributes().InsertString("k1", "v1")
				ilm := pdata.NewInstrumentationLibraryMetrics()
				ilm.InitEmpty()

				intHistogram := pdata.NewMetric()
				intHistogram.InitEmpty()
				intHistogram.SetName("int_histogram_with_dims")
				intHistogram.SetDataType(pdata.MetricDataTypeIntHistogram)
				intHistogram.IntHistogram().InitEmpty()
				intHistogramPt := pdata.NewIntHistogramDataPoint()
				intHistogramPt.InitEmpty()
				intHistogramPt.SetExplicitBounds(distributionBounds)
				intHistogramPt.SetBucketCounts(distributionCounts)
				intHistogramPt.SetCount(7)
				intHistogramPt.SetSum(23)
				intHistogramPt.SetTimestamp(pdata.TimestampUnixNano(tsUnix.UnixNano()))
				intHistogram.IntHistogram().DataPoints().Append(intHistogramPt)
				ilm.Metrics().Append(intHistogram)

				rm.InstrumentationLibraryMetrics().Append(ilm)
				return metrics
			},
			wantSplunkMetrics: []splunk.Event{
				{
					Host:       "unknown",
					Source:     "",
					SourceType: "",
					Event:      "metric",
					Time:       tsMSecs,
					Fields: map[string]interface{}{
						"k0":                                  "v0",
						"k1":                                  "v1",
						"metric_name:int_histogram_with_dims": int64(23),
						"metric_name:int_histogram_with_dims_1.000000": uint64(4),
						"metric_name:int_histogram_with_dims_2.000000": uint64(2),
						"metric_name:int_histogram_with_dims_4.000000": uint64(3),
						"metric_name:int_histogram_with_dims_count":    uint64(7),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := tt.metricsDataFn()
			gotMetrics, gotNumDroppedTimeSeries, err := metricDataToSplunk(logger, md, &Config{})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantNumDroppedTimeseries, gotNumDroppedTimeSeries)
			for i, want := range tt.wantSplunkMetrics {
				assert.Equal(t, &want, gotMetrics[i])
			}
		})
	}
}

func commonSplunkMetric(
	metricName string,
	ts float64,
	keys []string,
	values []interface{},
	val interface{},
) splunk.Event {
	fields := map[string]interface{}{fmt.Sprintf("metric_name:%s", metricName): val}

	for i, k := range keys {
		fields[k] = values[i]
	}

	return splunk.Event{
		Time:   ts,
		Host:   "unknown",
		Event:  "metric",
		Fields: fields,
	}
}

func TestTimestampFormat(t *testing.T) {
	tsUnix := time.Unix(32, 1000345)
	tsMSecs := timestampToEpochMilliseconds(pdata.TimestampUnixNano(tsUnix.UnixNano()))
	assert.Equal(t, 32.001, tsMSecs)
}

func TestTimestampFormatRounding(t *testing.T) {
	tsUnix := time.Unix(32, 1999345)
	tsMSecs := timestampToEpochMilliseconds(pdata.TimestampUnixNano(tsUnix.UnixNano()))
	assert.Equal(t, 32.002, tsMSecs)
}
