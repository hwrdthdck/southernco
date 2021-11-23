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

package internal

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/textparse"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

func isUsefulLabelPdata(mType pdata.MetricDataType, labelKey string) bool {
	switch labelKey {
	case model.MetricNameLabel, model.InstanceLabel, model.SchemeLabel, model.MetricsPathLabel, model.JobLabel:
		return false
	case model.BucketLabel:
		return mType != pdata.MetricDataTypeHistogram
	case model.QuantileLabel:
		return mType != pdata.MetricDataTypeSummary
	}
	return true
}

func getBoundaryPdata(metricType pdata.MetricDataType, labels labels.Labels) (float64, error) {
	labelName := ""
	switch metricType {
	case pdata.MetricDataTypeHistogram:
		labelName = model.BucketLabel
	case pdata.MetricDataTypeSummary:
		labelName = model.QuantileLabel
	default:
		return 0, errNoBoundaryLabel
	}

	v := labels.Get(labelName)
	if v == "" {
		return 0, errEmptyBoundaryLabel
	}

	return strconv.ParseFloat(v, 64)
}

func convToPdataMetricType(metricType textparse.MetricType) pdata.MetricDataType {
	switch metricType {
	case textparse.MetricTypeCounter: // metricspb.MetricDescriptor_CUMULATIVE_DOUBLE
		// always use float64, as it's the internal data type used in prometheus
		return pdata.MetricDataTypeSum
	// textparse.MetricTypeUnknown is converted to gauge by default to fix Prometheus untyped metrics from being dropped
	case textparse.MetricTypeGauge, textparse.MetricTypeUnknown: // metricspb.MetricDescriptor_GAUGE_DOUBLE
		return pdata.MetricDataTypeGauge
	case textparse.MetricTypeHistogram: // metricspb.MetricDescriptor_CUMULATIVE_DISTRIBUTION
		return pdata.MetricDataTypeHistogram
	// dropping support for gaugehistogram for now until we have an official spec of its implementation
	// a draft can be found in: https://docs.google.com/document/d/1KwV0mAXwwbvvifBvDKH_LU1YjyXE_wxCkHNoCGq1GX0/edit#heading=h.1cvzqd4ksd23
	// case textparse.MetricTypeGaugeHistogram:
	//	return metricspb.MetricDescriptor_GAUGE_DISTRIBUTION
	case textparse.MetricTypeSummary:
		return pdata.MetricDataTypeSummary
	default:
		// including: textparse.MetricTypeInfo, textparse.MetricTypeStateset
		return pdata.MetricDataTypeNone
	}
}

type metricBuilderPdata struct {
	metrics              pdata.MetricSlice
	families             map[string]MetricFamilyPdata
	hasData              bool
	hasInternalMetric    bool
	mc                   MetadataCache
	numTimeseries        int
	droppedTimeseries    int
	useStartTimeMetric   bool
	startTimeMetricRegex *regexp.Regexp
	startTime            float64
	intervalStartTimeMs  int64
	logger               *zap.Logger
}

// newMetricBuilder creates a MetricBuilder which is allowed to feed all the datapoints from a single prometheus
// scraped page by calling its AddDataPoint function, and turn them into a pdata.Metrics object.
// by calling its Build function
func newMetricBuilderPdata(mc MetadataCache, useStartTimeMetric bool, startTimeMetricRegex string, logger *zap.Logger, intervalStartTimeMs int64) *metricBuilderPdata {
	var regex *regexp.Regexp
	if startTimeMetricRegex != "" {
		regex, _ = regexp.Compile(startTimeMetricRegex)
	}
	return &metricBuilderPdata{
		metrics:              pdata.NewMetricSlice(),
		families:             map[string]MetricFamilyPdata{},
		mc:                   mc,
		logger:               logger,
		numTimeseries:        0,
		droppedTimeseries:    0,
		useStartTimeMetric:   useStartTimeMetric,
		startTimeMetricRegex: regex,
		intervalStartTimeMs:  intervalStartTimeMs,
	}
}

func (b *metricBuilderPdata) matchStartTimeMetric(metricName string) bool {
	if b.startTimeMetricRegex != nil {
		return b.startTimeMetricRegex.MatchString(metricName)
	}

	return metricName == startTimeMetricName
}

// AddDataPoint is for feeding prometheus data complexValue in its processing order
func (b *metricBuilderPdata) AddDataPoint(ls labels.Labels, t int64, v float64) error {
	// Any datapoint with duplicate labels MUST be rejected per:
	// * https://github.com/open-telemetry/wg-prometheus/issues/44
	// * https://github.com/open-telemetry/opentelemetry-collector/issues/3407
	// as Prometheus rejects such too as of version 2.16.0, released on 2020-02-13.
	seen := make(map[string]bool)
	dupLabels := make([]string, 0, len(ls))
	for _, label := range ls {
		if _, ok := seen[label.Name]; ok {
			dupLabels = append(dupLabels, label.Name)
		}
		seen[label.Name] = true
	}
	if len(dupLabels) != 0 {
		sort.Strings(dupLabels)
		return fmt.Errorf("invalid sample: non-unique label names: %q", dupLabels)
	}

	metricName := ls.Get(model.MetricNameLabel)
	switch {
	case metricName == "":
		b.numTimeseries++
		b.droppedTimeseries++
		return errMetricNameNotFound
	case isInternalMetric(metricName):
		b.hasInternalMetric = true
		lm := ls.Map()
		// See https://www.prometheus.io/docs/concepts/jobs_instances/#automatically-generated-labels-and-time-series
		// up: 1 if the instance is healthy, i.e. reachable, or 0 if the scrape failed.
		if metricName == scrapeUpMetricName && v != 1.0 {
			if v == 0.0 {
				b.logger.Warn("Failed to scrape Prometheus endpoint",
					zap.Int64("scrape_timestamp", t),
					zap.String("target_labels", fmt.Sprintf("%v", lm)))
			} else {
				b.logger.Warn("The 'up' metric contains invalid value",
					zap.Float64("value", v),
					zap.Int64("scrape_timestamp", t),
					zap.String("target_labels", fmt.Sprintf("%v", lm)))
			}
		}
	case b.useStartTimeMetric && b.matchStartTimeMetric(metricName):
		b.startTime = v
	}

	b.hasData = true

	familyName := normalizeMetricName(metricName)
	curMF, ok := b.families[familyName]
	if !ok {
		if mf, ok := b.families[metricName]; ok {
			curMF = mf
		} else {
			curMF = newMetricFamilyPdata(metricName, b.mc, b.logger, b.intervalStartTimeMs)
			b.families[familyName] = curMF
		}
	}

	return curMF.Add(metricName, ls, t, v)
}

// Build an pdata.MetricSlice based on all added data complexValue.
// The only error returned by this function is errNoDataToBuild.
func (b *metricBuilderPdata) Build() (*pdata.MetricSlice, int, int, error) {
	if !b.hasData {
		if b.hasInternalMetric {
			metricsL := pdata.NewMetricSlice()
			return &metricsL, 0, 0, nil
		}
		return nil, 0, 0, errNoDataToBuild
	}

	for _, mf := range b.families {
		ts, dts := mf.ToMetricPdata(&b.metrics)
		b.numTimeseries += ts
		b.droppedTimeseries += dts
	}

	return &b.metrics, b.numTimeseries, b.droppedTimeseries, nil
}
