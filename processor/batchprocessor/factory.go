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

package batchprocessor

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector/component"
	"github.com/open-telemetry/opentelemetry-collector/config/configerror"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/consumer"
)

const (
	// The value of "type" key in configuration.
	typeStr = "batch"
)

// Factory is the factory for batch processor.
type Factory struct {
}

// Type gets the type of the config created by this factory.
func (f *Factory) Type() string {
	return typeStr
}

// CreateDefaultConfig creates the default configuration for processor.
func (f *Factory) CreateDefaultConfig() configmodels.Processor {
	removeAfterTicks := int(defaultRemoveAfterCycles)
	sendBatchSize := int(defaultSendBatchSize)
	tickTime := defaultTickTime
	timeout := defaultTimeout

	return &Config{
		ProcessorSettings: configmodels.ProcessorSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		RemoveAfterTicks: &removeAfterTicks,
		SendBatchSize:    &sendBatchSize,
		NumTickers:       defaultNumTickers,
		TickTime:         &tickTime,
		Timeout:          &timeout,
	}
}

// CreateTraceProcessor creates a trace processor based on this config.
func (f *Factory) CreateTraceProcessor(
	logger *zap.Logger,
	nextConsumer consumer.TraceConsumerOld,
	c configmodels.Processor,
) (component.TraceProcessorOld, error) {
	cfg := c.(*Config)

	var batchingOptions []Option
	if cfg.Timeout != nil {
		batchingOptions = append(batchingOptions, WithTimeout(*cfg.Timeout))
	}
	if cfg.NumTickers > 0 {
		batchingOptions = append(
			batchingOptions, WithNumTickers(cfg.NumTickers),
		)
	}
	if cfg.TickTime != nil {
		batchingOptions = append(
			batchingOptions, WithTickTime(*cfg.TickTime),
		)
	}
	if cfg.SendBatchSize != nil {
		batchingOptions = append(
			batchingOptions, WithSendBatchSize(*cfg.SendBatchSize),
		)
	}
	if cfg.RemoveAfterTicks != nil {
		batchingOptions = append(
			batchingOptions, WithRemoveAfterTicks(*cfg.RemoveAfterTicks),
		)
	}

	return NewBatcher(cfg.NameVal, logger, nextConsumer, batchingOptions...), nil
}

// CreateMetricsProcessor creates a metrics processor based on this config.
func (f *Factory) CreateMetricsProcessor(
	logger *zap.Logger,
	nextConsumer consumer.MetricsConsumerOld,
	cfg configmodels.Processor,
) (component.MetricsProcessorOld, error) {
	return nil, configerror.ErrDataTypeIsNotSupported
}
