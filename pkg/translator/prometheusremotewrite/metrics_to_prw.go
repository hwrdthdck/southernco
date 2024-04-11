// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/multierr"

	prometheustranslator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus"
)

type Settings struct {
	Namespace           string
	ExternalLabels      map[string]string
	DisableTargetInfo   bool
	ExportCreatedMetric bool
	AddMetricSuffixes   bool
	SendMetadata        bool
}

// FromMetrics converts pmetric.Metrics to Prometheus remote write format.
func FromMetrics(md pmetric.Metrics, settings Settings) (map[string]*prompb.TimeSeries, error) {
	c := newPrometheusConverter()
	errs := c.FromMetrics(md, settings)
	tss := c.TimeSeries()
	out := make(map[string]*prompb.TimeSeries, len(tss))
	for i := range tss {
		out[strconv.Itoa(i)] = &tss[i]
	}

	return out, errs
}

// prometheusConverter converts from OTel write format to Prometheus write format.
type prometheusConverter struct {
	unique    map[uint64]*prompb.TimeSeries
	conflicts map[uint64][]*prompb.TimeSeries
}

func newPrometheusConverter() *prometheusConverter {
	return &prometheusConverter{
		unique:    map[uint64]*prompb.TimeSeries{},
		conflicts: map[uint64][]*prompb.TimeSeries{},
	}
}

// FromMetrics converts pmetric.Metrics to Prometheus remote write format.
func (c *prometheusConverter) FromMetrics(md pmetric.Metrics, settings Settings) (errs error) {
	resourceMetricsSlice := md.ResourceMetrics()
	for i := 0; i < resourceMetricsSlice.Len(); i++ {
		resourceMetrics := resourceMetricsSlice.At(i)
		resource := resourceMetrics.Resource()
		scopeMetricsSlice := resourceMetrics.ScopeMetrics()
		// keep track of the most recent timestamp in the ResourceMetrics for
		// use with the "target" info metric
		var mostRecentTimestamp pcommon.Timestamp
		for j := 0; j < scopeMetricsSlice.Len(); j++ {
			metricSlice := scopeMetricsSlice.At(j).Metrics()

			// TODO: decide if instrumentation library information should be exported as labels
			for k := 0; k < metricSlice.Len(); k++ {
				metric := metricSlice.At(k)
				mostRecentTimestamp = maxTimestamp(mostRecentTimestamp, mostRecentTimestampInMetric(metric))

				if !isValidAggregationTemporality(metric) {
					errs = multierr.Append(errs, fmt.Errorf("invalid temporality and type combination for metric %q", metric.Name()))
					continue
				}

				promName := prometheustranslator.BuildCompliantName(metric, settings.Namespace, settings.AddMetricSuffixes)

				// handle individual metrics based on type
				//exhaustive:enforce
				switch metric.Type() {
				case pmetric.MetricTypeGauge:
					dataPoints := metric.Gauge().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
						break
					}
					c.addGaugeNumberDataPoints(dataPoints, resource, settings, promName)
				case pmetric.MetricTypeSum:
					dataPoints := metric.Sum().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
						break
					}
					c.addSumNumberDataPoints(dataPoints, resource, metric, settings, promName)
				case pmetric.MetricTypeHistogram:
					dataPoints := metric.Histogram().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
						break
					}
					c.addHistogramDataPoints(dataPoints, resource, settings, promName)
				case pmetric.MetricTypeExponentialHistogram:
					dataPoints := metric.ExponentialHistogram().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
						break
					}
					errs = multierr.Append(errs, c.addExponentialHistogramDataPoints(
						dataPoints,
						resource,
						settings,
						promName,
					))
				case pmetric.MetricTypeSummary:
					dataPoints := metric.Summary().DataPoints()
					if dataPoints.Len() == 0 {
						errs = multierr.Append(errs, fmt.Errorf("empty data points. %s is dropped", metric.Name()))
						break
					}
					c.addSummaryDataPoints(dataPoints, resource, settings, promName)
				default:
					errs = multierr.Append(errs, errors.New("unsupported metric type"))
				}
			}
		}
		addResourceTargetInfo(resource, settings, mostRecentTimestamp, c)
	}

	return
}

// TimeSeries returns a slice of the prompb.TimeSeries that were converted from OTel format.
func (c *prometheusConverter) TimeSeries() []prompb.TimeSeries {
	conflicts := 0
	for _, ts := range c.conflicts {
		conflicts += len(ts)
	}
	allTS := make([]prompb.TimeSeries, 0, len(c.unique)+conflicts)
	for _, ts := range c.unique {
		allTS = append(allTS, *ts)
	}
	for _, cTS := range c.conflicts {
		for _, ts := range cTS {
			allTS = append(allTS, *ts)
		}
	}

	return allTS
}

func isSameMetric(ts *prompb.TimeSeries, lbls []prompb.Label) bool {
	if len(ts.Labels) != len(lbls) {
		return false
	}
	for i, l := range ts.Labels {
		if l.Name != ts.Labels[i].Name || l.Value != ts.Labels[i].Value {
			return false
		}
	}
	return true
}

// addExemplars adds exemplars for the dataPoint. For each exemplar, if it can find a bucket bound corresponding to its value,
// the exemplar is added to the bucket bound's time series, provided that the time series' has samples.
func (c *prometheusConverter) addExemplars(dataPoint pmetric.HistogramDataPoint, bucketBounds []bucketBoundsData) {
	if len(bucketBounds) == 0 {
		return
	}

	exemplars := getPromExemplars(dataPoint)
	if len(exemplars) == 0 {
		return
	}

	sort.Sort(byBucketBoundsData(bucketBounds))
	for _, exemplar := range exemplars {
		for _, bound := range bucketBounds {
			if len(bound.ts.Samples) > 0 && exemplar.Value <= bound.bound {
				bound.ts.Exemplars = append(bound.ts.Exemplars, exemplar)
				break
			}
		}
	}
}

// addSample finds a TimeSeries that corresponds to lbls, and adds sample to it.
// If there is no corresponding TimeSeries already, it's created.
// The corresponding TimeSeries is returned.
// If either lbls is nil/empty or sample is nil, nothing is done.
func (c *prometheusConverter) addSample(sample *prompb.Sample, lbls []prompb.Label) *prompb.TimeSeries {
	if sample == nil || len(lbls) == 0 {
		// This shouldn't happen
		return nil
	}

	ts, _ := c.getOrCreateTimeSeries(lbls)
	ts.Samples = append(ts.Samples, *sample)
	return ts
}
