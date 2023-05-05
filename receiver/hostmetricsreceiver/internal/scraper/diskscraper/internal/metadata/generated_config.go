// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
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

// MetricsSettings provides settings for hostmetricsreceiver/disk metrics.
type MetricsSettings struct {
	SystemDiskIo                MetricSettings `mapstructure:"system.disk.io"`
	SystemDiskIoTime            MetricSettings `mapstructure:"system.disk.io_time"`
	SystemDiskMerged            MetricSettings `mapstructure:"system.disk.merged"`
	SystemDiskOperationTime     MetricSettings `mapstructure:"system.disk.operation_time"`
	SystemDiskOperations        MetricSettings `mapstructure:"system.disk.operations"`
	SystemDiskPendingOperations MetricSettings `mapstructure:"system.disk.pending_operations"`
	SystemDiskWeightedIoTime    MetricSettings `mapstructure:"system.disk.weighted_io_time"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemDiskIo: MetricSettings{
			Enabled: true,
		},
		SystemDiskIoTime: MetricSettings{
			Enabled: true,
		},
		SystemDiskMerged: MetricSettings{
			Enabled: true,
		},
		SystemDiskOperationTime: MetricSettings{
			Enabled: true,
		},
		SystemDiskOperations: MetricSettings{
			Enabled: true,
		},
		SystemDiskPendingOperations: MetricSettings{
			Enabled: true,
		},
		SystemDiskWeightedIoTime: MetricSettings{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for hostmetricsreceiver/disk metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsSettings `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsSettings(),
	}
}
