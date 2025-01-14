// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"context"
	"errors"
	"runtime"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/internal/metadata"
)

// This file implements Factory for Process scraper.

// Type the value of "type" key in configuration.
var Type = component.MustNewType("process")

var (
	bootTimeCacheFeaturegateID = "hostmetrics.process.bootTimeCache"
	bootTimeCacheFeaturegate   = featuregate.GlobalRegistry().MustRegister(
		bootTimeCacheFeaturegateID,
		featuregate.StageBeta,
		featuregate.WithRegisterDescription("When enabled, all process scrapes will use the boot time value that is cached at the start of the process."),
		featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/28849"),
		featuregate.WithRegisterFromVersion("v0.98.0"),
	)
)

// Factory is the Factory for scraper.
type Factory struct{}

// CreateDefaultConfig creates the default configuration for the Scraper.
func (f *Factory) CreateDefaultConfig() component.Config {
	return &Config{
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

// CreateMetricsScraper creates a resource scraper based on provided config.
func (f *Factory) CreateMetricsScraper(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
) (scraper.Metrics, error) {
	if runtime.GOOS != "linux" && runtime.GOOS != "windows" && runtime.GOOS != "darwin" {
		return nil, errors.New("process scraper only available on Linux, Windows, or MacOS")
	}

	s, err := newProcessScraper(settings, cfg.(*Config))
	if err != nil {
		return nil, err
	}

	return scraper.NewMetrics(
		s.scrape,
		scraper.WithStart(s.start),
	)
}
