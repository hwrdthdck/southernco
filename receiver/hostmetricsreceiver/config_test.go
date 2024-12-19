// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostmetricsreceiver

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v4/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/networkscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/systemscraper"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	tests := []struct {
		id       component.ID
		expected component.Config
	}{
		{
			id: component.NewID(metadata.Type),
			expected: func() component.Config {
				cfg := createDefaultConfig().(*Config)
				cfg.Scrapers = map[string]internal.Config{
					cpuscraper.TypeStr: func() internal.Config {
						cfg := (&cpuscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
				}
				return cfg
			}(),
		},
		{
			id: component.NewIDWithName(metadata.Type, "customname"),
			expected: &Config{
				MetadataCollectionInterval: 5 * time.Minute,
				ControllerConfig: scraperhelper.ControllerConfig{
					CollectionInterval: 30 * time.Second,
					InitialDelay:       time.Second,
				},
				Scrapers: map[string]internal.Config{
					cpuscraper.TypeStr: func() internal.Config {
						cfg := (&cpuscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
					diskscraper.TypeStr: func() internal.Config {
						cfg := (&diskscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
					loadscraper.TypeStr: (func() internal.Config {
						cfg := (&loadscraper.Factory{}).CreateDefaultConfig()
						cfg.(*loadscraper.Config).CPUAverage = true
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					})(),
					filesystemscraper.TypeStr: func() internal.Config {
						cfg := (&filesystemscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
					memoryscraper.TypeStr: func() internal.Config {
						cfg := (&memoryscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
					networkscraper.TypeStr: (func() internal.Config {
						cfg := (&networkscraper.Factory{}).CreateDefaultConfig()
						cfg.(*networkscraper.Config).Include = networkscraper.MatchConfig{
							Interfaces: []string{"test1"},
							Config:     filterset.Config{MatchType: "strict"},
						}
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					})(),
					processesscraper.TypeStr: func() internal.Config {
						cfg := (&processesscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
					pagingscraper.TypeStr: func() internal.Config {
						cfg := (&pagingscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					}(),
					processscraper.TypeStr: (func() internal.Config {
						cfg := (&processscraper.Factory{}).CreateDefaultConfig()
						cfg.(*processscraper.Config).Include = processscraper.MatchConfig{
							Names:  []string{"test2", "test3"},
							Config: filterset.Config{MatchType: "regexp"},
						}
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					})(),
					systemscraper.TypeStr: (func() internal.Config {
						cfg := (&systemscraper.Factory{}).CreateDefaultConfig()
						cfg.SetEnvMap(common.EnvMap{})
						return cfg
					})(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, sub.Unmarshal(cfg))

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestLoadInvalidConfig_NoScrapers(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config-noscrapers.yaml"))
	require.NoError(t, err)

	require.NoError(t, cm.Unmarshal(cfg))
	require.ErrorContains(t, component.ValidateConfig(cfg), "must specify at least one scraper when using hostmetrics receiver")
}

func TestLoadInvalidConfig_InvalidScraperKey(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config-invalidscraperkey.yaml"))
	require.NoError(t, err)

	require.ErrorContains(t, cm.Unmarshal(cfg), "invalid scraper key: invalidscraperkey")
}
