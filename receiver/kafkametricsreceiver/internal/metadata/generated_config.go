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
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for kafkametrics metrics.
type MetricsConfig struct {
	KafkaBrokers                 MetricConfig `mapstructure:"kafka.brokers"`
	KafkaConsumerGroupLag        MetricConfig `mapstructure:"kafka.consumer_group.lag"`
	KafkaConsumerGroupLagSum     MetricConfig `mapstructure:"kafka.consumer_group.lag_sum"`
	KafkaConsumerGroupMembers    MetricConfig `mapstructure:"kafka.consumer_group.members"`
	KafkaConsumerGroupOffset     MetricConfig `mapstructure:"kafka.consumer_group.offset"`
	KafkaConsumerGroupOffsetSum  MetricConfig `mapstructure:"kafka.consumer_group.offset_sum"`
	KafkaPartitionCurrentOffset  MetricConfig `mapstructure:"kafka.partition.current_offset"`
	KafkaPartitionOldestOffset   MetricConfig `mapstructure:"kafka.partition.oldest_offset"`
	KafkaPartitionReplicas       MetricConfig `mapstructure:"kafka.partition.replicas"`
	KafkaPartitionReplicasInSync MetricConfig `mapstructure:"kafka.partition.replicas_in_sync"`
	KafkaTopicPartitions         MetricConfig `mapstructure:"kafka.topic.partitions"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		KafkaBrokers: MetricConfig{
			Enabled: true,
		},
		KafkaConsumerGroupLag: MetricConfig{
			Enabled: true,
		},
		KafkaConsumerGroupLagSum: MetricConfig{
			Enabled: true,
		},
		KafkaConsumerGroupMembers: MetricConfig{
			Enabled: true,
		},
		KafkaConsumerGroupOffset: MetricConfig{
			Enabled: true,
		},
		KafkaConsumerGroupOffsetSum: MetricConfig{
			Enabled: true,
		},
		KafkaPartitionCurrentOffset: MetricConfig{
			Enabled: true,
		},
		KafkaPartitionOldestOffset: MetricConfig{
			Enabled: true,
		},
		KafkaPartitionReplicas: MetricConfig{
			Enabled: true,
		},
		KafkaPartitionReplicasInSync: MetricConfig{
			Enabled: true,
		},
		KafkaTopicPartitions: MetricConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for kafkametrics metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
