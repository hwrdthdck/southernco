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

package goldendataset

import (
	"fmt"

	"go.opentelemetry.io/collector/consumer/pdata"
)

// Simple utilities for generating metrics for testing

// MetricCfg holds parameters for generating dummy metrics for testing. Set values on this struct to generate
// metrics with the corresponding number/type of attributes and pass into MetricDataFromCfg to generate metrics.
type MetricCfg struct {
	// The type of metric to generate
	MetricDescriptorType pdata.MetricDataType
	// If MetricDescriptorType is one of the Sum, this describes if the sum is monotonic or not.
	IsMonotonicSum bool
	// A prefix for every metric name
	MetricNamePrefix string
	// The number of instrumentation library metrics per resource
	NumILMPerResource int
	// The size of the MetricSlice and number of Metrics
	NumMetricsPerILM int
	// The number of labels on the LabelsMap associated with each point
	NumPtLabels int
	// The number of points to generate per Metric
	NumPtsPerMetric int
	// The number of Attributes to insert into each Resource's AttributesMap
	NumResourceAttrs int
	// The number of ResourceMetrics for the single MetricData generated
	NumResourceMetrics int
	// The base value for each point
	PtVal int
	// The start time for each point
	StartTime uint64
	// The duration of the steps between each generated point starting at StartTime
	StepSize uint64
}

// DefaultCfg produces a MetricCfg with default values. These should be good enough to produce sane
// (but boring) metrics, and can be used as a starting point for making alterations.
func DefaultCfg() MetricCfg {
	return MetricCfg{
		MetricDescriptorType: pdata.MetricDataTypeIntGauge,
		MetricNamePrefix:     "",
		NumILMPerResource:    1,
		NumMetricsPerILM:     1,
		NumPtLabels:          1,
		NumPtsPerMetric:      1,
		NumResourceAttrs:     1,
		NumResourceMetrics:   1,
		PtVal:                1,
		StartTime:            940000000000000000,
		StepSize:             42,
	}
}

// DefaultMetricData produces MetricData with a default config.
func DefaultMetricData() pdata.Metrics {
	return MetricDataFromCfg(DefaultCfg())
}

// MetricDataFromCfg produces MetricData with the passed-in config.
func MetricDataFromCfg(cfg MetricCfg) pdata.Metrics {
	return newMetricGenerator().genMetricDataFromCfg(cfg)
}

type metricGenerator struct {
	metricID int
}

func newMetricGenerator() *metricGenerator {
	return &metricGenerator{}
}

func (g *metricGenerator) genMetricDataFromCfg(cfg MetricCfg) pdata.Metrics {
	md := pdata.NewMetrics()
	rms := md.ResourceMetrics()
	rms.Resize(cfg.NumResourceMetrics)
	for i := 0; i < cfg.NumResourceMetrics; i++ {
		rm := rms.At(i)
		resource := rm.Resource()
		for j := 0; j < cfg.NumResourceAttrs; j++ {
			resource.Attributes().Insert(
				fmt.Sprintf("resource-attr-name-%d", j),
				pdata.NewAttributeValueString(fmt.Sprintf("resource-attr-val-%d", j)),
			)
		}
		g.populateIlm(cfg, rm)
	}
	return md
}

func (g *metricGenerator) populateIlm(cfg MetricCfg, rm pdata.ResourceMetrics) {
	ilms := rm.InstrumentationLibraryMetrics()
	ilms.Resize(cfg.NumILMPerResource)
	for i := 0; i < cfg.NumILMPerResource; i++ {
		ilm := ilms.At(i)
		g.populateMetrics(cfg, ilm)
	}
}

func (g *metricGenerator) populateMetrics(cfg MetricCfg, ilm pdata.InstrumentationLibraryMetrics) {
	metrics := ilm.Metrics()
	metrics.Resize(cfg.NumMetricsPerILM)
	for i := 0; i < cfg.NumMetricsPerILM; i++ {
		metric := metrics.At(i)
		g.populateMetricDesc(cfg, metric)
		switch cfg.MetricDescriptorType {
		case pdata.MetricDataTypeIntGauge:
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
			populateIntPoints(cfg, metric.IntGauge().DataPoints())
		case pdata.MetricDataTypeDoubleGauge:
			metric.SetDataType(pdata.MetricDataTypeDoubleGauge)
			populateDoublePoints(cfg, metric.DoubleGauge().DataPoints())
		case pdata.MetricDataTypeIntSum:
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			sum := metric.IntSum()
			sum.SetIsMonotonic(cfg.IsMonotonicSum)
			sum.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
			populateIntPoints(cfg, sum.DataPoints())
		case pdata.MetricDataTypeDoubleSum:
			metric.SetDataType(pdata.MetricDataTypeDoubleSum)
			sum := metric.DoubleSum()
			sum.SetIsMonotonic(cfg.IsMonotonicSum)
			sum.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
			populateDoublePoints(cfg, sum.DataPoints())
		case pdata.MetricDataTypeIntHistogram:
			metric.SetDataType(pdata.MetricDataTypeIntHistogram)
			histo := metric.IntHistogram()
			histo.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
			populateIntHistogram(cfg, histo)
		case pdata.MetricDataTypeDoubleHistogram:
			metric.SetDataType(pdata.MetricDataTypeDoubleHistogram)
			histo := metric.DoubleHistogram()
			histo.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
			populateDoubleHistogram(cfg, histo)
		}
	}
}

