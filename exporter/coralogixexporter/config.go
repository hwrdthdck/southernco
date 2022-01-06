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

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"fmt"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	typestr = "coralogix"
)

// Config defines by Coralogix.
type Config struct {
	config.ExporterSettings      `mapstructure:",squash"`
	exporterhelper.QueueSettings `mapstructure:"sending_queue"`
	exporterhelper.RetrySettings `mapstructure:"retry_on_failure"`

	// The Coralogix logs ingress endpoint
	configgrpc.GRPCClientSettings `mapstructure:",squash"`

	// Your Coralogix private key (sensitive) for authentication
	PrivateKey string `mapstructure:"private_key"`

	// Traces emitted by this OpenTelemetry exporter should be tagged
	// in Coralogix with the following application and subsystem names
	AppName   string `mapstructure:"application_name"`
	SubSystem string `mapstructure:"subsystem_name"`
}

func (c *Config) Validate() error {
	// validate each parameter and return specific error
	if c.GRPCClientSettings.Endpoint == "" {
		return fmt.Errorf("`endpoint` not specified, please fix the configuration file")
	}
	if c.PrivateKey == "" {
		return fmt.Errorf("`privateKey` not specified, please fix the configuration file")
	}
	if c.AppName == "" {
		return fmt.Errorf("`appName` not specified, please fix the configuration file")
	}
	if c.SubSystem == "" {
		return fmt.Errorf("`subSystem` not specified, please fix the configuration file")
	}

	// check if headers exists
	if len(c.GRPCClientSettings.Headers) == 0 {
		c.GRPCClientSettings.Headers = map[string]string{"ACCESS_TOKEN": c.PrivateKey, "appName": c.AppName, "subsystemName": c.SubSystem}
	} else {
		c.GRPCClientSettings.Headers["ACCESS_TOKEN"] = c.PrivateKey
		c.GRPCClientSettings.Headers["appName"] = c.AppName
		c.GRPCClientSettings.Headers["subsystemName"] = c.SubSystem
	}
	return nil
}
