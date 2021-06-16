// Copyright  OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package observiqexporter

import (
	"errors"
	"net/url"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config object for observIQ exporter
type Config struct {
	config.ExporterSettings        `mapstructure:",squash"`
	exporterhelper.TimeoutSettings `mapstructure:",squash"`
	exporterhelper.QueueSettings   `mapstructure:"sending_queue"`
	exporterhelper.RetrySettings   `mapstructure:"retry_on_failure"`
	// TLS Settings for http client to use when sending logs to endpoint
	TLSSetting configtls.TLSClientSetting `mapstructure:",squash"`
	// API key for authenticating with ingestion endpoint (required)
	APIKey string `mapstructure:"api_key"`
	// Endpoint URL; Defines the ingestion endpoint (optional)
	Endpoint string `mapstructure:"endpoint"`
	// Name that identifies this agent (optional)
	AgentID string `mapstructure:"agent_id"`
}

func (c *Config) validateConfig() error {
	if c.APIKey == "" {
		return errors.New("api_key must not be empty")
	}

	if c.Endpoint == "" {
		return errors.New("endpoint must not be empty")
	}

	url, urlParseError := url.Parse(c.Endpoint)

	if urlParseError != nil {
		return urlParseError
	}

	if url.Scheme != "http" && url.Scheme != "https" {
		return errors.New("url scheme must be http or https")
	}

	return nil
}
