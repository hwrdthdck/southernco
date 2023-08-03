// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordGitRepositoryBranchCountDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordGitRepositoryBranchTimeDataPoint(ts, 1, "attr-val", "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordGitRepositoryContributorCountDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordGitRepositoryCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordGitRepositoryPullRequestTimeDataPoint(ts, 1, "attr-val", "attr-val")

			metrics := mb.Emit(WithGitVendorName("attr-val"), WithOrganizationName("attr-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("git.vendor.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.GitVendorName.Enabled, ok)
			if mb.resourceAttributesConfig.GitVendorName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("organization.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.OrganizationName.Enabled, ok)
			if mb.resourceAttributesConfig.OrganizationName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 2)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "git.repository.branch.count":
					assert.False(t, validatedMetrics["git.repository.branch.count"], "Found a duplicate in the metrics slice: git.repository.branch.count")
					validatedMetrics["git.repository.branch.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of branches that exist in the repository", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "git.repository.branch.time":
					assert.False(t, validatedMetrics["git.repository.branch.time"], "Found a duplicate in the metrics slice: git.repository.branch.time")
					validatedMetrics["git.repository.branch.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time the branch has existed", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("branch.name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "git.repository.contributor.count":
					assert.False(t, validatedMetrics["git.repository.contributor.count"], "Found a duplicate in the metrics slice: git.repository.contributor.count")
					validatedMetrics["git.repository.contributor.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Total number of unique contributors to this repository", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "git.repository.count":
					assert.False(t, validatedMetrics["git.repository.count"], "Found a duplicate in the metrics slice: git.repository.count")
					validatedMetrics["git.repository.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of repositories that exist in an organization", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "git.repository.pull_request.time":
					assert.False(t, validatedMetrics["git.repository.pull_request.time"], "Found a duplicate in the metrics slice: git.repository.pull_request.time")
					validatedMetrics["git.repository.pull_request.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time the PR has been open", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("branch.name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				}
			}
		})
	}
}
