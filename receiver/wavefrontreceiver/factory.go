// Copyright 2019, OpenTelemetry Authors
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

package wavefrontreceiver

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector/config/configerror"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/consumer"
	"github.com/open-telemetry/opentelemetry-collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/transport"
)

// This file implements factory for the Wavefront receiver.

const (
	// The value of "type" key in configuration.
	typeStr = "wavefront"
)

// Factory is the factory for the Wavefront receiver.
type Factory struct {
}

var _ receiver.Factory = (*Factory)(nil)

// Type gets the type of the Receiver config created by this factory.
func (f *Factory) Type() string {
	return typeStr
}

// CustomUnmarshaler returns the custom function to handle any special settings
// used by the receiver.
func (f *Factory) CustomUnmarshaler() receiver.CustomUnmarshaler {
	return nil
}

// CreateDefaultConfig creates the default configuration for the Wavefront receiver.
func (f *Factory) CreateDefaultConfig() configmodels.Receiver {
	return &Config{
		ReceiverSettings: configmodels.ReceiverSettings{
			TypeVal:  typeStr,
			NameVal:  typeStr,
			Endpoint: "localhost:2003",
		},
		TCPIdleTimeout: transport.TCPIdleTimeoutDefault,
	}
}

// CreateTraceReceiver creates a trace receiver based on provided config.
func (f *Factory) CreateTraceReceiver(
	ctx context.Context,
	logger *zap.Logger,
	cfg configmodels.Receiver,
	consumer consumer.TraceConsumer,
) (receiver.TraceReceiver, error) {

	return nil, configerror.ErrDataTypeIsNotSupported
}

// CreateMetricsReceiver creates a metrics receiver based on provided config.
func (f *Factory) CreateMetricsReceiver(
	logger *zap.Logger,
	cfg configmodels.Receiver,
	consumer consumer.MetricsConsumer,
) (receiver.MetricsReceiver, error) {

	rCfg := cfg.(*Config)

	carbonCfg := carbonreceiver.Config{
		ReceiverSettings: rCfg.ReceiverSettings,
		Transport:        "tcp",
		TCPIdleTimeout:   rCfg.TCPIdleTimeout,
		Parser: &protocol.Config{
			Type: "plaintext", // TODO: update after other parsers are implemented for Carbon receiver.
			Config: &WavefrontParser{
				ExtractCollectdTags: rCfg.ExtractCollectdTags,
			},
		},
	}
	return carbonreceiver.New(logger, carbonCfg, consumer)
}
