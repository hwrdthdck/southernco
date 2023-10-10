// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/otlpencodingextension"
import "fmt"

type Config struct {
	Protocol string `mapstructure:"protocol"`
}

func (c *Config) validate() error {
	if c.Protocol != otlpProto && c.Protocol != otlpJSON {
		return fmt.Errorf("unsupported protocol: %q", c.Protocol)
	}

	return nil
}
