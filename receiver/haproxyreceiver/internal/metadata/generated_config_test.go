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
					HaproxyBytesInput:           MetricConfig{Enabled: true},
					HaproxyBytesOutput:          MetricConfig{Enabled: true},
					HaproxyClientsCanceled:      MetricConfig{Enabled: true},
					HaproxyCompressionBypass:    MetricConfig{Enabled: true},
					HaproxyCompressionCount:     MetricConfig{Enabled: true},
					HaproxyCompressionInput:     MetricConfig{Enabled: true},
					HaproxyCompressionOutput:    MetricConfig{Enabled: true},
					HaproxyConnectionsErrors:    MetricConfig{Enabled: true},
					HaproxyConnectionsRate:      MetricConfig{Enabled: true},
					HaproxyConnectionsRetries:   MetricConfig{Enabled: true},
					HaproxyConnectionsTotal:     MetricConfig{Enabled: true},
					HaproxyDowntime:             MetricConfig{Enabled: true},
					HaproxyFailedChecks:         MetricConfig{Enabled: true},
					HaproxyRequestsDenied:       MetricConfig{Enabled: true},
					HaproxyRequestsErrors:       MetricConfig{Enabled: true},
					HaproxyRequestsQueued:       MetricConfig{Enabled: true},
					HaproxyRequestsRate:         MetricConfig{Enabled: true},
					HaproxyRequestsRedispatched: MetricConfig{Enabled: true},
					HaproxyRequestsTotal:        MetricConfig{Enabled: true},
					HaproxyResponsesDenied:      MetricConfig{Enabled: true},
					HaproxyResponsesErrors:      MetricConfig{Enabled: true},
					HaproxyServerSelectedTotal:  MetricConfig{Enabled: true},
					HaproxySessionsAverage:      MetricConfig{Enabled: true},
					HaproxySessionsCount:        MetricConfig{Enabled: true},
					HaproxySessionsRate:         MetricConfig{Enabled: true},
					HaproxySessionsTotal:        MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					HaproxyAddr:        ResourceAttributeConfig{Enabled: true},
					HaproxyProxyName:   ResourceAttributeConfig{Enabled: true},
					HaproxyServiceName: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					HaproxyBytesInput:           MetricConfig{Enabled: false},
					HaproxyBytesOutput:          MetricConfig{Enabled: false},
					HaproxyClientsCanceled:      MetricConfig{Enabled: false},
					HaproxyCompressionBypass:    MetricConfig{Enabled: false},
					HaproxyCompressionCount:     MetricConfig{Enabled: false},
					HaproxyCompressionInput:     MetricConfig{Enabled: false},
					HaproxyCompressionOutput:    MetricConfig{Enabled: false},
					HaproxyConnectionsErrors:    MetricConfig{Enabled: false},
					HaproxyConnectionsRate:      MetricConfig{Enabled: false},
					HaproxyConnectionsRetries:   MetricConfig{Enabled: false},
					HaproxyConnectionsTotal:     MetricConfig{Enabled: false},
					HaproxyDowntime:             MetricConfig{Enabled: false},
					HaproxyFailedChecks:         MetricConfig{Enabled: false},
					HaproxyRequestsDenied:       MetricConfig{Enabled: false},
					HaproxyRequestsErrors:       MetricConfig{Enabled: false},
					HaproxyRequestsQueued:       MetricConfig{Enabled: false},
					HaproxyRequestsRate:         MetricConfig{Enabled: false},
					HaproxyRequestsRedispatched: MetricConfig{Enabled: false},
					HaproxyRequestsTotal:        MetricConfig{Enabled: false},
					HaproxyResponsesDenied:      MetricConfig{Enabled: false},
					HaproxyResponsesErrors:      MetricConfig{Enabled: false},
					HaproxyServerSelectedTotal:  MetricConfig{Enabled: false},
					HaproxySessionsAverage:      MetricConfig{Enabled: false},
					HaproxySessionsCount:        MetricConfig{Enabled: false},
					HaproxySessionsRate:         MetricConfig{Enabled: false},
					HaproxySessionsTotal:        MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					HaproxyAddr:        ResourceAttributeConfig{Enabled: false},
					HaproxyProxyName:   ResourceAttributeConfig{Enabled: false},
					HaproxyServiceName: ResourceAttributeConfig{Enabled: false},
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
				HaproxyAddr:        ResourceAttributeConfig{Enabled: true},
				HaproxyProxyName:   ResourceAttributeConfig{Enabled: true},
				HaproxyServiceName: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				HaproxyAddr:        ResourceAttributeConfig{Enabled: false},
				HaproxyProxyName:   ResourceAttributeConfig{Enabled: false},
				HaproxyServiceName: ResourceAttributeConfig{Enabled: false},
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
