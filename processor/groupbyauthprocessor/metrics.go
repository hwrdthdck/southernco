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

package groupbyauthprocessor

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/collector/obsreport"
)

var (
	mBatchesNoToken     = stats.Int64("processor_groupbyauth_batches_no_token", "Batches that had no token available", stats.UnitDimensionless)
	mBatchesWithToken   = stats.Int64("processor_groupbyauth_batches_with_token", "Batches that had token available", stats.UnitDimensionless)
	mNumTracesConf      = stats.Int64("processor_groupbyauth_conf_num_traces", "Maximum number of traces to hold in the internal storage", stats.UnitDimensionless)
	mNumEventsInQueue   = stats.Int64("processor_groupbyauth_num_events_in_queue", "Number of events currently in the queue", stats.UnitDimensionless)
	mNumTracesInMemory  = stats.Int64("processor_groupbyauth_num_traces_in_memory", "Number of traces currently in the in-memory storage", stats.UnitDimensionless)
	mTracesEvicted      = stats.Int64("processor_groupbyauth_traces_evicted", "Traces evicted from the internal buffer", stats.UnitDimensionless)
	mReleasedSpans      = stats.Int64("processor_groupbyauth_spans_released", "Spans released to the next consumer", stats.UnitDimensionless)
	mReleasedTraces     = stats.Int64("processor_groupbyauth_traces_released", "Traces released to the next consumer", stats.UnitDimensionless)
	mIncompleteReleases = stats.Int64("processor_groupbyauth_incomplete_releases", "Releases that are suspected to have been incomplete", stats.UnitDimensionless)
	mEventLatency       = stats.Int64("processor_groupbyauth_event_latency", "How long the queue events are taking to be processed", stats.UnitMilliseconds)
)

// MetricViews return the metrics views according to given telemetry level.
func MetricViews() []*view.View {
	legacyViews := []*view.View{
		{
			Name:        mBatchesNoToken.Name(),
			Measure:     mBatchesNoToken,
			Description: mBatchesNoToken.Description(),
			Aggregation: view.LastValue(),
		},
		{
			Name:        mBatchesWithToken.Name(),
			Measure:     mBatchesWithToken,
			Description: mBatchesWithToken.Description(),
			Aggregation: view.LastValue(),
		},
		{
			Name:        mNumTracesConf.Name(),
			Measure:     mNumTracesConf,
			Description: mNumTracesConf.Description(),
			Aggregation: view.LastValue(),
		},
		{
			Name:        mNumEventsInQueue.Name(),
			Measure:     mNumEventsInQueue,
			Description: mNumEventsInQueue.Description(),
			Aggregation: view.LastValue(),
		},
		{
			Name:        mNumTracesInMemory.Name(),
			Measure:     mNumTracesInMemory,
			Description: mNumTracesInMemory.Description(),
			Aggregation: view.LastValue(),
		},
		{
			Name:        mTracesEvicted.Name(),
			Measure:     mTracesEvicted,
			Description: mTracesEvicted.Description(),
			// sum allows us to start from 0, count will only show up if there's at least one eviction, which might take a while to happen (if ever!)
			Aggregation: view.Sum(),
		},
		{
			Name:        mReleasedSpans.Name(),
			Measure:     mReleasedSpans,
			Description: mReleasedSpans.Description(),
			Aggregation: view.Sum(),
		},
		{
			Name:        mReleasedTraces.Name(),
			Measure:     mReleasedTraces,
			Description: mReleasedTraces.Description(),
			Aggregation: view.Sum(),
		},
		{
			Name:        mIncompleteReleases.Name(),
			Measure:     mIncompleteReleases,
			Description: mIncompleteReleases.Description(),
			Aggregation: view.Sum(),
		},
		{
			Name:        mEventLatency.Name(),
			Measure:     mEventLatency,
			Description: mEventLatency.Description(),
			TagKeys: []tag.Key{
				tag.MustNewKey("event"),
			},
			Aggregation: view.Distribution(0, 5, 10, 20, 50, 100, 200, 500, 1000),
		},
	}

	return obsreport.ProcessorMetricViews(string(typeStr), legacyViews)
}
