// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func Test_DurationToMilli(t *testing.T) {
	tests := []struct {
		name     string
		duration ottl.StringGetter[interface{}]
		expected int64
	}{
		{
			name: "100 Milliseconds",
			duration: &ottl.StandardStringGetter[interface{}]{
				Getter: func(ctx context.Context, tCtx interface{}) (interface{}, error) {
					return "100ms", nil
				},
			},
			expected: 100,
		},
		{
			name: "1000 hour",
			duration: &ottl.StandardStringGetter[interface{}]{
				Getter: func(ctx context.Context, tCtx interface{}) (interface{}, error) {
					return "100h", nil
				},
			},
			expected: 360000000,
		},
		{
			name: "47 mins",
			duration: &ottl.StandardStringGetter[interface{}]{
				Getter: func(ctx context.Context, tCtx interface{}) (interface{}, error) {
					return "47m", nil
				},
			},
			expected: 2820000,
		},
		{
			name: "1 hour 40 mins 3 seconds 30 milliseconds",
			duration: &ottl.StandardStringGetter[interface{}]{
				Getter: func(ctx context.Context, tCtx interface{}) (interface{}, error) {
					return "1h40m3s30ms", nil
				},
			},
			expected: 6003030,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exprFunc, err := DurationToMilli(tt.duration)
			assert.NoError(t, err)
			result, err := exprFunc(nil, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
