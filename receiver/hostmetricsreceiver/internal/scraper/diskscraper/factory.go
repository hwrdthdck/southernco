// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package diskscraper

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/receiver/hostmetricsreceiver/internal"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// This file implements Factory for Disk scraper.

const (
	// The value of "type" key in configuration.
	TypeStr = "disk"
)

// Factory is the Factory for scraper.
type Factory struct {
}

// CreateDefaultConfig creates the default configuration for the Scraper.
func (f *Factory) CreateDefaultConfig() internal.Config {
	return &Config{}
}

// CreateMetricsScraper creates a scraper based on provided config.
func (f *Factory) CreateMetricsScraper(
	ctx context.Context,
	_ *zap.Logger,
	config internal.Config,
) (receiverhelper.MetricsScraper, error) {
	cfg := config.(*Config)
	s, err := newDiskScraper(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ms := receiverhelper.NewMetricsScraper(
		TypeStr,
		s.Scrape,
		receiverhelper.WithInitialize(s.Initialize),
	)

	return ms, nil
}
