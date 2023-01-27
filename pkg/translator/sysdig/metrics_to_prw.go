// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sysdig // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"errors"
	"fmt"

	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/multierr"
)

type Settings struct {
	Namespace         string
	ExternalLabels    map[string]string
	DisableTargetInfo bool
}

// FromMetrics converts pmetric.Metrics to prometheus remote write format.
func FromMetrics(md pmetric.Metrics, settings Settings) (tsMap map[string]*prompb.TimeSeries, mdMap map[string]*prompb.MetricMetadata, errs error) {
	tsMap = make(map[string]*prompb.TimeSeries)
	mdMap = make(map[string]*prompb.MetricMetadata)

	resourceMetricsSlice := md.ResourceMetrics()
	for i := 0; i < resourceMetricsSlice.Len(); i++ {
		resourceMetrics := resourceMetricsSlice.At(i)
		resource := resourceMetrics.Resource()
		scopeMetricsSlice := resourceMetrics.ScopeMetrics()
		// keep track of the most recent timestamp in the ResourceMetrics for
		// use with the "target" info metric
		var mostRecentTimestamp pcommon.Timestamp
		for j := 0; j < scopeMetricsSlice.Len(); j++ {
			scopeMetrics := scopeMetricsSlice.At(j)
			metricSlice := scopeMetrics.Metrics()

			// TODO: decide if instrumentation library information should be exported as labels
			for k := 0; k < metricSlice.Len(); k++ {
				metric := metricSlice.At(k)
				mostRecentTimestamp = maxTimestamp(mostRecentTimestamp, mostRecentTimestampInMetric(metric))

				if !isValidAggregationTemporality(metric) {
					errs = multierr.Append(errs, errors.New("invalid temporality and type combination"))
					continue
				}

				// handle individual metric based on type
				switch metric.Type() {
				case pmetric.MetricTypeGauge:
					dataPoints := metric.Gauge().DataPoints()
					if err := addNumberDataPointSlice(dataPoints, resource, metric, settings, tsMap, mdMap); err != nil {
						errs = multierr.Append(errs, err)
					}
				case pmetric.MetricTypeSum:
					if metric.Sum().AggregationTemporality() == pmetric.AggregationTemporalityDelta {
						metric.SetUnit("delta_counter:" + metric.Unit())
						if err := addNumberDataPointSlice(metric.Gauge().DataPoints(), resource, metric, settings, tsMap, mdMap); err != nil {
							errs = multierr.Append(errs, err)
						}
					} else {
						if err := addNumberDataPointSlice(metric.Sum().DataPoints(), resource, metric, settings, tsMap, mdMap); err != nil {
							errs = multierr.Append(errs, err)
						}
					}
				case pmetric.MetricTypeHistogram:
					dataPoints := metric.Histogram().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
					}
					for x := 0; x < dataPoints.Len(); x++ {
						addSingleHistogramDataPoint(dataPoints.At(x), resource, metric, settings, tsMap, mdMap)
					}
				case pmetric.MetricTypeSummary:
					dataPoints := metric.Summary().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
					}
					for x := 0; x < dataPoints.Len(); x++ {
						addSingleSummaryDataPoint(dataPoints.At(x), resource, metric, settings, tsMap, mdMap)
					}
				default:
					errs = multierr.Append(errs, errors.New("unsupported metric type"))
				}
			}
		}
		addResourceTargetInfo(resource, settings, mostRecentTimestamp, tsMap)
	}

	return
}

func addNumberDataPointSlice(
	dataPoints pmetric.NumberDataPointSlice,
	resource pcommon.Resource,
	metric pmetric.Metric,
	settings Settings,
	tsMap map[string]*prompb.TimeSeries,
	mdMap map[string]*prompb.MetricMetadata) error {
	if dataPoints.Len() == 0 {
		return fmt.Errorf("empty data points. %s is dropped", metric.Name())
	}
	for x := 0; x < dataPoints.Len(); x++ {
		addSingleNumberDataPoint(dataPoints.At(x), resource, metric, settings, tsMap, mdMap)
	}
	return nil
}
