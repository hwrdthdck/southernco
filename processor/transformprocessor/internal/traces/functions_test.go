// Copyright  The OpenTelemetry Authors
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

package traces

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common/testhelper"
)

func Test_newFunctionCall(t *testing.T) {
	input := ptrace.NewSpan()
	input.SetName("bear")
	attrs := pcommon.NewMap()
	attrs.InsertString("test", "hello world")
	attrs.InsertInt("test2", 3)
	attrs.InsertBool("test3", true)
	attrs.InsertString("test4", "hello")
	attrs.CopyTo(input.Attributes())

	tests := []struct {
		name string
		inv  common.Invocation
		want func(ptrace.Span)
	}{
		{
			name: "set name",
			inv: common.Invocation{
				Function: "set",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "name",
								},
							},
						},
					},
					{
						String: testhelper.Strp("cat"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.SetName("cat")
			},
		},
		{
			name: "set status.code",
			inv: common.Invocation{
				Function: "set",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "status",
								},
								{
									Name: "code",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(1),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Status().SetCode(ptrace.StatusCodeOk)
			},
		},
		{
			name: "keep_keys one",
			inv: common.Invocation{
				Function: "keep_keys",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						String: testhelper.Strp("test"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "keep_keys two",
			inv: common.Invocation{
				Function: "keep_keys",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						String: testhelper.Strp("test"),
					},
					{
						String: testhelper.Strp("test2"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "keep_keys none",
			inv: common.Invocation{
				Function: "keep_keys",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
			},
		},
		{
			name: "truncate attributes",
			inv: common.Invocation{
				Function: "truncate_all",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(1),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "h")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "h")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "truncate attributes with zero",
			inv: common.Invocation{
				Function: "truncate_all",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(0),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "truncate attributes nothing",
			inv: common.Invocation{
				Function: "truncate_all",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(100),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "truncate attributes exact",
			inv: common.Invocation{
				Function: "truncate_all",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(11),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "truncate resource attributes",
			inv: common.Invocation{
				Function: "truncate_all",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "resource",
								},
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(11),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "limit attributes",
			inv: common.Invocation{
				Function: "limit",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(1),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "limit attributes zero",
			inv: common.Invocation{
				Function: "limit",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(0),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "limit attributes nothing",
			inv: common.Invocation{
				Function: "limit",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(100),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "limit resource attributes",
			inv: common.Invocation{
				Function: "limit",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "resource",
								},
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						Int: testhelper.Intp(1),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "testing replace_match",
			inv: common.Invocation{
				Function: "replace_match",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name:   "attributes",
									MapKey: testhelper.Strp("test"),
								},
							},
						},
					},
					{
						String: testhelper.Strp("hello*"),
					},
					{
						String: testhelper.Strp("hello {universe}"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello {universe}")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "testing replace_match doesn't match",
			inv: common.Invocation{
				Function: "replace_match",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name:   "attributes",
									MapKey: testhelper.Strp("test"),
								},
							},
						},
					},
					{
						String: testhelper.Strp("goodbye*"),
					},
					{
						String: testhelper.Strp("goodbye {universe}"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "testing replace_all_matches",
			inv: common.Invocation{
				Function: "replace_all_matches",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						String: testhelper.Strp("hello*"),
					},
					{
						String: testhelper.Strp("hello {universe}"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello {universe}")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello {universe}")
				attrs.CopyTo(span.Attributes())
			},
		},
		{
			name: "testing replace_all_matches no matches",
			inv: common.Invocation{
				Function: "replace_all_matches",
				Arguments: []common.Value{
					{
						Path: &common.Path{
							Fields: []common.Field{
								{
									Name: "attributes",
								},
							},
						},
					},
					{
						String: testhelper.Strp("goodbye*"),
					},
					{
						String: testhelper.Strp("goodbye {universe}"),
					},
				},
			},
			want: func(span ptrace.Span) {
				input.CopyTo(span)
				span.Attributes().Clear()
				attrs := pcommon.NewMap()
				attrs.InsertString("test", "hello world")
				attrs.InsertInt("test2", 3)
				attrs.InsertBool("test3", true)
				attrs.InsertString("test4", "hello")
				attrs.CopyTo(span.Attributes())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			span := ptrace.NewSpan()
			input.CopyTo(span)

			evaluate, err := common.NewFunctionCall(tt.inv, DefaultFunctions(), ParsePath)
			assert.NoError(t, err)
			evaluate(spanTransformContext{
				span:     span,
				il:       pcommon.NewInstrumentationScope(),
				resource: pcommon.NewResource(),
			})

			expected := ptrace.NewSpan()
			tt.want(expected)
			assert.Equal(t, expected, span)
		})
	}
}
