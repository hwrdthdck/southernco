// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

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
				Metrics: MetricsConfig{
					MongodbActiveReads:            MetricConfig{Enabled: true},
					MongodbActiveWrites:           MetricConfig{Enabled: true},
					MongodbCacheOperations:        MetricConfig{Enabled: true},
					MongodbCollectionCount:        MetricConfig{Enabled: true},
					MongodbCommandsPerSec:         MetricConfig{Enabled: true},
					MongodbConnectionCount:        MetricConfig{Enabled: true},
					MongodbCursorCount:            MetricConfig{Enabled: true},
					MongodbCursorTimeoutCount:     MetricConfig{Enabled: true},
					MongodbDataSize:               MetricConfig{Enabled: true},
					MongodbDatabaseCount:          MetricConfig{Enabled: true},
					MongodbDeletesPerSec:          MetricConfig{Enabled: true},
					MongodbDocumentOperationCount: MetricConfig{Enabled: true},
					MongodbExtentCount:            MetricConfig{Enabled: true},
					MongodbFlushesPerSec:          MetricConfig{Enabled: true},
					MongodbGetmoresPerSec:         MetricConfig{Enabled: true},
					MongodbGlobalLockTime:         MetricConfig{Enabled: true},
					MongodbHealth:                 MetricConfig{Enabled: true},
					MongodbIndexAccessCount:       MetricConfig{Enabled: true},
					MongodbIndexCount:             MetricConfig{Enabled: true},
					MongodbIndexSize:              MetricConfig{Enabled: true},
					MongodbInsertsPerSec:          MetricConfig{Enabled: true},
					MongodbLockAcquireCount:       MetricConfig{Enabled: true},
					MongodbLockAcquireTime:        MetricConfig{Enabled: true},
					MongodbLockAcquireWaitCount:   MetricConfig{Enabled: true},
					MongodbLockDeadlockCount:      MetricConfig{Enabled: true},
					MongodbMemoryUsage:            MetricConfig{Enabled: true},
					MongodbNetworkIoReceive:       MetricConfig{Enabled: true},
					MongodbNetworkIoTransmit:      MetricConfig{Enabled: true},
					MongodbNetworkRequestCount:    MetricConfig{Enabled: true},
					MongodbObjectCount:            MetricConfig{Enabled: true},
					MongodbOperationCount:         MetricConfig{Enabled: true},
					MongodbOperationLatencyTime:   MetricConfig{Enabled: true},
					MongodbOperationReplCount:     MetricConfig{Enabled: true},
					MongodbOperationTime:          MetricConfig{Enabled: true},
					MongodbQueriesPerSec:          MetricConfig{Enabled: true},
					MongodbReplCommandsPerSec:     MetricConfig{Enabled: true},
					MongodbReplDeletesPerSec:      MetricConfig{Enabled: true},
					MongodbReplGetmoresPerSec:     MetricConfig{Enabled: true},
					MongodbReplInsertsPerSec:      MetricConfig{Enabled: true},
					MongodbReplQueriesPerSec:      MetricConfig{Enabled: true},
					MongodbReplUpdatesPerSec:      MetricConfig{Enabled: true},
					MongodbSessionCount:           MetricConfig{Enabled: true},
					MongodbStorageSize:            MetricConfig{Enabled: true},
					MongodbUpdatesPerSec:          MetricConfig{Enabled: true},
					MongodbUptime:                 MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					Database:      ResourceAttributeConfig{Enabled: true},
					ServerAddress: ResourceAttributeConfig{Enabled: true},
					ServerPort:    ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					MongodbActiveReads:            MetricConfig{Enabled: false},
					MongodbActiveWrites:           MetricConfig{Enabled: false},
					MongodbCacheOperations:        MetricConfig{Enabled: false},
					MongodbCollectionCount:        MetricConfig{Enabled: false},
					MongodbCommandsPerSec:         MetricConfig{Enabled: false},
					MongodbConnectionCount:        MetricConfig{Enabled: false},
					MongodbCursorCount:            MetricConfig{Enabled: false},
					MongodbCursorTimeoutCount:     MetricConfig{Enabled: false},
					MongodbDataSize:               MetricConfig{Enabled: false},
					MongodbDatabaseCount:          MetricConfig{Enabled: false},
					MongodbDeletesPerSec:          MetricConfig{Enabled: false},
					MongodbDocumentOperationCount: MetricConfig{Enabled: false},
					MongodbExtentCount:            MetricConfig{Enabled: false},
					MongodbFlushesPerSec:          MetricConfig{Enabled: false},
					MongodbGetmoresPerSec:         MetricConfig{Enabled: false},
					MongodbGlobalLockTime:         MetricConfig{Enabled: false},
					MongodbHealth:                 MetricConfig{Enabled: false},
					MongodbIndexAccessCount:       MetricConfig{Enabled: false},
					MongodbIndexCount:             MetricConfig{Enabled: false},
					MongodbIndexSize:              MetricConfig{Enabled: false},
					MongodbInsertsPerSec:          MetricConfig{Enabled: false},
					MongodbLockAcquireCount:       MetricConfig{Enabled: false},
					MongodbLockAcquireTime:        MetricConfig{Enabled: false},
					MongodbLockAcquireWaitCount:   MetricConfig{Enabled: false},
					MongodbLockDeadlockCount:      MetricConfig{Enabled: false},
					MongodbMemoryUsage:            MetricConfig{Enabled: false},
					MongodbNetworkIoReceive:       MetricConfig{Enabled: false},
					MongodbNetworkIoTransmit:      MetricConfig{Enabled: false},
					MongodbNetworkRequestCount:    MetricConfig{Enabled: false},
					MongodbObjectCount:            MetricConfig{Enabled: false},
					MongodbOperationCount:         MetricConfig{Enabled: false},
					MongodbOperationLatencyTime:   MetricConfig{Enabled: false},
					MongodbOperationReplCount:     MetricConfig{Enabled: false},
					MongodbOperationTime:          MetricConfig{Enabled: false},
					MongodbQueriesPerSec:          MetricConfig{Enabled: false},
					MongodbReplCommandsPerSec:     MetricConfig{Enabled: false},
					MongodbReplDeletesPerSec:      MetricConfig{Enabled: false},
					MongodbReplGetmoresPerSec:     MetricConfig{Enabled: false},
					MongodbReplInsertsPerSec:      MetricConfig{Enabled: false},
					MongodbReplQueriesPerSec:      MetricConfig{Enabled: false},
					MongodbReplUpdatesPerSec:      MetricConfig{Enabled: false},
					MongodbSessionCount:           MetricConfig{Enabled: false},
					MongodbStorageSize:            MetricConfig{Enabled: false},
					MongodbUpdatesPerSec:          MetricConfig{Enabled: false},
					MongodbUptime:                 MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					Database:      ResourceAttributeConfig{Enabled: false},
					ServerAddress: ResourceAttributeConfig{Enabled: false},
					ServerPort:    ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				Database:      ResourceAttributeConfig{Enabled: true},
				ServerAddress: ResourceAttributeConfig{Enabled: true},
				ServerPort:    ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				Database:      ResourceAttributeConfig{Enabled: false},
				ServerAddress: ResourceAttributeConfig{Enabled: false},
				ServerPort:    ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
