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
package containerinsight

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

func TestAggregateFields(t *testing.T) {
	fields := []map[string]interface{}{
		{
			"m1": float64(1),
			"m2": float64(2),
			"m3": float64(3),
		},
		{
			"m1": float64(2),
			"m3": float64(3),
		},
		{
			"m1": float64(2),
			"m2": float64(3),
		},
	}

	expected := map[string]float64{
		"m1": float64(1 + 2 + 2),
		"m2": float64(2 + 3),
		"m3": float64(3 + 3),
	}

	assert.Equal(t, expected, SumFields(fields))

	//test empty input
	assert.Nil(t, SumFields([]map[string]interface{}{}))

	//test single input
	fields = []map[string]interface{}{
		{
			"m1": float64(2),
			"m2": float64(3),
		},
	}
	expected = map[string]float64{
		"m1": float64(2),
		"m2": float64(3),
	}
	assert.Equal(t, expected, SumFields(fields))
}

func TestMetricName(t *testing.T) {
	assert.Equal(t, "instance_cpu_usage_total", MetricName(TypeInstance, "cpu_usage_total"))
	assert.Equal(t, "instance_filesystem_usage", MetricName(TypeInstanceFS, "filesystem_usage"))
	assert.Equal(t, "instance_interface_network_rx_bytes", MetricName(TypeInstanceNet, "network_rx_bytes"))
	assert.Equal(t, "instance_diskio_io_service_bytes_total", MetricName(TypeInstanceDiskIO, "diskio_io_service_bytes_total"))
	assert.Equal(t, "service_number_of_running_pods", MetricName(TypeService, "number_of_running_pods"))
	assert.Equal(t, "namespace_number_of_running_pods", MetricName(TypeClusterNamespace, "number_of_running_pods"))
	assert.Equal(t, "container_diskio_io_service_bytes_total", MetricName(TypeContainerDiskIO, "diskio_io_service_bytes_total"))
	assert.Equal(t, "unknown_metrics", MetricName("unknown_type", "unknown_metrics"))
}

func TestIsNode(t *testing.T) {
	assert.Equal(t, true, IsNode(TypeNode))
	assert.Equal(t, true, IsNode(TypeNodeNet))
	assert.Equal(t, true, IsNode(TypeNodeFS))
	assert.Equal(t, true, IsNode(TypeNodeDiskIO))
	assert.Equal(t, false, IsNode(TypePod))
}

func TestIsInstance(t *testing.T) {
	assert.Equal(t, true, IsInstance(TypeInstance))
	assert.Equal(t, true, IsInstance(TypeInstanceNet))
	assert.Equal(t, true, IsInstance(TypeInstanceFS))
	assert.Equal(t, true, IsInstance(TypeInstanceDiskIO))
	assert.Equal(t, false, IsInstance(TypePod))
}

func TestIsContainer(t *testing.T) {
	assert.Equal(t, true, IsContainer(TypeContainer))
	assert.Equal(t, true, IsContainer(TypeContainerDiskIO))
	assert.Equal(t, true, IsContainer(TypeContainerFS))
	assert.Equal(t, false, IsContainer(TypePod))
}

func TestIsPod(t *testing.T) {
	assert.Equal(t, true, IsPod(TypePod))
	assert.Equal(t, true, IsPod(TypePodNet))
	assert.Equal(t, false, IsPod(TypeInstance))
}

func convertToInt64(value interface{}) int64 {
	switch t := value.(type) {
	case int:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	default:
		valueType := fmt.Sprintf("%T", value)
		log.Printf("Detected unexpected type: %v", valueType)
	}
	return -1
}

func convertToFloat64(value interface{}) float64 {
	switch t := value.(type) {
	case uint:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return t
	default:
		valueType := fmt.Sprintf("%T", value)
		log.Printf("Detected unexpected type: %v", valueType)
	}
	return -1.0
}

