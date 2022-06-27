// Copyright The OpenTelemetry Authors
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

package skywalkingreceiver

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/ptrace"
	agentV3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

func TestSetInternalSpanStatus(t *testing.T) {
	tests := []struct {
		name   string
		swSpan *agentV3.SpanObject
		dest   ptrace.SpanStatus
		code   ptrace.StatusCode
	}{
		{
			name: "StatusCodeError",
			swSpan: &agentV3.SpanObject{
				IsError: true,
			},
			dest: generateTracesOneEmptyResourceSpans().Status(),
			code: ptrace.StatusCodeError,
		},
		{
			name: "StatusCodeOk",
			swSpan: &agentV3.SpanObject{
				IsError: false,
			},
			dest: generateTracesOneEmptyResourceSpans().Status(),
			code: ptrace.StatusCodeOk,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setInternalSpanStatus(test.swSpan, test.dest)
			assert.Equal(t, test.code, test.dest.Code())
		})
	}
}

func TestSwKvPairsToInternalAttributes(t *testing.T) {
	tests := []struct {
		name   string
		swSpan *agentV3.SegmentObject
		dest   ptrace.Span
	}{
		{
			name:   "mock-sw-swgment-1",
			swSpan: mockGrpcTraceSegment(1),
			dest:   generateTracesOneEmptyResourceSpans(),
		},
		{
			name:   "mock-sw-swgment-2",
			swSpan: mockGrpcTraceSegment(2),
			dest:   generateTracesOneEmptyResourceSpans(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			swKvPairsToInternalAttributes(test.swSpan.GetSpans()[0].Tags, test.dest.Attributes())
			assert.Equal(t, test.dest.Attributes().Len(), len(test.swSpan.GetSpans()[0].Tags))
			for _, tag := range test.swSpan.GetSpans()[0].Tags {
				value, _ := test.dest.Attributes().Get(tag.Key)
				assert.Equal(t, tag.Value, value.AsString())
			}
		})
	}
}

func TestSwProtoToTraces(t *testing.T) {
	tests := []struct {
		name   string
		swSpan *agentV3.SegmentObject
		dest   ptrace.Traces
		code   ptrace.StatusCode
	}{
		{
			name:   "mock-sw-swgment-1",
			swSpan: mockGrpcTraceSegment(1),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			td := SkywalkingToTraces(test.swSpan)
			assert.Equal(t, 1, td.ResourceSpans().Len())
			assert.Equal(t, 2, td.ResourceSpans().At(0).ScopeSpans().At(0).Spans().Len())
		})
	}
}

func TestSwReferencesToSpanLinks(t *testing.T) {
	tests := []struct {
		name   string
		swSpan *agentV3.SegmentObject
		dest   ptrace.Span
	}{
		{
			name:   "mock-sw-swgment-1",
			swSpan: mockGrpcTraceSegment(1),
			dest:   generateTracesOneEmptyResourceSpans(),
		},
		{
			name:   "mock-sw-swgment-2",
			swSpan: mockGrpcTraceSegment(2),
			dest:   generateTracesOneEmptyResourceSpans(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			swReferencesToSpanLinks(test.swSpan.GetSpans()[0].Refs, test.dest.Links())
			assert.Equal(t, 1, test.dest.Links().Len())
		})
	}
}

func TestSwLogsToSpanEvents(t *testing.T) {
	tests := []struct {
		name   string
		swSpan *agentV3.SegmentObject
		dest   ptrace.Span
	}{
		{
			name:   "mock-sw-swgment-0",
			swSpan: mockGrpcTraceSegment(0),
			dest:   generateTracesOneEmptyResourceSpans(),
		},
		{
			name:   "mock-sw-swgment-1",
			swSpan: mockGrpcTraceSegment(1),
			dest:   generateTracesOneEmptyResourceSpans(),
		},
	}
	for index, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			seq := strconv.Itoa(index)
			swLogsToSpanEvents(test.swSpan.GetSpans()[0].Logs, test.dest.Events())
			assert.Equal(t, 1, test.dest.Events().Len())
			assert.Equal(t, "logs", test.dest.Events().At(0).Name())
			logValue, _ := test.dest.Events().At(0).Attributes().Get("log-key" + seq)
			assert.Equal(t, "log-value"+seq, logValue.AsString())
		})
	}
}

func Test_stringToTraceID(t *testing.T) {
	type args struct {
		traceID string
	}
	tests := []struct {
		name          string
		segmentObject args
		want          [16]byte
	}{
		{
			name:          "mock-sw-normal-trace-id",
			segmentObject: args{traceID: "de5980b8-fce3-4a37-aab9-b4ac3af7eedd"},
			want:          [16]byte{222, 89, 128, 184, 252, 227, 74, 55, 170, 185, 180, 172, 58, 247, 238, 221},
		},
		{
			name:          "mock-sw-normal-trace-id",
			segmentObject: args{traceID: "de5980b8fce34a37aab9b4ac3af7eedd"},
			want:          [16]byte{222, 89, 128, 184, 252, 227, 74, 55, 170, 185, 180, 172, 58, 247, 238, 221},
		},
		{
			name:          "mock-sw-trace-id-length-shorter",
			segmentObject: args{traceID: "de59"},
			want:          [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:          "mock-sw-trace-id-length-overflow",
			segmentObject: args{traceID: "de5980b8fce34a37aab9b4ac3af7eeddddddde5980"},
			want:          [16]byte{222, 89, 128, 184, 252, 227, 74, 55, 170, 185, 180, 172, 58, 247, 238, 221},
		},
		{
			name:          "mock-sw-trace-id-illegal",
			segmentObject: args{traceID: ".,<>?/-=+MNop"},
			want:          [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringToTraceID(tt.segmentObject.traceID)
			assert.Equal(t, tt.want, got.Bytes())
		})
	}
}

func Test_segmentIdToSpanId(t *testing.T) {
	type args struct {
		segmentID string
		spanID    uint32
	}
	tests := []struct {
		name string
		args args
		want [8]byte
	}{
		{
			name: "mock-sw-span-id-normal",
			args: args{segmentID: "4f2f27748b8e44ecaf18fe0347194e86.33.16560607369950066", spanID: 123},
			want: [8]byte{123, 86, 6, 7, 54, 153, 80, 6},
		},
		{
			name: "mock-sw-span-id-short",
			args: args{segmentID: "16560607369950066", spanID: 12},
			want: [8]byte{12, 86, 6, 7, 54, 153, 80, 6},
		},
		{
			name: "mock-sw-span-id-illegal-1",
			args: args{segmentID: "1", spanID: 2},
			want: [8]byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "mock-sw-span-id-illegal-char",
			args: args{segmentID: ".,<>?/-=+MNop", spanID: 2},
			want: [8]byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := segmentIDToSpanID(tt.args.segmentID, tt.args.spanID)
			assert.Equal(t, tt.want, got.Bytes())
		})
	}
}

func generateTracesOneEmptyResourceSpans() ptrace.Span {
	td := ptrace.NewTraces()
	resourceSpan := td.ResourceSpans().AppendEmpty()
	il := resourceSpan.ScopeSpans().AppendEmpty()
	il.Spans().AppendEmpty()
	return il.Spans().At(0)
}
