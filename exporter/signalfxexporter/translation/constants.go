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
	// config.TranslationRules not specified explicitly.
	// Keep it in YAML format to be able to easily copy and paste it in config if modifications needed.
	DefaultTranslationRulesYaml = `
translation_rules:

- action: rename_dimension_keys
  mapping:
    k8s.cronjob.name: kubernetes_name
    k8s.cronjob.uid: kubernetes_uid
    k8s.daemonset.name: kubernetes_name
    k8s.daemonset.uid: kubernetes_uid
    k8s.deployment.name: kubernetes_name
    k8s.deployment.uid: kubernetes_uid
    k8s.hpa.name: kubernetes_name
    k8s.hpa.uid: kubernetes_uid
    k8s.job.name: kubernetes_name
    k8s.job.uid: kubernetes_uid
    k8s.replicaset.name: kubernetes_name
    k8s.replicaset.uid: kubernetes_uid
    k8s.replicationcontroller.name: kubernetes_name
    k8s.replicationcontroller.uid: kubernetes_uid
    k8s.resourcequota.uid: kubernetes_uid
    k8s.statefulset.name: kubernetes_name
    k8s.statefulset.uid: kubernetes_uid

- action: rename_metrics
  mapping:
    # kubeletstats container cpu needed for calculation below
    container.cpu.time: container_cpu_utilization

# compute cpu utilization metrics: cpu.utilization_per_core (excluded by default) and cpu.utilization
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
- action: copy_metrics
  mapping:
    system.cpu.delta: system.cpu.total
- action: aggregate_metric
  metric_name: system.cpu.total
  aggregation_method: sum
  without_dimensions:
  - state
- action: calculate_new_metric
  metric_name: cpu.utilization_per_core
  operand1_metric: system.cpu.usage
  operand2_metric: system.cpu.total
  operator: /
- action: copy_metrics
  mapping:
    cpu.utilization_per_core: cpu.utilization
- action: aggregate_metric
  metric_name: cpu.utilization
  aggregation_method: avg
  without_dimensions:
  - cpu

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

- action: aggregate_metric
  metric_name: cpu.idle
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.interrupt
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.system
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.user
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.steal
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.wait
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.softirq
  aggregation_method: sum
  without_dimensions:
  - cpu
- action: aggregate_metric
  metric_name: cpu.nice
  aggregation_method: sum
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
- action: copy_metrics
  mapping:
    system.memory.usage: system.memory.usage.copy

# memory.used needed to calculate memory.utilization
- action: split_metric
  metric_name: system.memory.usage.copy
  dimension_key: state
  mapping:
    used: memory.used

# Translations to derive filesystem metrics
## disk.total, required to compute disk.utilization
- action: copy_metrics
  mapping:
    system.filesystem.usage: disk.total
- action: aggregate_metric
  metric_name: disk.total
  aggregation_method: sum
  without_dimensions:
    - state

## disk.summary_total, required to compute disk.summary_utilization
- action: copy_metrics
  mapping:
    system.filesystem.usage: disk.summary_total
- action: aggregate_metric
  metric_name: disk.summary_total
  aggregation_method: avg
  without_dimensions:
    - mode
    - mountpoint
- action: aggregate_metric
  metric_name: disk.summary_total
  aggregation_method: sum
  without_dimensions:
    - state
    - device
    - type

## df_complex.used needed to calculate disk.utilization
- action: copy_metrics
  mapping:
    system.filesystem.usage: system.filesystem.usage.copy

- action: split_metric
  metric_name: system.filesystem.usage.copy
  dimension_key: state
  mapping:
    used: df_complex.used

## disk.utilization
- action: calculate_new_metric
  metric_name: disk.utilization
  operand1_metric: df_complex.used
  operand2_metric: disk.total
  operator: /
- action: multiply_float
  scale_factors_float:
    disk.utilization: 100


## disk.summary_utilization
- action: copy_metrics
  mapping:
    df_complex.used: df_complex.used_total

- action: aggregate_metric
  metric_name: df_complex.used_total
  aggregation_method: avg
  without_dimensions:
    - mode
    - mountpoint

- action: aggregate_metric
  metric_name: df_complex.used_total
  aggregation_method: sum
  without_dimensions:
  - device
  - type

- action: calculate_new_metric
  metric_name: disk.summary_utilization
  operand1_metric: df_complex.used_total
  operand2_metric: disk.summary_total
  operator: /
- action: multiply_float
  scale_factors_float:
    disk.summary_utilization: 100


# Translations to derive disk I/O metrics.

## Calculate extra system.disk.operations.total and system.disk.io.total metrics summing up read/write ops/IO across all devices.
- action: copy_metrics
  mapping:
    system.disk.operations: system.disk.operations.total
    system.disk.io: system.disk.io.total
- action: aggregate_metric
  metric_name: system.disk.operations.total
  aggregation_method: sum
  without_dimensions:
    - device
- action: aggregate_metric
  metric_name: system.disk.io.total
  aggregation_method: sum
  without_dimensions:
    - device

## Calculate an extra disk_ops.total metric as number all all read and write operations happened since the last report.
- action: copy_metrics
  mapping:
    system.disk.operations: disk.ops
- action: aggregate_metric
  metric_name: disk.ops
  aggregation_method: sum
  without_dimensions:
    - direction
    - device
- action: delta_metric
  mapping:
    disk.ops: disk_ops.total

## TODO remove after MEAT-2298
- action: rename_dimension_keys
  metric_names:
    system.disk.merged: true
    system.disk.io: true
    system.disk.operations: true
    system.disk.time: true
  mapping:
    device: disk
- action: delta_metric
  mapping:
    system.disk.pending_operations: disk_ops.pending

# Translations to derive Network I/O metrics.

## Calculate extra network I/O metrics system.network.packets.total and system.network.io.total.
- action: copy_metrics
  mapping:
    system.network.packets: system.network.packets.total
    system.network.io: system.network.io.total
- action: aggregate_metric
  metric_name: system.network.packets.total
  aggregation_method: sum
  without_dimensions:
  - device
- action: aggregate_metric
  metric_name: system.network.io.total
  aggregation_method: sum
  without_dimensions:
  - device

## Calculate extra network.total metric.
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
  - device

## TODO remove after MEAT-2298
## Rename dimension device to interface.
- action: rename_dimension_keys
  metric_names:
    system.network.dropped: true
    system.network.errors: true
    system.network.io: true
    system.network.packets: true
  mapping:
    device: interface

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

# Virtual memory metrics
- action: split_metric
  metric_name: system.paging.operations
  dimension_key: direction
  mapping:
    page_in: system.paging.operations.page_in
    page_out: system.paging.operations.page_out

- action: split_metric
  metric_name: system.paging.operations.page_in
  dimension_key: type
  mapping:
    major: vmpage_io.swap.in
    minor: vmpage_io.memory.in

- action: split_metric
  metric_name: system.paging.operations.page_out
  dimension_key: type
  mapping:
    major: vmpage_io.swap.out
    minor: vmpage_io.memory.out

# remove redundant metrics
- action: drop_metrics
  metric_names:
    system.filesystem.usage.copy: true
    df_complex.used: true
    df_complex.used_total: true
    disk.ops: true
    disk.summary_total: true
    disk.total: true
    system.cpu.usage: true
    system.cpu.total: true
    system.cpu.delta: true
    system.paging.operations.page_in: true
    system.paging.operations.page_out: true
    system.memory.usage.copy: true
    memory.used: true
`
)
