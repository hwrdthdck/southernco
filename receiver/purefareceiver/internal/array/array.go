// Copyright 2022 The OpenTelemetry Authors
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

package array // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver/internal/array"

import (
	"context"
	"fmt"
	"net/url"
	"time"

	configutil "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery"
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver/internal"
)

type arrScraper struct {
	endpoint       string
	arrays         []internal.ScraperConfig
	scrapeInterval time.Duration
}

func NewScraper(ctx context.Context,
	endpoint string,
	arrs []internal.ScraperConfig,
	scrapeInterval time.Duration,
) internal.Scraper {
	return &arrScraper{
		endpoint:       endpoint,
		arrays:         arrs,
		scrapeInterval: scrapeInterval,
	}
}

func (a *arrScraper) ToPrometheusReceiverConfig(host component.Host, fact component.ReceiverFactory) ([]*config.ScrapeConfig, error) {
	scrapeCfgs := []*config.ScrapeConfig{}

	for _, arr := range a.arrays {
		u, err := url.Parse(a.endpoint)
		if err != nil {
			return nil, err
		}

		bearerToken, err := internal.RetrieveBearerToken(arr.Auth, host.GetExtensions())
		if err != nil {
			return nil, err
		}

		httpConfig := configutil.HTTPClientConfig{}
		httpConfig.BearerToken = configutil.Secret(bearerToken)

		scrapeConfig := &config.ScrapeConfig{
			HTTPClientConfig: httpConfig,
			ScrapeInterval:   model.Duration(a.scrapeInterval),
			ScrapeTimeout:    model.Duration(a.scrapeInterval),
			JobName:          fmt.Sprintf("%s/%s/%s", "purefa", "arrays", arr.Address),
			HonorTimestamps:  true,
			Scheme:           u.Scheme,
			MetricsPath:      "/metrics/array",
			Params: url.Values{
				"endpoint": {arr.Address},
			},

			ServiceDiscoveryConfigs: discovery.Configs{
				&discovery.StaticConfig{
					{
						Targets: []model.LabelSet{
							{model.AddressLabel: model.LabelValue(u.Host)},
						},
					},
				},
			},
		}

		scrapeCfgs = append(scrapeCfgs, scrapeConfig)
	}

	return scrapeCfgs, nil
}
