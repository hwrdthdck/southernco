// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"

import (
	"context"
	"strconv"
	"strings"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
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
	tv := th.TValue()
	if tv == "" {
		tc.clearThreshold()
		return nil
	}
	return tc.W3CTraceState.OTelValue().UpdateTValueWithSampling(th, tv)
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
	tsc := &tracestateCarrier{
		span: s,
	}

	var err error
	tsc.W3CTraceState, err = sampling.NewW3CTraceState(s.TraceState().AsRaw())
	if err != nil {
		return rnd, nil, err
	}

	// If the tracestate contains a proper R-value or T-value, we
	// have to leave it alone.  The user should not be using this
	// sampler mode if they are using specified forms of consistent
	// sampling in OTel.
	if _, has := tsc.explicitRandomness(); has {
		err = ErrRandomnessInUse
	} else if _, has := tsc.threshold(); has {
		err = ErrThresholdInUse
	} else {
		// When no sampling information is present, add a
		// Randomness value.
		tsc.setExplicitRandomness(rnd)
	}
	return rnd, tsc, err
}

func (ctc *consistentTracestateCommon) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	rawts := s.TraceState().AsRaw()
	rnd := newMissingRandomnessMethod()
	tsc := &tracestateCarrier{
		span: s,
	}

	// Parse the arriving TraceState.
	var err error
	tsc.W3CTraceState, err = sampling.NewW3CTraceState(rawts)
	if err != nil {
		tsc = nil
	} else if rv, has := tsc.W3CTraceState.OTelValue().RValueRandomness(); has {
		// When the tracestate is OK and has r-value, use it.
		rnd = newSamplingRandomnessMethod(rv)
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
				rnd, carrier, err := tp.sampler.randomnessFromSpan(s)

				if err == nil {
					err = consistencyCheck(rnd, carrier)
				}
				var threshold sampling.Threshold
				if err != nil {
					if _, is := err.(samplerError); is {
						tp.logger.Info(err.Error())
					} else {
						tp.logger.Error("trace sampler", zap.Error(err))
					}
					if tp.failClosed {
						threshold = sampling.NeverSampleThreshold
					} else {
						threshold = sampling.AlwaysSampleThreshold
					}
				} else {
					threshold = tp.sampler.decide(carrier)
				}

				switch parseSpanSamplingPriority(s) {
				case doNotSampleSpan:
					// The OpenTelemetry mentions this as a "hint" we take a stronger
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

				sampled := threshold.ShouldSample(rnd.randomness())

				if sampled && carrier != nil {
					if err := carrier.updateThreshold(threshold); err != nil {
						tp.logger.Warn("tracestate", zap.Error(err))
					}
					if err := carrier.reserialize(); err != nil {
						tp.logger.Debug("tracestate serialize", zap.Error(err))
					}
				}

				_ = stats.RecordWithTags(
					ctx,
					[]tag.Mutator{tag.Upsert(tagPolicyKey, rnd.policyName()), tag.Upsert(tagSampledKey, strconv.FormatBool(sampled))},
					statCountTracesSampled.M(int64(1)),
				)

				return !sampled
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
