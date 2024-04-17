// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processortest"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/idutils"
)

// defaultHashSeed is used throughout to ensure that the HashSeed is real
// and does not fall back to proportional-mode sampling due to HashSeed == 0.
const defaultHashSeed = 4312

func TestHashBucketsLog2(t *testing.T) {
	require.Equal(t, numHashBuckets, 1<<numHashBucketsLg2)
}

func TestNewTracesProcessor(t *testing.T) {
	tests := []struct {
		name         string
		nextConsumer consumer.Traces
		cfg          *Config
		wantErr      bool
	}{
		{
			name:         "happy_path_default",
			nextConsumer: consumertest.NewNop(),
			cfg: &Config{
				SamplingPercentage: 15.5,
			},
		},
		{
			name:         "happy_path_hash_seed",
			nextConsumer: consumertest.NewNop(),
			cfg: &Config{
				SamplingPercentage: 13.33,
				HashSeed:           defaultHashSeed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTracesProcessor(context.Background(), processortest.NewNopCreateSettings(), tt.cfg, tt.nextConsumer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

// Test_tracesamplerprocessor_SamplingPercentageRange checks for different sampling rates and ensures
// that they are within acceptable deltas.
func Test_tracesamplerprocessor_SamplingPercentageRange(t *testing.T) {
	tests := []struct {
		name              string
		cfg               *Config
		numBatches        int
		numTracesPerBatch int
		acceptableDelta   float64
	}{
		{
			name: "random_sampling_tiny",
			cfg: &Config{
				SamplingPercentage: 0.03,
			},
			numBatches:        1e5,
			numTracesPerBatch: 2,
			acceptableDelta:   0.02,
		},
		{
			name: "random_sampling_small",
			cfg: &Config{
				SamplingPercentage: 5,
			},
			numBatches:        1e5,
			numTracesPerBatch: 2,
			acceptableDelta:   0.1,
		},
		{
			name: "random_sampling_medium",
			cfg: &Config{
				SamplingPercentage: 50.0,
			},
			numBatches:        1e5,
			numTracesPerBatch: 4,
			acceptableDelta:   0.2,
		},
		{
			name: "random_sampling_high",
			cfg: &Config{
				SamplingPercentage: 90.0,
			},
			numBatches:        1e5,
			numTracesPerBatch: 1,
			acceptableDelta:   0.2,
		},
		{
			name: "random_sampling_all",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			numBatches:        1e5,
			numTracesPerBatch: 1,
			acceptableDelta:   0.0,
		},
	}
	const testSvcName = "test-svc"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cfg.HashSeed = defaultHashSeed

			sink := newAssertTraces(t, testSvcName)
			tsp, err := newTracesProcessor(context.Background(), processortest.NewNopCreateSettings(), tt.cfg, sink)
			if err != nil {
				t.Errorf("error when creating traceSamplerProcessor: %v", err)
				return
			}
			for _, td := range genRandomTestData(tt.numBatches, tt.numTracesPerBatch, testSvcName, 1) {
				assert.NoError(t, tsp.ConsumeTraces(context.Background(), td))
			}
			sampled := sink.spanCount
			actualPercentageSamplingPercentage := float32(sampled) / float32(tt.numBatches*tt.numTracesPerBatch) * 100.0
			delta := math.Abs(float64(actualPercentageSamplingPercentage - tt.cfg.SamplingPercentage))
			if delta > tt.acceptableDelta {
				t.Errorf(
					"got %f percentage sampling rate, want %f (allowed delta is %f but got %f)",
					actualPercentageSamplingPercentage,
					tt.cfg.SamplingPercentage,
					tt.acceptableDelta,
					delta,
				)
			}
		})
	}
}

// Test_tracesamplerprocessor_SamplingPercentageRange_MultipleResourceSpans checks for number of spans sent to xt consumer. This is to avoid duplicate spans
func Test_tracesamplerprocessor_SamplingPercentageRange_MultipleResourceSpans(t *testing.T) {
	tests := []struct {
		name                 string
		cfg                  *Config
		numBatches           int
		numTracesPerBatch    int
		acceptableDelta      float64
		resourceSpanPerTrace int
	}{
		{
			name: "single_batch_single_trace_two_resource_spans",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			numBatches:           1,
			numTracesPerBatch:    1,
			acceptableDelta:      0.0,
			resourceSpanPerTrace: 2,
		},
		{
			name: "single_batch_two_traces_two_resource_spans",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			numBatches:           1,
			numTracesPerBatch:    2,
			acceptableDelta:      0.0,
			resourceSpanPerTrace: 2,
		},
	}
	const testSvcName = "test-svc"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cfg.HashSeed = defaultHashSeed

			sink := new(consumertest.TracesSink)
			tsp, err := newTracesProcessor(context.Background(), processortest.NewNopCreateSettings(), tt.cfg, sink)
			if err != nil {
				t.Errorf("error when creating traceSamplerProcessor: %v", err)
				return
			}

			for _, td := range genRandomTestData(tt.numBatches, tt.numTracesPerBatch, testSvcName, tt.resourceSpanPerTrace) {
				assert.NoError(t, tsp.ConsumeTraces(context.Background(), td))
				assert.Equal(t, tt.resourceSpanPerTrace*tt.numTracesPerBatch, sink.SpanCount())
				sink.Reset()
			}

		})
	}
}

