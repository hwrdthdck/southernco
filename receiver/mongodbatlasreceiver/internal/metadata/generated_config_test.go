// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesSettings
	}{
		{
			name: "default",
			want: DefaultResourceAttributesSettings(),
		},
		{
			name: "all_set",
			want: ResourceAttributesSettings{
				MongodbAtlasClusterName:     ResourceAttributeSettings{Enabled: true},
				MongodbAtlasDbName:          ResourceAttributeSettings{Enabled: true},
				MongodbAtlasDiskPartition:   ResourceAttributeSettings{Enabled: true},
				MongodbAtlasHostName:        ResourceAttributeSettings{Enabled: true},
				MongodbAtlasOrgName:         ResourceAttributeSettings{Enabled: true},
				MongodbAtlasProcessID:       ResourceAttributeSettings{Enabled: true},
				MongodbAtlasProcessPort:     ResourceAttributeSettings{Enabled: true},
				MongodbAtlasProcessTypeName: ResourceAttributeSettings{Enabled: true},
				MongodbAtlasProjectID:       ResourceAttributeSettings{Enabled: true},
				MongodbAtlasProjectName:     ResourceAttributeSettings{Enabled: true},
				MongodbAtlasUserAlias:       ResourceAttributeSettings{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesSettings{
				MongodbAtlasClusterName:     ResourceAttributeSettings{Enabled: false},
				MongodbAtlasDbName:          ResourceAttributeSettings{Enabled: false},
				MongodbAtlasDiskPartition:   ResourceAttributeSettings{Enabled: false},
				MongodbAtlasHostName:        ResourceAttributeSettings{Enabled: false},
				MongodbAtlasOrgName:         ResourceAttributeSettings{Enabled: false},
				MongodbAtlasProcessID:       ResourceAttributeSettings{Enabled: false},
				MongodbAtlasProcessPort:     ResourceAttributeSettings{Enabled: false},
				MongodbAtlasProcessTypeName: ResourceAttributeSettings{Enabled: false},
				MongodbAtlasProjectID:       ResourceAttributeSettings{Enabled: false},
				MongodbAtlasProjectName:     ResourceAttributeSettings{Enabled: false},
				MongodbAtlasUserAlias:       ResourceAttributeSettings{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesSettings(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeSettings{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsSettings{
					MongodbatlasDbCounts:                                  MetricSettings{Enabled: true},
					MongodbatlasDbSize:                                    MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionIopsAverage:                  MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionIopsMax:                      MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionLatencyAverage:               MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionLatencyMax:                   MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionSpaceAverage:                 MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionSpaceMax:                     MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionUsageAverage:                 MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionUsageMax:                     MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionUtilizationAverage:           MetricSettings{Enabled: true},
					MongodbatlasDiskPartitionUtilizationMax:               MetricSettings{Enabled: true},
					MongodbatlasProcessAsserts:                            MetricSettings{Enabled: true},
					MongodbatlasProcessBackgroundFlush:                    MetricSettings{Enabled: true},
					MongodbatlasProcessCacheIo:                            MetricSettings{Enabled: true},
					MongodbatlasProcessCacheSize:                          MetricSettings{Enabled: true},
					MongodbatlasProcessConnections:                        MetricSettings{Enabled: true},
					MongodbatlasProcessCPUChildrenNormalizedUsageAverage:  MetricSettings{Enabled: true},
					MongodbatlasProcessCPUChildrenNormalizedUsageMax:      MetricSettings{Enabled: true},
					MongodbatlasProcessCPUChildrenUsageAverage:            MetricSettings{Enabled: true},
					MongodbatlasProcessCPUChildrenUsageMax:                MetricSettings{Enabled: true},
					MongodbatlasProcessCPUNormalizedUsageAverage:          MetricSettings{Enabled: true},
					MongodbatlasProcessCPUNormalizedUsageMax:              MetricSettings{Enabled: true},
					MongodbatlasProcessCPUUsageAverage:                    MetricSettings{Enabled: true},
					MongodbatlasProcessCPUUsageMax:                        MetricSettings{Enabled: true},
					MongodbatlasProcessCursors:                            MetricSettings{Enabled: true},
					MongodbatlasProcessDbDocumentRate:                     MetricSettings{Enabled: true},
					MongodbatlasProcessDbOperationsRate:                   MetricSettings{Enabled: true},
					MongodbatlasProcessDbOperationsTime:                   MetricSettings{Enabled: true},
					MongodbatlasProcessDbQueryExecutorScanned:             MetricSettings{Enabled: true},
					MongodbatlasProcessDbQueryTargetingScannedPerReturned: MetricSettings{Enabled: true},
					MongodbatlasProcessDbStorage:                          MetricSettings{Enabled: true},
					MongodbatlasProcessGlobalLock:                         MetricSettings{Enabled: true},
					MongodbatlasProcessIndexBtreeMissRatio:                MetricSettings{Enabled: true},
					MongodbatlasProcessIndexCounters:                      MetricSettings{Enabled: true},
					MongodbatlasProcessJournalingCommits:                  MetricSettings{Enabled: true},
					MongodbatlasProcessJournalingDataFiles:                MetricSettings{Enabled: true},
					MongodbatlasProcessJournalingWritten:                  MetricSettings{Enabled: true},
					MongodbatlasProcessMemoryUsage:                        MetricSettings{Enabled: true},
					MongodbatlasProcessNetworkIo:                          MetricSettings{Enabled: true},
					MongodbatlasProcessNetworkRequests:                    MetricSettings{Enabled: true},
					MongodbatlasProcessOplogRate:                          MetricSettings{Enabled: true},
					MongodbatlasProcessOplogTime:                          MetricSettings{Enabled: true},
					MongodbatlasProcessPageFaults:                         MetricSettings{Enabled: true},
					MongodbatlasProcessRestarts:                           MetricSettings{Enabled: true},
					MongodbatlasProcessTickets:                            MetricSettings{Enabled: true},
					MongodbatlasSystemCPUNormalizedUsageAverage:           MetricSettings{Enabled: true},
					MongodbatlasSystemCPUNormalizedUsageMax:               MetricSettings{Enabled: true},
					MongodbatlasSystemCPUUsageAverage:                     MetricSettings{Enabled: true},
					MongodbatlasSystemCPUUsageMax:                         MetricSettings{Enabled: true},
					MongodbatlasSystemFtsCPUNormalizedUsage:               MetricSettings{Enabled: true},
					MongodbatlasSystemFtsCPUUsage:                         MetricSettings{Enabled: true},
					MongodbatlasSystemFtsDiskUsed:                         MetricSettings{Enabled: true},
					MongodbatlasSystemFtsMemoryUsage:                      MetricSettings{Enabled: true},
					MongodbatlasSystemMemoryUsageAverage:                  MetricSettings{Enabled: true},
					MongodbatlasSystemMemoryUsageMax:                      MetricSettings{Enabled: true},
					MongodbatlasSystemNetworkIoAverage:                    MetricSettings{Enabled: true},
					MongodbatlasSystemNetworkIoMax:                        MetricSettings{Enabled: true},
					MongodbatlasSystemPagingIoAverage:                     MetricSettings{Enabled: true},
					MongodbatlasSystemPagingIoMax:                         MetricSettings{Enabled: true},
					MongodbatlasSystemPagingUsageAverage:                  MetricSettings{Enabled: true},
					MongodbatlasSystemPagingUsageMax:                      MetricSettings{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesSettings{
					MongodbAtlasClusterName:     ResourceAttributeSettings{Enabled: true},
					MongodbAtlasDbName:          ResourceAttributeSettings{Enabled: true},
					MongodbAtlasDiskPartition:   ResourceAttributeSettings{Enabled: true},
					MongodbAtlasHostName:        ResourceAttributeSettings{Enabled: true},
					MongodbAtlasOrgName:         ResourceAttributeSettings{Enabled: true},
					MongodbAtlasProcessID:       ResourceAttributeSettings{Enabled: true},
					MongodbAtlasProcessPort:     ResourceAttributeSettings{Enabled: true},
					MongodbAtlasProcessTypeName: ResourceAttributeSettings{Enabled: true},
					MongodbAtlasProjectID:       ResourceAttributeSettings{Enabled: true},
					MongodbAtlasProjectName:     ResourceAttributeSettings{Enabled: true},
					MongodbAtlasUserAlias:       ResourceAttributeSettings{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsSettings{
					MongodbatlasDbCounts:                                  MetricSettings{Enabled: false},
					MongodbatlasDbSize:                                    MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionIopsAverage:                  MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionIopsMax:                      MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionLatencyAverage:               MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionLatencyMax:                   MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionSpaceAverage:                 MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionSpaceMax:                     MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionUsageAverage:                 MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionUsageMax:                     MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionUtilizationAverage:           MetricSettings{Enabled: false},
					MongodbatlasDiskPartitionUtilizationMax:               MetricSettings{Enabled: false},
					MongodbatlasProcessAsserts:                            MetricSettings{Enabled: false},
					MongodbatlasProcessBackgroundFlush:                    MetricSettings{Enabled: false},
					MongodbatlasProcessCacheIo:                            MetricSettings{Enabled: false},
					MongodbatlasProcessCacheSize:                          MetricSettings{Enabled: false},
					MongodbatlasProcessConnections:                        MetricSettings{Enabled: false},
					MongodbatlasProcessCPUChildrenNormalizedUsageAverage:  MetricSettings{Enabled: false},
					MongodbatlasProcessCPUChildrenNormalizedUsageMax:      MetricSettings{Enabled: false},
					MongodbatlasProcessCPUChildrenUsageAverage:            MetricSettings{Enabled: false},
					MongodbatlasProcessCPUChildrenUsageMax:                MetricSettings{Enabled: false},
					MongodbatlasProcessCPUNormalizedUsageAverage:          MetricSettings{Enabled: false},
					MongodbatlasProcessCPUNormalizedUsageMax:              MetricSettings{Enabled: false},
					MongodbatlasProcessCPUUsageAverage:                    MetricSettings{Enabled: false},
					MongodbatlasProcessCPUUsageMax:                        MetricSettings{Enabled: false},
					MongodbatlasProcessCursors:                            MetricSettings{Enabled: false},
					MongodbatlasProcessDbDocumentRate:                     MetricSettings{Enabled: false},
					MongodbatlasProcessDbOperationsRate:                   MetricSettings{Enabled: false},
					MongodbatlasProcessDbOperationsTime:                   MetricSettings{Enabled: false},
					MongodbatlasProcessDbQueryExecutorScanned:             MetricSettings{Enabled: false},
					MongodbatlasProcessDbQueryTargetingScannedPerReturned: MetricSettings{Enabled: false},
					MongodbatlasProcessDbStorage:                          MetricSettings{Enabled: false},
					MongodbatlasProcessGlobalLock:                         MetricSettings{Enabled: false},
					MongodbatlasProcessIndexBtreeMissRatio:                MetricSettings{Enabled: false},
					MongodbatlasProcessIndexCounters:                      MetricSettings{Enabled: false},
					MongodbatlasProcessJournalingCommits:                  MetricSettings{Enabled: false},
					MongodbatlasProcessJournalingDataFiles:                MetricSettings{Enabled: false},
					MongodbatlasProcessJournalingWritten:                  MetricSettings{Enabled: false},
					MongodbatlasProcessMemoryUsage:                        MetricSettings{Enabled: false},
					MongodbatlasProcessNetworkIo:                          MetricSettings{Enabled: false},
					MongodbatlasProcessNetworkRequests:                    MetricSettings{Enabled: false},
					MongodbatlasProcessOplogRate:                          MetricSettings{Enabled: false},
					MongodbatlasProcessOplogTime:                          MetricSettings{Enabled: false},
					MongodbatlasProcessPageFaults:                         MetricSettings{Enabled: false},
					MongodbatlasProcessRestarts:                           MetricSettings{Enabled: false},
					MongodbatlasProcessTickets:                            MetricSettings{Enabled: false},
					MongodbatlasSystemCPUNormalizedUsageAverage:           MetricSettings{Enabled: false},
					MongodbatlasSystemCPUNormalizedUsageMax:               MetricSettings{Enabled: false},
					MongodbatlasSystemCPUUsageAverage:                     MetricSettings{Enabled: false},
					MongodbatlasSystemCPUUsageMax:                         MetricSettings{Enabled: false},
					MongodbatlasSystemFtsCPUNormalizedUsage:               MetricSettings{Enabled: false},
					MongodbatlasSystemFtsCPUUsage:                         MetricSettings{Enabled: false},
					MongodbatlasSystemFtsDiskUsed:                         MetricSettings{Enabled: false},
					MongodbatlasSystemFtsMemoryUsage:                      MetricSettings{Enabled: false},
					MongodbatlasSystemMemoryUsageAverage:                  MetricSettings{Enabled: false},
					MongodbatlasSystemMemoryUsageMax:                      MetricSettings{Enabled: false},
					MongodbatlasSystemNetworkIoAverage:                    MetricSettings{Enabled: false},
					MongodbatlasSystemNetworkIoMax:                        MetricSettings{Enabled: false},
					MongodbatlasSystemPagingIoAverage:                     MetricSettings{Enabled: false},
					MongodbatlasSystemPagingIoMax:                         MetricSettings{Enabled: false},
					MongodbatlasSystemPagingUsageAverage:                  MetricSettings{Enabled: false},
					MongodbatlasSystemPagingUsageMax:                      MetricSettings{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesSettings{
					MongodbAtlasClusterName:     ResourceAttributeSettings{Enabled: false},
					MongodbAtlasDbName:          ResourceAttributeSettings{Enabled: false},
					MongodbAtlasDiskPartition:   ResourceAttributeSettings{Enabled: false},
					MongodbAtlasHostName:        ResourceAttributeSettings{Enabled: false},
					MongodbAtlasOrgName:         ResourceAttributeSettings{Enabled: false},
					MongodbAtlasProcessID:       ResourceAttributeSettings{Enabled: false},
					MongodbAtlasProcessPort:     ResourceAttributeSettings{Enabled: false},
					MongodbAtlasProcessTypeName: ResourceAttributeSettings{Enabled: false},
					MongodbAtlasProjectID:       ResourceAttributeSettings{Enabled: false},
					MongodbAtlasProjectName:     ResourceAttributeSettings{Enabled: false},
					MongodbAtlasUserAlias:       ResourceAttributeSettings{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricSettings{}, ResourceAttributeSettings{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesSettings(t *testing.T, name string) ResourceAttributesSettings {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesSettings()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
