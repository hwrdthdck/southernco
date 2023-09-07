// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetHostArch("host.arch-val")
			rb.SetHostID("host.id-val")
			rb.SetHostName("host.name-val")
			rb.SetOsDescription("os.description-val")
			rb.SetOsType("os.type-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 2, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 5, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("host.arch")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.arch-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.id")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "host.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("os.description")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "os.description-val", val.Str())
			}
			val, ok = res.Attributes().Get("os.type")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "os.type-val", val.Str())
			}
		})
	}
}
