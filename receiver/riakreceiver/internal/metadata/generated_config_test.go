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
				RiakNodeName: ResourceAttributeSettings{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesSettings{
				RiakNodeName: ResourceAttributeSettings{Enabled: false},
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
					RiakMemoryLimit:              MetricSettings{Enabled: true},
					RiakNodeOperationCount:       MetricSettings{Enabled: true},
					RiakNodeOperationTimeMean:    MetricSettings{Enabled: true},
					RiakNodeReadRepairCount:      MetricSettings{Enabled: true},
					RiakVnodeIndexOperationCount: MetricSettings{Enabled: true},
					RiakVnodeOperationCount:      MetricSettings{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesSettings{
					RiakNodeName: ResourceAttributeSettings{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsSettings{
					RiakMemoryLimit:              MetricSettings{Enabled: false},
					RiakNodeOperationCount:       MetricSettings{Enabled: false},
					RiakNodeOperationTimeMean:    MetricSettings{Enabled: false},
					RiakNodeReadRepairCount:      MetricSettings{Enabled: false},
					RiakVnodeIndexOperationCount: MetricSettings{Enabled: false},
					RiakVnodeOperationCount:      MetricSettings{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesSettings{
					RiakNodeName: ResourceAttributeSettings{Enabled: false},
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
