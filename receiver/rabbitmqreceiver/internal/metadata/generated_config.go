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

// MetricsSettings provides settings for rabbitmqreceiver metrics.
type MetricsSettings struct {
	RabbitmqConsumerCount       MetricSettings `mapstructure:"rabbitmq.consumer.count"`
	RabbitmqMessageAcknowledged MetricSettings `mapstructure:"rabbitmq.message.acknowledged"`
	RabbitmqMessageCurrent      MetricSettings `mapstructure:"rabbitmq.message.current"`
	RabbitmqMessageDelivered    MetricSettings `mapstructure:"rabbitmq.message.delivered"`
	RabbitmqMessageDropped      MetricSettings `mapstructure:"rabbitmq.message.dropped"`
	RabbitmqMessagePublished    MetricSettings `mapstructure:"rabbitmq.message.published"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		RabbitmqConsumerCount: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageAcknowledged: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageCurrent: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageDelivered: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageDropped: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessagePublished: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular resource attribute.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesSettings provides settings for rabbitmqreceiver resource attributes.
type ResourceAttributesSettings struct {
	RabbitmqNodeName  ResourceAttributeSettings `mapstructure:"rabbitmq.node.name"`
	RabbitmqQueueName ResourceAttributeSettings `mapstructure:"rabbitmq.queue.name"`
	RabbitmqVhostName ResourceAttributeSettings `mapstructure:"rabbitmq.vhost.name"`
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{
		RabbitmqNodeName: ResourceAttributeSettings{
			Enabled: true,
		},
		RabbitmqQueueName: ResourceAttributeSettings{
			Enabled: true,
		},
		RabbitmqVhostName: ResourceAttributeSettings{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for rabbitmqreceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsSettings            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesSettings `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsSettings(),
		ResourceAttributes: DefaultResourceAttributesSettings(),
	}
}