func Test_tracessamplerprocessor_MissingRandomness(t *testing.T) {
	type test struct {
		pct        float32
		failClosed bool
		sampled    bool
	}

	for _, tt := range []test{
		// When the TraceID is empty and failClosed==true, the span is not sampled.
		{0, true, false},
		{62, true, false},
		{100, true, false},

		// When the TraceID is empty and failClosed==false, the span is sampled when pct != 0.
		{0, false, false},
		{62, false, true},
		{100, false, true},
	} {
		t.Run(fmt.Sprint(tt.pct, "_", tt.failClosed), func(t *testing.T) {

			ctx := context.Background()
			traces := ptrace.NewTraces()
			span := traces.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty()
			span.SetTraceID(pcommon.TraceID{})                     // invalid TraceID
			span.SetSpanID(pcommon.SpanID{1, 2, 3, 4, 5, 6, 7, 8}) // valid SpanID
			span.SetName("testing")

			cfg := &Config{
				SamplingPercentage: tt.pct,
				HashSeed:           defaultHashSeed,
				FailClosed:         tt.failClosed,
			}

			sink := new(consumertest.TracesSink)

			set := processortest.NewNopCreateSettings()
			logger, observed := observer.New(zap.InfoLevel)
			set.Logger = zap.New(logger)

			tsp, err := newTracesProcessor(ctx, set, cfg, sink)
			require.NoError(t, err)

			err = tsp.ConsumeTraces(ctx, traces)
			require.NoError(t, err)

			sampledData := sink.AllTraces()
			if tt.sampled {
				require.Equal(t, 1, len(sampledData))
				assert.Equal(t, 1, sink.SpanCount())
			} else {
				require.Equal(t, 0, len(sampledData))
				assert.Equal(t, 0, sink.SpanCount())
			}

			if tt.pct != 0 {
				// pct==0 bypasses the randomness check
				require.Equal(t, 1, len(observed.All()), "should have one log: %v", observed.All())
				require.Contains(t, observed.All()[0].Message, "traces sampler")
				require.Contains(t, observed.All()[0].Context[0].Interface.(error).Error(), "missing randomness")
			} else {
				require.Equal(t, 0, len(observed.All()), "should have no logs: %v", observed.All())
			}
		})
	}
}

