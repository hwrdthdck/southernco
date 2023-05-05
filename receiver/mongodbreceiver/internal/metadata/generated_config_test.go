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
				Database: ResourceAttributeSettings{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesSettings{
				Database: ResourceAttributeSettings{Enabled: false},
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
					MongodbCacheOperations:        MetricSettings{Enabled: true},
					MongodbCollectionCount:        MetricSettings{Enabled: true},
					MongodbConnectionCount:        MetricSettings{Enabled: true},
					MongodbCursorCount:            MetricSettings{Enabled: true},
					MongodbCursorTimeoutCount:     MetricSettings{Enabled: true},
					MongodbDataSize:               MetricSettings{Enabled: true},
					MongodbDatabaseCount:          MetricSettings{Enabled: true},
					MongodbDocumentOperationCount: MetricSettings{Enabled: true},
					MongodbExtentCount:            MetricSettings{Enabled: true},
					MongodbGlobalLockTime:         MetricSettings{Enabled: true},
					MongodbHealth:                 MetricSettings{Enabled: true},
					MongodbIndexAccessCount:       MetricSettings{Enabled: true},
					MongodbIndexCount:             MetricSettings{Enabled: true},
					MongodbIndexSize:              MetricSettings{Enabled: true},
					MongodbLockAcquireCount:       MetricSettings{Enabled: true},
					MongodbLockAcquireTime:        MetricSettings{Enabled: true},
					MongodbLockAcquireWaitCount:   MetricSettings{Enabled: true},
					MongodbLockDeadlockCount:      MetricSettings{Enabled: true},
					MongodbMemoryUsage:            MetricSettings{Enabled: true},
					MongodbNetworkIoReceive:       MetricSettings{Enabled: true},
					MongodbNetworkIoTransmit:      MetricSettings{Enabled: true},
					MongodbNetworkRequestCount:    MetricSettings{Enabled: true},
					MongodbObjectCount:            MetricSettings{Enabled: true},
					MongodbOperationCount:         MetricSettings{Enabled: true},
					MongodbOperationLatencyTime:   MetricSettings{Enabled: true},
					MongodbOperationReplCount:     MetricSettings{Enabled: true},
					MongodbOperationTime:          MetricSettings{Enabled: true},
					MongodbSessionCount:           MetricSettings{Enabled: true},
					MongodbStorageSize:            MetricSettings{Enabled: true},
					MongodbUptime:                 MetricSettings{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesSettings{
					Database: ResourceAttributeSettings{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsSettings{
					MongodbCacheOperations:        MetricSettings{Enabled: false},
					MongodbCollectionCount:        MetricSettings{Enabled: false},
					MongodbConnectionCount:        MetricSettings{Enabled: false},
					MongodbCursorCount:            MetricSettings{Enabled: false},
					MongodbCursorTimeoutCount:     MetricSettings{Enabled: false},
					MongodbDataSize:               MetricSettings{Enabled: false},
					MongodbDatabaseCount:          MetricSettings{Enabled: false},
					MongodbDocumentOperationCount: MetricSettings{Enabled: false},
					MongodbExtentCount:            MetricSettings{Enabled: false},
					MongodbGlobalLockTime:         MetricSettings{Enabled: false},
					MongodbHealth:                 MetricSettings{Enabled: false},
					MongodbIndexAccessCount:       MetricSettings{Enabled: false},
					MongodbIndexCount:             MetricSettings{Enabled: false},
					MongodbIndexSize:              MetricSettings{Enabled: false},
					MongodbLockAcquireCount:       MetricSettings{Enabled: false},
					MongodbLockAcquireTime:        MetricSettings{Enabled: false},
					MongodbLockAcquireWaitCount:   MetricSettings{Enabled: false},
					MongodbLockDeadlockCount:      MetricSettings{Enabled: false},
					MongodbMemoryUsage:            MetricSettings{Enabled: false},
					MongodbNetworkIoReceive:       MetricSettings{Enabled: false},
					MongodbNetworkIoTransmit:      MetricSettings{Enabled: false},
					MongodbNetworkRequestCount:    MetricSettings{Enabled: false},
					MongodbObjectCount:            MetricSettings{Enabled: false},
					MongodbOperationCount:         MetricSettings{Enabled: false},
					MongodbOperationLatencyTime:   MetricSettings{Enabled: false},
					MongodbOperationReplCount:     MetricSettings{Enabled: false},
					MongodbOperationTime:          MetricSettings{Enabled: false},
					MongodbSessionCount:           MetricSettings{Enabled: false},
					MongodbStorageSize:            MetricSettings{Enabled: false},
					MongodbUptime:                 MetricSettings{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesSettings{
					Database: ResourceAttributeSettings{Enabled: false},
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
