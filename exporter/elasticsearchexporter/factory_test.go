// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestFactory_CreateLogsExporter(t *testing.T) {
	factory := NewFactory()
	cfg := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoints = []string{"http://test:9200"}
	})
	params := exportertest.NewNopSettings()
	exporter, err := factory.CreateLogsExporter(context.Background(), params, cfg)
	require.NoError(t, err)
	require.NotNil(t, exporter)
	require.NoError(t, exporter.Start(context.Background(), componenttest.NewNopHost()))

	require.NoError(t, exporter.Shutdown(context.Background()))
}

func TestFactory_CreateLogsExporter_Fail(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	params := exportertest.NewNopSettings()
	_, err := factory.CreateLogsExporter(context.Background(), params, cfg)
	require.Error(t, err, "expected an error when creating a logs exporter")
	assert.EqualError(t, err, "cannot configure Elasticsearch exporter: exactly one of [endpoint, endpoints, cloudid] must be specified")
}

func TestFactory_CreateMetricsExporter(t *testing.T) {
	factory := NewFactory()
	cfg := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoints = []string{"http://test:9200"}
	})
	params := exportertest.NewNopSettings()
	exporter, err := factory.CreateMetricsExporter(context.Background(), params, cfg)
	require.NoError(t, err)
	require.NotNil(t, exporter)

	require.NoError(t, exporter.Shutdown(context.Background()))
}

func TestFactory_CreateTracesExporter(t *testing.T) {
	factory := NewFactory()
	cfg := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoints = []string{"http://test:9200"}
	})
	params := exportertest.NewNopSettings()
	exporter, err := factory.CreateTracesExporter(context.Background(), params, cfg)
	require.NoError(t, err)
	require.NotNil(t, exporter)
	require.NoError(t, exporter.Start(context.Background(), componenttest.NewNopHost()))

	require.NoError(t, exporter.Shutdown(context.Background()))
}

func TestFactory_CreateTracesExporter_Fail(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	params := exportertest.NewNopSettings()
	_, err := factory.CreateTracesExporter(context.Background(), params, cfg)
	require.Error(t, err, "expected an error when creating a traces exporter")
	assert.EqualError(t, err, "cannot configure Elasticsearch exporter: exactly one of [endpoint, endpoints, cloudid] must be specified")
}

func TestFactory_CreateLogsAndTracesExporterWithDeprecatedIndexOption(t *testing.T) {
	factory := NewFactory()
	cfg := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoints = []string{"http://test:9200"}
		cfg.Index = "test_index"
	})
	params := exportertest.NewNopSettings()
	logsExporter, err := factory.CreateLogsExporter(context.Background(), params, cfg)
	require.NoError(t, err)
	require.NotNil(t, logsExporter)
	require.NoError(t, logsExporter.Start(context.Background(), componenttest.NewNopHost()))
	require.NoError(t, logsExporter.Shutdown(context.Background()))

	tracesExporter, err := factory.CreateTracesExporter(context.Background(), params, cfg)
	require.NoError(t, err)
	require.NotNil(t, tracesExporter)
	require.NoError(t, tracesExporter.Start(context.Background(), componenttest.NewNopHost()))
	require.NoError(t, tracesExporter.Shutdown(context.Background()))
}

func TestFactory_DedupDeprecated(t *testing.T) {
	factory := NewFactory()
	cfg := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoint = "http://testing.invalid:9200"
		cfg.Mapping.Dedup = false
		cfg.Mapping.Dedot = false // avoid dedot warnings
	})

	loggerCore, logObserver := observer.New(zap.WarnLevel)
	set := exportertest.NewNopSettings()
	set.Logger = zap.New(loggerCore)

	logsExporter, err := factory.CreateLogsExporter(context.Background(), set, cfg)
	require.NoError(t, err)
	require.NoError(t, logsExporter.Shutdown(context.Background()))

	tracesExporter, err := factory.CreateTracesExporter(context.Background(), set, cfg)
	require.NoError(t, err)
	require.NoError(t, tracesExporter.Shutdown(context.Background()))

	metricsExporter, err := factory.CreateMetricsExporter(context.Background(), set, cfg)
	require.NoError(t, err)
	require.NoError(t, metricsExporter.Shutdown(context.Background()))

	records := logObserver.AllUntimed()
	var cnt int
	for _, record := range records {
		if record.Message == "dedup has been deprecated, and will always be enabled in future" {
			cnt++
		}
	}
	assert.Equal(t, 3, cnt)
}

func TestFactory_DedotDeprecated(t *testing.T) {
	loggerCore, logObserver := observer.New(zap.WarnLevel)
	set := exportertest.NewNopSettings()
	set.Logger = zap.New(loggerCore)

	cfgNoDedotECS := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoint = "http://testing.invalid:9200"
		cfg.Mapping.Dedot = false
		cfg.Mapping.Mode = "ecs"
	})

	cfgDedotRaw := withDefaultConfig(func(cfg *Config) {
		cfg.Endpoint = "http://testing.invalid:9200"
		cfg.Mapping.Dedot = true
		cfg.Mapping.Mode = "raw"
	})

	for _, cfg := range []*Config{cfgNoDedotECS, cfgDedotRaw} {
		factory := NewFactory()
		logsExporter, err := factory.CreateLogsExporter(context.Background(), set, cfg)
		require.NoError(t, err)
		require.NoError(t, logsExporter.Shutdown(context.Background()))

		tracesExporter, err := factory.CreateTracesExporter(context.Background(), set, cfg)
		require.NoError(t, err)
		require.NoError(t, tracesExporter.Shutdown(context.Background()))

		metricsExporter, err := factory.CreateMetricsExporter(context.Background(), set, cfg)
		require.NoError(t, err)
		require.NoError(t, metricsExporter.Shutdown(context.Background()))
	}

	records := logObserver.AllUntimed()
	var cnt int
	for _, record := range records {
		if record.Message == "dedot has been deprecated: in the future, dedotting will always be performed in ECS mode only" {
			cnt++
		}
	}
	assert.Equal(t, 6, cnt)
}