// Test_tracesamplerprocessor_SpanSamplingPriority checks if handling of "sampling.priority" is correct.
func Test_tracesamplerprocessor_SpanSamplingPriority(t *testing.T) {
	singleSpanWithAttrib := func(key string, attribValue pcommon.Value) ptrace.Traces {
		traces := ptrace.NewTraces()
		initSpanWithAttribute(key, attribValue, traces.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty())
		return traces
	}
	tests := []struct {
		name    string
		cfg     *Config
		td      ptrace.Traces
		sampled bool
	}{
		{
			name: "must_sample",
			cfg: &Config{
				SamplingPercentage: 0.0,
			},
			td: singleSpanWithAttrib(
				"sampling.priority",
				pcommon.NewValueInt(2)),
			sampled: true,
		},
		{
			name: "must_sample_double",
			cfg: &Config{
				SamplingPercentage: 0.0,
			},
			td: singleSpanWithAttrib(
				"sampling.priority",
				pcommon.NewValueDouble(1)),
			sampled: true,
		},
		{
			name: "must_sample_string",
			cfg: &Config{
				SamplingPercentage: 0.0,
			},
			td: singleSpanWithAttrib(
				"sampling.priority",
				pcommon.NewValueStr("1")),
			sampled: true,
		},
		{
			name: "must_not_sample",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			td: singleSpanWithAttrib(
				"sampling.priority",
				pcommon.NewValueInt(0)),
		},
		{
			name: "must_not_sample_double",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			td: singleSpanWithAttrib(
				"sampling.priority",
				pcommon.NewValueDouble(0)),
		},
		{
			name: "must_not_sample_string",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			td: singleSpanWithAttrib(
				"sampling.priority",
				pcommon.NewValueStr("0")),
		},
		{
			name: "defer_sample_expect_not_sampled",
			cfg: &Config{
				SamplingPercentage: 0.0,
			},
			td: singleSpanWithAttrib(
				"no.sampling.priority",
				pcommon.NewValueInt(2)),
		},
		{
			name: "defer_sample_expect_sampled",
			cfg: &Config{
				SamplingPercentage: 100.0,
			},
			td: singleSpanWithAttrib(
				"no.sampling.priority",
				pcommon.NewValueInt(2)),
			sampled: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.name), func(t *testing.T) {

			sink := new(consumertest.TracesSink)
			cfg := &Config{}
			if tt.cfg != nil {
				*cfg = *tt.cfg
			}

			cfg.HashSeed = defaultHashSeed
			tsp, err := newTracesProcessor(context.Background(), processortest.NewNopCreateSettings(), cfg, sink)
			require.NoError(t, err)

			err = tsp.ConsumeTraces(context.Background(), tt.td)
			require.NoError(t, err)

			sampledData := sink.AllTraces()
			if tt.sampled {
				require.Equal(t, 1, len(sampledData))
				assert.Equal(t, 1, sink.SpanCount())
			} else {
				require.Equal(t, 0, len(sampledData))
				assert.Equal(t, 0, sink.SpanCount())
			}
		})
	}
}

// Test_parseSpanSamplingPriority ensures that the function parsing the attributes is taking "sampling.priority"
// attribute correctly.
func Test_parseSpanSamplingPriority(t *testing.T) {
	tests := []struct {
		name string
		span ptrace.Span
		want samplingPriority
	}{
		{
			name: "nil_span",
			span: ptrace.NewSpan(),
			want: deferDecision,
		},
		{
			name: "nil_attributes",
			span: ptrace.NewSpan(),
			want: deferDecision,
		},
		{
			name: "no_sampling_priority",
			span: getSpanWithAttributes("key", pcommon.NewValueBool(true)),
			want: deferDecision,
		},
		{
			name: "sampling_priority_int_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueInt(0)),
			want: doNotSampleSpan,
		},
		{
			name: "sampling_priority_int_gt_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueInt(1)),
			want: mustSampleSpan,
		},
		{
			name: "sampling_priority_int_lt_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueInt(-1)),
			want: deferDecision,
		},
		{
			name: "sampling_priority_double_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueDouble(0)),
			want: doNotSampleSpan,
		},
		{
			name: "sampling_priority_double_gt_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueDouble(1)),
			want: mustSampleSpan,
		},
		{
			name: "sampling_priority_double_lt_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueDouble(-1)),
			want: deferDecision,
		},
		{
			name: "sampling_priority_string_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueStr("0.0")),
			want: doNotSampleSpan,
		},
		{
			name: "sampling_priority_string_gt_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueStr("0.5")),
			want: mustSampleSpan,
		},
		{
			name: "sampling_priority_string_lt_zero",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueStr("-0.5")),
			want: deferDecision,
		},
		{
			name: "sampling_priority_string_NaN",
			span: getSpanWithAttributes("sampling.priority", pcommon.NewValueStr("NaN")),
			want: deferDecision,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseSpanSamplingPriority(tt.span))
		})
	}
}

