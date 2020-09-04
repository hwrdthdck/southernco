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
	"time"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.opentelemetry.io/collector/consumer/consumerdata"
)

type metricDataAccumulator struct {
	md []*consumerdata.MetricsData
}

const (
	taskPrefix      = "ecs.task."
	containerPrefix = "ecs.container."
	cpusInVCpu      = 1024
)

func (acc *metricDataAccumulator) getMetricsData(containerStatsMap map[string]ContainerStats, metadata TaskMetadata) {

	var taskMemLimit uint64
	var taskCPULimit float64
	taskMetrics := ECSMetrics{}

	for _, containerMetadata := range metadata.Containers {
		stats := containerStatsMap[containerMetadata.DockerId]
		containerMetrics := getContainerMetrics(stats)
		containerMetrics.MemoryReserved = *containerMetadata.Limits.Memory
		containerMetrics.CPUReserved = *containerMetadata.Limits.CPU

		taskMemLimit += *containerMetadata.Limits.Memory
		taskCPULimit += *containerMetadata.Limits.CPU

		containerResource := containerResource(containerMetadata)
		labelKeys, labelValues := containerLabelKeysAndValues(containerMetadata)

		acc.accumulate(
			timestampProto(time.Now()),
			containerResource,

			convertToOTMetrics(containerPrefix, containerMetrics, labelKeys, labelValues),
		)

		aggregateTaskMetrics(&taskMetrics, containerMetrics)
	}

	// Overwrite Memory limit with task level limit
	if metadata.Limits.Memory != nil {
		taskMetrics.MemoryReserved = *metadata.Limits.Memory
	}

	taskMetrics.CPUReserved = taskMetrics.CPUReserved / cpusInVCpu

	// Overwrite CPU limit with task level limit
	if metadata.Limits.CPU != nil {
		taskMetrics.CPUReserved = *metadata.Limits.CPU
	}

	taskResource := taskResource(metadata)
	taskLabelKeys, taskLabelValues := taskLabelKeysAndValues(metadata)
	acc.accumulate(
		timestampProto(time.Now()),
		taskResource,
		convertToOTMetrics(taskPrefix, taskMetrics, taskLabelKeys, taskLabelValues),
	)
}

func (acc *metricDataAccumulator) accumulate(
	startTime *timestamp.Timestamp,
	r *resourcepb.Resource,
	m ...[]*metricspb.Metric,
) {
	var resourceMetrics []*metricspb.Metric
	for _, metrics := range m {
		for _, metric := range metrics {
			if metric != nil {
				metric.Timeseries[0].StartTimestamp = startTime
				resourceMetrics = append(resourceMetrics, metric)
			}
		}
	}

	acc.md = append(acc.md, &consumerdata.MetricsData{
		Metrics:  resourceMetrics,
		Resource: r,
	})
}
