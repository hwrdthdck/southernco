// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Type is the component type name.
const Type config.Type = "kubeletstatsreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pmetric.Metric
	Init(metric pmetric.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pmetric.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pmetric.Metric {
	metric := pmetric.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pmetric.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	ContainerCPUTime               MetricIntf
	ContainerCPUUtilization        MetricIntf
	ContainerFilesystemAvailable   MetricIntf
	ContainerFilesystemCapacity    MetricIntf
	ContainerFilesystemUsage       MetricIntf
	ContainerMemoryAvailable       MetricIntf
	ContainerMemoryMajorPageFaults MetricIntf
	ContainerMemoryPageFaults      MetricIntf
	ContainerMemoryRss             MetricIntf
	ContainerMemoryUsage           MetricIntf
	ContainerMemoryWorkingSet      MetricIntf
	K8sNodeCPUTime                 MetricIntf
	K8sNodeCPUUtilization          MetricIntf
	K8sNodeFilesystemAvailable     MetricIntf
	K8sNodeFilesystemCapacity      MetricIntf
	K8sNodeFilesystemUsage         MetricIntf
	K8sNodeMemoryAvailable         MetricIntf
	K8sNodeMemoryMajorPageFaults   MetricIntf
	K8sNodeMemoryPageFaults        MetricIntf
	K8sNodeMemoryRss               MetricIntf
	K8sNodeMemoryUsage             MetricIntf
	K8sNodeMemoryWorkingSet        MetricIntf
	K8sNodeNetworkErrors           MetricIntf
	K8sNodeNetworkIo               MetricIntf
	K8sPodCPUTime                  MetricIntf
	K8sPodCPUUtilization           MetricIntf
	K8sPodFilesystemAvailable      MetricIntf
	K8sPodFilesystemCapacity       MetricIntf
	K8sPodFilesystemUsage          MetricIntf
	K8sPodMemoryAvailable          MetricIntf
	K8sPodMemoryMajorPageFaults    MetricIntf
	K8sPodMemoryPageFaults         MetricIntf
	K8sPodMemoryRss                MetricIntf
	K8sPodMemoryUsage              MetricIntf
	K8sPodMemoryWorkingSet         MetricIntf
	K8sPodNetworkErrors            MetricIntf
	K8sPodNetworkIo                MetricIntf
	K8sVolumeAvailable             MetricIntf
	K8sVolumeCapacity              MetricIntf
	K8sVolumeInodes                MetricIntf
	K8sVolumeInodesFree            MetricIntf
	K8sVolumeInodesUsed            MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"container.cpu.time",
		"container.cpu.utilization",
		"container.filesystem.available",
		"container.filesystem.capacity",
		"container.filesystem.usage",
		"container.memory.available",
		"container.memory.major_page_faults",
		"container.memory.page_faults",
		"container.memory.rss",
		"container.memory.usage",
		"container.memory.working_set",
		"k8s.node.cpu.time",
		"k8s.node.cpu.utilization",
		"k8s.node.filesystem.available",
		"k8s.node.filesystem.capacity",
		"k8s.node.filesystem.usage",
		"k8s.node.memory.available",
		"k8s.node.memory.major_page_faults",
		"k8s.node.memory.page_faults",
		"k8s.node.memory.rss",
		"k8s.node.memory.usage",
		"k8s.node.memory.working_set",
		"k8s.node.network.errors",
		"k8s.node.network.io",
		"k8s.pod.cpu.time",
		"k8s.pod.cpu.utilization",
		"k8s.pod.filesystem.available",
		"k8s.pod.filesystem.capacity",
		"k8s.pod.filesystem.usage",
		"k8s.pod.memory.available",
		"k8s.pod.memory.major_page_faults",
		"k8s.pod.memory.page_faults",
		"k8s.pod.memory.rss",
		"k8s.pod.memory.usage",
		"k8s.pod.memory.working_set",
		"k8s.pod.network.errors",
		"k8s.pod.network.io",
		"k8s.volume.available",
		"k8s.volume.capacity",
		"k8s.volume.inodes",
		"k8s.volume.inodes.free",
		"k8s.volume.inodes.used",
	}
}