func checkMetricsAreExpected(t *testing.T, md pdata.Metrics, fields map[string]interface{}, tags map[string]string,
	expectedUnits map[string]string) {

	rms := md.ResourceMetrics()
	assert.Equal(t, 1, rms.Len())

	//check the attributes are expected
	rm := rms.At(0)
	attributes := rm.Resource().Attributes()
	assert.Equal(t, len(tags), attributes.Len())
	var timeUnixNano uint64
	for key, val := range tags {
		log.Printf("key=%v value=%v", key, val)
		attr, ok := attributes.Get(key)
		assert.Equal(t, true, ok)
		if key == Timestamp {
			timeUnixNano, _ = strconv.ParseUint(val, 10, 64)
			val = strconv.FormatUint(timeUnixNano/uint64(time.Millisecond), 10)
		}
		assert.Equal(t, val, attr.StringVal())
	}

	//check the metrics are expected
	ilms := rm.InstrumentationLibraryMetrics()
	for j := 0; j < ilms.Len(); j++ {
		ilm := ilms.At(j)
		ms := ilm.Metrics()
		for k := 0; k < ms.Len(); k++ {
			m := ms.At(k)
			metricName := m.Name()
			log.Printf("metric=%v", metricName)
			assert.Equal(t, expectedUnits[metricName], m.Unit(), "Wrong unit for metric: "+metricName)
			switch m.DataType() {
			//we only need to worry about gauge types for container insights metrics
			case pdata.MetricDataTypeIntGauge:
				dps := m.IntGauge().DataPoints()
				assert.Equal(t, 1, dps.Len())
				dp := dps.At(0)
				assert.Equal(t, convertToInt64(fields[metricName]), dp.Value())
				assert.Equal(t, pdata.Timestamp(timeUnixNano), dp.Timestamp())
			case pdata.MetricDataTypeDoubleGauge:
				dps := m.DoubleGauge().DataPoints()
				assert.Equal(t, 1, dps.Len())
				dp := dps.At(0)
				assert.Equal(t, convertToFloat64(fields[metricName]), dp.Value())
				assert.Equal(t, pdata.Timestamp(timeUnixNano), dp.Timestamp())
			}
		}
	}
}

func TestConvertToOTLPMetricsForInvalidMetrics(t *testing.T) {
	var fields map[string]interface{}
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"node_cpu_limit": "an invalid value",
	}

	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "ci-demo",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"Type":                 "Node",
		"Version":              "0",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	rm := md.ResourceMetrics().At(0)
	ilms := rm.InstrumentationLibraryMetrics()
	assert.Equal(t, 0, ilms.Len())
}

func TestConvertToOTLPMetricsForClusterMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test cluster-level metrics
	fields = map[string]interface{}{
		"cluster_failed_node_count": int64(1),
		"cluster_node_count":        int64(3),
	}
	expectedUnits = map[string]string{
		"cluster_failed_node_count": UnitCount,
		"cluster_node_count":        UnitCount,
	}
	tags = map[string]string{
		"ClusterName": "test-cluster",
		"Type":        TypeCluster,
		"Timestamp":   timestamp,
		"Version":     "0",
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)

	//test cluster namespace metrics
	fields = map[string]interface{}{
		"namespace_number_of_running_pods": int64(8),
	}
	expectedUnits = map[string]string{
		"namespace_number_of_running_pods": UnitCount,
	}
	tags = map[string]string{
		"ClusterName": "test-cluster",
		"Type":        TypeClusterNamespace,
		"Timestamp":   timestamp,
		"Version":     "0",
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)

	//test cluster service metrics
	fields = map[string]interface{}{
		"service_number_of_running_pods": int64(8),
	}
	expectedUnits = map[string]string{
		"service_number_of_running_pods": UnitCount,
	}
	tags = map[string]string{
		"ClusterName": "test-cluster",
		"Type":        TypeClusterService,
		"Timestamp":   timestamp,
		"Version":     "0",
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)

}

func TestConvertToOTLPMetricsForContainerMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"container_cpu_limit":                      int64(200),
		"container_cpu_request":                    int64(200),
		"container_cpu_usage_system":               float64(2.7662289817161336),
		"container_cpu_usage_total":                float64(9.140206205783091),
		"container_cpu_usage_user":                 float64(2.9638167661244292),
		"container_cpu_utilization":                float64(0.22850515514457728),
		"container_memory_cache":                   int64(3244032),
		"container_memory_failcnt":                 int64(0),
		"container_memory_hierarchical_pgfault":    float64(0),
		"container_memory_hierarchical_pgmajfault": float64(0),
		"container_memory_limit":                   int64(209715200),
		"container_memory_mapped_file":             int64(0),
		"container_memory_max_usage":               int64(44482560),
		"container_memory_pgfault":                 float64(0),
		"container_memory_pgmajfault":              float64(0),
		"container_memory_request":                 int64(209715200),
		"container_memory_rss":                     int64(38686720),
		"container_memory_swap":                    int64(0),
		"container_memory_usage":                   int64(44257280),
		"container_memory_utilization":             float64(0.16909561488772057),
		"container_memory_working_set":             int64(28438528),
	}
	expectedUnits = map[string]string{
		"container_cpu_limit":                      "",
		"container_cpu_request":                    "",
		"container_cpu_usage_system":               "",
		"container_cpu_usage_total":                "",
		"container_cpu_usage_user":                 "",
		"container_cpu_utilization":                UnitPercent,
		"container_memory_cache":                   UnitBytes,
		"container_memory_failcnt":                 UnitCount,
		"container_memory_hierarchical_pgfault":    UnitCountPerSec,
		"container_memory_hierarchical_pgmajfault": UnitCountPerSec,
		"container_memory_limit":                   UnitBytes,
		"container_memory_mapped_file":             UnitBytes,
		"container_memory_max_usage":               UnitBytes,
		"container_memory_pgfault":                 UnitCountPerSec,
		"container_memory_pgmajfault":              UnitCountPerSec,
		"container_memory_request":                 UnitBytes,
		"container_memory_rss":                     UnitBytes,
		"container_memory_swap":                    UnitBytes,
		"container_memory_usage":                   UnitBytes,
		"container_memory_utilization":             UnitPercent,
		"container_memory_working_set":             UnitBytes,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "ci-demo",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"Namespace":            "aws-otel-eks",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"PodName":              "aws-otel-eks-ci",
		"Type":                 "Container",
		"Version":              "0",
		"container_status":     "Running",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)

	//test container filesystem metrics
	fields = map[string]interface{}{
		"container_filesystem_available":   int64(0),
		"container_filesystem_capacity":    int64(21462233088),
		"container_filesystem_usage":       int64(36864),
		"container_filesystem_utilization": float64(0.00017176218266221077),
	}
	expectedUnits = map[string]string{
		"container_filesystem_available":   UnitBytes,
		"container_filesystem_capacity":    UnitBytes,
		"container_filesystem_usage":       UnitBytes,
		"container_filesystem_utilization": UnitPercent,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "ci-demo",
		"EBSVolumeId":          "aws://us-east-1a/vol-09e13990fbaa644fe",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"Namespace":            "amazon-cloudwatch",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"PodName":              "cloudwatch-agent",
		"Type":                 "ContainerFS",
		"Timestamp":            timestamp,
		"Version":              "0",
		"device":               "/dev/xvda1",
		"fstype":               "vfs",
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}

func TestConvertToOTLPMetricsForNodeMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"node_cpu_limit":                      int64(4000),
		"node_cpu_request":                    int64(610),
		"node_cpu_reserved_capacity":          float64(15.25),
		"node_cpu_usage_system":               float64(31.93421003691368),
		"node_cpu_usage_total":                float64(120.37540542432477),
		"node_cpu_usage_user":                 float64(57.347822212886875),
		"node_cpu_utilization":                float64(3.0093851356081194),
		"node_memory_cache":                   int64(3810631680),
		"node_memory_failcnt":                 int64(0),
		"node_memory_hierarchical_pgfault":    float64(0),
		"node_memory_hierarchical_pgmajfault": float64(0),
		"node_memory_limit":                   int64(16818016256),
		"node_memory_mapped_file":             int64(527560704),
		"node_memory_max_usage":               int64(4752089088),
		"node_memory_pgfault":                 float64(0),
		"node_memory_pgmajfault":              float64(0),
		"node_memory_request":                 int64(492830720),
		"node_memory_reserved_capacity":       float64(2.9303736689169724),
		"node_memory_rss":                     int64(410030080),
		"node_memory_swap":                    uint32(0),
		"node_memory_usage":                   int64(4220661760),
		"node_memory_utilization":             float64(7.724233133242133),
		"node_memory_working_set":             int64(1299062784),
		"node_network_rx_bytes":               float64(13032.150482284136),
		"node_network_rx_dropped":             float64(0),
		"node_network_rx_errors":              float32(0),
		"node_network_rx_packets":             float64(39.27406250089541),
		"node_network_total_bytes":            float64(27124.366262458552),
		"node_network_tx_bytes":               float64(14092.215780174418),
		"node_network_tx_dropped":             uint64(0),
		"node_network_tx_errors":              float64(0),
		"node_network_tx_packets":             float64(37.802748111760124),
		"node_number_of_running_containers":   int32(7),
		"node_number_of_running_pods":         int64(7),
	}
	expectedUnits = map[string]string{
		"node_cpu_limit":                      "",
		"node_cpu_request":                    "",
		"node_cpu_reserved_capacity":          UnitPercent,
		"node_cpu_usage_system":               "",
		"node_cpu_usage_total":                "",
		"node_cpu_usage_user":                 "",
		"node_cpu_utilization":                UnitPercent,
		"node_memory_cache":                   UnitBytes,
		"node_memory_failcnt":                 UnitCount,
		"node_memory_hierarchical_pgfault":    UnitCountPerSec,
		"node_memory_hierarchical_pgmajfault": UnitCountPerSec,
		"node_memory_limit":                   UnitBytes,
		"node_memory_mapped_file":             UnitBytes,
		"node_memory_max_usage":               UnitBytes,
		"node_memory_pgfault":                 UnitCountPerSec,
		"node_memory_pgmajfault":              UnitCountPerSec,
		"node_memory_request":                 UnitBytes,
		"node_memory_reserved_capacity":       UnitPercent,
		"node_memory_rss":                     UnitBytes,
		"node_memory_swap":                    UnitBytes,
		"node_memory_usage":                   UnitBytes,
		"node_memory_utilization":             UnitPercent,
		"node_memory_working_set":             UnitBytes,
		"node_network_rx_bytes":               UnitBytesPerSec,
		"node_network_rx_dropped":             UnitCountPerSec,
		"node_network_rx_errors":              UnitCountPerSec,
		"node_network_rx_packets":             UnitCountPerSec,
		"node_network_total_bytes":            UnitBytesPerSec,
		"node_network_tx_bytes":               UnitBytesPerSec,
		"node_network_tx_dropped":             UnitCountPerSec,
		"node_network_tx_errors":              UnitCountPerSec,
		"node_network_tx_packets":             UnitCountPerSec,
		"node_number_of_running_containers":   UnitCount,
		"node_number_of_running_pods":         UnitCount,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "ci-demo",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"Type":                 "Node",
		"Version":              "0",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}

