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
				Metrics: MetricsSettings{
					SystemDiskIo:                MetricSettings{Enabled: true},
					SystemDiskIoTime:            MetricSettings{Enabled: true},
					SystemDiskMerged:            MetricSettings{Enabled: true},
					SystemDiskOperationTime:     MetricSettings{Enabled: true},
					SystemDiskOperations:        MetricSettings{Enabled: true},
					SystemDiskPendingOperations: MetricSettings{Enabled: true},
					SystemDiskWeightedIoTime:    MetricSettings{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesSettings{},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsSettings{
					SystemDiskIo:                MetricSettings{Enabled: false},
					SystemDiskIoTime:            MetricSettings{Enabled: false},
					SystemDiskMerged:            MetricSettings{Enabled: false},
					SystemDiskOperationTime:     MetricSettings{Enabled: false},
					SystemDiskOperations:        MetricSettings{Enabled: false},
					SystemDiskPendingOperations: MetricSettings{Enabled: false},
					SystemDiskWeightedIoTime:    MetricSettings{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesSettings{},
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

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