func getSpanWithAttributes(key string, value pcommon.Value) ptrace.Span {
	span := ptrace.NewSpan()
	initSpanWithAttribute(key, value, span)
	return span
}

func initSpanWithAttribute(key string, value pcommon.Value, dest ptrace.Span) {
	dest.SetName("spanName")
	value.CopyTo(dest.Attributes().PutEmpty(key))

	// ensure a non-empty trace ID with a deterministic value, one that has
	// all zero bits for the w3c randomness portion.  this value, if sampled
	// with the OTel specification, has R-value 0 and sampled only at 100%.
	dest.SetTraceID(pcommon.TraceID{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

// genRandomTestData generates a slice of ptrace.Traces with the numBatches elements which one with
// numTracesPerBatch spans (ie.: each span has a different trace ID). All spans belong to the specified
// serviceName.
func genRandomTestData(numBatches, numTracesPerBatch int, serviceName string, resourceSpanCount int) (tdd []ptrace.Traces) {
	r := rand.New(rand.NewSource(1))
	var traceBatches []ptrace.Traces
	for i := 0; i < numBatches; i++ {
		traces := ptrace.NewTraces()
		traces.ResourceSpans().EnsureCapacity(resourceSpanCount)
		for j := 0; j < resourceSpanCount; j++ {
			rs := traces.ResourceSpans().AppendEmpty()
			rs.Resource().Attributes().PutStr("service.name", serviceName)
			rs.Resource().Attributes().PutBool("bool", true)
			rs.Resource().Attributes().PutStr("string", "yes")
			rs.Resource().Attributes().PutInt("int64", 10000000)
			ils := rs.ScopeSpans().AppendEmpty()
			ils.Spans().EnsureCapacity(numTracesPerBatch)

			for k := 0; k < numTracesPerBatch; k++ {
				span := ils.Spans().AppendEmpty()
				span.SetTraceID(idutils.UInt64ToTraceID(r.Uint64(), r.Uint64()))
				span.SetSpanID(idutils.UInt64ToSpanID(r.Uint64()))
				span.Attributes().PutInt(conventions.AttributeHTTPStatusCode, 404)
				span.Attributes().PutStr("http.status_text", "Not Found")
			}
		}
		traceBatches = append(traceBatches, traces)
	}

	return traceBatches
}

// assertTraces is a traces consumer.Traces
type assertTraces struct {
	*testing.T
	testName  string
	traceIDs  map[[16]byte]bool
	spanCount int
}

var _ consumer.Traces = &assertTraces{}

func (a *assertTraces) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{}
}

func (a *assertTraces) ConsumeTraces(_ context.Context, data ptrace.Traces) error {
	tt := [1]ptrace.Traces{
		data,
	}
	a.onSampledData(tt[:])
	return nil
}

func newAssertTraces(t *testing.T, name string) *assertTraces {
	return &assertTraces{
		T:         t,
		testName:  name,
		traceIDs:  map[[16]byte]bool{},
		spanCount: 0,
	}
}

// onSampledData checks for no repeated traceIDs and counts the number of spans on the sampled data for
// the given service.
func (a *assertTraces) onSampledData(sampled []ptrace.Traces) {
	for _, td := range sampled {
		rspans := td.ResourceSpans()
		for i := 0; i < rspans.Len(); i++ {
			rspan := rspans.At(i)
			ilss := rspan.ScopeSpans()
			for j := 0; j < ilss.Len(); j++ {
				ils := ilss.At(j)
				if svcNameAttr, _ := rspan.Resource().Attributes().Get("service.name"); svcNameAttr.Str() != a.testName {
					continue
				}
				for k := 0; k < ils.Spans().Len(); k++ {
					a.spanCount++
					span := ils.Spans().At(k)
					key := span.TraceID()
					if a.traceIDs[key] {
						a.Errorf("same traceID used more than once %q", key)
						return
					}
					a.traceIDs[key] = true
				}
			}
		}
	}
	return
}
