// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusreceiver

import (
	"math"
	"testing"

	promcfg "github.com/prometheus/prometheus/config"
	dto "github.com/prometheus/prometheus/prompb/io/prometheus/client"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestScrapeViaProtobuf(t *testing.T) {
	mf := &dto.MetricFamily{
		Name: "test_counter",
		Type: dto.MetricType_COUNTER,
		Metric: []dto.Metric{
			{
				Label: []dto.LabelPair{
					{
						Name:  "foo",
						Value: "bar",
					},
				},
				Counter: &dto.Counter{
					Value: 1234,
				},
			},
		},
	}
	buffer := prometheusMetricFamilyToProtoBuf(t, nil, mf)

	mf = &dto.MetricFamily{
		Name: "test_gauge",
		Type: dto.MetricType_GAUGE,
		Metric: []dto.Metric{
			{
				Gauge: &dto.Gauge{
					Value: 400.8,
				},
			},
		},
	}
	prometheusMetricFamilyToProtoBuf(t, buffer, mf)

	mf = &dto.MetricFamily{
		Name: "test_summary",
		Type: dto.MetricType_SUMMARY,
		Metric: []dto.Metric{
			{
				Summary: &dto.Summary{
					SampleCount: 1213,
					SampleSum:   456,
					Quantile: []dto.Quantile{
						{
							Quantile: 0.5,
							Value:    789,
						},
						{
							Quantile: 0.9,
							Value:    1011,
						},
					},
				},
			},
		},
	}
	prometheusMetricFamilyToProtoBuf(t, buffer, mf)

	mf = &dto.MetricFamily{
		Name: "test_histogram",
		Type: dto.MetricType_HISTOGRAM,
		Metric: []dto.Metric{
			{
				Histogram: &dto.Histogram{
					SampleCount: 1213,
					SampleSum:   456,
					Bucket: []dto.Bucket{
						{
							UpperBound:      0.5,
							CumulativeCount: 789,
						},
						{
							UpperBound:      10,
							CumulativeCount: 1011,
						},
						{
							UpperBound:      math.Inf(1),
							CumulativeCount: 1213,
						},
					},
				},
			},
		},
	}
	prometheusMetricFamilyToProtoBuf(t, buffer, mf)

	expectations := []testExpectation{
		assertMetricPresent(
			"test_counter",
			compareMetricType(pmetric.MetricTypeSum),
			compareMetricUnit(""),
			[]dataPointExpectation{{
				numberPointComparator: []numberPointComparator{
					compareDoubleValue(1234),
				},
			}},
		),
		assertMetricPresent(
			"test_gauge",
			compareMetricType(pmetric.MetricTypeGauge),
			compareMetricUnit(""),
			[]dataPointExpectation{{
				numberPointComparator: []numberPointComparator{
					compareDoubleValue(400.8),
				},
			}},
		),
		assertMetricPresent(
			"test_summary",
			compareMetricType(pmetric.MetricTypeSummary),
			compareMetricUnit(""),
			[]dataPointExpectation{{
				summaryPointComparator: []summaryPointComparator{
					compareSummary(1213, 456, [][]float64{{0.5, 789}, {0.9, 1011}}),
				},
			}},
		),
		assertMetricPresent(
			"test_histogram",
			compareMetricType(pmetric.MetricTypeHistogram),
			compareMetricUnit(""),
			[]dataPointExpectation{{
				histogramPointComparator: []histogramPointComparator{
					compareHistogram(1213, 456, []float64{0.5, 10}, []uint64{789, 222, 202}),
				},
			}},
		),
	}

	targets := []*testData{
		{
			name: "target1",
			pages: []mockPrometheusResponse{
				{code: 200, useProtoBuf: true, buf: buffer.Bytes()},
			},
			validateFunc: func(t *testing.T, td *testData, result []pmetric.ResourceMetrics) {
				verifyNumValidScrapeResults(t, td, result)
				doCompare(t, "target1", td.attributes, result[0], expectations)
			},
		},
	}

	testComponent(t, targets, func(c *Config) {
		c.EnableProtobufNegotiation = true
	})
}

func TestNativeVsClassicHistogramScrapeViaProtobuf(t *testing.T) {
	classicHistogram := &dto.MetricFamily{
		Name: "test_classic_histogram",
		Type: dto.MetricType_HISTOGRAM,
		Metric: []dto.Metric{
			{
				Histogram: &dto.Histogram{
					SampleCount: 1213,
					SampleSum:   456,
					Bucket: []dto.Bucket{
						{
							UpperBound:      0.5,
							CumulativeCount: 789,
						},
						{
							UpperBound:      10,
							CumulativeCount: 1011,
						},
						{
							UpperBound:      math.Inf(1),
							CumulativeCount: 1213,
						},
					},
				},
			},
		},
	}
	buffer := prometheusMetricFamilyToProtoBuf(t, nil, classicHistogram)

	mixedHistogram := &dto.MetricFamily{
		Name: "test_mixed_histogram",
		Type: dto.MetricType_HISTOGRAM,
		Metric: []dto.Metric{
			{
				Histogram: &dto.Histogram{
					SampleCount: 1213,
					SampleSum:   456,
					Bucket: []dto.Bucket{
						{
							UpperBound:      0.5,
							CumulativeCount: 789,
						},
						{
							UpperBound:      10,
							CumulativeCount: 1011,
						},
						{
							UpperBound:      math.Inf(1),
							CumulativeCount: 1213,
						},
					},
					// Integer counter histogram definition
					Schema:        3,
					ZeroThreshold: 0.001,
					ZeroCount:     2,
					NegativeSpan: []dto.BucketSpan{
						{Offset: 0, Length: 1},
						{Offset: 1, Length: 1},
					},
					NegativeDelta: []int64{1, 1},
					PositiveSpan: []dto.BucketSpan{
						{Offset: -2, Length: 1},
						{Offset: 1, Length: 1},
					},
					PositiveDelta: []int64{1, 0},
				},
			},
		},
	}
	prometheusMetricFamilyToProtoBuf(t, buffer, mixedHistogram)

	nativeHistogram := &dto.MetricFamily{
		Name: "test_native_histogram",
		Type: dto.MetricType_HISTOGRAM,
		Metric: []dto.Metric{
			{
				Histogram: &dto.Histogram{
					SampleCount: 1213,
					SampleSum:   456,
					// Integer counter histogram definition
					Schema:        3,
					ZeroThreshold: 0.001,
					ZeroCount:     2,
					NegativeSpan: []dto.BucketSpan{
						{Offset: 0, Length: 1},
						{Offset: 1, Length: 1},
					},
					NegativeDelta: []int64{1, 1},
					PositiveSpan: []dto.BucketSpan{
						{Offset: -2, Length: 1},
						{Offset: 2, Length: 1},
					},
					PositiveDelta: []int64{1, 0},
				},
			},
		},
	}
	prometheusMetricFamilyToProtoBuf(t, buffer, nativeHistogram)

	testCases := map[string]struct {
		mutCfg   func(*promcfg.Config)
		expected []testExpectation
	}{
		"native only": {
			expected: []testExpectation{
				assertMetricPresent( // Scrape classic only histograms as is.
					"test_classic_histogram",
					compareMetricType(pmetric.MetricTypeHistogram),
					compareMetricUnit(""),
					[]dataPointExpectation{{
						histogramPointComparator: []histogramPointComparator{
							compareHistogram(1213, 456, []float64{0.5, 10}, []uint64{789, 222, 202}),
						},
					}},
				),
				assertMetricPresent( // Only scrape native buckets from mixed histograms.
					"test_mixed_histogram",
					compareMetricType(pmetric.MetricTypeExponentialHistogram),
					compareMetricUnit(""),
					[]dataPointExpectation{{
						exponentialHistogramComparator: []exponentialHistogramComparator{
							compareExponentialHistogram(1213, 456, 2, 0, []uint64{1, 0, 2}, -2, []uint64{1, 0, 0, 1}),
						},
					}},
				),
				assertMetricPresent( // Scrape native only histograms as is.
					"test_native_histogram",
					compareMetricType(pmetric.MetricTypeExponentialHistogram),
					compareMetricUnit(""),
					[]dataPointExpectation{{
						exponentialHistogramComparator: []exponentialHistogramComparator{
							compareExponentialHistogram(1213, 456, 2, 0, []uint64{1, 0, 2}, -2, []uint64{1, 0, 0, 1}),
						},
					}},
				),
			},
		},
		"classic only": {
			mutCfg: func(cfg *promcfg.Config) {
				for _, scrapeConfig := range cfg.ScrapeConfigs {
					scrapeConfig.ScrapeClassicHistograms = true
				}
			},
			expected: []testExpectation{
				assertMetricPresent( // Scrape classic only histograms as is.
					"test_classic_histogram",
					compareMetricType(pmetric.MetricTypeHistogram),
					compareMetricUnit(""),
					[]dataPointExpectation{{
						histogramPointComparator: []histogramPointComparator{
							compareHistogram(1213, 456, []float64{0.5, 10}, []uint64{789, 222, 202}),
						},
					}},
				),
				assertMetricPresent( // Only scrape classic buckets from mixed histograms.
					"test_mixed_histogram",
					compareMetricType(pmetric.MetricTypeHistogram),
					compareMetricUnit(""),
					[]dataPointExpectation{{
						histogramPointComparator: []histogramPointComparator{
							compareHistogram(1213, 456, []float64{0.5, 10}, []uint64{789, 222, 202}),
						},
					}},
				),
				assertMetricPresent( // Scrape native only histograms as is.
					"test_native_histogram",
					compareMetricType(pmetric.MetricTypeExponentialHistogram),
					compareMetricUnit(""),
					[]dataPointExpectation{{
						exponentialHistogramComparator: []exponentialHistogramComparator{
							compareExponentialHistogram(1213, 456, 2, 0, []uint64{1, 0, 2}, -2, []uint64{1, 0, 0, 1}),
						},
					}},
				),
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			targets := []*testData{
				{
					name: "target1",
					pages: []mockPrometheusResponse{
						{code: 200, useProtoBuf: true, buf: buffer.Bytes()},
					},
					validateFunc: func(t *testing.T, td *testData, result []pmetric.ResourceMetrics) {
						verifyNumValidScrapeResults(t, td, result)
						doCompare(t, "target1", td.attributes, result[0], tc.expected)
					},
				},
			}
			mutCfg := tc.mutCfg
			if mutCfg == nil {
				mutCfg = func(*promcfg.Config) {}
			}
			testComponent(t, targets, func(c *Config) {
				c.EnableProtobufNegotiation = true
			}, mutCfg)
		})
	}
}
