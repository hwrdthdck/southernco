// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ottlspanevent

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottltest"
)

var (
	traceID  = [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	traceID2 = [16]byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	spanID   = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	spanID2  = [8]byte{8, 7, 6, 5, 4, 3, 2, 1}
)

func Test_newPathGetSetter(t *testing.T) {
	refSpanEvent, refSpan, refIS, refResource := createTelemetry()

	newAttrs := pcommon.NewMap()
	newAttrs.PutStr("hello", "world")

	newEvents := ptrace.NewSpanEventSlice()
	newEvents.AppendEmpty().SetName("new event")

	newLinks := ptrace.NewSpanLinkSlice()
	newLinks.AppendEmpty().SetSpanID(spanID2)

	newStatus := ptrace.NewStatus()
	newStatus.SetMessage("new status")

	tests := []struct {
		name     string
		path     []ottl.Field
		orig     interface{}
		newVal   interface{}
		modified func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource)
	}{
		{
			name: "name",
			path: []ottl.Field{
				{
					Name: "name",
				},
			},
			orig:   "bear",
			newVal: "cat",
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.SetName("cat")
			},
		},
		{
			name: "time_unix_nano",
			path: []ottl.Field{
				{
					Name: "time_unix_nano",
				},
			},
			orig:   int64(100_000_000),
			newVal: int64(200_000_000),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.SetTimestamp(pcommon.NewTimestampFromTime(time.UnixMilli(200)))
			},
		},
		{
			name: "attributes",
			path: []ottl.Field{
				{
					Name: "attributes",
				},
			},
			orig:   refSpanEvent.Attributes(),
			newVal: newAttrs,
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				newAttrs.CopyTo(spanEvent.Attributes())
			},
		},
		{
			name: "attributes string",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("str"),
				},
			},
			orig:   "val",
			newVal: "newVal",
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutStr("str", "newVal")
			},
		},
		{
			name: "attributes bool",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("bool"),
				},
			},
			orig:   true,
			newVal: false,
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutBool("bool", false)
			},
		},
		{
			name: "attributes int",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("int"),
				},
			},
			orig:   int64(10),
			newVal: int64(20),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutInt("int", 20)
			},
		},
		{
			name: "attributes float",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("double"),
				},
			},
			orig:   float64(1.2),
			newVal: float64(2.4),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutDouble("double", 2.4)
			},
		},
		{
			name: "attributes bytes",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("bytes"),
				},
			},
			orig:   []byte{1, 3, 2},
			newVal: []byte{2, 3, 4},
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutEmptyBytes("bytes").FromRaw([]byte{2, 3, 4})
			},
		},
		{
			name: "attributes array string",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("arr_str"),
				},
			},
			orig: func() pcommon.Slice {
				val, _ := refSpanEvent.Attributes().Get("arr_str")
				return val.Slice()
			}(),
			newVal: []string{"new"},
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutEmptySlice("arr_str").AppendEmpty().SetStr("new")
			},
		},
		{
			name: "attributes array bool",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("arr_bool"),
				},
			},
			orig: func() pcommon.Slice {
				val, _ := refSpanEvent.Attributes().Get("arr_bool")
				return val.Slice()
			}(),
			newVal: []bool{false},
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutEmptySlice("arr_bool").AppendEmpty().SetBool(false)
			},
		},
		{
			name: "attributes array int",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("arr_int"),
				},
			},
			orig: func() pcommon.Slice {
				val, _ := refSpanEvent.Attributes().Get("arr_int")
				return val.Slice()
			}(),
			newVal: []int64{20},
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutEmptySlice("arr_int").AppendEmpty().SetInt(20)
			},
		},
		{
			name: "attributes array float",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("arr_float"),
				},
			},
			orig: func() pcommon.Slice {
				val, _ := refSpanEvent.Attributes().Get("arr_float")
				return val.Slice()
			}(),
			newVal: []float64{2.0},
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutEmptySlice("arr_float").AppendEmpty().SetDouble(2.0)
			},
		},
		{
			name: "attributes array bytes",
			path: []ottl.Field{
				{
					Name:   "attributes",
					MapKey: ottltest.Strp("arr_bytes"),
				},
			},
			orig: func() pcommon.Slice {
				val, _ := refSpanEvent.Attributes().Get("arr_bytes")
				return val.Slice()
			}(),
			newVal: [][]byte{{9, 6, 4}},
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.Attributes().PutEmptySlice("arr_bytes").AppendEmpty().SetEmptyBytes().FromRaw([]byte{9, 6, 4})
			},
		},
		{
			name: "dropped_attributes_count",
			path: []ottl.Field{
				{
					Name: "dropped_attributes_count",
				},
			},
			orig:   int64(10),
			newVal: int64(20),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				spanEvent.SetDroppedAttributesCount(20)
			},
		},
		{
			name: "instrumentation_scope",
			path: []ottl.Field{
				{
					Name: "instrumentation_scope",
				},
			},
			orig:   refIS,
			newVal: pcommon.NewInstrumentationScope(),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				pcommon.NewInstrumentationScope().CopyTo(il)
			},
		},
		{
			name: "resource",
			path: []ottl.Field{
				{
					Name: "resource",
				},
			},
			orig:   refResource,
			newVal: pcommon.NewResource(),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				pcommon.NewResource().CopyTo(resource)
			},
		},
		{
			name: "span",
			path: []ottl.Field{
				{
					Name: "span",
				},
			},
			orig:   refSpan,
			newVal: ptrace.NewSpan(),
			modified: func(spanEvent ptrace.SpanEvent, span ptrace.Span, il pcommon.InstrumentationScope, resource pcommon.Resource) {
				ptrace.NewSpan().CopyTo(span)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessor, err := newPathGetSetter(tt.path)
			assert.NoError(t, err)

			spanEvent, span, il, resource := createTelemetry()

			got, err := accessor.Get(context.Background(), NewTransformContext(spanEvent, span, il, resource))
			assert.NoError(t, err)
			assert.Equal(t, tt.orig, got)

			err = accessor.Set(context.Background(), NewTransformContext(spanEvent, span, il, resource), tt.newVal)
			assert.NoError(t, err)

			exSpanEvent, exSpan, exIl, exRes := createTelemetry()
			tt.modified(exSpanEvent, exSpan, exIl, exRes)

			assert.Equal(t, exSpan, span)
			assert.Equal(t, exIl, il)
			assert.Equal(t, exRes, resource)
		})
	}
}

