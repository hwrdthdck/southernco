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

// SetHaproxyAddr sets provided value as "haproxy.addr" attribute.
func (rb *ResourceBuilder) SetHaproxyAddr(val string) {
	if rb.config.HaproxyAddr.Enabled {
		rb.res.Attributes().PutStr("haproxy.addr", val)
	}
}

// SetHaproxyAlgo sets provided value as "haproxy.algo" attribute.
func (rb *ResourceBuilder) SetHaproxyAlgo(val string) {
	if rb.config.HaproxyAlgo.Enabled {
		rb.res.Attributes().PutStr("haproxy.algo", val)
	}
}

// SetHaproxyIid sets provided value as "haproxy.iid" attribute.
func (rb *ResourceBuilder) SetHaproxyIid(val string) {
	if rb.config.HaproxyIid.Enabled {
		rb.res.Attributes().PutStr("haproxy.iid", val)
	}
}

// SetHaproxyPid sets provided value as "haproxy.pid" attribute.
func (rb *ResourceBuilder) SetHaproxyPid(val string) {
	if rb.config.HaproxyPid.Enabled {
		rb.res.Attributes().PutStr("haproxy.pid", val)
	}
}

// SetHaproxySid sets provided value as "haproxy.sid" attribute.
func (rb *ResourceBuilder) SetHaproxySid(val string) {
	if rb.config.HaproxySid.Enabled {
		rb.res.Attributes().PutStr("haproxy.sid", val)
	}
}

// SetHaproxyType sets provided value as "haproxy.type" attribute.
func (rb *ResourceBuilder) SetHaproxyType(val string) {
	if rb.config.HaproxyType.Enabled {
		rb.res.Attributes().PutStr("haproxy.type", val)
	}
}

// SetHaproxyURL sets provided value as "haproxy.url" attribute.
func (rb *ResourceBuilder) SetHaproxyURL(val string) {
	if rb.config.HaproxyURL.Enabled {
		rb.res.Attributes().PutStr("haproxy.url", val)
	}
}

// SetProxyName sets provided value as "proxy_name" attribute.
func (rb *ResourceBuilder) SetProxyName(val string) {
	if rb.config.ProxyName.Enabled {
		rb.res.Attributes().PutStr("proxy_name", val)
	}
}

// SetServiceName sets provided value as "service_name" attribute.
func (rb *ResourceBuilder) SetServiceName(val string) {
	if rb.config.ServiceName.Enabled {
		rb.res.Attributes().PutStr("service_name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
