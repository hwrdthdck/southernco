// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package intervalprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/intervalprocessor"

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor/processortest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
)

func TestAggregation(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"basic_aggregation",
		"non_monotonic_sums_are_passed_through",
		"summaries_are_passed_through",
		"histograms_are_aggregated",
		"exp_histograms_are_aggregated",
		"gauges_are_aggregated",
		"all_delta_metrics_are_passed_through",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &Config{Interval: time.Second}

	for _, tc := range testCases {
		testName := tc

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			// next stores the results of the filter metric processor
			next := &consumertest.MetricsSink{}

			factory := NewFactory()
			mgp, err := factory.CreateMetricsProcessor(
				context.Background(),
				processortest.NewNopSettings(),
				config,
				next,
			)
			require.NoError(t, err)

			dir := filepath.Join("testdata", testName)

			md, err := golden.ReadMetrics(filepath.Join(dir, "input.yaml"))
			require.NoError(t, err)

			// Test that ConsumeMetrics works
			err = mgp.ConsumeMetrics(ctx, md)
			require.NoError(t, err)

			require.IsType(t, &Processor{}, mgp)
			processor := mgp.(*Processor)

			// Pretend we hit the interval timer and call export
			processor.exportMetrics()

			// All the lookup tables should now be empty
			require.Empty(t, processor.rmLookup)
			require.Empty(t, processor.smLookup)
			require.Empty(t, processor.mLookup)
			require.Empty(t, processor.numberLookup)
			require.Empty(t, processor.histogramLookup)
			require.Empty(t, processor.expHistogramLookup)

			// Exporting again should return nothing
			processor.exportMetrics()

			// Next should have gotten three data sets:
			// 1. Anything left over from ConsumeMetrics()
			// 2. Anything exported from exportMetrics()
			// 3. An empty entry for the second call to exportMetrics()
			allMetrics := next.AllMetrics()
			require.Len(t, allMetrics, 3)

			nextData := allMetrics[0]
			exportData := allMetrics[1]
			secondExportData := allMetrics[2]

			expectedNextData, err := golden.ReadMetrics(filepath.Join(dir, "next.yaml"))
			require.NoError(t, err)
			require.NoError(t, pmetrictest.CompareMetrics(expectedNextData, nextData))

			expectedExportData, err := golden.ReadMetrics(filepath.Join(dir, "output.yaml"))
			require.NoError(t, err)
			require.NoError(t, pmetrictest.CompareMetrics(expectedExportData, exportData))

			require.NoError(t, pmetrictest.CompareMetrics(pmetric.NewMetrics(), secondExportData), "the second export data should be empty")
		})
	}
}
