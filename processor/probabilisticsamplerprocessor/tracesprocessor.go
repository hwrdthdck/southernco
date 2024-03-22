// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"

import (
	"context"
	"strconv"
	"strings"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.uber.org/zap"
)

// samplingPriority has the semantic result of parsing the "sampling.priority"
// attribute per OpenTracing semantic conventions.
type samplingPriority int

const (
	// deferDecision means that the decision if a span will be "sampled" (ie.:
	// forwarded by the collector) is made by hashing the trace ID according
	// to the configured sampling rate.
	deferDecision samplingPriority = iota
	// mustSampleSpan indicates that the span had a "sampling.priority" attribute
	// greater than zero and it is going to be sampled, ie.: forwarded by the
	// collector.
	mustSampleSpan
	// doNotSampleSpan indicates that the span had a "sampling.priority" attribute
	// equal zero and it is NOT going to be sampled, ie.: it won't be forwarded
	// by the collector.
	doNotSampleSpan
)

type traceProcessor struct {
	sampler    dataSampler
	failClosed bool
	logger     *zap.Logger
}

type tracestateCarrier struct {
	span ptrace.Span
	sampling.W3CTraceState
}

var _ samplingCarrier = &tracestateCarrier{}

func newTracestateCarrier(s ptrace.Span) (samplingCarrier, error) {
	var err error
	tsc := &tracestateCarrier{
		span: s,
	}
	tsc.W3CTraceState, err = sampling.NewW3CTraceState(s.TraceState().AsRaw())
	return tsc, err
}

func (tc *tracestateCarrier) threshold() (sampling.Threshold, bool) {
	return tc.W3CTraceState.OTelValue().TValueThreshold()
}

func (tc *tracestateCarrier) explicitRandomness() (randomnessNamer, bool) {
	rnd, ok := tc.W3CTraceState.OTelValue().RValueRandomness()
	if !ok {
		return newMissingRandomnessMethod(), false
	}
	return newSamplingRandomnessMethod(rnd), true
}

func (tc *tracestateCarrier) updateThreshold(th sampling.Threshold) error {
	return tc.W3CTraceState.OTelValue().UpdateTValueWithSampling(th)
}

func (tc *tracestateCarrier) setExplicitRandomness(rnd randomnessNamer) {
	tc.W3CTraceState.OTelValue().SetRValue(rnd.randomness())
}

func (tc *tracestateCarrier) clearThreshold() {
	tc.W3CTraceState.OTelValue().ClearTValue()
}

func (tc *tracestateCarrier) reserialize() error {
	var w strings.Builder
	err := tc.W3CTraceState.Serialize(&w)
	if err == nil {
		tc.span.TraceState().FromRaw(w.String())
	}
	return err
}

// newTracesProcessor returns a processor.TracesProcessor that will
// perform intermediate span sampling according to the given
// configuration.
func newTracesProcessor(ctx context.Context, set processor.CreateSettings, cfg *Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	tp := &traceProcessor{
		sampler:    makeSampler(cfg, false),
		failClosed: cfg.FailClosed,
		logger:     set.Logger,
	}

	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		tp.processTraces,
		processorhelper.WithCapabilities(consumer.Capabilities{MutatesData: true}))
}

func (th *hashingSampler) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	tid := s.TraceID()
	// Note: this admits empty TraceIDs.
	rnd := newTraceIDHashingMethod(randomnessFromBytes(tid[:], th.hashSeed))
	tsc, err := newTracestateCarrier(s)

	// If the tracestate contains a proper R-value or T-value, we
	// have to leave it alone.  The user should not be using this
	// sampler mode if they are using specified forms of consistent
	// sampling in OTel.
	if err != nil {
		return rnd, nil, err
	} else if _, has := tsc.explicitRandomness(); has {
		err = ErrRandomnessInUse
		tsc = nil
	} else if _, has := tsc.threshold(); has {
		err = ErrThresholdInUse
		tsc = nil
	} else {
		// When no sampling information is present, add a
		// Randomness value.
		tsc.setExplicitRandomness(rnd)
	}
	return rnd, tsc, err
}

