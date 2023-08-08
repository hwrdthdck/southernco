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

// SetPostgresqlDatabaseName sets provided value as "postgresql.database.name" attribute.
func (rb *ResourceBuilder) SetPostgresqlDatabaseName(val string) {
	if rb.config.PostgresqlDatabaseName.Enabled {
		rb.res.Attributes().PutStr("postgresql.database.name", val)
	}
}

// SetPostgresqlIndexName sets provided value as "postgresql.index.name" attribute.
func (rb *ResourceBuilder) SetPostgresqlIndexName(val string) {
	if rb.config.PostgresqlIndexName.Enabled {
		rb.res.Attributes().PutStr("postgresql.index.name", val)
	}
}

// SetPostgresqlTableName sets provided value as "postgresql.table.name" attribute.
func (rb *ResourceBuilder) SetPostgresqlTableName(val string) {
	if rb.config.PostgresqlTableName.Enabled {
		rb.res.Attributes().PutStr("postgresql.table.name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