func TestConvertToOTLPMetricsForNodeDiskIOMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"node_diskio_io_service_bytes_async": float64(6704.018980016907),
		"node_diskio_io_service_bytes_read":  float64(0),
		"node_diskio_io_service_bytes_sync":  float64(284.2693560431197),
		"node_diskio_io_service_bytes_total": float64(6988.288336060026),
		"node_diskio_io_service_bytes_write": float64(6988.288336060026),
		"node_diskio_io_serviced_async":      float64(1.326343566607438),
		"node_diskio_io_serviced_read":       float64(0),
		"node_diskio_io_serviced_sync":       float64(0.04626779883514318),
		"node_diskio_io_serviced_total":      float64(1.372611365442581),
		"node_diskio_io_serviced_write":      float64(1.372611365442581),
	}
	expectedUnits = map[string]string{
		"node_diskio_io_service_bytes_async": UnitBytesPerSec,
		"node_diskio_io_service_bytes_read":  UnitBytesPerSec,
		"node_diskio_io_service_bytes_sync":  UnitBytesPerSec,
		"node_diskio_io_service_bytes_total": UnitBytesPerSec,
		"node_diskio_io_service_bytes_write": UnitBytesPerSec,
		"node_diskio_io_serviced_async":      UnitCountPerSec,
		"node_diskio_io_serviced_read":       UnitCountPerSec,
		"node_diskio_io_serviced_sync":       UnitCountPerSec,
		"node_diskio_io_serviced_total":      UnitCountPerSec,
		"node_diskio_io_serviced_write":      UnitCountPerSec,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "eks-aoc",
		"EBSVolumeId":          "aws://us-east-1a/vol-09e13990fbaa644fe",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"Type":                 "NodeDiskIO",
		"Version":              "0",
		"device":               "/dev/xvda",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}

func TestConvertToOTLPMetricsForNodeFSMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"node_filesystem_available":   int64(4271607808),
		"node_filesystem_capacity":    int64(21462233088),
		"node_filesystem_inodes":      int64(8450312),
		"node_filesystem_inodes_free": int64(8345085),
		"node_filesystem_usage":       int64(17190625280),
		"node_filesystem_utilization": float64(80.09709525339025),
	}
	expectedUnits = map[string]string{
		"node_filesystem_available":   UnitBytes,
		"node_filesystem_capacity":    UnitBytes,
		"node_filesystem_inodes":      UnitCount,
		"node_filesystem_inodes_free": UnitCount,
		"node_filesystem_usage":       UnitBytes,
		"node_filesystem_utilization": UnitPercent,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "eks-aoc",
		"EBSVolumeId":          "aws://us-east-1a/vol-09e13990fbaa644fe",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"Type":                 "NodeFS",
		"Version":              "0",
		"device":               "/dev/xvda",
		"fstype":               "vfs",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}

