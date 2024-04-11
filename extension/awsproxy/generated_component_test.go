// Code generated by mdatagen. DO NOT EDIT.

package awsproxy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/extension/extensiontest"
)

func TestComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	cm, err := confmaptest.LoadConf("metadata.yaml")
	require.NoError(t, err)
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub("tests::config")
	require.NoError(t, err)
	require.NoError(t, component.UnmarshalConfig(sub, cfg))
	t.Run("shutdown", func(t *testing.T) {
		e, err := factory.CreateExtension(context.Background(), extensiontest.NewNopCreateSettings(), cfg)
		require.NoError(t, err)
		err = e.Shutdown(context.Background())
		require.NoError(t, err)
	})
}