var metricsByName = map[string]MetricIntf{
	"container.cpu.time":                 Metrics.ContainerCPUTime,
	"container.cpu.utilization":          Metrics.ContainerCPUUtilization,
	"container.filesystem.available":     Metrics.ContainerFilesystemAvailable,
	"container.filesystem.capacity":      Metrics.ContainerFilesystemCapacity,
	"container.filesystem.usage":         Metrics.ContainerFilesystemUsage,
	"container.memory.available":         Metrics.ContainerMemoryAvailable,
	"container.memory.major_page_faults": Metrics.ContainerMemoryMajorPageFaults,
	"container.memory.page_faults":       Metrics.ContainerMemoryPageFaults,
	"container.memory.rss":               Metrics.ContainerMemoryRss,
	"container.memory.usage":             Metrics.ContainerMemoryUsage,
	"container.memory.working_set":       Metrics.ContainerMemoryWorkingSet,
	"k8s.node.cpu.time":                  Metrics.K8sNodeCPUTime,
	"k8s.node.cpu.utilization":           Metrics.K8sNodeCPUUtilization,
	"k8s.node.filesystem.available":      Metrics.K8sNodeFilesystemAvailable,
	"k8s.node.filesystem.capacity":       Metrics.K8sNodeFilesystemCapacity,
	"k8s.node.filesystem.usage":          Metrics.K8sNodeFilesystemUsage,
	"k8s.node.memory.available":          Metrics.K8sNodeMemoryAvailable,
	"k8s.node.memory.major_page_faults":  Metrics.K8sNodeMemoryMajorPageFaults,
	"k8s.node.memory.page_faults":        Metrics.K8sNodeMemoryPageFaults,
	"k8s.node.memory.rss":                Metrics.K8sNodeMemoryRss,
	"k8s.node.memory.usage":              Metrics.K8sNodeMemoryUsage,
	"k8s.node.memory.working_set":        Metrics.K8sNodeMemoryWorkingSet,
	"k8s.node.network.errors":            Metrics.K8sNodeNetworkErrors,
	"k8s.node.network.io":                Metrics.K8sNodeNetworkIo,
	"k8s.pod.cpu.time":                   Metrics.K8sPodCPUTime,
	"k8s.pod.cpu.utilization":            Metrics.K8sPodCPUUtilization,
	"k8s.pod.filesystem.available":       Metrics.K8sPodFilesystemAvailable,
	"k8s.pod.filesystem.capacity":        Metrics.K8sPodFilesystemCapacity,
	"k8s.pod.filesystem.usage":           Metrics.K8sPodFilesystemUsage,
	"k8s.pod.memory.available":           Metrics.K8sPodMemoryAvailable,
	"k8s.pod.memory.major_page_faults":   Metrics.K8sPodMemoryMajorPageFaults,
	"k8s.pod.memory.page_faults":         Metrics.K8sPodMemoryPageFaults,
	"k8s.pod.memory.rss":                 Metrics.K8sPodMemoryRss,
	"k8s.pod.memory.usage":               Metrics.K8sPodMemoryUsage,
	"k8s.pod.memory.working_set":         Metrics.K8sPodMemoryWorkingSet,
	"k8s.pod.network.errors":             Metrics.K8sPodNetworkErrors,
	"k8s.pod.network.io":                 Metrics.K8sPodNetworkIo,
	"k8s.volume.available":               Metrics.K8sVolumeAvailable,
	"k8s.volume.capacity":                Metrics.K8sVolumeCapacity,
	"k8s.volume.inodes":                  Metrics.K8sVolumeInodes,
	"k8s.volume.inodes.free":             Metrics.K8sVolumeInodesFree,
	"k8s.volume.inodes.used":             Metrics.K8sVolumeInodesUsed,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"container.cpu.time",
		func(metric pmetric.Metric) {
			metric.SetName("container.cpu.time")
			metric.SetDescription("Container CPU time")
			metric.SetUnit("s")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"container.cpu.utilization",
		func(metric pmetric.Metric) {
			metric.SetName("container.cpu.utilization")
			metric.SetDescription("Container CPU utilization")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.filesystem.available",
		func(metric pmetric.Metric) {
			metric.SetName("container.filesystem.available")
			metric.SetDescription("Container filesystem available")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.filesystem.capacity",
		func(metric pmetric.Metric) {
			metric.SetName("container.filesystem.capacity")
			metric.SetDescription("Container filesystem capacity")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.filesystem.usage",
		func(metric pmetric.Metric) {
			metric.SetName("container.filesystem.usage")
			metric.SetDescription("Container filesystem usage")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.memory.available",
		func(metric pmetric.Metric) {
			metric.SetName("container.memory.available")
			metric.SetDescription("Container memory available")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.memory.major_page_faults",
		func(metric pmetric.Metric) {
			metric.SetName("container.memory.major_page_faults")
			metric.SetDescription("Container memory major_page_faults")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.memory.page_faults",
		func(metric pmetric.Metric) {
			metric.SetName("container.memory.page_faults")
			metric.SetDescription("Container memory page_faults")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.memory.rss",
		func(metric pmetric.Metric) {
			metric.SetName("container.memory.rss")
			metric.SetDescription("Container memory rss")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.memory.usage",
		func(metric pmetric.Metric) {
			metric.SetName("container.memory.usage")
			metric.SetDescription("Container memory usage")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"container.memory.working_set",
		func(metric pmetric.Metric) {
			metric.SetName("container.memory.working_set")
			metric.SetDescription("Container memory working_set")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.cpu.time",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.cpu.time")
			metric.SetDescription("Node CPU time")
			metric.SetUnit("s")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"k8s.node.cpu.utilization",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.cpu.utilization")
			metric.SetDescription("Node CPU utilization")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.filesystem.available",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.filesystem.available")
			metric.SetDescription("Node filesystem available")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.filesystem.capacity",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.filesystem.capacity")
			metric.SetDescription("Node filesystem capacity")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.filesystem.usage",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.filesystem.usage")
			metric.SetDescription("Node filesystem usage")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.memory.available",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.memory.available")
			metric.SetDescription("Node memory available")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.memory.major_page_faults",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.memory.major_page_faults")
			metric.SetDescription("Node memory major_page_faults")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.memory.page_faults",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.memory.page_faults")
			metric.SetDescription("Node memory page_faults")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.memory.rss",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.memory.rss")
			metric.SetDescription("Node memory rss")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.memory.usage",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.memory.usage")
			metric.SetDescription("Node memory usage")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.memory.working_set",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.memory.working_set")
			metric.SetDescription("Node memory working_set")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.node.network.errors",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.network.errors")
			metric.SetDescription("Node network errors")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"k8s.node.network.io",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.node.network.io")
			metric.SetDescription("Node network IO")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"k8s.pod.cpu.time",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.cpu.time")
			metric.SetDescription("Pod CPU time")
			metric.SetUnit("s")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"k8s.pod.cpu.utilization",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.cpu.utilization")
			metric.SetDescription("Pod CPU utilization")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.filesystem.available",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.filesystem.available")
			metric.SetDescription("Pod filesystem available")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.filesystem.capacity",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.filesystem.capacity")
			metric.SetDescription("Pod filesystem capacity")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.filesystem.usage",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.filesystem.usage")
			metric.SetDescription("Pod filesystem usage")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.memory.available",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.memory.available")
			metric.SetDescription("Pod memory available")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.memory.major_page_faults",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.memory.major_page_faults")
			metric.SetDescription("Pod memory major_page_faults")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.memory.page_faults",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.memory.page_faults")
			metric.SetDescription("Pod memory page_faults")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.memory.rss",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.memory.rss")
			metric.SetDescription("Pod memory rss")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.memory.usage",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.memory.usage")
			metric.SetDescription("Pod memory usage")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.memory.working_set",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.memory.working_set")
			metric.SetDescription("Pod memory working_set")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.pod.network.errors",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.network.errors")
			metric.SetDescription("Pod network errors")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"k8s.pod.network.io",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.pod.network.io")
			metric.SetDescription("Pod network IO")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"k8s.volume.available",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.volume.available")
			metric.SetDescription("The number of available bytes in the volume.")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.volume.capacity",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.volume.capacity")
			metric.SetDescription("The total capacity in bytes of the volume.")
			metric.SetUnit("By")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.volume.inodes",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.volume.inodes")
			metric.SetDescription("The total inodes in the filesystem.")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.volume.inodes.free",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.volume.inodes.free")
			metric.SetDescription("The free inodes in the filesystem.")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"k8s.volume.inodes.used",
		func(metric pmetric.Metric) {
			metric.SetName("k8s.volume.inodes.used")
			metric.SetDescription("The inodes used by the filesystem. This may not equal inodes - free because filesystem may share inodes with other filesystems.")
			metric.SetUnit("1")
			metric.SetDataType(pmetric.MetricDataTypeGauge)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Direction (Direction of flow of bytes/operations (receive or transmit).)
	Direction string
	// Interface (Name of the network interface.)
	Interface string
}{
	"direction",
	"interface",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	Receive  string
	Transmit string
}{
	"receive",
	"transmit",
}
