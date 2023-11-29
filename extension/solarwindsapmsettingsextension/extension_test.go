package solarwindsapmsettingsextension

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"
	"testing"
)

func TestCreateExtension(t *testing.T) {
	conf := &Config{
		Endpoint: "apm-testcollector.click:443",
		Key:      "valid:unittest",
		Interval: "1s",
	}
	ex := createAnExtension(conf, t)
	ex.Shutdown(context.TODO())
}

func TestCreateExtensionWrongEndpoint(t *testing.T) {
	conf := &Config{
		Endpoint: "apm-testcollector.nothing:443",
		Key:      "valid:unittest",
		Interval: "1s",
	}
	ex := createAnExtension(conf, t)
	ex.Shutdown(context.TODO())
}

func TestCreateExtensionWrongKey(t *testing.T) {
	conf := &Config{
		Endpoint: "apm-testcollector.click:443",
		Key:      "invalid",
		Interval: "1s",
	}
	ex := createAnExtension(conf, t)
	ex.Shutdown(context.TODO())
}

// create extension
func createAnExtension(c *Config, t *testing.T) extension.Extension {
	logger, err := zap.NewProduction()
	ex, err := newSolarwindsApmSettingsExtension(c, logger)
	require.NoError(t, err)
	err = ex.Start(context.TODO(), nil)
	require.NoError(t, err)
	return ex
}
