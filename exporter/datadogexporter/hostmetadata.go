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

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"fmt"
	"strings"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata"
)

// getHostTags gets the host tags extracted from the configuration.
func getHostTags(t *config.TagsConfig) []string {
	tags := t.Tags

	if len(tags) == 0 {
		//lint:ignore SA1019 Will be removed when environment variable detection is removed
		tags = strings.Split(t.EnvVarTags, " ")
	}

	if t.Env != "none" {
		tags = append(tags, fmt.Sprintf("env:%s", t.Env))
	}
	return tags
}

// newMetadataConfigfromConfig creates a new metadata pusher config from the main config.
func newMetadataConfigfromConfig(cfg *config.Config) metadata.PusherConfig {
	return metadata.PusherConfig{
		ConfigHostname:      cfg.Hostname,
		ConfigTags:          getHostTags(&cfg.TagsConfig),
		MetricsEndpoint:     cfg.Metrics.Endpoint,
		APIKey:              cfg.API.Key,
		UseResourceMetadata: cfg.UseResourceMetadata,
		InsecureSkipVerify:  cfg.TLSSetting.InsecureSkipVerify,
		TimeoutSettings:     cfg.TimeoutSettings,
		RetrySettings:       cfg.RetrySettings,
	}
}
