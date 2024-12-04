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

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
		{
			name:        "filter_set_include",
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "filter_set_exclude",
			resAttrsSet: testDataSetAll,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryChangeCountDataPoint(ts, 1, "vcs.repository.url.full-val", AttributeChangeStateOpen, "repository.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryChangeTimeOpenDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryChangeTimeToApprovalDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryChangeTimeToMergeDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val")

			allMetricsCount++
			mb.RecordVcsRepositoryContributorCountDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryRefCountDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", AttributeRefTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryRefLinesAddedDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val", AttributeRefTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryRefLinesDeletedDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val", AttributeRefTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryRefRevisionsAheadDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val", AttributeRefTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryRefRevisionsBehindDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val", AttributeRefTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryRefTimeDataPoint(ts, 1, "vcs.repository.url.full-val", "repository.name-val", "ref.name-val", AttributeRefTypeBranch)

			rb := mb.NewResourceBuilder()
			rb.SetOrganizationName("organization.name-val")
			rb.SetVcsVendorName("vcs.vendor.name-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "vcs.repository.change.count":
					assert.False(t, validatedMetrics["vcs.repository.change.count"], "Found a duplicate in the metrics slice: vcs.repository.change.count")
					validatedMetrics["vcs.repository.change.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of changes (pull requests) in a repository, categorized by their state (either open or merged).", ms.At(i).Description())
					assert.Equal(t, "{change}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("change.state")
					assert.True(t, ok)
					assert.EqualValues(t, "open", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
				case "vcs.repository.change.time_open":
					assert.False(t, validatedMetrics["vcs.repository.change.time_open"], "Found a duplicate in the metrics slice: vcs.repository.change.time_open")
					validatedMetrics["vcs.repository.change.time_open"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of time a change (pull request) has been open.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
				case "vcs.repository.change.time_to_approval":
					assert.False(t, validatedMetrics["vcs.repository.change.time_to_approval"], "Found a duplicate in the metrics slice: vcs.repository.change.time_to_approval")
					validatedMetrics["vcs.repository.change.time_to_approval"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of time it took a change (pull request) to go from open to approved.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
				case "vcs.repository.change.time_to_merge":
					assert.False(t, validatedMetrics["vcs.repository.change.time_to_merge"], "Found a duplicate in the metrics slice: vcs.repository.change.time_to_merge")
					validatedMetrics["vcs.repository.change.time_to_merge"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of time it took a change (pull request) to go from open to merged.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
				case "vcs.repository.contributor.count":
					assert.False(t, validatedMetrics["vcs.repository.contributor.count"], "Found a duplicate in the metrics slice: vcs.repository.contributor.count")
					validatedMetrics["vcs.repository.contributor.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of unique contributors to a repository.", ms.At(i).Description())
					assert.Equal(t, "{contributor}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
				case "vcs.repository.count":
					assert.False(t, validatedMetrics["vcs.repository.count"], "Found a duplicate in the metrics slice: vcs.repository.count")
					validatedMetrics["vcs.repository.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of repositories in an organization.", ms.At(i).Description())
					assert.Equal(t, "{repository}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "vcs.repository.ref.count":
					assert.False(t, validatedMetrics["vcs.repository.ref.count"], "Found a duplicate in the metrics slice: vcs.repository.ref.count")
					validatedMetrics["vcs.repository.ref.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of refs of type branch in a repository.", ms.At(i).Description())
					assert.Equal(t, "{ref}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.type")
					assert.True(t, ok)
					assert.EqualValues(t, "branch", attrVal.Str())
				case "vcs.repository.ref.lines_added":
					assert.False(t, validatedMetrics["vcs.repository.ref.lines_added"], "Found a duplicate in the metrics slice: vcs.repository.ref.lines_added")
					validatedMetrics["vcs.repository.ref.lines_added"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of lines added in a ref (branch) relative to the default branch (trunk).", ms.At(i).Description())
					assert.Equal(t, "{line}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.type")
					assert.True(t, ok)
					assert.EqualValues(t, "branch", attrVal.Str())
				case "vcs.repository.ref.lines_deleted":
					assert.False(t, validatedMetrics["vcs.repository.ref.lines_deleted"], "Found a duplicate in the metrics slice: vcs.repository.ref.lines_deleted")
					validatedMetrics["vcs.repository.ref.lines_deleted"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of lines deleted in a ref (branch) relative to the default branch (trunk).", ms.At(i).Description())
					assert.Equal(t, "{line}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.type")
					assert.True(t, ok)
					assert.EqualValues(t, "branch", attrVal.Str())
				case "vcs.repository.ref.revisions_ahead":
					assert.False(t, validatedMetrics["vcs.repository.ref.revisions_ahead"], "Found a duplicate in the metrics slice: vcs.repository.ref.revisions_ahead")
					validatedMetrics["vcs.repository.ref.revisions_ahead"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of revisions (commits) a ref (branch) is ahead of the default branch (trunk).", ms.At(i).Description())
					assert.Equal(t, "{revision}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.type")
					assert.True(t, ok)
					assert.EqualValues(t, "branch", attrVal.Str())
				case "vcs.repository.ref.revisions_behind":
					assert.False(t, validatedMetrics["vcs.repository.ref.revisions_behind"], "Found a duplicate in the metrics slice: vcs.repository.ref.revisions_behind")
					validatedMetrics["vcs.repository.ref.revisions_behind"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of revisions (commits) a ref (branch) is behind the default branch (trunk).", ms.At(i).Description())
					assert.Equal(t, "{revision}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.type")
					assert.True(t, ok)
					assert.EqualValues(t, "branch", attrVal.Str())
				case "vcs.repository.ref.time":
					assert.False(t, validatedMetrics["vcs.repository.ref.time"], "Found a duplicate in the metrics slice: vcs.repository.ref.time")
					validatedMetrics["vcs.repository.ref.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.EqualValues(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("repository.name")
					assert.True(t, ok)
					assert.EqualValues(t, "repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.name")
					assert.True(t, ok)
					assert.EqualValues(t, "ref.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("ref.type")
					assert.True(t, ok)
					assert.EqualValues(t, "branch", attrVal.Str())
				}
			}
		})
	}
}
