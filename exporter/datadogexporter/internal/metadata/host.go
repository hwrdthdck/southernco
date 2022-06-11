// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/internal/azure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/internal/ec2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/internal/gcp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/internal/system"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/provider"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/valid"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/utils/cache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/docker"
)

// UsePreviewHostnameLogic decides whether to use the preview hostname logic or not.
const UsePreviewHostnameLogic = false

func buildPreviewProvider(set component.TelemetrySettings, configHostname string) (provider.HostnameProvider, error) {
	dockerProvider, err := docker.NewProvider()
	if err != nil {
		return nil, err
	}

	chain, err := provider.Chain(
		set.Logger,
		map[string]provider.HostnameProvider{
			"config": provider.Config(configHostname),
			"docker": dockerProvider,
			"azure":  azure.NewProvider(),
			"ec2":    ec2.NewProvider(set.Logger),
			"gcp":    gcp.NewProvider(),
			"system": system.NewProvider(set.Logger),
		},
		[]string{"config", "docker", "azure", "ec2", "gcp", "system"},
	)

	if err != nil {
		return nil, err
	}

	return provider.Once(chain), nil
}

func buildCurrentProvider(set component.TelemetrySettings, configHostname string) (provider.HostnameProvider, error) {
	return &currentProvider{
		logger:         set.Logger,
		configHostname: configHostname,
		systemProvider: system.NewProvider(set.Logger),
		ec2Provider:    ec2.NewProvider(set.Logger),
	}, nil
}

func GetHostnameProvider(set component.TelemetrySettings, configHostname string) (provider.HostnameProvider, error) {
	if UsePreviewHostnameLogic {
		return buildPreviewProvider(set, configHostname)
	}

	return buildCurrentProvider(set, configHostname)
}

var _ provider.HostnameProvider = (*currentProvider)(nil)

type currentProvider struct {
	logger         *zap.Logger
	configHostname string
	systemProvider *system.Provider
	ec2Provider    *ec2.Provider
}

// Hostname gets the hostname according to configuration.
// It checks in the following order
// 1. Configuration
// 2. Cache
// 3. EC2 instance metadata
// 4. System
func (c *currentProvider) Hostname(ctx context.Context) (string, error) {
	if c.configHostname != "" {
		return c.configHostname, nil
	}

	if cacheVal, ok := cache.Cache.Get(cache.CanonicalHostnameKey); ok {
		return cacheVal.(string), nil
	}

	ec2Info := c.ec2Provider.HostInfo()
	hostname := ec2Info.GetHostname(c.logger)

	if hostname == "" {
		// Get system hostname
		var err error
		hostname, err = c.systemProvider.Hostname(ctx)
		if err != nil {
			c.logger.Debug("system provider is unavailable", zap.Error(err))
		}
	}

	if err := valid.Hostname(hostname); err != nil {
		// If invalid log but continue
		c.logger.Error("Detected hostname is not valid", zap.Error(err))
	}

	c.logger.Debug("Canonical hostname automatically set", zap.String("hostname", hostname))
	cache.Cache.Set(cache.CanonicalHostnameKey, hostname, cache.NoExpiration)
	return hostname, nil
}
