// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package libhoneyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver"

import (
	"errors"
	"fmt"
	"net/url"
	"path"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/simplespan"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	HTTP           *HTTPConfig               `mapstructure:"http"`
	AuthAPI        string                    `mapstructure:"auth_api"`
	Wrapper        string                    `mapstructure:"wrapper"`
	FieldMapConfig simplespan.FieldMapConfig `mapstructure:"fields"`
}

type HTTPConfig struct {
	*confighttp.ServerConfig `mapstructure:",squash"`

	// The URL path to receive traces on. If omitted "/" will be used.
	TracesURLPaths []string `mapstructure:"traces_url_paths,omitempty"`
}

func (cfg *Config) Validate() error {
	if cfg.HTTP == nil {
		return errors.New("must specify at least one protocol when using the arbitrary JSON receiver")
	}
	return nil
}

func (cfg *Config) Unmarshal(conf *confmap.Conf) error {
	// first load the config normally
	err := conf.Unmarshal(cfg)
	if err != nil {
		return err
	}

	if !conf.IsSet("http") {
		cfg.HTTP = nil
	} else {
		var err error

		for idx := range cfg.HTTP.TracesURLPaths {
			if cfg.HTTP.TracesURLPaths[idx], err = sanitizeURLPath(cfg.HTTP.TracesURLPaths[idx]); err != nil {
				return err
			}
		}
	}
	if cleanURL, err := url.Parse(cfg.AuthAPI); err != nil {
		cfg.AuthAPI = cleanURL.String()
	} else {
		return err
	}

	return nil
}

func sanitizeURLPath(urlPath string) (string, error) {
	u, err := url.Parse(urlPath)
	if err != nil {
		return "", fmt.Errorf("invalid HTTP URL path set for signal: %w", err)
	}

	if !path.IsAbs(u.Path) {
		u.Path = "/" + u.Path
	}
	return u.Path, nil
}
