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
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

type MetricGroup string

// Values for MetricGroup enum.
const (
	ContainerMetricGroup = MetricGroup("container")
	PodMetricGroup       = MetricGroup("pod")
	NodeMetricGroup      = MetricGroup("node")
	VolumeMetricGroup    = MetricGroup("volume")
)

// ValidMetricGroups map of valid metrics.
var ValidMetricGroups = map[MetricGroup]bool{
	ContainerMetricGroup: true,
	PodMetricGroup:       true,
	NodeMetricGroup:      true,
	VolumeMetricGroup:    true,
}

type metricDataAccumulator struct {
	m                     []pmetric.Metrics
	metadata              Metadata
	logger                *zap.Logger
	metricGroupsToCollect map[MetricGroup]bool
	time                  time.Time
	mb                    *metadata.MetricsBuilder
}

func (a *metricDataAccumulator) nodeStats(s stats.NodeStats) {
	if !a.metricGroupsToCollect[NodeMetricGroup] {
		return
	}

	currentTime := pcommon.NewTimestampFromTime(a.time)
	addCPUMetrics(a.mb, metadata.NodeCPUMetrics, s.CPU, currentTime)
	addMemoryMetrics(a.mb, metadata.NodeMemoryMetrics, s.Memory, currentTime)
	addFilesystemMetrics(a.mb, metadata.NodeFilesystemMetrics, s.Fs, currentTime)
	addNetworkMetrics(a.mb, metadata.NodeNetworkMetrics, s.Network, currentTime)
	// todo s.Runtime.ImageFs

	a.m = append(a.m, a.mb.Emit(getNodeResourceOptions(s)...))
}

func (a *metricDataAccumulator) podStats(s stats.PodStats) {
	if !a.metricGroupsToCollect[PodMetricGroup] {
		return
	}

	currentTime := pcommon.NewTimestampFromTime(a.time)
	addCPUMetrics(a.mb, metadata.PodCPUMetrics, s.CPU, currentTime)
	addMemoryMetrics(a.mb, metadata.PodMemoryMetrics, s.Memory, currentTime)
	addFilesystemMetrics(a.mb, metadata.PodFilesystemMetrics, s.EphemeralStorage, currentTime)
	addNetworkMetrics(a.mb, metadata.PodNetworkMetrics, s.Network, currentTime)

	a.m = append(a.m, a.mb.Emit(getPodResourceOptions(s)...))
}

func (a *metricDataAccumulator) containerStats(sPod stats.PodStats, s stats.ContainerStats) {
	if !a.metricGroupsToCollect[ContainerMetricGroup] {
		return
	}

	ro, err := getContainerResourceOptions(sPod, s, a.metadata)
	if err != nil {
		a.logger.Warn(
			"failed to fetch container metrics",
			zap.String("pod", sPod.PodRef.Name),
			zap.String("container", s.Name),
			zap.Error(err))
		return
	}

	currentTime := pcommon.NewTimestampFromTime(a.time)
	addCPUMetrics(a.mb, metadata.ContainerCPUMetrics, s.CPU, currentTime)
	addMemoryMetrics(a.mb, metadata.ContainerMemoryMetrics, s.Memory, currentTime)
	addFilesystemMetrics(a.mb, metadata.ContainerFilesystemMetrics, s.Rootfs, currentTime)

	a.m = append(a.m, a.mb.Emit(ro...))
}

func (a *metricDataAccumulator) volumeStats(sPod stats.PodStats, s stats.VolumeStats) {
	if !a.metricGroupsToCollect[VolumeMetricGroup] {
		return
	}

	ro, err := getVolumeResourceOptions(sPod, s, a.metadata)
	if err != nil {
		a.logger.Warn(
			"Failed to gather additional volume metadata. Skipping metric collection.",
			zap.String("pod", sPod.PodRef.Name),
			zap.String("volume", s.Name),
			zap.Error(err))
		return
	}

	currentTime := pcommon.NewTimestampFromTime(a.time)
	addVolumeMetrics(a.mb, metadata.K8sVolumeMetrics, s, currentTime)

	a.m = append(a.m, a.mb.Emit(ro...))
}
