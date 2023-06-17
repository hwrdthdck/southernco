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

// MetricsConfig provides config for k8s/namespace metrics.
type MetricsConfig struct {
	K8sNamespacePhase MetricConfig `mapstructure:"k8s.namespace.phase"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		K8sNamespacePhase: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for k8s/namespace resource attributes.
type ResourceAttributesConfig struct {
	K8sNamespaceName       ResourceAttributeConfig `mapstructure:"k8s.namespace.name"`
	K8sNamespaceUID        ResourceAttributeConfig `mapstructure:"k8s.namespace.uid"`
	OpencensusResourcetype ResourceAttributeConfig `mapstructure:"opencensus.resourcetype"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		K8sNamespaceName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNamespaceUID: ResourceAttributeConfig{
			Enabled: true,
		},
		OpencensusResourcetype: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for k8s/namespace metrics builder.
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
