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

package purefareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"

// This file implements Factory for Array scraper.
import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory creates a factory for Pure Storage FlashArray receiver.
const (
	typeStr   = "purefa"
	stability = component.StabilityLevelDevelopment

	defaultEncoding = "otlp_proto"
)

// FactoryOption applies changes to purefaReceiverFactory.
type FactoryOption func(factory *purefaReceiverFactory)

func NewFactory(options ...FactoryOption) receiver.Factory {

	f := &purefaReceiverFactory{
		logsUnmarshalers: defaultLogsUnmarshalers(),
	}

	for _, o := range options {
		o(f)
	}
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, stability),
		receiver.WithLogs(f.createLogsReceiver, stability))
}

func createDefaultConfig() component.Config {
	return &Config{
		HTTPClientSettings: confighttp.HTTPClientSettings{},
		Settings: &Settings{
			ReloadIntervals: &ReloadIntervals{
				Array:       15 * time.Second,
				Host:        15 * time.Second,
				Directories: 15 * time.Second,
				Pods:        15 * time.Second,
				Volumes:     15 * time.Second,
			},
		},
	}
}

type purefaReceiverFactory struct {
	logsUnmarshalers map[string]LogsUnmarshaler
}

// WithLogsUnmarshalers adds LogsUnmarshalers.
func WithLogsUnmarshalers(logsUnmarshalers ...LogsUnmarshaler) FactoryOption {
	return func(factory *purefaReceiverFactory) {
		for _, unmarshaler := range logsUnmarshalers {
			factory.logsUnmarshalers[unmarshaler.Encoding()] = unmarshaler
		}
	}
}

func createMetricsReceiver(
	_ context.Context,
	set receiver.CreateSettings,
	rCfg component.Config,
	next consumer.Metrics,
) (receiver.Metrics, error) {
	cfg, ok := rCfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("a purefa receiver config was expected by the receiver factory, but got %T", rCfg)
	}
	return newMetricsReceiver(cfg, set, next), nil
}

func (f *purefaReceiverFactory) createLogsReceiver(
	_ context.Context,
	set receiver.CreateSettings,
	rCfg component.Config,
	next consumer.Logs,
) (receiver.Logs, error) {
	cfg, ok := rCfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("a purefa receiver config was expected by the receiver factory, but got %T", rCfg)
	}
	return newLogsReceiver(cfg, set, next), nil
}