func TestConvertToOTLPMetricsForNodeNetMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"node_interface_network_rx_bytes":    float64(294.8620421098953),
		"node_interface_network_rx_dropped":  float64(0),
		"node_interface_network_rx_errors":   float64(0),
		"node_interface_network_rx_packets":  float64(2.69744105680903),
		"node_interface_network_total_bytes": float64(1169.5469730310588),
		"node_interface_network_tx_bytes":    float64(874.6849309211634),
		"node_interface_network_tx_dropped":  float64(0),
		"node_interface_network_tx_errors":   float64(0),
		"node_interface_network_tx_packets":  float64(2.713308357143201),
	}
	expectedUnits = map[string]string{
		"node_interface_network_rx_bytes":    UnitBytesPerSec,
		"node_interface_network_rx_dropped":  UnitCountPerSec,
		"node_interface_network_rx_errors":   UnitCountPerSec,
		"node_interface_network_rx_packets":  UnitCountPerSec,
		"node_interface_network_total_bytes": UnitBytesPerSec,
		"node_interface_network_tx_bytes":    UnitBytesPerSec,
		"node_interface_network_tx_dropped":  UnitCountPerSec,
		"node_interface_network_tx_errors":   UnitCountPerSec,
		"node_interface_network_tx_packets":  UnitCountPerSec,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "eks-aoc",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"Type":                 "NodeNet",
		"Version":              "0",
		"interface":            "eni7cce1b61ea4",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}

func TestConvertToOTLPMetricsForPodMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	fields = map[string]interface{}{
		"pod_cpu_limit":                         int64(200),
		"pod_cpu_request":                       int64(200),
		"pod_cpu_reserved_capacity":             float64(5),
		"pod_cpu_usage_system":                  float64(1.2750419580493375),
		"pod_cpu_usage_total":                   float64(5.487638329191601),
		"pod_cpu_usage_user":                    float64(1.8214885114990538),
		"pod_cpu_utilization":                   float64(0.13719095822979002),
		"pod_cpu_utilization_over_pod_limit":    float64(2.7438191645958003),
		"pod_memory_cache":                      int64(811008),
		"pod_memory_failcnt":                    int64(0),
		"pod_memory_hierarchical_pgfault":       float64(0),
		"pod_memory_hierarchical_pgmajfault":    float64(0),
		"pod_memory_limit":                      int64(209715200),
		"pod_memory_mapped_file":                int64(135168),
		"pod_memory_max_usage":                  int64(37670912),
		"pod_memory_pgfault":                    float64(0),
		"pod_memory_pgmajfault":                 float64(0),
		"pod_memory_request":                    int64(209715200),
		"pod_memory_reserved_capacity":          float64(1.2469675186880733),
		"pod_memory_rss":                        int64(32845824),
		"pod_memory_swap":                       int64(0),
		"pod_memory_usage":                      int64(37543936),
		"pod_memory_utilization":                float64(0.1477851348320162),
		"pod_memory_utilization_over_pod_limit": float64(11.8515625),
		"pod_memory_working_set":                int64(24854528),
		"pod_network_rx_bytes":                  float64(3364.5791007424054),
		"pod_network_rx_dropped":                float64(0),
		"pod_network_rx_errors":                 float64(0),
		"pod_network_rx_packets":                float64(2.3677681271483983),
		"pod_network_total_bytes":               float64(3777.9099024135485),
		"pod_network_tx_bytes":                  float64(413.3308016711429),
		"pod_network_tx_dropped":                float64(0),
		"pod_network_tx_errors":                 float64(0),
		"pod_network_tx_packets":                float64(2.3677681271483983),
		"pod_number_of_container_restarts":      int(0),
		"pod_number_of_containers":              uint(1),
		"pod_number_of_running_containers":      uint(1),
	}
	expectedUnits = map[string]string{
		"pod_cpu_limit":                         "",
		"pod_cpu_request":                       "",
		"pod_cpu_reserved_capacity":             UnitPercent,
		"pod_cpu_usage_system":                  "",
		"pod_cpu_usage_total":                   "",
		"pod_cpu_usage_user":                    "",
		"pod_cpu_utilization":                   UnitPercent,
		"pod_cpu_utilization_over_pod_limit":    UnitPercent,
		"pod_memory_cache":                      UnitBytes,
		"pod_memory_failcnt":                    UnitCount,
		"pod_memory_hierarchical_pgfault":       UnitCountPerSec,
		"pod_memory_hierarchical_pgmajfault":    UnitCountPerSec,
		"pod_memory_limit":                      UnitBytes,
		"pod_memory_mapped_file":                UnitBytes,
		"pod_memory_max_usage":                  UnitBytes,
		"pod_memory_pgfault":                    UnitCountPerSec,
		"pod_memory_pgmajfault":                 UnitCountPerSec,
		"pod_memory_request":                    UnitBytes,
		"pod_memory_reserved_capacity":          UnitPercent,
		"pod_memory_rss":                        UnitBytes,
		"pod_memory_swap":                       UnitBytes,
		"pod_memory_usage":                      UnitBytes,
		"pod_memory_utilization":                UnitPercent,
		"pod_memory_utilization_over_pod_limit": UnitPercent,
		"pod_memory_working_set":                UnitBytes,
		"pod_network_rx_bytes":                  UnitBytesPerSec,
		"pod_network_rx_dropped":                UnitCountPerSec,
		"pod_network_rx_errors":                 UnitCountPerSec,
		"pod_network_rx_packets":                UnitCountPerSec,
		"pod_network_total_bytes":               UnitBytesPerSec,
		"pod_network_tx_bytes":                  UnitBytesPerSec,
		"pod_network_tx_dropped":                UnitCountPerSec,
		"pod_network_tx_errors":                 UnitCountPerSec,
		"pod_network_tx_packets":                UnitCountPerSec,
		"pod_number_of_container_restarts":      UnitCount,
		"pod_number_of_containers":              UnitCount,
		"pod_number_of_running_containers":      UnitCount,
	}
	tags = map[string]string{
		"ClusterName":  "eks-aoc",
		"InstanceId":   "i-01bf9fb097cbf3205",
		"InstanceType": "t2.xlarge",
		"Namespace":    "amazon-cloudwatch",
		"NodeName":     "ip-192-168-12-170.ec2.internal",
		"PodName":      "cloudwatch-agent",
		"Type":         "Pod",
		"Version":      "0",
		"Timestamp":    timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}

