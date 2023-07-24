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
					K8sJobActivePods:            MetricConfig{Enabled: true},
					K8sJobDesiredSuccessfulPods: MetricConfig{Enabled: true},
					K8sJobFailedPods:            MetricConfig{Enabled: true},
					K8sJobMaxParallelPods:       MetricConfig{Enabled: true},
					K8sJobSuccessfulPods:        MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					K8sJobName:             ResourceAttributeConfig{Enabled: true},
					K8sJobUID:              ResourceAttributeConfig{Enabled: true},
					K8sNamespaceName:       ResourceAttributeConfig{Enabled: true},
					OpencensusResourcetype: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sJobActivePods:            MetricConfig{Enabled: false},
					K8sJobDesiredSuccessfulPods: MetricConfig{Enabled: false},
					K8sJobFailedPods:            MetricConfig{Enabled: false},
					K8sJobMaxParallelPods:       MetricConfig{Enabled: false},
					K8sJobSuccessfulPods:        MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					K8sJobName:             ResourceAttributeConfig{Enabled: false},
					K8sJobUID:              ResourceAttributeConfig{Enabled: false},
					K8sNamespaceName:       ResourceAttributeConfig{Enabled: false},
					OpencensusResourcetype: ResourceAttributeConfig{Enabled: false},
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
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
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
				K8sJobName:             ResourceAttributeConfig{Enabled: true},
				K8sJobUID:              ResourceAttributeConfig{Enabled: true},
				K8sNamespaceName:       ResourceAttributeConfig{Enabled: true},
				OpencensusResourcetype: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				K8sJobName:             ResourceAttributeConfig{Enabled: false},
				K8sJobUID:              ResourceAttributeConfig{Enabled: false},
				K8sNamespaceName:       ResourceAttributeConfig{Enabled: false},
				OpencensusResourcetype: ResourceAttributeConfig{Enabled: false},
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
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