func createTelemetry() (ptrace.SpanEvent, ptrace.Span, pcommon.InstrumentationScope, pcommon.Resource) {
	spanEvent := ptrace.NewSpanEvent()

	spanEvent.SetName("bear")
	spanEvent.SetTimestamp(pcommon.NewTimestampFromTime(time.UnixMilli(100)))
	spanEvent.SetDroppedAttributesCount(10)

	spanEvent.Attributes().PutStr("str", "val")
	spanEvent.Attributes().PutBool("bool", true)
	spanEvent.Attributes().PutInt("int", 10)
	spanEvent.Attributes().PutDouble("double", 1.2)
	spanEvent.Attributes().PutEmptyBytes("bytes").FromRaw([]byte{1, 3, 2})

	arrStr := spanEvent.Attributes().PutEmptySlice("arr_str")
	arrStr.AppendEmpty().SetStr("one")
	arrStr.AppendEmpty().SetStr("two")

	arrBool := spanEvent.Attributes().PutEmptySlice("arr_bool")
	arrBool.AppendEmpty().SetBool(true)
	arrBool.AppendEmpty().SetBool(false)

	arrInt := spanEvent.Attributes().PutEmptySlice("arr_int")
	arrInt.AppendEmpty().SetInt(2)
	arrInt.AppendEmpty().SetInt(3)

	arrFloat := spanEvent.Attributes().PutEmptySlice("arr_float")
	arrFloat.AppendEmpty().SetDouble(1.0)
	arrFloat.AppendEmpty().SetDouble(2.0)

	arrBytes := spanEvent.Attributes().PutEmptySlice("arr_bytes")
	arrBytes.AppendEmpty().SetEmptyBytes().FromRaw([]byte{1, 2, 3})
	arrBytes.AppendEmpty().SetEmptyBytes().FromRaw([]byte{2, 3, 4})

	span := ptrace.NewSpan()
	span.SetName("test")

	il := pcommon.NewInstrumentationScope()
	il.SetName("library")
	il.SetVersion("version")

	resource := pcommon.NewResource()
	span.Attributes().CopyTo(resource.Attributes())

	return spanEvent, span, il, resource
}

func Test_ParseEnum(t *testing.T) {
	tests := []struct {
		name string
		want ottl.Enum
	}{
		{
			name: "SPAN_KIND_UNSPECIFIED",
			want: ottl.Enum(ptrace.SpanKindUnspecified),
		},
		{
			name: "SPAN_KIND_INTERNAL",
			want: ottl.Enum(ptrace.SpanKindInternal),
		},
		{
			name: "SPAN_KIND_SERVER",
			want: ottl.Enum(ptrace.SpanKindServer),
		},
		{
			name: "SPAN_KIND_CLIENT",
			want: ottl.Enum(ptrace.SpanKindClient),
		},
		{
			name: "SPAN_KIND_PRODUCER",
			want: ottl.Enum(ptrace.SpanKindProducer),
		},
		{
			name: "SPAN_KIND_CONSUMER",
			want: ottl.Enum(ptrace.SpanKindConsumer),
		},
		{
			name: "STATUS_CODE_UNSET",
			want: ottl.Enum(ptrace.StatusCodeUnset),
		},
		{
			name: "STATUS_CODE_OK",
			want: ottl.Enum(ptrace.StatusCodeOk),
		},
		{
			name: "STATUS_CODE_ERROR",
			want: ottl.Enum(ptrace.StatusCodeError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseEnum((*ottl.EnumSymbol)(ottltest.Strp(tt.name)))
			assert.NoError(t, err)
			assert.Equal(t, *actual, tt.want)
		})
	}
}

func Test_ParseEnum_False(t *testing.T) {
	tests := []struct {
		name       string
		enumSymbol *ottl.EnumSymbol
	}{
		{
			name:       "unknown enum symbol",
			enumSymbol: (*ottl.EnumSymbol)(ottltest.Strp("not an enum")),
		},
		{
			name:       "nil enum symbol",
			enumSymbol: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseEnum(tt.enumSymbol)
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	}
}