func TestConvertToOTLPMetricsForPodNetMetrics(t *testing.T) {
	var fields map[string]interface{}
	var expectedUnits map[string]string
	var tags map[string]string
	var md pdata.Metrics
	now := time.Now()
	timestamp := strconv.FormatInt(now.UnixNano(), 10)

	//test container metrics
	fields = map[string]interface{}{
		"node_interface_network_rx_bytes":    float64(294.8620421098953),
		"node_interface_network_rx_dropped":  float64(0),
		"node_interface_network_rx_errors":   float64(0),
		"node_interface_network_rx_packets":  float64(2.69744105680903),
		"node_interface_network_total_bytes": float64(1169.5469730310588),
		"node_interface_network_tx_bytes":    float64(874.6849309211634),
		"node_interface_network_tx_dropped":  float64(0),
		"node_interface_network_tx_errors":   float64(0),
		"node_interface_network_tx_packets":  float64(2.713308357143201),
	}
	expectedUnits = map[string]string{
		"pod_interface_network_rx_bytes":    UnitBytesPerSec,
		"pod_interface_network_rx_dropped":  UnitCountPerSec,
		"pod_interface_network_rx_errors":   UnitCountPerSec,
		"pod_interface_network_rx_packets":  UnitCountPerSec,
		"pod_interface_network_total_bytes": UnitBytesPerSec,
		"pod_interface_network_tx_bytes":    UnitBytesPerSec,
		"pod_interface_network_tx_dropped":  UnitCountPerSec,
		"pod_interface_network_tx_errors":   UnitCountPerSec,
		"pod_interface_network_tx_packets":  UnitCountPerSec,
	}
	tags = map[string]string{
		"AutoScalingGroupName": "eks-a6bb9db9-267c-401c-db55-df8ef645b06f",
		"ClusterName":          "eks-aoc",
		"InstanceId":           "i-01bf9fb097cbf3205",
		"InstanceType":         "t2.xlarge",
		"Namespace":            "default",
		"NodeName":             "ip-192-168-12-170.ec2.internal",
		"PodName":              "aws-otel-eks-ci",
		"Type":                 "PodNet",
		"Version":              "0",
		"interface":            "eth0",
		"Timestamp":            timestamp,
	}
	md = ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	checkMetricsAreExpected(t, md, fields, tags, expectedUnits)
}
