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
			rb.SetBigipNodeIPAddress("bigip.node.ip_address-val")
			rb.SetBigipNodeName("bigip.node.name-val")
			rb.SetBigipPoolName("bigip.pool.name-val")
			rb.SetBigipPoolMemberIPAddress("bigip.pool_member.ip_address-val")
			rb.SetBigipPoolMemberName("bigip.pool_member.name-val")
			rb.SetBigipVirtualServerDestination("bigip.virtual_server.destination-val")
			rb.SetBigipVirtualServerName("bigip.virtual_server.name-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 7, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 7, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("bigip.node.ip_address")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.node.ip_address-val", val.Str())
			}
			val, ok = res.Attributes().Get("bigip.node.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.node.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("bigip.pool.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.pool.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("bigip.pool_member.ip_address")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.pool_member.ip_address-val", val.Str())
			}
			val, ok = res.Attributes().Get("bigip.pool_member.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.pool_member.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("bigip.virtual_server.destination")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.virtual_server.destination-val", val.Str())
			}
			val, ok = res.Attributes().Get("bigip.virtual_server.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "bigip.virtual_server.name-val", val.Str())
			}
		})
	}
}
