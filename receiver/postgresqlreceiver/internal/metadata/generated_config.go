// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for postgresql metrics.
type MetricsConfig struct {
	PostgresqlBackends                 MetricConfig `mapstructure:"postgresql.backends"`
	PostgresqlBgwriterBuffersAllocated MetricConfig `mapstructure:"postgresql.bgwriter.buffers.allocated"`
	PostgresqlBgwriterBuffersWrites    MetricConfig `mapstructure:"postgresql.bgwriter.buffers.writes"`
	PostgresqlBgwriterCheckpointCount  MetricConfig `mapstructure:"postgresql.bgwriter.checkpoint.count"`
	PostgresqlBgwriterDuration         MetricConfig `mapstructure:"postgresql.bgwriter.duration"`
	PostgresqlBgwriterMaxwritten       MetricConfig `mapstructure:"postgresql.bgwriter.maxwritten"`
	PostgresqlBlocksRead               MetricConfig `mapstructure:"postgresql.blocks_read"`
	PostgresqlCommits                  MetricConfig `mapstructure:"postgresql.commits"`
	PostgresqlConnectionMax            MetricConfig `mapstructure:"postgresql.connection.max"`
	PostgresqlDatabaseCount            MetricConfig `mapstructure:"postgresql.database.count"`
	PostgresqlDatabaseLocks            MetricConfig `mapstructure:"postgresql.database.locks"`
	PostgresqlDbSize                   MetricConfig `mapstructure:"postgresql.db_size"`
	PostgresqlDeadlocks                MetricConfig `mapstructure:"postgresql.deadlocks"`
	PostgresqlIndexScans               MetricConfig `mapstructure:"postgresql.index.scans"`
	PostgresqlIndexSize                MetricConfig `mapstructure:"postgresql.index.size"`
	PostgresqlOperations               MetricConfig `mapstructure:"postgresql.operations"`
	PostgresqlReplicationDataDelay     MetricConfig `mapstructure:"postgresql.replication.data_delay"`
	PostgresqlRollbacks                MetricConfig `mapstructure:"postgresql.rollbacks"`
	PostgresqlRows                     MetricConfig `mapstructure:"postgresql.rows"`
	PostgresqlSequentialScans          MetricConfig `mapstructure:"postgresql.sequential_scans"`
	PostgresqlTableCount               MetricConfig `mapstructure:"postgresql.table.count"`
	PostgresqlTableSize                MetricConfig `mapstructure:"postgresql.table.size"`
	PostgresqlTableVacuumCount         MetricConfig `mapstructure:"postgresql.table.vacuum.count"`
	PostgresqlTempFiles                MetricConfig `mapstructure:"postgresql.temp_files"`
	PostgresqlWalAge                   MetricConfig `mapstructure:"postgresql.wal.age"`
	PostgresqlWalLag                   MetricConfig `mapstructure:"postgresql.wal.lag"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		PostgresqlBackends: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterBuffersAllocated: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterBuffersWrites: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterCheckpointCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterDuration: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterMaxwritten: MetricConfig{
			Enabled: true,
		},
		PostgresqlBlocksRead: MetricConfig{
			Enabled: true,
		},
		PostgresqlCommits: MetricConfig{
			Enabled: true,
		},
		PostgresqlConnectionMax: MetricConfig{
			Enabled: true,
		},
		PostgresqlDatabaseCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlDatabaseLocks: MetricConfig{
			Enabled: false,
		},
		PostgresqlDbSize: MetricConfig{
			Enabled: true,
		},
		PostgresqlDeadlocks: MetricConfig{
			Enabled: false,
		},
		PostgresqlIndexScans: MetricConfig{
			Enabled: true,
		},
		PostgresqlIndexSize: MetricConfig{
			Enabled: true,
		},
		PostgresqlOperations: MetricConfig{
			Enabled: true,
		},
		PostgresqlReplicationDataDelay: MetricConfig{
			Enabled: true,
		},
		PostgresqlRollbacks: MetricConfig{
			Enabled: true,
		},
		PostgresqlRows: MetricConfig{
			Enabled: true,
		},
		PostgresqlSequentialScans: MetricConfig{
			Enabled: false,
		},
		PostgresqlTableCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlTableSize: MetricConfig{
			Enabled: true,
		},
		PostgresqlTableVacuumCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlTempFiles: MetricConfig{
			Enabled: false,
		},
		PostgresqlWalAge: MetricConfig{
			Enabled: true,
		},
		PostgresqlWalLag: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for postgresql resource attributes.
type ResourceAttributesConfig struct {
	PostgresqlDatabaseName ResourceAttributeConfig `mapstructure:"postgresql.database.name"`
	PostgresqlIndexName    ResourceAttributeConfig `mapstructure:"postgresql.index.name"`
	PostgresqlTableName    ResourceAttributeConfig `mapstructure:"postgresql.table.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		PostgresqlDatabaseName: ResourceAttributeConfig{
			Enabled: true,
		},
		PostgresqlIndexName: ResourceAttributeConfig{
			Enabled: true,
		},
		PostgresqlTableName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for postgresql metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
