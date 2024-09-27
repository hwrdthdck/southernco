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
			rb.SetAerospikeNamespace("aerospike.namespace-val")
			rb.SetAerospikeNodeName("aerospike.node.name-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch tt {
			case "default":
				assert.Equal(t, 2, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 2, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", tt)
			}

			val, ok := res.Attributes().Get("aerospike.namespace")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "aerospike.namespace-val", val.Str())
			}
			val, ok = res.Attributes().Get("aerospike.node.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "aerospike.node.name-val", val.Str())
			}
		})
	}
}
