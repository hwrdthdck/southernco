// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package vcenterreceiver // import github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/debug"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/scraperinttest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal"
)

func TestIntegration(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		pw, set := simulator.DefaultLogin.Password()
		require.True(t, set)

		s := session.NewManager(c)
		newVcenterClient = func(cfg *Config) *vcenterClient {
			client := &vcenterClient{
				cfg: cfg,
				moClient: &govmomi.Client{
					Client:         c,
					SessionManager: s,
				},
			}
			require.NoError(t, client.EnsureConnection(context.Background()))
			client.vimDriver = c
			client.finder = find.NewFinder(c)
			// Performance metrics rely on time based publishing so this is inherently flaky for an
			// integration test, so setting the performance manager to nil to not attempt to compare
			// performance metrics. Coverage for this is encompassed in ./scraper_test.go
			client.pm = nil
			return client
		}
		defer func() {
			newVcenterClient = defaultNewVcenterClient
		}()

		scraperinttest.NewIntegrationTest(
			NewFactory(),
			scraperinttest.WithCustomConfig(
				func(t *testing.T, cfg component.Config, ci *scraperinttest.ContainerInfo) {
					rCfg := cfg.(*Config)
					rCfg.CollectionInterval = 2 * time.Second
					rCfg.Endpoint = fmt.Sprintf("%s://%s", c.URL().Scheme, c.URL().Host)
					rCfg.Username = simulator.DefaultLogin.Username()
					rCfg.Password = configopaque.String(pw)
					rCfg.TLSClientSetting = configtls.TLSClientSetting{
						Insecure: true,
					}
				}),
			scraperinttest.WithCompareOptions(
				pmetrictest.IgnoreResourceAttributeValue("vcenter.host.name"),
				pmetrictest.IgnoreTimestamp(),
				pmetrictest.IgnoreResourceMetricsOrder(),
				pmetrictest.IgnoreMetricsOrder(),
				pmetrictest.IgnoreStartTimestamp(),
				pmetrictest.IgnoreMetricValues(),
			),
		).Run(t)
	})
}

func TestIntegrationDebug(t *testing.T) {
	// As debug.SetProvider is a global configuration setting
	// it must be set before vim25.Client is instantiated
	l := logrus.WithFields(logrus.Fields{})
	debug.SetProvider(internal.NewLogProvider(l))
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		pw, set := simulator.DefaultLogin.Password()
		require.True(t, set)

		s := session.NewManager(c)
		newVcenterClient = func(cfg *Config) *vcenterClient {
			client := &vcenterClient{
				cfg: cfg,
				moClient: &govmomi.Client{
					Client:         c,
					SessionManager: s,
				},
			}
			require.NoError(t, client.EnsureConnection(context.Background()))
			client.vimDriver = c
			client.finder = find.NewFinder(c)
			// Performance metrics rely on time based publishing so this is inherently flaky for an
			// integration test, so setting the performance manager to nil to not attempt to compare
			// performance metrics. Coverage for this is encompassed in ./scraper_test.go
			client.pm = nil
			return client
		}
		defer func() {
			newVcenterClient = defaultNewVcenterClient
		}()

		scraperinttest.NewIntegrationTest(
			NewFactory(),
			scraperinttest.WithCustomConfig(
				func(t *testing.T, cfg component.Config, ci *scraperinttest.ContainerInfo) {
					rCfg := cfg.(*Config)
					rCfg.CollectionInterval = 2 * time.Second
					rCfg.Endpoint = fmt.Sprintf("%s://%s", c.URL().Scheme, c.URL().Host)
					rCfg.Username = simulator.DefaultLogin.Username()
					rCfg.Password = configopaque.String(pw)
					rCfg.TLSClientSetting = configtls.TLSClientSetting{
						Insecure: true,
					}
				}),
			scraperinttest.WithCompareOptions(
				pmetrictest.IgnoreResourceAttributeValue("vcenter.host.name"),
				pmetrictest.IgnoreTimestamp(),
				pmetrictest.IgnoreResourceMetricsOrder(),
				pmetrictest.IgnoreMetricsOrder(),
				pmetrictest.IgnoreStartTimestamp(),
				pmetrictest.IgnoreMetricValues(),
			),
		).Run(t)
	})
}
