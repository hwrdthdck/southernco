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

// MetricsConfig provides config for k8s/job metrics.
type MetricsConfig struct {
	K8sDaemonsetCurrentScheduledNodes MetricConfig `mapstructure:"k8s.daemonset.current_scheduled_nodes"`
	K8sDaemonsetDesiredScheduledNodes MetricConfig `mapstructure:"k8s.daemonset.desired_scheduled_nodes"`
	K8sDaemonsetMisscheduledNodes     MetricConfig `mapstructure:"k8s.daemonset.misscheduled_nodes"`
	K8sDaemonsetReadyNodes            MetricConfig `mapstructure:"k8s.daemonset.ready_nodes"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		K8sDaemonsetCurrentScheduledNodes: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetDesiredScheduledNodes: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetMisscheduledNodes: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetReadyNodes: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for k8s/job resource attributes.
type ResourceAttributesConfig struct {
	K8sDaemonsetName       ResourceAttributeConfig `mapstructure:"k8s.daemonset.name"`
	K8sDaemonsetUID        ResourceAttributeConfig `mapstructure:"k8s.daemonset.uid"`
	K8sNamespaceName       ResourceAttributeConfig `mapstructure:"k8s.namespace.name"`
	OpencensusResourcetype ResourceAttributeConfig `mapstructure:"opencensus.resourcetype"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		K8sDaemonsetName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sDaemonsetUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNamespaceName: ResourceAttributeConfig{
			Enabled: true,
		},
		OpencensusResourcetype: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for k8s/job metrics builder.
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
