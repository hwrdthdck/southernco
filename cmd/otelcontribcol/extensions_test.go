// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/ballastextension"
	"go.opentelemetry.io/collector/extension/extensiontest"
	"go.opentelemetry.io/collector/extension/zpagesextension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/googleclientauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckv2extension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarderextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/testutil"
)

func TestDefaultExtensions(t *testing.T) {
	allFactories, err := components()
	require.NoError(t, err)

	extFactories := allFactories.Extensions
	endpoint := testutil.GetAvailableLocalAddress(t)

	tests := []struct {
		getConfigFn   getExtensionConfigFn
		extension     component.Type
		skipLifecycle bool
	}{
		{
			extension: "health_check",
			getConfigFn: func() component.Config {
				cfg := extFactories["health_check"].CreateDefaultConfig().(*healthcheckextension.Config)
				cfg.Endpoint = endpoint
				return cfg
			},
		},
		{
			extension: "healthcheckv2",
			getConfigFn: func() component.Config {
				cfg := extFactories["healthcheckv2"].CreateDefaultConfig().(*healthcheckv2extension.Config)
				cfg.Endpoint = endpoint
				return cfg
			},
		},
		{
			extension: "pprof",
			getConfigFn: func() component.Config {
				cfg := extFactories["pprof"].CreateDefaultConfig().(*pprofextension.Config)
				cfg.TCPAddr.Endpoint = endpoint
				return cfg
			},
		},
		{
			extension: "sigv4auth",
			getConfigFn: func() component.Config {
				cfg := extFactories["sigv4auth"].CreateDefaultConfig().(*sigv4authextension.Config)
				return cfg
			},
		},
		{
			extension: "zpages",
			getConfigFn: func() component.Config {
				cfg := extFactories["zpages"].CreateDefaultConfig().(*zpagesextension.Config)
				cfg.TCPAddr.Endpoint = endpoint
				return cfg
			},
		},
		{
			extension: "basicauth",
			getConfigFn: func() component.Config {
				cfg := extFactories["basicauth"].CreateDefaultConfig().(*basicauthextension.Config)
				// No need to clean up, t.TempDir will be deleted entirely.
				fileName := filepath.Join(t.TempDir(), "random.file")
				require.NoError(t, os.WriteFile(fileName, []byte("username:password"), 0600))

				cfg.Htpasswd = &basicauthextension.HtpasswdSettings{
					File:   fileName,
					Inline: "username:password",
				}
				return cfg
			},
		},
		{
			extension: "bearertokenauth",
			getConfigFn: func() component.Config {
				cfg := extFactories["bearertokenauth"].CreateDefaultConfig().(*bearertokenauthextension.Config)
				cfg.BearerToken = "sometoken"
				return cfg
			},
		},
		{
			extension: "memory_ballast",
			getConfigFn: func() component.Config {
				cfg := extFactories["memory_ballast"].CreateDefaultConfig().(*ballastextension.Config)
				return cfg
			},
		},
		{
			extension: "asapclient",
			getConfigFn: func() component.Config {
				cfg := extFactories["asapclient"].CreateDefaultConfig().(*asapauthextension.Config)
				cfg.KeyID = "test_issuer/test_kid"
				cfg.Issuer = "test_issuer"
				cfg.Audience = []string{"some_service"}
				cfg.TTL = 10 * time.Second
				// Valid PEM data required for successful initialisation. Key not actually used anywhere.
				cfg.PrivateKey = "data:application/pkcs8;kid=test;base64,MIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgE" +
					"AAkEA0ZPr5JeyVDoB8RyZqQsx6qUD+9gMFg1/0hgdAvmytWBMXQJYdwkK2dFJwwZcWJVhJGcOJBDfB/8tcbdJd34KZQIDAQ" +
					"ABAkBZD20tJTHJDSWKGsdJyNIbjqhUu4jXTkFFPK4Hd6jz3gV3fFvGnaolsD5Bt50dTXAiSCpFNSb9M9GY6XUAAdlBAiEA6" +
					"MccfdZRfVapxKtAZbjXuAgMvnPtTvkVmwvhWLT5Wy0CIQDmfE8Et/pou0Jl6eM0eniT8/8oRzBWgy9ejDGfj86PGQIgWePq" +
					"IL4OofRBgu0O5TlINI0HPtTNo12U9lbUIslgMdECICXT2RQpLcvqj+cyD7wZLZj6vrHZnTFVrnyR/cL2UyxhAiBswe/MCcD" +
					"7T7J4QkNrCG+ceQGypc7LsxlIxQuKh5GWYA=="
				return cfg
			},
		},
		{
			extension:     "ecs_observer",
			skipLifecycle: true,
		},
		{
			extension: "ecs_task_observer",
			getConfigFn: func() component.Config {
				cfg := extFactories["ecs_task_observer"].CreateDefaultConfig().(*ecstaskobserver.Config)
				cfg.Endpoint = "http://localhost"
				return cfg
			},
		},
		{
			extension:     "awsproxy",
			skipLifecycle: true, // Requires EC2 metadata service to be running
		},
		{
			extension: "http_forwarder",
			getConfigFn: func() component.Config {
				cfg := extFactories["http_forwarder"].CreateDefaultConfig().(*httpforwarderextension.Config)
				cfg.Egress.Endpoint = "http://" + endpoint
				cfg.Ingress.Endpoint = testutil.GetAvailableLocalAddress(t)
				return cfg
			},
		},
		{
			extension: "oauth2client",
			getConfigFn: func() component.Config {
				cfg := extFactories["oauth2client"].CreateDefaultConfig().(*oauth2clientauthextension.Config)
				cfg.ClientID = "otel-extension"
				cfg.ClientSecret = "testsarehard"
				cfg.TokenURL = "http://" + endpoint
				return cfg
			},
		},
		{
			extension:     "oidc",
			skipLifecycle: true, // Requires a running OIDC server in order to complete life cycle testing
		},
		{
			extension: "db_storage",
			getConfigFn: func() component.Config {
				cfg := extFactories["db_storage"].CreateDefaultConfig().(*dbstorage.Config)
				cfg.DriverName = "sqlite3"
				cfg.DataSource = filepath.Join(t.TempDir(), "foo.db")
				return cfg
			},
		},
		{
			extension: "file_storage",
			getConfigFn: func() component.Config {
				cfg := extFactories["file_storage"].CreateDefaultConfig().(*filestorage.Config)
				cfg.Directory = t.TempDir()
				return cfg
			},
		},
		{
			extension: "host_observer",
			getConfigFn: func() component.Config {
				cfg := extFactories["host_observer"].CreateDefaultConfig().(*hostobserver.Config)
				return cfg
			},
		},
		{
			extension:     "k8s_observer",
			skipLifecycle: true, // Requires a K8s api to interfact with and validate
		},
		{
			extension:     "docker_observer",
			skipLifecycle: true, // Requires a docker api to interface and validate.
		},
		{
			extension: "headers_setter",
			getConfigFn: func() component.Config {
				cfg := extFactories["headers_setter"].CreateDefaultConfig().(*headerssetterextension.Config)
				return cfg
			},
		},
		{
			extension:     "jaegerremotesampling",
			skipLifecycle: true,
			getConfigFn: func() component.Config {
				return extFactories["jaegerremotesampling"].CreateDefaultConfig().(*jaegerremotesampling.Config)
			},
		},
		{
			extension: "otlp_encoding",
		},
		{
			extension: "text_encoding",
		},
		{
			extension: "jaeger_encoding",
		},
		{
			extension: "json_log_encoding",
		},
		{
			extension: "zipkin_encoding",
		},
		{
			extension: "remotetap",
			getConfigFn: func() component.Config {
				return extFactories["remotetap"].CreateDefaultConfig().(*remotetapextension.Config)
			},
		},
		{
			extension: "opamp",
			getConfigFn: func() component.Config {
				cfg := extFactories["opamp"].CreateDefaultConfig().(*opampextension.Config)
				cfg.Server.WS.Endpoint = "wss://" + endpoint
				return cfg
			},
		},
		{
			extension:     "solarwindsapmsettings",
			skipLifecycle: true, // Requires Solarwinds APM endpoint and token
		},
		{
			extension: "ackextension",
			getConfigFn: func() component.Config {
				return extFactories["ackextension"].CreateDefaultConfig().(*ackextension.Config)
			},
		},
		{
			extension: "googleclientauthextension",
			getConfigFn: func() component.Config {
				return extFactories["googleclientauthextension"].CreateDefaultConfig().(*googleclientauthextension.Config)
			},
			skipLifecycle: true,
		},
	}

	extensionCount := 0
	expectedExtensions := map[component.Type]struct{}{}
	for k := range extFactories {
		expectedExtensions[k] = struct{}{}
	}
	for _, tt := range tests {
		_, ok := extFactories[tt.extension]
		if !ok {
			// not part of the distro, skipping.
			continue
		}
		delete(expectedExtensions, tt.extension)
		extensionCount++
		t.Run(string(tt.extension), func(t *testing.T) {
			factory := extFactories[tt.extension]
			assert.Equal(t, tt.extension, factory.Type())

			t.Run("shutdown", func(t *testing.T) {
				verifyExtensionShutdown(t, factory, tt.getConfigFn)
			})
			t.Run("lifecycle", func(t *testing.T) {
				if tt.skipLifecycle {
					t.SkipNow()
				}
				verifyExtensionLifecycle(t, factory, tt.getConfigFn)
			})

		})
	}
	assert.Len(t, extFactories, extensionCount, "All extensions must be added to the lifecycle tests", expectedExtensions)
}

