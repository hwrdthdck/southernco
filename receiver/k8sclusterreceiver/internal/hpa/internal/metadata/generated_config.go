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

// MetricsConfig provides config for k8s/hpa metrics.
type MetricsConfig struct {
	K8sHpaCurrentReplicas MetricConfig `mapstructure:"k8s.hpa.current_replicas"`
	K8sHpaDesiredReplicas MetricConfig `mapstructure:"k8s.hpa.desired_replicas"`
	K8sHpaMaxReplicas     MetricConfig `mapstructure:"k8s.hpa.max_replicas"`
	K8sHpaMinReplicas     MetricConfig `mapstructure:"k8s.hpa.min_replicas"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		K8sHpaCurrentReplicas: MetricConfig{
			Enabled: true,
		},
		K8sHpaDesiredReplicas: MetricConfig{
			Enabled: true,
		},
		K8sHpaMaxReplicas: MetricConfig{
			Enabled: true,
		},
		K8sHpaMinReplicas: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for k8s/hpa resource attributes.
type ResourceAttributesConfig struct {
	K8sHpaName       ResourceAttributeConfig `mapstructure:"k8s.hpa.name"`
	K8sHpaUID        ResourceAttributeConfig `mapstructure:"k8s.hpa.uid"`
	K8sNamespaceName ResourceAttributeConfig `mapstructure:"k8s.namespace.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		K8sHpaName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sHpaUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNamespaceName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for k8s/hpa metrics builder.
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
