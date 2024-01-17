// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package headerless_jarray // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jarray"

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

	"github.com/valyala/fastjson"
)

const operatorType = "headerless_jarray_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new jarray parser config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new jarray parser config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		ParserConfig: helper.NewParserConfig(operatorID, operatorType),
	}
}

// Config is the configuration of a jarray parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`

	Header          string `mapstructure:"header"`
	HeaderDelimiter string `mapstructure:"header_delimiter"`
	HeaderAttribute string `mapstructure:"header_attribute"`
}

// Build will build a jarray parser operator.
func (c Config) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	parserOperator, err := c.ParserConfig.Build(logger)
	if err != nil {
		return nil, err
	}

	if c.HeaderDelimiter == "" {
		c.HeaderDelimiter = ","
	}

	headerDelimiter := []rune(c.HeaderDelimiter)[0]

	if len([]rune(c.HeaderDelimiter)) != 1 {
		return nil, fmt.Errorf("invalid 'header_delimiter': '%s'", c.HeaderDelimiter)
	}

	var headers []string
	switch {
	case c.Header == "" && c.HeaderAttribute == "":
		return nil, errors.New("missing required field 'header' or 'header_attribute'")
	case c.Header != "" && c.HeaderAttribute != "":
		return nil, errors.New("only one header parameter can be set: 'header' or 'header_attribute'")
	case c.Header != "" && !strings.Contains(c.Header, c.HeaderDelimiter):
		return nil, errors.New("missing field delimiter in header")
	case c.Header != "":
		headers = strings.Split(c.Header, c.HeaderDelimiter)
	}

	pp := &fastjson.ParserPool{}

	return &Parser{
		ParserOperator:  parserOperator,
		header:          headers,
		headerAttribute: c.HeaderAttribute,
		headerDelimiter: headerDelimiter,
		parse:           generateJarrayParseFunc(headers, pp),
		pp:              pp,
	}, nil
}

// Parser is an operator that parses jarray in an entry.
type Parser struct {
	helper.ParserOperator
	headerDelimiter rune
	header          []string
	headerAttribute string
	parse           parseFunc
	pp              *fastjson.ParserPool
}

type parseFunc func(any) (any, error)

// Process will parse an entry for jarray.
func (r *Parser) Process(ctx context.Context, e *entry.Entry) error {
	parse := r.parse

	// If we have a headerAttribute set we need to dynamically generate our parser function
	if r.headerAttribute != "" {
		h, ok := e.Attributes[r.headerAttribute]
		if !ok {
			err := fmt.Errorf("failed to read dynamic header attribute %s", r.headerAttribute)
			r.Error(err)
			return err
		}
		headerString, ok := h.(string)
		if !ok {
			err := fmt.Errorf("header is expected to be a string but is %T", h)
			r.Error(err)
			return err
		}
		headers := strings.Split(headerString, string([]rune{r.headerDelimiter}))
		parse = generateJarrayParseFunc(headers, r.pp)
	}

	return r.ParserOperator.ProcessWith(ctx, e, parse)
}

func generateJarrayParseFunc(headers []string, pp *fastjson.ParserPool) parseFunc {
	return func(value any) (any, error) {
		jArrayLine, err := valueAsString(value)
		if err != nil {
			return nil, err
		}

		p := pp.Get()
		v, err := p.Parse(jArrayLine)
		pp.Put(p)
		if err != nil {
			return nil, errors.New("failed to parse entry")
		}

		jArray := v.GetArray() // a is a []*Value slice
		if len(jArray) != len(headers) {
			return nil, fmt.Errorf("wrong number of fields: expected %d, found %d", len(headers), len(jArray))
		}
		parsedValues := make(map[string]any)
		for i := range jArray {
			switch jArray[i].Type() {
			case fastjson.TypeNumber:
				parsedValues[headers[i]] = jArray[i].GetInt64()
			case fastjson.TypeString:
				parsedValues[headers[i]] = string(jArray[i].GetStringBytes())
			case fastjson.TypeTrue:
				parsedValues[headers[i]] = true
			case fastjson.TypeFalse:
				parsedValues[headers[i]] = false
			case fastjson.TypeNull:
				parsedValues[headers[i]] = nil
			case fastjson.TypeObject:
				// Nested objects handled as a string since this parser doesn't support nested headers
				parsedValues[headers[i]] = jArray[i].String()
			default:
				return nil, errors.New("failed to parse entry: " + string(jArray[i].MarshalTo(nil)))
			}
		}

		return parsedValues, nil
	}
}

// valueAsString interprets the given value as a string.
func valueAsString(value any) (string, error) {
	var s string
	switch t := value.(type) {
	case string:
		s += t
	case []byte:
		s += string(t)
	default:
		return s, fmt.Errorf("type '%T' cannot be parsed as jarray", value)
	}

	return s, nil
}
