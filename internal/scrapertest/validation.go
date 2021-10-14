package scrapertest

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/model/otlp"
	"go.opentelemetry.io/collector/model/pdata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/scraperhelper"
)

// ValidateScraper validates That a scraper extracts metrics correctly.
// it runs the scrape function and the compares the scraped metrics against
// a given json file that represents an expected pdata.Metrics object.
func ValidateScraper(t *testing.T, actualScraper scraperhelper.ScrapeResourceMetrics, expectedFilePath string) {
	scrapedRMS, err := actualScraper(context.Background())
	require.NoError(t, err)

	expectedFileBytes, err := ioutil.ReadFile(expectedFilePath)
	require.NoError(t, err)
	unmarshaller := otlp.NewJSONMetricsUnmarshaler()
	expectedMetrics, err := unmarshaller.UnmarshalMetrics(expectedFileBytes)
	require.NoError(t, err)

	eMetricSlice := expectedMetrics.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics()
	aMetricSlice := scrapedRMS.At(0).InstrumentationLibraryMetrics().At(0).Metrics()

	require.NoError(t, CompareMetrics(eMetricSlice, aMetricSlice))
}

// CompareMetrics compares each part of two given metric slices and returns
// an error if they don't match. The error describes what didn't match.
func CompareMetrics(expectedAll, actualAll pdata.MetricSlice) error {
	if actualAll.Len() != expectedAll.Len() {
		return fmt.Errorf("metric slices not of same length")
	}

	lessFunc := func(a, b pdata.Metric) bool {
		return a.Name() < b.Name()
	}

	actualMetrics := actualAll.Sort(lessFunc)
	expectedMetrics := expectedAll.Sort(lessFunc)

	for i := 0; i < actualMetrics.Len(); i++ {
		actual := actualMetrics.At(i)
		expected := expectedMetrics.At(i)

		if actual.Name() != expected.Name() {
			return fmt.Errorf("metric name does not match expected: %s, actual: %s", expected.Name(), actual.Name())
		}
		if actual.DataType() != expected.DataType() {
			return fmt.Errorf("metric datatype does not match expected: %s, actual: %s", expected.DataType(), actual.DataType())
		}
		if actual.Description() != expected.Description() {
			return fmt.Errorf("metric description does not match expected: %s, actual: %s", expected.Description(), actual.Description())
		}
		if actual.Unit() != expected.Unit() {
			return fmt.Errorf("metric Unit does not match expected: %s, actual: %s", expected.Unit(), actual.Unit())
		}

		var actualDataPoints pdata.NumberDataPointSlice
		var expectedDataPoints pdata.NumberDataPointSlice

		switch actual.DataType() {
		case pdata.MetricDataTypeGauge:
			actualDataPoints = actual.Gauge().DataPoints()
			expectedDataPoints = expected.Gauge().DataPoints()
		case pdata.MetricDataTypeSum:
			if actual.Sum().AggregationTemporality() != expected.Sum().AggregationTemporality() {
				return fmt.Errorf("metric AggregationTemporality does not match expected: %s, actual: %s", expected.Sum().AggregationTemporality(), actual.Sum().AggregationTemporality())
			}
			if actual.Sum().IsMonotonic() != expected.Sum().IsMonotonic() {
				return fmt.Errorf("metric IsMonotonic does not match expected: %t, actual: %t", expected.Sum().IsMonotonic(), actual.Sum().IsMonotonic())
			}
			actualDataPoints = actual.Sum().DataPoints()
			expectedDataPoints = expected.Sum().DataPoints()
		}

		if actualDataPoints.Len() != expectedDataPoints.Len() {
			return fmt.Errorf("length of datapoints don't match, metric name: %s", actual.Name())
		}

		dataPointMatches := 0
		for j := 0; j < expectedDataPoints.Len(); j++ {
			edp := expectedDataPoints.At(j)
			for k := 0; k < actualDataPoints.Len(); k++ {
				adp := actualDataPoints.At(k)
				adpAttributes := adp.Attributes()
				labelMatches := true

				if edp.Attributes().Len() != adpAttributes.Len() {
					break
				}
				edp.Attributes().Range(func(k string, v pdata.AttributeValue) bool {
					if attributeVal, ok := adpAttributes.Get(k); ok && attributeVal.StringVal() == v.StringVal() {
						return true
					}
					labelMatches = false
					return false
				})
				if !labelMatches {
					continue
				}
				if edp.Type() != adp.Type() {
					return fmt.Errorf("metric datapoint types don't match: expected type: %v, actual type: %v", edp.Type(), adp.Type())
				}
				if edp.IntVal() != adp.IntVal() {
					return fmt.Errorf("metric datapoint IntVal doesn't match expected: %d, actual: %d", edp.IntVal(), adp.IntVal())
				}
				if edp.DoubleVal() != adp.DoubleVal() {
					return fmt.Errorf("metric datapoint DoubleVal doesn't match expected: %f, actual: %f", edp.DoubleVal(), adp.DoubleVal())
				}
				dataPointMatches++
				break
			}
		}
		if dataPointMatches != expectedDataPoints.Len() {
			return fmt.Errorf("missing datapoints for: metric name: %s", actual.Name())
		}
	}
	return nil
}
