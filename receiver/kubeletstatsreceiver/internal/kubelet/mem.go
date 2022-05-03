// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

func addMemoryMetrics(dest pmetric.MetricSlice, memoryMetrics metadata.MemoryMetrics, s *stats.MemoryStats, currentTime pcommon.Timestamp) {
	if s == nil {
		return
	}

	addIntGauge(dest, memoryMetrics.Available, s.AvailableBytes, currentTime)
	addIntGauge(dest, memoryMetrics.Usage, s.UsageBytes, currentTime)
	addIntGauge(dest, memoryMetrics.Rss, s.RSSBytes, currentTime)
	addIntGauge(dest, memoryMetrics.WorkingSet, s.WorkingSetBytes, currentTime)
	addIntGauge(dest, memoryMetrics.PageFaults, s.PageFaults, currentTime)
	addIntGauge(dest, memoryMetrics.MajorPageFaults, s.MajorPageFaults, currentTime)
}
