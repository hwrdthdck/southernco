// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdout // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/output/stdout"

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Stdout is a global handle to standard output
var Stdout io.Writer = os.Stdout

var operatorType = component.MustNewType("stdout")

func init() {
	operator.RegisterFactory(NewFactory())
}

// Deprecated [v0.97.0] Use Factory.NewDefaultConfig instead.
func NewConfig(operatorID string) *Config {
	return NewFactory().NewDefaultConfig(operatorID).(*Config)
}

// Config is the configuration of the Stdout operator
type Config struct {
	helper.OutputConfig `mapstructure:",squash"`
}

// Deprecated [v0.97.0] Use Factory.CreateOperator instead.
func (c Config) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	set := component.TelemetrySettings{}
	if logger != nil {
		set.Logger = logger.Desugar()
	}
	return NewFactory().CreateOperator(&c, set)
}

// NewFactory creates a new factory.
func NewFactory() operator.Factory {
	return operator.NewFactory(operatorType, newDefaultConfig, createOperator)
}

func newDefaultConfig(operatorID string) component.Config {
	return &Config{
		OutputConfig: helper.NewOutputConfig(operatorID, "stdout"),
	}
}

func createOperator(cfg component.Config, set component.TelemetrySettings) (operator.Operator, error) {
	c := cfg.(*Config)
	outputOperator, err := helper.NewOutputOperator(c.OutputConfig, set)
	if err != nil {
		return nil, err
	}

	return &Output{
		OutputOperator: outputOperator,
		encoder:        json.NewEncoder(Stdout),
	}, nil
}

// Output is an operator that logs entries using stdout.
type Output struct {
	helper.OutputOperator
	encoder *json.Encoder
	mux     sync.Mutex
}

// Process will log entries received.
func (o *Output) Process(_ context.Context, entry *entry.Entry) error {
	o.mux.Lock()
	err := o.encoder.Encode(entry)
	if err != nil {
		o.mux.Unlock()
		o.Errorf("Failed to process entry: %s", err)
		return err
	}
	o.mux.Unlock()
	return nil
}
