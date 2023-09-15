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
					PostgresqlBackends:                 MetricConfig{Enabled: true},
					PostgresqlBgwriterBuffersAllocated: MetricConfig{Enabled: true},
					PostgresqlBgwriterBuffersWrites:    MetricConfig{Enabled: true},
					PostgresqlBgwriterCheckpointCount:  MetricConfig{Enabled: true},
					PostgresqlBgwriterDuration:         MetricConfig{Enabled: true},
					PostgresqlBgwriterMaxwritten:       MetricConfig{Enabled: true},
					PostgresqlBlocksRead:               MetricConfig{Enabled: true},
					PostgresqlCommits:                  MetricConfig{Enabled: true},
					PostgresqlConnectionMax:            MetricConfig{Enabled: true},
					PostgresqlDatabaseCount:            MetricConfig{Enabled: true},
					PostgresqlDatabaseLocks:            MetricConfig{Enabled: true},
					PostgresqlDbSize:                   MetricConfig{Enabled: true},
					PostgresqlDeadlocks:                MetricConfig{Enabled: true},
					PostgresqlIndexScans:               MetricConfig{Enabled: true},
					PostgresqlIndexSize:                MetricConfig{Enabled: true},
					PostgresqlOperations:               MetricConfig{Enabled: true},
					PostgresqlReplicationDataDelay:     MetricConfig{Enabled: true},
					PostgresqlRollbacks:                MetricConfig{Enabled: true},
					PostgresqlRows:                     MetricConfig{Enabled: true},
					PostgresqlSequentialScans:          MetricConfig{Enabled: true},
					PostgresqlTableCount:               MetricConfig{Enabled: true},
					PostgresqlTableSize:                MetricConfig{Enabled: true},
					PostgresqlTableVacuumCount:         MetricConfig{Enabled: true},
					PostgresqlTempFiles:                MetricConfig{Enabled: true},
					PostgresqlWalAge:                   MetricConfig{Enabled: true},
					PostgresqlWalLag:                   MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					PostgresqlDatabaseName: ResourceAttributeConfig{Enabled: true},
					PostgresqlIndexName:    ResourceAttributeConfig{Enabled: true},
					PostgresqlTableName:    ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					PostgresqlBackends:                 MetricConfig{Enabled: false},
					PostgresqlBgwriterBuffersAllocated: MetricConfig{Enabled: false},
					PostgresqlBgwriterBuffersWrites:    MetricConfig{Enabled: false},
					PostgresqlBgwriterCheckpointCount:  MetricConfig{Enabled: false},
					PostgresqlBgwriterDuration:         MetricConfig{Enabled: false},
					PostgresqlBgwriterMaxwritten:       MetricConfig{Enabled: false},
					PostgresqlBlocksRead:               MetricConfig{Enabled: false},
					PostgresqlCommits:                  MetricConfig{Enabled: false},
					PostgresqlConnectionMax:            MetricConfig{Enabled: false},
					PostgresqlDatabaseCount:            MetricConfig{Enabled: false},
					PostgresqlDatabaseLocks:            MetricConfig{Enabled: false},
					PostgresqlDbSize:                   MetricConfig{Enabled: false},
					PostgresqlDeadlocks:                MetricConfig{Enabled: false},
					PostgresqlIndexScans:               MetricConfig{Enabled: false},
					PostgresqlIndexSize:                MetricConfig{Enabled: false},
					PostgresqlOperations:               MetricConfig{Enabled: false},
					PostgresqlReplicationDataDelay:     MetricConfig{Enabled: false},
					PostgresqlRollbacks:                MetricConfig{Enabled: false},
					PostgresqlRows:                     MetricConfig{Enabled: false},
					PostgresqlSequentialScans:          MetricConfig{Enabled: false},
					PostgresqlTableCount:               MetricConfig{Enabled: false},
					PostgresqlTableSize:                MetricConfig{Enabled: false},
					PostgresqlTableVacuumCount:         MetricConfig{Enabled: false},
					PostgresqlTempFiles:                MetricConfig{Enabled: false},
					PostgresqlWalAge:                   MetricConfig{Enabled: false},
					PostgresqlWalLag:                   MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					PostgresqlDatabaseName: ResourceAttributeConfig{Enabled: false},
					PostgresqlIndexName:    ResourceAttributeConfig{Enabled: false},
					PostgresqlTableName:    ResourceAttributeConfig{Enabled: false},
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
				PostgresqlDatabaseName: ResourceAttributeConfig{Enabled: true},
				PostgresqlIndexName:    ResourceAttributeConfig{Enabled: true},
				PostgresqlTableName:    ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				PostgresqlDatabaseName: ResourceAttributeConfig{Enabled: false},
				PostgresqlIndexName:    ResourceAttributeConfig{Enabled: false},
				PostgresqlTableName:    ResourceAttributeConfig{Enabled: false},
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