// getExtensionConfigFn is used customize the configuration passed to the verification.
// This is used to change ports or provide values required but not provided by the
// default configuration.
type getExtensionConfigFn func() component.Config

// verifyExtensionLifecycle is used to test if an extension type can handle the typical
// lifecycle of a component. The getConfigFn parameter only need to be specified if
// the test can't be done with the default configuration for the component.
func verifyExtensionLifecycle(t *testing.T, factory extension.Factory, getConfigFn getExtensionConfigFn) {
	ctx := context.Background()
	host := componenttest.NewNopHost()
	extCreateSet := extensiontest.NewNopSettings()
	extCreateSet.ReportStatus = func(event *component.StatusEvent) {
		require.NoError(t, event.Err())
	}

	if getConfigFn == nil {
		getConfigFn = factory.CreateDefaultConfig
	}

	firstExt, err := factory.CreateExtension(ctx, extCreateSet, getConfigFn())
	require.NoError(t, err)
	require.NoError(t, firstExt.Start(ctx, host))
	require.NoError(t, firstExt.Shutdown(ctx))

	secondExt, err := factory.CreateExtension(ctx, extCreateSet, getConfigFn())
	require.NoError(t, err)
	require.NoError(t, secondExt.Start(ctx, host))
	require.NoError(t, secondExt.Shutdown(ctx))
}

// verifyExtensionShutdown is used to test if an extension type can be shutdown without being started first.
func verifyExtensionShutdown(tb testing.TB, factory extension.Factory, getConfigFn getExtensionConfigFn) {
	ctx := context.Background()
	extCreateSet := extensiontest.NewNopSettings()

	if getConfigFn == nil {
		getConfigFn = factory.CreateDefaultConfig
	}

	e, err := factory.CreateExtension(ctx, extCreateSet, getConfigFn())
	if errors.Is(err, component.ErrDataTypeIsNotSupported) {
		return
	}
	if e == nil {
		return
	}

	assert.NotPanics(tb, func() {
		assert.NoError(tb, e.Shutdown(ctx))
	})
}
