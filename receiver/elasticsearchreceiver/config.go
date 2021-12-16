// Copyright  The OpenTelemetry Authors
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

package elasticsearchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"

import (
	"errors"
	"fmt"
	"net/url"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"
)

// Config is the configuration for the elasticsearch receiver
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	confighttp.HTTPClientSettings           `mapstructure:",squash"`
	// Username is the username used when making REST calls to elasticsearch. Must be specified if Password is. Not required.
	Username string `mapstructure:"username"`
	// Password is the password used when making REST calls to elasticsearch. Must be specified if Username is. Not required.
	Password string `mapstructure:"password"`
}

var (
	defaultEndpoint = "http://localhost:9200"
)

var (
	errEndpointBadScheme    = errors.New("endpoint scheme must be http or https")
	errUsernameNotSpecified = errors.New("password was specified, but not username")
	errPasswordNotSpecified = errors.New("username was specified, but not password")
	errEmptyEndpoint        = errors.New("endpoint must be specified")
)

var validSchemes = []string{
	"http",
	"https",
}

// Validate validates the given config, returning an error specifying any issues with the config.
func (cfg *Config) Validate() error {
	var combinedErr error
	if err := invalidCredentials(cfg.Username, cfg.Password); err != nil {
		combinedErr = multierr.Append(combinedErr, err)
	}

	if cfg.Endpoint == "" {
		return multierr.Append(combinedErr, errEmptyEndpoint)
	}

	u, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return multierr.Append(
			combinedErr,
			fmt.Errorf("invalid endpoint '%s': %w", cfg.Endpoint, err),
		)
	}

	if !validScheme(u.Scheme) {
		return multierr.Append(combinedErr, errEndpointBadScheme)
	}

	return combinedErr
}

// validScheme checks if the given scheme string is one of the valid schemes.
func validScheme(scheme string) bool {
	for _, s := range validSchemes {
		if s == scheme {
			return true
		}
	}

	return false
}

// invalidCredentials returns true if only one username or password is not empty.
func invalidCredentials(username, password string) error {
	if username == "" && password != "" {
		return errUsernameNotSpecified
	}

	if password == "" && username != "" {
		return errPasswordNotSpecified
	}
	return nil
}
