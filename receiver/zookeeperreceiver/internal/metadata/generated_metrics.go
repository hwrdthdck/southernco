// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer/pdata"
)

// Type is the component type name.
const Type config.Type = "zookeeperreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	ZookeeperApproximateDateSize   MetricIntf
	ZookeeperConnectionsAlive      MetricIntf
	ZookeeperEphemeralNodes        MetricIntf
	ZookeeperFollowers             MetricIntf
	ZookeeperFsyncThresholdExceeds MetricIntf
	ZookeeperLatencyAvg            MetricIntf
	ZookeeperLatencyMax            MetricIntf
	ZookeeperLatencyMin            MetricIntf
	ZookeeperMaxFileDescriptors    MetricIntf
	ZookeeperOpenFileDescriptors   MetricIntf
	ZookeeperOutstandingRequests   MetricIntf
	ZookeeperPacketsReceived       MetricIntf
	ZookeeperPacketsSent           MetricIntf
	ZookeeperPendingSyncs          MetricIntf
	ZookeeperSyncedFollowers       MetricIntf
	ZookeeperWatches               MetricIntf
	ZookeeperZnodes                MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"zookeeper.approximate_date_size",
		"zookeeper.connections_alive",
		"zookeeper.ephemeral_nodes",
		"zookeeper.followers",
		"zookeeper.fsync_threshold_exceeds",
		"zookeeper.latency.avg",
		"zookeeper.latency.max",
		"zookeeper.latency.min",
		"zookeeper.max_file_descriptors",
		"zookeeper.open_file_descriptors",
		"zookeeper.outstanding_requests",
		"zookeeper.packets.received",
		"zookeeper.packets.sent",
		"zookeeper.pending_syncs",
		"zookeeper.synced_followers",
		"zookeeper.watches",
		"zookeeper.znodes",
	}
}