func (ctc *consistentTracestateCommon) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	rnd := newMissingRandomnessMethod()
	tsc, err := newTracestateCarrier(s)
	if err != nil {
		tsc = nil
	} else if rv, has := tsc.explicitRandomness(); has {
		// When the tracestate is OK and has r-value, use it.
		rnd = rv
	} else if s.TraceID().IsEmpty() {
		// If the TraceID() is all zeros, which W3C calls an invalid TraceID.
		// rnd continues to be missing.
	} else {
		rnd = newTraceIDW3CSpecMethod(sampling.TraceIDToRandomness(s.TraceID()))
	}

	return rnd, tsc, err
}

func (tp *traceProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	td.ResourceSpans().RemoveIf(func(rs ptrace.ResourceSpans) bool {
		rs.ScopeSpans().RemoveIf(func(ils ptrace.ScopeSpans) bool {
			ils.Spans().RemoveIf(func(s ptrace.Span) bool {
				return commonSamplingLogic(
					ctx,
					s,
					tp.sampler,
					tp.failClosed,
					tp.sampler.randomnessFromSpan,
					tp.priorityFunc,
					"traces sampler",
					tp.logger,
				)
			})
			// Filter out empty ScopeMetrics
			return ils.Spans().Len() == 0
		})
		// Filter out empty ResourceMetrics
		return rs.ScopeSpans().Len() == 0
	})
	if td.ResourceSpans().Len() == 0 {
		return td, processorhelper.ErrSkipProcessingData
	}
	return td, nil
}

func (tp *traceProcessor) priorityFunc(s ptrace.Span, rnd randomnessNamer, threshold sampling.Threshold) (randomnessNamer, sampling.Threshold) {
	switch parseSpanSamplingPriority(s) {
	case doNotSampleSpan:
		// OpenTracing mentions this as a "hint". We take a stronger
		// approach and do not sample the span since some may use it to
		// remove specific spans from traces.
		threshold = sampling.NeverSampleThreshold
		rnd = newSamplingPriorityMethod(rnd.randomness()) // override policy name
	case mustSampleSpan:
		threshold = sampling.AlwaysSampleThreshold
		rnd = newSamplingPriorityMethod(rnd.randomness()) // override policy name
	case deferDecision:
		// Note that the logs processor has very different logic here,
		// but that in tracing the priority can only force to never or
		// always.
	}
	return rnd, threshold
}

// parseSpanSamplingPriority checks if the span has the "sampling.priority" tag to
// decide if the span should be sampled or not. The usage of the tag follows the
// OpenTracing semantic tags:
// https://github.com/opentracing/specification/blob/main/semantic_conventions.md#span-tags-table
func parseSpanSamplingPriority(span ptrace.Span) samplingPriority {
	attribMap := span.Attributes()
	if attribMap.Len() <= 0 {
		return deferDecision
	}

	samplingPriorityAttrib, ok := attribMap.Get("sampling.priority")
	if !ok {
		return deferDecision
	}

	// By default defer the decision.
	decision := deferDecision

	// Try check for different types since there are various client libraries
	// using different conventions regarding "sampling.priority". Besides the
	// client libraries it is also possible that the type was lost in translation
	// between different formats.
	switch samplingPriorityAttrib.Type() {
	case pcommon.ValueTypeInt:
		value := samplingPriorityAttrib.Int()
		if value == 0 {
			decision = doNotSampleSpan
		} else if value > 0 {
			decision = mustSampleSpan
		}
	case pcommon.ValueTypeDouble:
		value := samplingPriorityAttrib.Double()
		if value == 0.0 {
			decision = doNotSampleSpan
		} else if value > 0.0 {
			decision = mustSampleSpan
		}
	case pcommon.ValueTypeStr:
		attribVal := samplingPriorityAttrib.Str()
		if value, err := strconv.ParseFloat(attribVal, 64); err == nil {
			if value == 0.0 {
				decision = doNotSampleSpan
			} else if value > 0.0 {
				decision = mustSampleSpan
			}
		}
	}

	return decision
}
