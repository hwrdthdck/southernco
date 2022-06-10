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

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common/testhelper"
)

func Test_truncateAll(t *testing.T) {
	input := pcommon.NewMap()
	input.InsertString("test", "hello world")
	input.InsertInt("test2", 3)
	input.InsertBool("test3", true)

	target := &testGetSetter{
		getter: func(ctx TransformContext) interface{} {
			return ctx.GetItem()
		},
		setter: func(ctx TransformContext, val interface{}) {
			ctx.GetItem().(pcommon.Map).Clear()
			val.(pcommon.Map).CopyTo(ctx.GetItem().(pcommon.Map))
		},
	}

	tests := []struct {
		name   string
		target GetSetter
		limit  int64
		want   func(pcommon.Map)
	}{
		{
			name:   "truncate map",
			target: target,
			limit:  1,
			want: func(expectedMap pcommon.Map) {
				expectedMap.Clear()
				expectedMap.InsertString("test", "h")
				expectedMap.InsertInt("test2", 3)
				expectedMap.InsertBool("test3", true)
			},
		},
		{
			name:   "truncate map to zero",
			target: target,
			limit:  0,
			want: func(expectedMap pcommon.Map) {
				expectedMap.Clear()
				expectedMap.InsertString("test", "")
				expectedMap.InsertInt("test2", 3)
				expectedMap.InsertBool("test3", true)
			},
		},
		{
			name:   "truncate nothing",
			target: target,
			limit:  100,
			want: func(expectedMap pcommon.Map) {
				expectedMap.Clear()
				expectedMap.InsertString("test", "hello world")
				expectedMap.InsertInt("test2", 3)
				expectedMap.InsertBool("test3", true)
			},
		},
		{
			name:   "truncate exact",
			target: target,
			limit:  11,
			want: func(expectedMap pcommon.Map) {
				expectedMap.Clear()
				expectedMap.InsertString("test", "hello world")
				expectedMap.InsertInt("test2", 3)
				expectedMap.InsertBool("test3", true)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scenarioMap := pcommon.NewMap()
			input.CopyTo(scenarioMap)

			ctx := testhelper.TestTransformContext{
				Item: scenarioMap,
			}

			exprFunc, _ := truncateAll(tt.target, tt.limit)
			exprFunc(ctx)

			expected := pcommon.NewMap()
			tt.want(expected)

			assert.Equal(t, expected, scenarioMap)
		})
	}
}

func Test_truncateAll_validation(t *testing.T) {
	tests := []struct {
		name   string
		target GetSetter
		limit  int64
	}{
		{
			name:   "limit less than zero",
			target: &testGetSetter{},
			limit:  int64(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := truncateAll(tt.target, tt.limit)
			assert.Error(t, err, "invalid limit for truncate_all function, -1 cannot be negative")
		})
	}
}

func Test_truncateAll_bad_input(t *testing.T) {
	input := pcommon.NewValueString("not a map")
	ctx := testhelper.TestTransformContext{
		Item: input,
	}

	target := &testGetSetter{
		getter: func(ctx TransformContext) interface{} {
			return ctx.GetItem()
		},
		setter: func(ctx TransformContext, val interface{}) {
			t.Errorf("nothing should be set in this scenario")
		},
	}

	exprFunc, _ := truncateAll(target, 1)
	exprFunc(ctx)

	assert.Equal(t, pcommon.NewValueString("not a map"), input)
}

func Test_truncateAll_get_nil(t *testing.T) {
	ctx := testhelper.TestTransformContext{
		Item: nil,
	}

	target := &testGetSetter{
		getter: func(ctx TransformContext) interface{} {
			return ctx.GetItem()
		},
		setter: func(ctx TransformContext, val interface{}) {
			t.Errorf("nothing should be set in this scenario")
		},
	}

	exprFunc, _ := truncateAll(target, 1)
	exprFunc(ctx)
}
