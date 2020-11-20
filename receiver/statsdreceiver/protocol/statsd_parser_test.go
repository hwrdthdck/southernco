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

package protocol

import (
	"errors"
	"testing"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_ParseMessageToMetric(t *testing.T) {

	tests := []struct {
		name       string
		input      string
		wantMetric *statsDMetric
		err        error
	}{
		{
			name:  "empty input string",
			input: "",
			err:   errors.New("invalid message format: "),
		},
		{
			name:  "missing metric value",
			input: "test.metric|c",
			err:   errors.New("invalid <name>:<value> format: test.metric"),
		},
		{
			name:  "empty metric name",
			input: ":42|c",
			err:   errors.New("empty metric name"),
		},
		{
			name:  "empty metric value",
			input: "test.metric:|c",
			err:   errors.New("empty metric value"),
		},
		{
			name:  "invalid sample rate value",
			input: "test.metric:42|c|@1.0a",
			err:   errors.New("parse sample rate: 1.0a"),
		},
		{
			name:  "invalid tag format",
			input: "test.metric:42|c|#key1",
			err:   errors.New("invalid tag format: [key1]"),
		},
		{
			name:  "unrecognized message part",
			input: "test.metric:42|c|$extra",
			err:   errors.New("unrecognized message part: $extra"),
		},
		{
			name:  "integer counter",
			input: "test.metric:42|c",
			wantMetric: testStatsDMetric("test.metric",
				"42",
				42,
				0,
				"test.metric",
				false,
				"c",
				"", 1, 0, nil, nil),
		},
		{
			name:  "invalid  counter metric value",
			input: "test.metric:42.abc|c",
			err:   errors.New("counter: parse metric value string: 42.abc"),
		},
		{
			name:  "unhandled metric type",
			input: "test.metric:42|unhandled_type",
			err:   errors.New("unsupported metric type: unhandled_type"),
		},
		{
			name:  "counter metric with sample rate and tag",
			input: "test.metric:42|c|@0.1|#key:value",
			wantMetric: testStatsDMetric("test.metric",
				"42",
				420,
				0,
				"test.metrickey=value",
				false,
				"c",
				"",
				1,
				0.1,
				[]*metricspb.LabelKey{
					{
						Key: "key",
					},
				},
				[]*metricspb.LabelValue{
					{
						Value:    "value",
						HasValue: true,
					},
				}),
		},
		{
			name:  "counter metric with sample rate(not divisible) and tag",
			input: "test.metric:42|c|@0.8|#key:value",
			wantMetric: testStatsDMetric("test.metric",
				"42",
				52,
				0,
				"test.metrickey=value",
				false,
				"c",
				"",
				1,
				0.8,
				[]*metricspb.LabelKey{
					{
						Key: "key",
					},
				},
				[]*metricspb.LabelValue{
					{
						Value:    "value",
						HasValue: true,
					},
				}),
		},
		{
			name:  "counter metric with sample rate(not divisible) and two tags",
			input: "test.metric:42|c|@0.8|#key:value,key2:value2",
			wantMetric: testStatsDMetric("test.metric",
				"42",
				52,
				0,
				"test.metrickey=valuekey2=value2",
				false,
				"c",
				"",
				1,
				0.8,
				[]*metricspb.LabelKey{
					{
						Key: "key",
					},
					{
						Key: "key2",
					},
				},
				[]*metricspb.LabelValue{
					{
						Value:    "value",
						HasValue: true,
					},
					{
						Value:    "value2",
						HasValue: true,
					},
				}),
		},
		{
			name:  "double gauge",
			input: "test.metric:42.0|g",
			wantMetric: testStatsDMetric("test.metric",
				"42.0",
				0,
				42,
				"test.metric",
				false,
				"g",
				"", 2, 0, nil, nil),
		},
		{
			name:  "int gauge",
			input: "test.metric:42|g",
			wantMetric: testStatsDMetric("test.metric",
				"42",
				0,
				42,
				"test.metric",
				false,
				"g",
				"", 2, 0, nil, nil),
		},
		{
			name:  "invalid gauge metric value",
			input: "test.metric:42.abc|g",
			err:   errors.New("gauge: parse metric value string: 42.abc"),
		},
		{
			name:  "gauge metric with sample rate and tag",
			input: "test.metric:11|g|@0.1|#key:value",
			wantMetric: testStatsDMetric("test.metric",
				"11",
				0,
				11,
				"test.metrickey=value",
				false,
				"g",
				"",
				2,
				0.1,
				[]*metricspb.LabelKey{
					{
						Key: "key",
					},
				},
				[]*metricspb.LabelValue{
					{
						Value:    "value",
						HasValue: true,
					},
				}),
		},
		{
			name:  "gauge metric with sample rate and two tags",
			input: "test.metric:11|g|@0.8|#key:value,key2:value2",
			wantMetric: testStatsDMetric("test.metric",
				"11",
				0,
				11,
				"test.metrickey=valuekey2=value2",
				false,
				"g",
				"",
				2,
				0.8,
				[]*metricspb.LabelKey{
					{
						Key: "key",
					},
					{
						Key: "key2",
					},
				},
				[]*metricspb.LabelValue{
					{
						Value:    "value",
						HasValue: true,
					},
					{
						Value:    "value2",
						HasValue: true,
					},
				}),
		},
		{
			name:  "double gauge plus",
			input: "test.metric:+42.0|g",
			wantMetric: testStatsDMetric("test.metric",
				"+42.0",
				0,
				42,
				"test.metric",
				true,
				"g",
				"", 2, 0, nil, nil),
		},
		{
			name:  "double gauge minus",
			input: "test.metric:-42.0|g",
			wantMetric: testStatsDMetric("test.metric",
				"-42.0",
				0,
				-42,
				"test.metric",
				true,
				"g",
				"", 2, 0, nil, nil),
		},
		{
			name:  "int gauge plus",
			input: "test.metric:+42|g",
			wantMetric: testStatsDMetric("test.metric",
				"+42",
				0,
				42,
				"test.metric",
				true,
				"g",
				"", 2, 0, nil, nil),
		},
		{
			name:  "int gauge minus",
			input: "test.metric:-42|g",
			wantMetric: testStatsDMetric("test.metric",
				"-42",
				0,
				-42,
				"test.metric",
				true,
				"g",
				"", 2, 0, nil, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := parseMessageToMetric(tt.input)

			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantMetric, got)
			}
		})
	}
}

