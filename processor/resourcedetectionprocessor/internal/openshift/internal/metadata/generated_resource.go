// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetCloudPlatform sets provided value as "cloud.platform" attribute.
func (rb *ResourceBuilder) SetCloudPlatform(val string) {
	if rb.config.CloudPlatform.Enabled {
		rb.res.Attributes().PutStr("cloud.platform", val)
	}
}

// SetCloudProvider sets provided value as "cloud.provider" attribute.
func (rb *ResourceBuilder) SetCloudProvider(val string) {
	if rb.config.CloudProvider.Enabled {
		rb.res.Attributes().PutStr("cloud.provider", val)
	}
}

// SetCloudRegion sets provided value as "cloud.region" attribute.
func (rb *ResourceBuilder) SetCloudRegion(val string) {
	if rb.config.CloudRegion.Enabled {
		rb.res.Attributes().PutStr("cloud.region", val)
	}
}

// SetK8sClusterName sets provided value as "k8s.cluster.name" attribute.
func (rb *ResourceBuilder) SetK8sClusterName(val string) {
	if rb.config.K8sClusterName.Enabled {
		rb.res.Attributes().PutStr("k8s.cluster.name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