func (g *metricGenerator) populateMetricDesc(cfg MetricCfg, metric pdata.Metric) {
	metric.SetName(fmt.Sprintf("%smetric_%d", cfg.MetricNamePrefix, g.metricID))
	g.metricID++
	metric.SetDescription("my-md-description")
	metric.SetUnit("my-md-units")
}

func populateIntPoints(cfg MetricCfg, pts pdata.IntDataPointSlice) {
	pts.Resize(cfg.NumPtsPerMetric)
	for i := 0; i < cfg.NumPtsPerMetric; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.Timestamp(cfg.StartTime))
		pt.SetTimestamp(getTimestamp(cfg.StartTime, cfg.StepSize, i))
		pt.SetValue(int64(cfg.PtVal + i))
		populatePtLabels(cfg, pt.LabelsMap())
	}
}

func populateDoublePoints(cfg MetricCfg, pts pdata.DoubleDataPointSlice) {
	pts.Resize(cfg.NumPtsPerMetric)
	for i := 0; i < cfg.NumPtsPerMetric; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.Timestamp(cfg.StartTime))
		pt.SetTimestamp(getTimestamp(cfg.StartTime, cfg.StepSize, i))
		pt.SetValue(float64(cfg.PtVal + i))
		populatePtLabels(cfg, pt.LabelsMap())
	}
}

func populateDoubleHistogram(cfg MetricCfg, dh pdata.DoubleHistogram) {
	pts := dh.DataPoints()
	pts.Resize(cfg.NumPtsPerMetric)
	for i := 0; i < cfg.NumPtsPerMetric; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.Timestamp(cfg.StartTime))
		ts := getTimestamp(cfg.StartTime, cfg.StepSize, i)
		pt.SetTimestamp(ts)
		populatePtLabels(cfg, pt.LabelsMap())
		setDoubleHistogramBounds(pt, 1, 2, 3, 4, 5)
		addDoubleHistogramVal(pt, 1)
		for i := 0; i < cfg.PtVal; i++ {
			addDoubleHistogramVal(pt, 3)
		}
		addDoubleHistogramVal(pt, 5)
	}
}

func setDoubleHistogramBounds(hdp pdata.DoubleHistogramDataPoint, bounds ...float64) {
	hdp.SetBucketCounts(make([]uint64, len(bounds)))
	hdp.SetExplicitBounds(bounds)
}

func addDoubleHistogramVal(hdp pdata.DoubleHistogramDataPoint, val float64) {
	hdp.SetCount(hdp.Count() + 1)
	hdp.SetSum(hdp.Sum() + val)
	buckets := hdp.BucketCounts()
	bounds := hdp.ExplicitBounds()
	for i := 0; i < len(bounds); i++ {
		bound := bounds[i]
		if val <= bound {
			buckets[i]++
			break
		}
	}
}

func populateIntHistogram(cfg MetricCfg, dh pdata.IntHistogram) {
	pts := dh.DataPoints()
	pts.Resize(cfg.NumPtsPerMetric)
	for i := 0; i < cfg.NumPtsPerMetric; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.Timestamp(cfg.StartTime))
		ts := getTimestamp(cfg.StartTime, cfg.StepSize, i)
		pt.SetTimestamp(ts)
		populatePtLabels(cfg, pt.LabelsMap())
		setIntHistogramBounds(pt, 1, 2, 3, 4, 5)
		addIntHistogramVal(pt, 1)
		for i := 0; i < cfg.PtVal; i++ {
			addIntHistogramVal(pt, 3)
		}
		addIntHistogramVal(pt, 5)
	}
}

func setIntHistogramBounds(hdp pdata.IntHistogramDataPoint, bounds ...float64) {
	hdp.SetBucketCounts(make([]uint64, len(bounds)))
	hdp.SetExplicitBounds(bounds)
}

func addIntHistogramVal(hdp pdata.IntHistogramDataPoint, val int64) {
	hdp.SetCount(hdp.Count() + 1)
	hdp.SetSum(hdp.Sum() + val)
	buckets := hdp.BucketCounts()
	bounds := hdp.ExplicitBounds()
	for i := 0; i < len(bounds); i++ {
		bound := bounds[i]
		if float64(val) <= bound {
			buckets[i]++
			break
		}
	}
}

func populatePtLabels(cfg MetricCfg, lm pdata.StringMap) {
	for i := 0; i < cfg.NumPtLabels; i++ {
		k := fmt.Sprintf("pt-label-key-%d", i)
		v := fmt.Sprintf("pt-label-val-%d", i)
		lm.Insert(k, v)
	}
}

func getTimestamp(startTime uint64, stepSize uint64, i int) pdata.Timestamp {
	return pdata.Timestamp(startTime + (stepSize * uint64(i+1)))
}
