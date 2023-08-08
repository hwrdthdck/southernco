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

// SetContainerCommandLine sets provided value as "container.command_line" attribute.
func (rb *ResourceBuilder) SetContainerCommandLine(val string) {
	if rb.config.ContainerCommandLine.Enabled {
		rb.res.Attributes().PutStr("container.command_line", val)
	}
}

// SetContainerHostname sets provided value as "container.hostname" attribute.
func (rb *ResourceBuilder) SetContainerHostname(val string) {
	if rb.config.ContainerHostname.Enabled {
		rb.res.Attributes().PutStr("container.hostname", val)
	}
}

// SetContainerID sets provided value as "container.id" attribute.
func (rb *ResourceBuilder) SetContainerID(val string) {
	if rb.config.ContainerID.Enabled {
		rb.res.Attributes().PutStr("container.id", val)
	}
}

// SetContainerImageID sets provided value as "container.image.id" attribute.
func (rb *ResourceBuilder) SetContainerImageID(val string) {
	if rb.config.ContainerImageID.Enabled {
		rb.res.Attributes().PutStr("container.image.id", val)
	}
}

// SetContainerImageName sets provided value as "container.image.name" attribute.
func (rb *ResourceBuilder) SetContainerImageName(val string) {
	if rb.config.ContainerImageName.Enabled {
		rb.res.Attributes().PutStr("container.image.name", val)
	}
}

// SetContainerName sets provided value as "container.name" attribute.
func (rb *ResourceBuilder) SetContainerName(val string) {
	if rb.config.ContainerName.Enabled {
		rb.res.Attributes().PutStr("container.name", val)
	}
}

// SetContainerRuntime sets provided value as "container.runtime" attribute.
func (rb *ResourceBuilder) SetContainerRuntime(val string) {
	if rb.config.ContainerRuntime.Enabled {
		rb.res.Attributes().PutStr("container.runtime", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
