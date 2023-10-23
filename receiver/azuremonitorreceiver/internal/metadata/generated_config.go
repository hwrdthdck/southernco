// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

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

// ResourceAttributesConfig provides config for azuremonitor resource attributes.
type ResourceAttributesConfig struct {
	AzuremonitorSubscriptionID ResourceAttributeConfig `mapstructure:"azuremonitor.subscription_id"`
	AzuremonitorTenantID       ResourceAttributeConfig `mapstructure:"azuremonitor.tenant_id"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		AzuremonitorSubscriptionID: ResourceAttributeConfig{
			Enabled: false,
		},
		AzuremonitorTenantID: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}
