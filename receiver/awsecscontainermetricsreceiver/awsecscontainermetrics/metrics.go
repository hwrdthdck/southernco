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

package awsecscontainermetrics

import (
	"go.opentelemetry.io/collector/consumer/consumerdata"
)

func MetricsData(containerStatsMap map[string]ContainerStats, metadata TaskMetadata, typeStr string) []*consumerdata.MetricsData {
	acc := &metricDataAccumulator{}
	acc.getStats(containerStatsMap, metadata)

	for _, md := range acc.md {
		// TODO this should prob go in core
		md.Resource.Labels["receiver"] = typeStr
	}
	return acc.md
}
