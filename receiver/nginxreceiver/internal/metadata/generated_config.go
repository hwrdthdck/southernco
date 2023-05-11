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

// MetricsConfig provides config for nginxreceiver metrics.
type MetricsConfig struct {
	NginxConnectionsAccepted MetricConfig `mapstructure:"nginx.connections_accepted"`
	NginxConnectionsCurrent  MetricConfig `mapstructure:"nginx.connections_current"`
	NginxConnectionsHandled  MetricConfig `mapstructure:"nginx.connections_handled"`
	NginxRequests            MetricConfig `mapstructure:"nginx.requests"`
	TempConnectionsCurrent   MetricConfig `mapstructure:"temp.connections_current"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		NginxConnectionsAccepted: MetricConfig{
			Enabled: true,
		},
		NginxConnectionsCurrent: MetricConfig{
			Enabled: true,
		},
		NginxConnectionsHandled: MetricConfig{
			Enabled: true,
		},
		NginxRequests: MetricConfig{
			Enabled: true,
		},
		TempConnectionsCurrent: MetricConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for nginxreceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
