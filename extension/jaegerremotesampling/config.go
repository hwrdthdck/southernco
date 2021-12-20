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

package jaegerremotesampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"

import (
	"time"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
)

// Config has the configuration for the extension enabling the health check
// extension, used to report the health status of the service.
type Config struct {
	config.ExtensionSettings      `mapstructure:",squash"`
	confighttp.HTTPServerSettings `mapstructure:"http"`
	configgrpc.GRPCServerSettings `mapstructure:"grpc"`

	// Source configures the source for the stratagies file
	Source Source `mapstructure:"source"`
}

type Source struct {
	// Remote defines the remote location for the file
	Remote configgrpc.GRPCClientSettings `mapstructure:"remote"`

	// File specifies a local file as the strategies source
	File string `mapstructure:"file"`

	// ReloadInterval determines the periodicity to refresh the strategies
	ReloadInterval time.Duration `mapstructure:"reload_interval"`
}

var _ config.Extension = (*Config)(nil)

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error {
	return nil
}
