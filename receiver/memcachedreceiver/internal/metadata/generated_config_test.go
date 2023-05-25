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
					MemcachedBytes:              MetricConfig{Enabled: true},
					MemcachedCommands:           MetricConfig{Enabled: true},
					MemcachedConnectionsCurrent: MetricConfig{Enabled: true},
					MemcachedConnectionsTotal:   MetricConfig{Enabled: true},
					MemcachedCPUUsage:           MetricConfig{Enabled: true},
					MemcachedCurrentItems:       MetricConfig{Enabled: true},
					MemcachedEvictions:          MetricConfig{Enabled: true},
					MemcachedNetwork:            MetricConfig{Enabled: true},
					MemcachedOperationHitRatio:  MetricConfig{Enabled: true},
					MemcachedOperations:         MetricConfig{Enabled: true},
					MemcachedThreads:            MetricConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					MemcachedBytes:              MetricConfig{Enabled: false},
					MemcachedCommands:           MetricConfig{Enabled: false},
					MemcachedConnectionsCurrent: MetricConfig{Enabled: false},
					MemcachedConnectionsTotal:   MetricConfig{Enabled: false},
					MemcachedCPUUsage:           MetricConfig{Enabled: false},
					MemcachedCurrentItems:       MetricConfig{Enabled: false},
					MemcachedEvictions:          MetricConfig{Enabled: false},
					MemcachedNetwork:            MetricConfig{Enabled: false},
					MemcachedOperationHitRatio:  MetricConfig{Enabled: false},
					MemcachedOperations:         MetricConfig{Enabled: false},
					MemcachedThreads:            MetricConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{})); diff != "" {
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
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
