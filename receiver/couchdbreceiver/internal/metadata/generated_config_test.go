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
					CouchdbAverageRequestTime: MetricConfig{Enabled: true},
					CouchdbDatabaseOpen:       MetricConfig{Enabled: true},
					CouchdbDatabaseOperations: MetricConfig{Enabled: true},
					CouchdbFileDescriptorOpen: MetricConfig{Enabled: true},
					CouchdbHttpdBulkRequests:  MetricConfig{Enabled: true},
					CouchdbHttpdRequests:      MetricConfig{Enabled: true},
					CouchdbHttpdResponses:     MetricConfig{Enabled: true},
					CouchdbHttpdViews:         MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					CouchdbNodeName: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					CouchdbAverageRequestTime: MetricConfig{Enabled: false},
					CouchdbDatabaseOpen:       MetricConfig{Enabled: false},
					CouchdbDatabaseOperations: MetricConfig{Enabled: false},
					CouchdbFileDescriptorOpen: MetricConfig{Enabled: false},
					CouchdbHttpdBulkRequests:  MetricConfig{Enabled: false},
					CouchdbHttpdRequests:      MetricConfig{Enabled: false},
					CouchdbHttpdResponses:     MetricConfig{Enabled: false},
					CouchdbHttpdViews:         MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					CouchdbNodeName: ResourceAttributeConfig{Enabled: false},
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
