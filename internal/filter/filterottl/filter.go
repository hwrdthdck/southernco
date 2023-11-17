// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

// NewBoolExprForSpan creates a BoolExpr[ottlspan.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlspan.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForSpan(conditions []string, functions map[string]ottl.Factory[ottlspan.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (expr.BoolExpr[ottlspan.TransformContext], error) {
	match := newMatchFactory[ottlspan.TransformContext]()
	if _, ok := functions[match.Name()]; !ok {
		functions[match.Name()] = match
	}
	parser, err := ottlspan.NewParser(functions, set)
	if err != nil {
		return nil, err
	}
	statements, err := parser.ParseConditions(conditions)
	if err != nil {
		return nil, err
	}
	c := ottlspan.NewConditionSequence(statements, errorMode, set)
	return &c, nil
}

// NewBoolExprForSpanEvent creates a BoolExpr[ottlspanevent.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlspanevent.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForSpanEvent(conditions []string, functions map[string]ottl.Factory[ottlspanevent.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (expr.BoolExpr[ottlspanevent.TransformContext], error) {
	match := newMatchFactory[ottlspanevent.TransformContext]()
	if _, ok := functions[match.Name()]; !ok {
		functions[match.Name()] = match
	}
	parser, err := ottlspanevent.NewParser(functions, set)
	if err != nil {
		return nil, err
	}
	statements, err := parser.ParseConditions(conditions)
	if err != nil {
		return nil, err
	}
	c := ottlspanevent.NewConditionSequence(statements, errorMode, set)
	return &c, nil
}

// NewBoolExprForMetric creates a BoolExpr[ottlmetric.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlmetric.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForMetric(conditions []string, functions map[string]ottl.Factory[ottlmetric.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (expr.BoolExpr[ottlmetric.TransformContext], error) {
	match := newMatchFactory[ottlmetric.TransformContext]()
	if _, ok := functions[match.Name()]; !ok {
		functions[match.Name()] = match
	}
	parser, err := ottlmetric.NewParser(functions, set)
	if err != nil {
		return nil, err
	}
	statements, err := parser.ParseConditions(conditions)
	if err != nil {
		return nil, err
	}
	c := ottlmetric.NewConditionSequence(statements, errorMode, set)
	return &c, nil
}

// NewBoolExprForDataPoint creates a BoolExpr[ottldatapoint.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottldatapoint.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForDataPoint(conditions []string, functions map[string]ottl.Factory[ottldatapoint.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (expr.BoolExpr[ottldatapoint.TransformContext], error) {
	match := newMatchFactory[ottldatapoint.TransformContext]()
	if _, ok := functions[match.Name()]; !ok {
		functions[match.Name()] = match
	}
	parser, err := ottldatapoint.NewParser(functions, set)
	if err != nil {
		return nil, err
	}
	statements, err := parser.ParseConditions(conditions)
	if err != nil {
		return nil, err
	}
	c := ottldatapoint.NewConditionSequence(statements, errorMode, set)
	return &c, nil
}

// NewBoolExprForLog creates a BoolExpr[ottllog.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottllog.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForLog(conditions []string, functions map[string]ottl.Factory[ottllog.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (expr.BoolExpr[ottllog.TransformContext], error) {
	match := newMatchFactory[ottllog.TransformContext]()
	if _, ok := functions[match.Name()]; !ok {
		functions[match.Name()] = match
	}
	parser, err := ottllog.NewParser(functions, set)
	if err != nil {
		return nil, err
	}
	statements, err := parser.ParseConditions(conditions)
	if err != nil {
		return nil, err
	}
	c := ottllog.NewConditionSequence(statements, errorMode, set)
	return &c, nil
}

// NewBoolExprForResource creates a BoolExpr[ottlresource.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlresource.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForResource(conditions []string, functions map[string]ottl.Factory[ottlresource.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (expr.BoolExpr[ottlresource.TransformContext], error) {
	match := newMatchFactory[ottlresource.TransformContext]()
	if _, ok := functions[match.Name()]; !ok {
		functions[match.Name()] = match
	}
	parser, err := ottlresource.NewParser(functions, set)
	if err != nil {
		return nil, err
	}
	statements, err := parser.ParseConditions(conditions)
	if err != nil {
		return nil, err
	}
	c := ottlresource.NewConditionSequence(statements, errorMode, set)
	return &c, nil
}
