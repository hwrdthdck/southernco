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

package internaldata

import (
	"testing"

	occommon "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	ocmetrics "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	ocresource "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/consumer/consumerdata"
	"go.opentelemetry.io/collector/internal/data"
	"go.opentelemetry.io/collector/internal/data/testdata"
)

func TestOCToMetrics(t *testing.T) {
	// From OC we never generate Int Histograms, will generate Double Histogram always.
	allTypesNoDataPoints := testdata.GenerateMetricsAllTypesNoDataPoints()
	dh := allTypesNoDataPoints.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(4)
	ih := allTypesNoDataPoints.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(5)
	ih.SetDoubleHistogramData(dh.DoubleHistogramData())

	sampleMetricData := testdata.GenerateMetricsWithCountersHistograms()
	dh = sampleMetricData.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(2)
	ih = sampleMetricData.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(3)
	ih.SetDoubleHistogramData(dh.DoubleHistogramData())

	tests := []struct {
		name     string
		oc       consumerdata.MetricsData
		internal data.MetricData
	}{
		{
			name:     "empty",
			oc:       consumerdata.MetricsData{},
			internal: testdata.GenerateMetricsEmpty(),
		},

		{
			name: "one-empty-resource-metrics",
			oc: consumerdata.MetricsData{
				Node:     &occommon.Node{},
				Resource: &ocresource.Resource{},
			},
			internal: wrapMetricsWithEmptyResource(testdata.GenerateMetricsOneEmptyResourceMetrics()),
		},

		{
			name:     "no-libraries",
			oc:       generateOCTestDataNoMetrics(),
			internal: testdata.GenerateMetricsNoLibraries(),
		},

		{
			name:     "all-types-no-data-points",
			oc:       generateOCTestDataNoPoints(),
			internal: allTypesNoDataPoints,
		},

		{
			name:     "one-metric-no-labels",
			oc:       generateOCTestDataNoLabels(),
			internal: testdata.GenerateMetricsOneMetricNoLabels(),
		},

		{
			name:     "one-metric",
			oc:       generateOCTestDataMetricsOneMetric(),
			internal: testdata.GenerateMetricsOneMetric(),
		},

		{
			name: "one-metric-one-summary",
			oc: consumerdata.MetricsData{
				Resource: generateOCTestResource(),
				Metrics: []*ocmetrics.Metric{
					generateOCTestMetricInt(),
					generateOCTestMetricSummary(),
				},
			},
			internal: testdata.GenerateMetricsOneMetricOneNil(),
		},

		{
			name:     "one-metric-one-nil",
			oc:       generateOCTestDataMetricsOneMetricOneNil(),
			internal: testdata.GenerateMetricsOneMetric(),
		},

		{
			name:     "one-metric-one-nil-timeseries",
			oc:       generateOCTestDataMetricsOneMetricOneNilTimeseries(),
			internal: testdata.GenerateMetricsOneMetric(),
		},

		{
			name:     "one-metric-one-nil-point",
			oc:       generateOCTestDataMetricsOneMetricOneNilPoint(),
			internal: testdata.GenerateMetricsOneMetric(),
		},

		{
			name:     "one-metric-one-nil-point",
			oc:       generateOCTestDataMetricsOneMetricOneNilPoint(),
			internal: testdata.GenerateMetricsOneMetric(),
		},

		{
			name: "sample-metric",
			oc: consumerdata.MetricsData{
				Resource: generateOCTestResource(),
				Metrics: []*ocmetrics.Metric{
					generateOCTestMetricInt(),
					generateOCTestMetricDouble(),
					generateOCTestMetricDoubleHistogram(),
					generateOCTestMetricIntHistogram(),
				},
			},
			internal: sampleMetricData,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := OCToMetrics(test.oc)
			assert.EqualValues(t, test.internal, got)

			ocslice := []consumerdata.MetricsData{
				test.oc,
				test.oc,
			}
			wantSlice := data.NewMetricData()
			// Double the ResourceMetrics only if not empty.
			if test.internal.ResourceMetrics().Len() != 0 {
				test.internal.Clone().ResourceMetrics().MoveAndAppendTo(wantSlice.ResourceMetrics())
				test.internal.Clone().ResourceMetrics().MoveAndAppendTo(wantSlice.ResourceMetrics())
			}
			gotSlice := OCSliceToMetrics(ocslice)
			assert.EqualValues(t, wantSlice, gotSlice)
		})
	}
}

// TODO: Try to avoid unnecessary Resource object allocation.
func wrapMetricsWithEmptyResource(md data.MetricData) data.MetricData {
	md.ResourceMetrics().At(0).Resource().InitEmpty()
	return md
}

func TestOCToMetrics_ResourceInMetric(t *testing.T) {
	internal := testdata.GenerateMetricsOneMetric()
	want := data.NewMetricData()
	internal.Clone().ResourceMetrics().MoveAndAppendTo(want.ResourceMetrics())
	internal.Clone().ResourceMetrics().MoveAndAppendTo(want.ResourceMetrics())
	want.ResourceMetrics().At(1).Resource().Attributes().UpsertString("resource-attr", "another-value")
	oc := generateOCTestDataMetricsOneMetric()
	oc2 := generateOCTestDataMetricsOneMetric()
	oc.Metrics = append(oc.Metrics, oc2.Metrics...)
	oc.Metrics[1].Resource = oc2.Resource
	oc.Metrics[1].Resource.Labels["resource-attr"] = "another-value"
	got := OCToMetrics(oc)
	assert.EqualValues(t, want, got)
}

func BenchmarkMetricIntOCToMetrics(b *testing.B) {
	ocMetric := consumerdata.MetricsData{
		Resource: generateOCTestResource(),
		Metrics: []*ocmetrics.Metric{
			generateOCTestMetricInt(),
			generateOCTestMetricInt(),
			generateOCTestMetricInt(),
		},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		OCToMetricData(ocMetric)
	}
}

func BenchmarkMetricDoubleOCToMetrics(b *testing.B) {
	ocMetric := consumerdata.MetricsData{
		Resource: generateOCTestResource(),
		Metrics: []*ocmetrics.Metric{
			generateOCTestMetricDouble(),
			generateOCTestMetricDouble(),
			generateOCTestMetricDouble(),
		},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		OCToMetricData(ocMetric)
	}
}

func BenchmarkMetricHistogramOCToMetrics(b *testing.B) {
	ocMetric := consumerdata.MetricsData{
		Resource: generateOCTestResource(),
		Metrics: []*ocmetrics.Metric{
			generateOCTestMetricDoubleHistogram(),
			generateOCTestMetricDoubleHistogram(),
			generateOCTestMetricDoubleHistogram(),
		},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		OCToMetricData(ocMetric)
	}
}

func generateOCTestResource() *ocresource.Resource {
	return &ocresource.Resource{
		Labels: map[string]string{
			"resource-attr": "resource-attr-val-1",
		},
	}
}
