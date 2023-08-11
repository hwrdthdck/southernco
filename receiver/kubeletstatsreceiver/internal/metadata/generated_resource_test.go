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
			rb.SetAwsVolumeID("aws.volume.id-val")
			rb.SetContainerID("container.id-val")
			rb.SetCsiDriver("csi.driver-val")
			rb.SetCsiVolumeHandle("csi.volume.handle-val")
			rb.SetFsType("fs.type-val")
			rb.SetGcePdName("gce.pd.name-val")
			rb.SetGlusterfsEndpointsName("glusterfs.endpoints.name-val")
			rb.SetGlusterfsPath("glusterfs.path-val")
			rb.SetK8sContainerName("k8s.container.name-val")
			rb.SetK8sNamespaceName("k8s.namespace.name-val")
			rb.SetK8sNodeName("k8s.node.name-val")
			rb.SetK8sPersistentvolumeclaimName("k8s.persistentvolumeclaim.name-val")
			rb.SetK8sPodName("k8s.pod.name-val")
			rb.SetK8sPodUID("k8s.pod.uid-val")
			rb.SetK8sVolumeName("k8s.volume.name-val")
			rb.SetK8sVolumeType("k8s.volume.type-val")
			rb.SetPartition("partition-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return 0

			switch test {
			case "default":
				assert.Equal(t, 17, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 17, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("aws.volume.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "aws.volume.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "container.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("csi.driver")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "csi.driver-val", val.Str())
			}
			val, ok = res.Attributes().Get("csi.volume.handle")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "csi.volume.handle-val", val.Str())
			}
			val, ok = res.Attributes().Get("fs.type")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "fs.type-val", val.Str())
			}
			val, ok = res.Attributes().Get("gce.pd.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "gce.pd.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("glusterfs.endpoints.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "glusterfs.endpoints.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("glusterfs.path")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "glusterfs.path-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.container.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.container.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.namespace.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.namespace.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.node.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.node.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.persistentvolumeclaim.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.persistentvolumeclaim.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.pod.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.pod.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.pod.uid")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.pod.uid-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.volume.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.volume.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.volume.type")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.volume.type-val", val.Str())
			}
			val, ok = res.Attributes().Get("partition")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "partition-val", val.Str())
			}
		})
	}
}
