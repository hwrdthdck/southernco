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
					ProcessContextSwitches:     MetricConfig{Enabled: true},
					ProcessCPUTime:             MetricConfig{Enabled: true},
					ProcessCPUUtilization:      MetricConfig{Enabled: true},
					ProcessDiskIo:              MetricConfig{Enabled: true},
					ProcessDiskOperations:      MetricConfig{Enabled: true},
					ProcessHandles:             MetricConfig{Enabled: true},
					ProcessMemoryUsage:         MetricConfig{Enabled: true},
					ProcessMemoryUtilization:   MetricConfig{Enabled: true},
					ProcessMemoryVirtual:       MetricConfig{Enabled: true},
					ProcessOpenFileDescriptors: MetricConfig{Enabled: true},
					ProcessPagingFaults:        MetricConfig{Enabled: true},
					ProcessSignalsPending:      MetricConfig{Enabled: true},
					ProcessThreads:             MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ProcessCgroup:         ResourceAttributeConfig{Enabled: true},
					ProcessCommand:        ResourceAttributeConfig{Enabled: true},
					ProcessCommandLine:    ResourceAttributeConfig{Enabled: true},
					ProcessExecutableName: ResourceAttributeConfig{Enabled: true},
					ProcessExecutablePath: ResourceAttributeConfig{Enabled: true},
					ProcessOwner:          ResourceAttributeConfig{Enabled: true},
					ProcessParentPid:      ResourceAttributeConfig{Enabled: true},
					ProcessPid:            ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ProcessContextSwitches:     MetricConfig{Enabled: false},
					ProcessCPUTime:             MetricConfig{Enabled: false},
					ProcessCPUUtilization:      MetricConfig{Enabled: false},
					ProcessDiskIo:              MetricConfig{Enabled: false},
					ProcessDiskOperations:      MetricConfig{Enabled: false},
					ProcessHandles:             MetricConfig{Enabled: false},
					ProcessMemoryUsage:         MetricConfig{Enabled: false},
					ProcessMemoryUtilization:   MetricConfig{Enabled: false},
					ProcessMemoryVirtual:       MetricConfig{Enabled: false},
					ProcessOpenFileDescriptors: MetricConfig{Enabled: false},
					ProcessPagingFaults:        MetricConfig{Enabled: false},
					ProcessSignalsPending:      MetricConfig{Enabled: false},
					ProcessThreads:             MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ProcessCgroup:         ResourceAttributeConfig{Enabled: false},
					ProcessCommand:        ResourceAttributeConfig{Enabled: false},
					ProcessCommandLine:    ResourceAttributeConfig{Enabled: false},
					ProcessExecutableName: ResourceAttributeConfig{Enabled: false},
					ProcessExecutablePath: ResourceAttributeConfig{Enabled: false},
					ProcessOwner:          ResourceAttributeConfig{Enabled: false},
					ProcessParentPid:      ResourceAttributeConfig{Enabled: false},
					ProcessPid:            ResourceAttributeConfig{Enabled: false},
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
				ProcessCgroup:         ResourceAttributeConfig{Enabled: true},
				ProcessCommand:        ResourceAttributeConfig{Enabled: true},
				ProcessCommandLine:    ResourceAttributeConfig{Enabled: true},
				ProcessExecutableName: ResourceAttributeConfig{Enabled: true},
				ProcessExecutablePath: ResourceAttributeConfig{Enabled: true},
				ProcessOwner:          ResourceAttributeConfig{Enabled: true},
				ProcessParentPid:      ResourceAttributeConfig{Enabled: true},
				ProcessPid:            ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				ProcessCgroup:         ResourceAttributeConfig{Enabled: false},
				ProcessCommand:        ResourceAttributeConfig{Enabled: false},
				ProcessCommandLine:    ResourceAttributeConfig{Enabled: false},
				ProcessExecutableName: ResourceAttributeConfig{Enabled: false},
				ProcessExecutablePath: ResourceAttributeConfig{Enabled: false},
				ProcessOwner:          ResourceAttributeConfig{Enabled: false},
				ProcessParentPid:      ResourceAttributeConfig{Enabled: false},
				ProcessPid:            ResourceAttributeConfig{Enabled: false},
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
