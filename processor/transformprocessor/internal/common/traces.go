// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

var _ consumer.Traces = &traceStatements{}

type traceStatements struct {
	ottl.StatementSequence[ottlspan.TransformContext]
	expr.BoolExpr[ottlspan.TransformContext]
}

func (t traceStatements) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: true,
	}
}

func (t traceStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	for i := 0; i < td.ResourceSpans().Len(); i++ {
		rspans := td.ResourceSpans().At(i)
		for j := 0; j < rspans.ScopeSpans().Len(); j++ {
			sspans := rspans.ScopeSpans().At(j)
			spans := sspans.Spans()
			for k := 0; k < spans.Len(); k++ {
				tCtx := ottlspan.NewTransformContext(spans.At(k), sspans.Scope(), rspans.Resource(), sspans, rspans)
				condition, err := t.BoolExpr.Eval(ctx, tCtx)
				if err != nil {
					return err
				}
				if condition {
					err := t.Execute(ctx, tCtx)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

var _ consumer.Traces = &spanEventStatements{}

type spanEventStatements struct {
	ottl.StatementSequence[ottlspanevent.TransformContext]
	expr.BoolExpr[ottlspanevent.TransformContext]
}

func (s spanEventStatements) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: true,
	}
}

func (s spanEventStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	for i := 0; i < td.ResourceSpans().Len(); i++ {
		rspans := td.ResourceSpans().At(i)
		for j := 0; j < rspans.ScopeSpans().Len(); j++ {
			sspans := rspans.ScopeSpans().At(j)
			spans := sspans.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				spanEvents := span.Events()
				for n := 0; n < spanEvents.Len(); n++ {
					tCtx := ottlspanevent.NewTransformContext(spanEvents.At(n), span, sspans.Scope(), rspans.Resource(), sspans, rspans)
					condition, err := s.BoolExpr.Eval(ctx, tCtx)
					if err != nil {
						return err
					}
					if condition {
						err := s.Execute(ctx, tCtx)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

type TraceParserCollection ottl.ParserCollection[ContextStatements, consumer.Traces]

type TraceParserCollectionOption ottl.ParserCollectionOption[ContextStatements, consumer.Traces]

func WithSpanParser(functions map[string]ottl.Factory[ottlspan.TransformContext]) TraceParserCollectionOption {
	return func(pc *ottl.ParserCollection[ContextStatements, consumer.Traces]) error {
		parser, err := ottlspan.NewParser(functions, pc.Settings, ottlspan.WithPathContextNames())
		if err != nil {
			return err
		}
		return ottl.WithParserCollectionContext(ottlspan.PathContextName, &parser, convertSpanStatements)(pc)
	}
}

func convertSpanStatements(collection *ottl.ParserCollection[ContextStatements, consumer.Traces], _ *ottl.Parser[ottlspan.TransformContext], _ string, statements ContextStatements, parsedStatements []*ottl.Statement[ottlspan.TransformContext]) (consumer.Traces, error) {
	globalExpr, errGlobalBoolExpr := parseGlobalExpr(filterottl.NewBoolExprForSpan, statements.Conditions, collection.ErrorMode, collection.Settings, filterottl.StandardSpanFuncs())
	if errGlobalBoolExpr != nil {
		return nil, errGlobalBoolExpr
	}
	sStatements := ottlspan.NewStatementSequence(parsedStatements, collection.Settings, ottlspan.WithStatementSequenceErrorMode(collection.ErrorMode))
	return traceStatements{sStatements, globalExpr}, nil
}

func WithSpanEventParser(functions map[string]ottl.Factory[ottlspanevent.TransformContext]) TraceParserCollectionOption {
	return func(pc *ottl.ParserCollection[ContextStatements, consumer.Traces]) error {
		parser, err := ottlspanevent.NewParser(functions, pc.Settings, ottlspanevent.WithPathContextNames())
		if err != nil {
			return err
		}
		return ottl.WithParserCollectionContext(ottlspanevent.PathContextName, &parser, convertSpanEventStatements)(pc)
	}
}

func convertSpanEventStatements(collection *ottl.ParserCollection[ContextStatements, consumer.Traces], _ *ottl.Parser[ottlspanevent.TransformContext], _ string, statements ContextStatements, parsedStatements []*ottl.Statement[ottlspanevent.TransformContext]) (consumer.Traces, error) {
	globalExpr, errGlobalBoolExpr := parseGlobalExpr(filterottl.NewBoolExprForSpanEvent, statements.Conditions, collection.ErrorMode, collection.Settings, filterottl.StandardSpanEventFuncs())
	if errGlobalBoolExpr != nil {
		return nil, errGlobalBoolExpr
	}
	seStatements := ottlspanevent.NewStatementSequence(parsedStatements, collection.Settings, ottlspanevent.WithStatementSequenceErrorMode(collection.ErrorMode))
	return spanEventStatements{seStatements, globalExpr}, nil
}

func WithTraceErrorMode(errorMode ottl.ErrorMode) TraceParserCollectionOption {
	return TraceParserCollectionOption(ottl.WithParserCollectionErrorMode[ContextStatements, consumer.Traces](errorMode))
}

func NewTraceParserCollection(settings component.TelemetrySettings, options ...TraceParserCollectionOption) (*TraceParserCollection, error) {
	pcOptions := []ottl.ParserCollectionOption[ContextStatements, consumer.Traces]{
		withCommonContextParsers[consumer.Traces](),
		ottl.EnableParserCollectionModifiedStatementLogging[ContextStatements, consumer.Traces](true),
	}

	for _, option := range options {
		pcOptions = append(pcOptions, ottl.ParserCollectionOption[ContextStatements, consumer.Traces](option))
	}

	pc, err := ottl.NewParserCollection(settings, pcOptions...)
	if err != nil {
		return nil, err
	}

	tpc := TraceParserCollection(*pc)
	return &tpc, nil
}

func (tpc *TraceParserCollection) ParseContextStatements(contextStatements ContextStatements) (consumer.Traces, error) {
	pc := ottl.ParserCollection[ContextStatements, consumer.Traces](*tpc)
	if contextStatements.Context != "" {
		return pc.ParseStatementsWithContext(string(contextStatements.Context), contextStatements, true)
	}
	return pc.ParseStatements(contextStatements)
}
