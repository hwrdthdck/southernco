// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, tt := range []string{"default", "all_set", "none_set"} {
		t.Run(tt, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt)
			rb := NewResourceBuilder(cfg)
			rb.SetPostgresqlDatabaseName("postgresql.database.name-val")
			rb.SetPostgresqlIndexName("postgresql.index.name-val")
			rb.SetPostgresqlSchemaName("postgresql.schema.name-val")
			rb.SetPostgresqlTableName("postgresql.table.name-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch tt {
			case "default":
				assert.Equal(t, 4, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 4, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", tt)
			}

			val, ok := res.Attributes().Get("postgresql.database.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "postgresql.database.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("postgresql.index.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "postgresql.index.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("postgresql.schema.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "postgresql.schema.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("postgresql.table.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "postgresql.table.name-val", val.Str())
			}
		})
	}
}
