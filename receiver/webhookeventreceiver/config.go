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

package webhookeventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.uber.org/multierr"
)

var (
    errMissingEndpointFromConfig = errors.New("Missing receiver server endpoint from config.")
)

// Config defines configuration for the Generic Webhook receiver.
type Config struct {
	confighttp.HTTPServerSettings `mapstructure:",squash"`      // squash ensures fields are correctly decoded in embedded struct
    ReadTimeout  string              `mapstructure:"read_timeout"`      // wait time for reading request headers in ms. Default is twenty seconds.
    WriteTimeout string             `mapstructure:"write_timeout"`      // wait time for writing request response in ms. Default is twenty seconds.
    Path         string           `mapstructure:"path"`        // path for data collection. Default is <host>:<port>/services/collector
    HealthPath   string           `mapstructure:"health_path"` // path for health check api. Default is /services/collector/health
}

func (cfg *Config) Validate() error {
    var errs error

    if cfg.HTTPServerSettings.Endpoint == "" {
        errs = multierr.Append(errs, errMissingEndpointFromConfig)
    }

    return errs
}