var metricsByName = map[string]MetricIntf{
	"zookeeper.approximate_date_size":   Metrics.ZookeeperApproximateDateSize,
	"zookeeper.connections_alive":       Metrics.ZookeeperConnectionsAlive,
	"zookeeper.ephemeral_nodes":         Metrics.ZookeeperEphemeralNodes,
	"zookeeper.followers":               Metrics.ZookeeperFollowers,
	"zookeeper.fsync_threshold_exceeds": Metrics.ZookeeperFsyncThresholdExceeds,
	"zookeeper.latency.avg":             Metrics.ZookeeperLatencyAvg,
	"zookeeper.latency.max":             Metrics.ZookeeperLatencyMax,
	"zookeeper.latency.min":             Metrics.ZookeeperLatencyMin,
	"zookeeper.max_file_descriptors":    Metrics.ZookeeperMaxFileDescriptors,
	"zookeeper.open_file_descriptors":   Metrics.ZookeeperOpenFileDescriptors,
	"zookeeper.outstanding_requests":    Metrics.ZookeeperOutstandingRequests,
	"zookeeper.packets.received":        Metrics.ZookeeperPacketsReceived,
	"zookeeper.packets.sent":            Metrics.ZookeeperPacketsSent,
	"zookeeper.pending_syncs":           Metrics.ZookeeperPendingSyncs,
	"zookeeper.synced_followers":        Metrics.ZookeeperSyncedFollowers,
	"zookeeper.watches":                 Metrics.ZookeeperWatches,
	"zookeeper.znodes":                  Metrics.ZookeeperZnodes,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

func (m *metricStruct) FactoriesByName() map[string]func(pdata.Metric) {
	return map[string]func(pdata.Metric){
		Metrics.ZookeeperApproximateDateSize.Name():   Metrics.ZookeeperApproximateDateSize.Init,
		Metrics.ZookeeperConnectionsAlive.Name():      Metrics.ZookeeperConnectionsAlive.Init,
		Metrics.ZookeeperEphemeralNodes.Name():        Metrics.ZookeeperEphemeralNodes.Init,
		Metrics.ZookeeperFollowers.Name():             Metrics.ZookeeperFollowers.Init,
		Metrics.ZookeeperFsyncThresholdExceeds.Name(): Metrics.ZookeeperFsyncThresholdExceeds.Init,
		Metrics.ZookeeperLatencyAvg.Name():            Metrics.ZookeeperLatencyAvg.Init,
		Metrics.ZookeeperLatencyMax.Name():            Metrics.ZookeeperLatencyMax.Init,
		Metrics.ZookeeperLatencyMin.Name():            Metrics.ZookeeperLatencyMin.Init,
		Metrics.ZookeeperMaxFileDescriptors.Name():    Metrics.ZookeeperMaxFileDescriptors.Init,
		Metrics.ZookeeperOpenFileDescriptors.Name():   Metrics.ZookeeperOpenFileDescriptors.Init,
		Metrics.ZookeeperOutstandingRequests.Name():   Metrics.ZookeeperOutstandingRequests.Init,
		Metrics.ZookeeperPacketsReceived.Name():       Metrics.ZookeeperPacketsReceived.Init,
		Metrics.ZookeeperPacketsSent.Name():           Metrics.ZookeeperPacketsSent.Init,
		Metrics.ZookeeperPendingSyncs.Name():          Metrics.ZookeeperPendingSyncs.Init,
		Metrics.ZookeeperSyncedFollowers.Name():       Metrics.ZookeeperSyncedFollowers.Init,
		Metrics.ZookeeperWatches.Name():               Metrics.ZookeeperWatches.Init,
		Metrics.ZookeeperZnodes.Name():                Metrics.ZookeeperZnodes.Init,
	}
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"zookeeper.approximate_date_size",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.approximate_date_size")
			metric.SetDescription("Size of data in bytes that a ZooKeeper server has in its data tree.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.connections_alive",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.connections_alive")
			metric.SetDescription("Number of active clients connected to a ZooKeeper server.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.ephemeral_nodes",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.ephemeral_nodes")
			metric.SetDescription("Number of ephemeral nodes that a ZooKeeper server has in its data tree.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.followers",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.followers")
			metric.SetDescription("The number of followers in sync with the leader. Only exposed by the leader.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.fsync_threshold_exceeds",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.fsync_threshold_exceeds")
			metric.SetDescription("Number of times fsync duration has exceeded warning threshold.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			metric.IntSum().SetIsMonotonic(true)
			metric.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"zookeeper.latency.avg",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.latency.avg")
			metric.SetDescription("Average time in milliseconds for requests to be processed.")
			metric.SetUnit("ms")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.latency.max",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.latency.max")
			metric.SetDescription("Maximum time in milliseconds for requests to be processed.")
			metric.SetUnit("ms")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.latency.min",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.latency.min")
			metric.SetDescription("Minimum time in milliseconds for requests to be processed.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.max_file_descriptors",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.max_file_descriptors")
			metric.SetDescription("Maximum number of file descriptors that a ZooKeeper server can open.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.open_file_descriptors",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.open_file_descriptors")
			metric.SetDescription("Number of file descriptors that a ZooKeeper server has open.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.outstanding_requests",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.outstanding_requests")
			metric.SetDescription("Number of currently executing requests.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.packets.received",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.packets.received")
			metric.SetDescription("Number of ZooKeeper packets received by a server.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			metric.IntSum().SetIsMonotonic(true)
			metric.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"zookeeper.packets.sent",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.packets.sent")
			metric.SetDescription("Number of ZooKeeper packets sent by a server.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			metric.IntSum().SetIsMonotonic(true)
			metric.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"zookeeper.pending_syncs",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.pending_syncs")
			metric.SetDescription("The number of pending syncs from the followers. Only exposed by the leader.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.synced_followers",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.synced_followers")
			metric.SetDescription("The number of followers in sync with the leader. Only exposed by the leader.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.watches",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.watches")
			metric.SetDescription("Number of watches placed on Z-Nodes on a ZooKeeper server.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
	&metricImpl{
		"zookeeper.znodes",
		func(metric pdata.Metric) {
			metric.SetName("zookeeper.znodes")
			metric.SetDescription("Number of z-nodes that a ZooKeeper server has in its data tree.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// ServerState (State of the Zookeeper server (leader, standalone or follower).)
	ServerState string
	// ZkVersion (Zookeeper version of the instance.)
	ZkVersion string
}{
	"server.state",
	"zk.version",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels
