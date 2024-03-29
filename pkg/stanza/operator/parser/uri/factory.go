// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package uri // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/uri"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

var operatorType = component.MustNewType("uri_parser")

func init() {
	operator.RegisterFactory(NewFactory())
}

// NewFactory creates a new factory.
func NewFactory() operator.Factory {
	return operator.NewFactory(operatorType, newDefaultConfig, createOperator)
}

func newDefaultConfig(operatorID string) component.Config {
	return &Config{
		ParserConfig: helper.NewParserConfig(operatorID, operatorType.String()),
	}
}

func createOperator(cfg component.Config, set component.TelemetrySettings) (operator.Operator, error) {
	c := cfg.(*Config)
	parserOperator, err := helper.NewParser(c.ParserConfig, set)
	if err != nil {
		return nil, err
	}

	return &Parser{
		ParserOperator: parserOperator,
	}, nil
}
