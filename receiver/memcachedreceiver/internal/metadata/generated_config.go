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

// MetricsConfig provides config for memcachedreceiver metrics.
type MetricsConfig struct {
	MemcachedBytes              MetricConfig `mapstructure:"memcached.bytes"`
	MemcachedCommands           MetricConfig `mapstructure:"memcached.commands"`
	MemcachedConnectionsCurrent MetricConfig `mapstructure:"memcached.connections.current"`
	MemcachedConnectionsTotal   MetricConfig `mapstructure:"memcached.connections.total"`
	MemcachedCPUUsage           MetricConfig `mapstructure:"memcached.cpu.usage"`
	MemcachedCurrentItems       MetricConfig `mapstructure:"memcached.current_items"`
	MemcachedEvictions          MetricConfig `mapstructure:"memcached.evictions"`
	MemcachedNetwork            MetricConfig `mapstructure:"memcached.network"`
	MemcachedOperationHitRatio  MetricConfig `mapstructure:"memcached.operation_hit_ratio"`
	MemcachedOperations         MetricConfig `mapstructure:"memcached.operations"`
	MemcachedThreads            MetricConfig `mapstructure:"memcached.threads"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		MemcachedBytes: MetricConfig{
			Enabled: true,
		},
		MemcachedCommands: MetricConfig{
			Enabled: true,
		},
		MemcachedConnectionsCurrent: MetricConfig{
			Enabled: true,
		},
		MemcachedConnectionsTotal: MetricConfig{
			Enabled: true,
		},
		MemcachedCPUUsage: MetricConfig{
			Enabled: true,
		},
		MemcachedCurrentItems: MetricConfig{
			Enabled: true,
		},
		MemcachedEvictions: MetricConfig{
			Enabled: true,
		},
		MemcachedNetwork: MetricConfig{
			Enabled: true,
		},
		MemcachedOperationHitRatio: MetricConfig{
			Enabled: true,
		},
		MemcachedOperations: MetricConfig{
			Enabled: true,
		},
		MemcachedThreads: MetricConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for memcachedreceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
