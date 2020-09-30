// Copyright 2019, OpenTelemetry Authors
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

package translation

const (
	// DefaultTranslationRulesYaml defines default translation rules that will be applied to metrics if
	// config.SendCompatibleMetrics set to true and config.TranslationRules not specified explicitly.
	// Keep it in YAML format to be able to easily copy and paste it in config if modifications needed.
	DefaultTranslationRulesYaml = `
translation_rules:

- action: rename_dimension_keys
  mapping:

    # dimensions
    container.image.name: container_image
    k8s.container.name: container_spec_name
    k8s.cluster.name: kubernetes_cluster
    k8s.daemonset.name: kubernetes_name
    k8s.daemonset.uid: kubernetes_uid
    k8s.deployment.name: kubernetes_name
    k8s.deployment.uid: kubernetes_uid
    k8s.hpa.name: kubernetes_name
    k8s.hpa.uid: kubernetes_uid
    k8s.namespace.name: kubernetes_namespace
    k8s.node.name: kubernetes_node
    k8s.node.uid: kubernetes_node_uid
    k8s.pod.name: kubernetes_pod_name
    k8s.pod.uid: kubernetes_pod_uid
    k8s.replicaset.name: kubernetes_name
    k8s.replicaset.uid: kubernetes_uid
    k8s.replicationcontroller.name: kubernetes_name
    k8s.replicationcontroller.uid: kubernetes_uid
    k8s.resourcequota.name: quota_name
    k8s.resourcequota.uid: kubernetes_uid
    k8s.statefulset.name: kubernetes_name
    k8s.statefulset.uid: kubernetes_uid
    host.name: host

    # properties
    cronjob_uid: cronJob_uid
    cronjob: cronJob
    daemonset_uid: daemonSet_uid
    daemonset: daemonSet
    k8s.workload.kind: kubernetes_workload
    k8s.workload.name: kubernetes_workload_name
    replicaset_uid: replicaSet_uid
    replicaset: replicaSet
    statefulset_uid: statefulSet_uid
    statefulset: statefulSet

- action: rename_metrics
  mapping:

    # kubeletstats receiver metrics
    container.cpu.time: container_cpu_utilization
    container.filesystem.available: container_fs_available_bytes
    container.filesystem.capacity: container_fs_capacity_bytes
    container.filesystem.usage: container_fs_usage_bytes
    container.memory.available: container_memory_available_bytes
    container.memory.major_page_faults: container_memory_major_page_faults
    container.memory.page_faults: container_memory_page_faults
    container.memory.rss: container_memory_rss_bytes
    container.memory.usage: container_memory_usage_bytes
    container.memory.working_set: container_memory_working_set_bytes

    # k8s cluster receiver metrics
    k8s.container.cpu_limit: kubernetes.container_cpu_limit
    k8s.container.cpu_request: kubernetes.container_cpu_request
    k8s.container.ephemeral-storage_limit: kubernetes.container_ephemeral_storage_limit
    k8s.container.ephemeral-storage_request: kubernetes.container_ephemeral_storage_request
    k8s.container.memory_limit: kubernetes.container_memory_limit
    k8s.container.memory_request: kubernetes.container_memory_request
    k8s.container.ready: kubernetes.container_ready
    k8s.container.restarts: kubernetes.container_restart_count
    k8s.cronjob.active_jobs: kubernetes.cronjob.active
    k8s.daemonset.current_scheduled_nodes: kubernetes.daemon_set.current_scheduled
    k8s.daemonset.desired_scheduled_nodes: kubernetes.daemon_set.desired_scheduled
    k8s.daemonset.misscheduled_nodes: kubernetes.daemon_set.misscheduled
    k8s.daemonset.ready_nodes: kubernetes.daemon_set.ready
    k8s.deployment.available: kubernetes.deployment.available
    k8s.deployment.desired: kubernetes.deployment.desired
    k8s.job.active_pods: kubernetes.job.active
    k8s.job.desired_successful_pods: kubernetes.job.completions
    k8s.job.failed_pods: kubernetes.job.failed
    k8s.job.max_parallel_pods: kubernetes.job.parallelism
    k8s.job.successful_pods: kubernetes.job.succeeded
    k8s.hpa.current_replicas: kubernetes.hpa.status.current_replicas
    k8s.hpa.desired_replicas: kubernetes.hpa.status.desired_replicas
    k8s.hpa.max_replicas: kubernetes.hpa.spec.max_replicas
    k8s.hpa.min_replicas: kubernetes.hpa.spec.min_replicas
    k8s.namespace.phase: kubernetes.namespace_phase
    k8s.node.condition_memory_pressure: kubernetes.node_memory_pressure
    k8s.node.condition_network_unavailable: kubernetes.node_network_unavailable
    k8s.node.condition_out_of_disk: kubernetes.node_out_of_disk
    k8s.node.condition_p_i_d_pressure: kubernetes.node_p_i_d_pressure
    k8s.node.condition_ready: kubernetes.node_ready
    k8s.pod.phase: kubernetes.pod_phase
    k8s.replicaset.available: kubernetes.replica_set.available
    k8s.replicaset.desired: kubernetes.replica_set.desired
    k8s.replication_controller.available: kubernetes.replication_controller.available
    k8s.replication_controller.desired: kubernetes.replication_controller.desired
    k8s.resource_quota.hard_limit: kubernetes.resource_quota_hard
    k8s.resource_quota.used: kubernetes.resource_quota_used
    k8s.statefulset.current_pods: kubernetes.stateful_set.current
    k8s.statefulset.desired_pods: kubernetes.stateful_set.desired
    k8s.statefulset.ready_pods: kubernetes.stateful_set.ready
    k8s.statefulset.updated_pods: kubernetes.stateful_set.updated

    # load metrics
    system.cpu.load_average.15m: load.longterm
    system.cpu.load_average.5m: load.midterm
    system.cpu.load_average.1m: load.shortterm

# container network metrics
- action: split_metric
  metric_name: k8s.pod.network.io
  dimension_key: direction
  mapping: 
    receive: pod_network_receive_bytes_total
    transmit: pod_network_transmit_bytes_total
- action: split_metric
  metric_name: k8s.pod.network.errors
  dimension_key: direction
  mapping: 
    receive: pod_network_receive_errors_total
    transmit: pod_network_transmit_errors_total

# compute cpu utilization
- action: delta_metric
  mapping:
    system.cpu.time: system.cpu.delta
- action: copy_metrics
  mapping:
    system.cpu.delta: system.cpu.usage
  dimension_key: state
  dimension_values:
    interrupt: true
    nice: true
    softirq: true
    steal: true
    system: true
    user: true
    wait: true
- action: aggregate_metric
  metric_name: system.cpu.usage
  aggregation_method: sum
  without_dimensions:
  - state
  - cpu
- action: copy_metrics
  mapping:
    system.cpu.delta: system.cpu.total
- action: aggregate_metric
  metric_name: system.cpu.total
  aggregation_method: sum
  without_dimensions:
  - state
  - cpu
- action: calculate_new_metric
  metric_name: cpu.utilization
  operand1_metric: system.cpu.usage
  operand2_metric: system.cpu.total
  operator: /

# convert cpu metrics
- action: split_metric
  metric_name: system.cpu.time
  dimension_key: state
  mapping:
    idle: cpu.idle
    interrupt: cpu.interrupt
    system: cpu.system
    user: cpu.user
    steal: cpu.steal
    wait: cpu.wait
    softirq: cpu.softirq
    nice: cpu.nice
- action: multiply_float
  scale_factors_float:
    container_cpu_utilization: 100
    cpu.idle: 100
    cpu.interrupt: 100
    cpu.system: 100
    cpu.user: 100
    cpu.steal: 100
    cpu.wait: 100
    cpu.softirq: 100
    cpu.nice: 100
- action: convert_values
  types_mapping:
    container_cpu_utilization: int
    cpu.idle: int
    cpu.interrupt: int
    cpu.system: int
    cpu.user: int
    cpu.steal: int
    cpu.wait: int
    cpu.softirq: int
    cpu.nice: int

# compute cpu.num_processors
- action: copy_metrics
  mapping:
    cpu.idle: cpu.num_processors
- action: aggregate_metric
  metric_name: cpu.num_processors
  aggregation_method: count
  without_dimensions: 
  - cpu

# compute memory.total
- action: copy_metrics
  mapping:
    system.memory.usage: memory.total
  dimension_key: state
  dimension_values:
    buffered: true
    cached: true
    free: true
    used: true
- action: aggregate_metric
  metric_name: memory.total
  aggregation_method: sum
  without_dimensions: 
  - state

# convert memory metrics
- action: split_metric
  metric_name: system.memory.usage
  dimension_key: state
  mapping:
    buffered: memory.buffered
    cached: memory.cached
    free: memory.free
    slab_reclaimable: memory.slab_recl
    slab_unreclaimable: memory.slab_unrecl
    used: memory.used

# calculate disk.total
- action: copy_metrics
  mapping:
    system.filesystem.usage: disk.total
- action: aggregate_metric
  metric_name: disk.total
  aggregation_method: sum
  without_dimensions:
    - state

# calculate disk.summary_total
- action: copy_metrics
  mapping:
    system.filesystem.usage: disk.summary_total
- action: aggregate_metric
  metric_name: disk.summary_total
  aggregation_method: sum
  without_dimensions:
    - state
    - device

# convert filesystem metrics
- action: split_metric
  metric_name: system.filesystem.usage
  dimension_key: state
  mapping:
    free: df_complex.free
    reserved: df_complex.reserved
    used: df_complex.used
- action: split_metric
  metric_name: system.filesystem.inodes.usage
  dimension_key: state
  mapping:
    free: df_inodes.free
    used: df_inodes.used

# df_complex.used_total
- action: copy_metrics
  mapping:
    df_complex.used: df_complex.used_total 
- action: aggregate_metric
  metric_name: df_complex.used_total
  aggregation_method: sum
  without_dimensions:
  - device

# disk utilization
- action: calculate_new_metric
  metric_name: disk.utilization
  operand1_metric: df_complex.used
  operand2_metric: disk.total
  operator: /
- action: multiply_float
  scale_factors_float:
    disk.utilization: 100

- action: calculate_new_metric
  metric_name: disk.summary_utilization
  operand1_metric: df_complex.used_total
  operand2_metric: disk.summary_total
  operator: /
- action: multiply_float
  scale_factors_float:
    disk.summary_utilization: 100

# convert disk I/O metrics
- action: copy_metrics
  mapping:
    system.disk.ops: disk.ops
- action: aggregate_metric
  metric_name: disk.ops
  aggregation_method: sum
  without_dimensions:
   - direction
   - device
- action: delta_metric
  mapping:
    disk.ops: disk_ops.total
- action: rename_dimension_keys
  metric_names:
    system.disk.merged: true
    system.disk.io: true
    system.disk.ops: true
    system.disk.time: true
  mapping:
    device: disk
- action: split_metric
  metric_name: system.disk.merged
  dimension_key: direction
  mapping:
    read: disk_merged.read
    write: disk_merged.write
- action: split_metric
  metric_name: system.disk.io
  dimension_key: direction
  mapping:
    read: disk_octets.read
    write: disk_octets.write
- action: split_metric
  metric_name: system.disk.ops
  dimension_key: direction
  mapping:
    read: disk_ops.read
    write: disk_ops.write
- action: split_metric
  metric_name: system.disk.time
  dimension_key: direction
  mapping:
    read: disk_time.read
    write: disk_time.write
- action: delta_metric
  mapping:
    system.disk.pending_operations: disk_ops.pending

# convert network I/O metrics
- action: copy_metrics
  mapping:
    system.network.io: network.total
  dimension_key: direction
  dimension_values:
    receive: true
    transmit: true
- action: aggregate_metric
  metric_name: network.total
  aggregation_method: sum
  without_dimensions: 
  - direction
  - interface
- action: split_metric
  metric_name: system.network.dropped_packets
  dimension_key: direction
  mapping:
    receive: if_dropped.rx
    transmit: if_dropped.tx
- action: split_metric
  metric_name: system.network.errors
  dimension_key: direction
  mapping:
    receive: if_errors.rx
    transmit: if_errors.tx
- action: split_metric
  metric_name: system.network.io
  dimension_key: direction
  mapping:
    receive: if_octets.rx
    transmit: if_octets.tx
- action: split_metric
  metric_name: system.network.packets
  dimension_key: direction
  mapping:
    receive: if_packets.rx
    transmit: if_packets.tx

# memory utilization
- action: calculate_new_metric
  metric_name: memory.utilization
  operand1_metric: memory.used
  operand2_metric: memory.total
  operator: /

- action: multiply_float
  scale_factors_float:
    memory.utilization: 100
    cpu.utilization: 100

# remove redundant metrics
- action: drop_metrics
  metric_names:
    df_complex.used_total: true
    disk.ops: true
    disk.summary_total: true
    disk.total: true
    system.cpu.usage: true
    system.cpu.total: true
    system.cpu.delta: true
`
)