func testStatsDMetric(name string,
	value string, intValue int64,
	floatValue float64, hash string,
	addition bool, statsdMetricType string,
	unit string, metricType metricspb.MetricDescriptor_Type,
	sampleRate float64, labelKeys []*metricspb.LabelKey,
	labelValue []*metricspb.LabelValue) *statsDMetric {
	return &statsDMetric{
		name:             name,
		value:            value,
		intvalue:         intValue,
		floatvalue:       floatValue,
		hash:             hash,
		addition:         addition,
		statsdMetricType: statsdMetricType,
		unit:             unit,
		metricType:       metricType,
		sampleRate:       sampleRate,
		labelKeys:        labelKeys,
		labelValues:      labelValue,
	}
}

func NewTestInitialization() {
	gauges = make(map[string]*metricspb.Metric)
	counters = make(map[string]*metricspb.Metric)
}

func TestStatsDParser_Aggregate(t *testing.T) {
	timeNowFunc = func() int64 {
		return 0
	}

	tests := []struct {
		name             string
		input            []string
		expectedGauges   map[string]*metricspb.Metric
		expectedCounters map[string]*metricspb.Metric
		err              error
	}{
		{
			name: "parsedMetric error: empty metric value",
			input: []string{
				"test.metric:|c",
			},
			err: errors.New("empty metric value"),
		},
		{
			name: "parsedMetric error: empty metric name",
			input: []string{
				":42|c",
			},
			err: errors.New("empty metric name"),
		},
		{
			name: "gauge plus",
			input: []string{
				"statsdTestMetric1:1|g|#mykey:myvalue",
				"statsdTestMetric2:2|g|#mykey:myvalue",
				"statsdTestMetric1:+1|g|#mykey:myvalue",
				"statsdTestMetric1:+100|g|#mykey:myvalue",
				"statsdTestMetric1:+10000|g|#mykey:myvalue",
				"statsdTestMetric2:+5|g|#mykey:myvalue",
				"statsdTestMetric2:+500|g|#mykey:myvalue",
			},
			expectedGauges: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 10102,
						},
					}),
				"statsdTestMetric2mykey=myvalue": testMetric("statsdTestMetric2",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 507,
						},
					}),
			},
			expectedCounters: map[string]*metricspb.Metric{},
		},
		{
			name: "gauge minus",
			input: []string{
				"statsdTestMetric1:5000|g|#mykey:myvalue",
				"statsdTestMetric2:10|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric2:-5|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric1:-10|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric1:-100|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
			},
			expectedGauges: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 4885,
						},
					}),
				"statsdTestMetric2mykey=myvalue": testMetric("statsdTestMetric2",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 5,
						},
					}),
			},
			expectedCounters: map[string]*metricspb.Metric{},
		},
		{
			name: "gauge plus and minus",
			input: []string{
				"statsdTestMetric1:5000|g|#mykey:myvalue",
				"statsdTestMetric1:4000|g|#mykey:myvalue",
				"statsdTestMetric1:+500|g|#mykey:myvalue",
				"statsdTestMetric1:-400|g|#mykey:myvalue",
				"statsdTestMetric1:+2|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric2:365|g|#mykey:myvalue",
				"statsdTestMetric2:+300|g|#mykey:myvalue",
				"statsdTestMetric2:-200|g|#mykey:myvalue",
				"statsdTestMetric2:200|g|#mykey:myvalue",
			},
			expectedGauges: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 4101,
						},
					}),
				"statsdTestMetric2mykey=myvalue": testMetric("statsdTestMetric2",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 200,
						},
					}),
			},
			expectedCounters: map[string]*metricspb.Metric{},
		},
		{
			name: "counter with increment and sample rate",
			input: []string{
				"statsdTestMetric1:3000|c|#mykey:myvalue",
				"statsdTestMetric1:4000|c|#mykey:myvalue",
				"statsdTestMetric2:20|c|@0.8|#mykey:myvalue",
				"statsdTestMetric2:20|c|@0.8|#mykey:myvalue",
			},
			expectedGauges: map[string]*metricspb.Metric{},
			expectedCounters: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_INT64,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_Int64Value{
							Int64Value: 7000,
						},
					}),
				"statsdTestMetric2mykey=myvalue": testMetric("statsdTestMetric2",
					metricspb.MetricDescriptor_GAUGE_INT64,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_Int64Value{
							Int64Value: 50,
						},
					}),
			},
		},
		{
			name: "counter and gauge: one gauge and two counters",
			input: []string{
				"statsdTestMetric1:3000|c|#mykey:myvalue",
				"statsdTestMetric1:500|g|#mykey:myvalue",
				"statsdTestMetric1:400|g|#mykey:myvalue",
				"statsdTestMetric1:+20|g|#mykey:myvalue",
				"statsdTestMetric1:4000|c|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric2:20|c|@0.8|#mykey:myvalue",
				"statsdTestMetric1:+2|g|#mykey:myvalue",
				"statsdTestMetric2:20|c|@0.8|#mykey:myvalue",
			},
			expectedGauges: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 421,
						},
					}),
			},
			expectedCounters: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_INT64,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_Int64Value{
							Int64Value: 7000,
						},
					}),
				"statsdTestMetric2mykey=myvalue": testMetric("statsdTestMetric2",
					metricspb.MetricDescriptor_GAUGE_INT64,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_Int64Value{
							Int64Value: 50,
						},
					}),
			},
		},
		{
			name: "counter and gauge: 2 gauges and 2 counters",
			input: []string{
				"statsdTestMetric1:500|g|#mykey:myvalue",
				"statsdTestMetric1:400|g|#mykey:myvalue1",
				"statsdTestMetric1:300|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue1",
				"statsdTestMetric1:+20|g|#mykey:myvalue",
				"statsdTestMetric1:-1|g|#mykey:myvalue",
				"statsdTestMetric1:20|c|@0.1|#mykey:myvalue",
				"statsdTestMetric2:50|c|#mykey:myvalue",
				"statsdTestMetric1:15|c|#mykey:myvalue",
				"statsdTestMetric2:5|c|@0.2|#mykey:myvalue",
			},
			expectedGauges: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 319,
						},
					}),
				"statsdTestMetric1mykey=myvalue1": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_DOUBLE,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue1",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_DoubleValue{
							DoubleValue: 399,
						},
					}),
			},
			expectedCounters: map[string]*metricspb.Metric{
				"statsdTestMetric1mykey=myvalue": testMetric("statsdTestMetric1",
					metricspb.MetricDescriptor_GAUGE_INT64,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_Int64Value{
							Int64Value: 215,
						},
					}),
				"statsdTestMetric2mykey=myvalue": testMetric("statsdTestMetric2",
					metricspb.MetricDescriptor_GAUGE_INT64,
					[]*metricspb.LabelKey{
						{
							Key: "mykey",
						},
					},
					[]*metricspb.LabelValue{
						{
							Value:    "myvalue",
							HasValue: true,
						},
					},
					"",
					&metricspb.Point{
						Timestamp: &timestamppb.Timestamp{
							Seconds: 0,
						},
						Value: &metricspb.Point_Int64Value{
							Int64Value: 75,
						},
					}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewTestInitialization()
			var err error
			for _, line := range tt.input {
				p := &StatsDParser{}
				err = p.Aggregate(line)
			}
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Equal(t, tt.expectedGauges, gauges)
				assert.Equal(t, tt.expectedCounters, counters)
			}
		})
	}
}

func testMetric(metricName string,
	metricType metricspb.MetricDescriptor_Type,
	lableKeys []*metricspb.LabelKey,
	labelValues []*metricspb.LabelValue,
	unit string,
	point *metricspb.Point) *metricspb.Metric {
	return &metricspb.Metric{
		MetricDescriptor: &metricspb.MetricDescriptor{
			Name:      metricName,
			Type:      metricType,
			LabelKeys: lableKeys,
			Unit:      unit,
		},
		Timeseries: []*metricspb.TimeSeries{
			{
				LabelValues: labelValues,
				Points: []*metricspb.Point{
					point,
				},
			},
		},
	}
}

func Test_contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		expected bool
	}{
		{
			name: "contain 1",
			slice: []string{
				"m",
				"g",
			},
			element:  "m",
			expected: true,
		},
		{
			name: "contain 2",
			slice: []string{
				"m",
				"g",
			},
			element:  "g",
			expected: true,
		},
		{
			name: "does not contain",
			slice: []string{
				"m",
				"g",
			},
			element:  "t",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := contains(tt.slice, tt.element)
			assert.Equal(t, tt.expected, answer)
		})
	}
}

func TestStatsDParser_Initialize(t *testing.T) {
	p := &StatsDParser{}
	p.Initialize()
	gauges["test"] = &metricspb.Metric{}
	counters["test"] = &metricspb.Metric{}
	assert.Equal(t, 1, len(gauges))
	assert.Equal(t, 1, len(counters))
}

func TestStatsDParser_GetMetrics(t *testing.T) {
	p := &StatsDParser{}
	p.Initialize()
	gauges["testGauge1"] = testMetric("testGauge1",
		metricspb.MetricDescriptor_GAUGE_DOUBLE,
		nil,
		nil,
		"",
		&metricspb.Point{
			Timestamp: &timestamppb.Timestamp{
				Seconds: 0,
			},
			Value: &metricspb.Point_DoubleValue{
				DoubleValue: 1,
			},
		})
	gauges["testGauge2"] = testMetric("testGauge2",
		metricspb.MetricDescriptor_GAUGE_DOUBLE,
		nil,
		nil,
		"",
		&metricspb.Point{
			Timestamp: &timestamppb.Timestamp{
				Seconds: 0,
			},
			Value: &metricspb.Point_DoubleValue{
				DoubleValue: 2,
			},
		})
	counters["testCounter1"] = testMetric("testCounter1",
		metricspb.MetricDescriptor_GAUGE_INT64,
		nil,
		nil,
		"",
		&metricspb.Point{
			Timestamp: &timestamppb.Timestamp{
				Seconds: 0,
			},
			Value: &metricspb.Point_Int64Value{
				Int64Value: 1,
			},
		})
	metrics := p.GetMetrics()
	assert.Equal(t, 3, len(metrics))
}
