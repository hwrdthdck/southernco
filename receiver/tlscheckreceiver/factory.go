// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tlscheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver/internal/metadata"
)

var errConfigNotTLSCheck = errors.New(`invalid config`)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		newDefaultConfig,
		receiver.WithMetrics(newReceiver, metadata.MetricsStability))
}

func newDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()

	return &Config{
		ControllerConfig:     cfg,
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
		Targets:              []*confignet.TCPAddrConfig{},
	}
}

func newReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	tlsCheckConfig, ok := cfg.(*Config)
	if !ok {
		return nil, errConfigNotTLSCheck
	}

	mp := newScraper(tlsCheckConfig, settings, getConnectionState)
	s, err := scraper.NewMetrics(mp.scrape)
	if err != nil {
		return nil, err
	}
	opt := scraperhelper.AddScraper(metadata.Type, s)

	return scraperhelper.NewScraperControllerReceiver(
		&tlsCheckConfig.ControllerConfig,
		settings,
		consumer,
		opt,
	)
}
