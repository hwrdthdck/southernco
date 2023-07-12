// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type TruncateArguments[K any] struct {
	Target ottl.GetSetter[K] `ottlarg:"0"`
	Limit  int64             `ottlarg:"1"`
}

func NewTruncateFactory[K any]() ottl.Factory[K] {
	return ottl.NewFactory("truncate", &TruncateArguments[K]{}, createTruncateFunction[K])
}

func createTruncateFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	args, ok := oArgs.(*TruncateArguments[K])

	if !ok {
		return nil, fmt.Errorf("TruncateFactory args must be of type *TruncateArguments[K]")
	}

	return truncate(args.Target, args.Limit)
}

func truncate[K any](target ottl.GetSetter[K], limit int64) (ottl.ExprFunc[K], error) {
	if limit < 0 {
		return nil, fmt.Errorf("invalid limit for truncate function, %d cannot be negative", limit)
	}
	return func(ctx context.Context, tCtx K) (interface{}, error) {
		val, err := target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}
		if valStr, ok := val.(string); ok {
			if truncatedVal, isTruncated := maybeTruncate(valStr, limit); isTruncated {
				return nil, target.Set(ctx, tCtx, truncatedVal)
			}
		}
		return nil, nil
	}, nil
}

func maybeTruncate(value string, limit int64) (string, bool) {
	if int64(len(value)) > limit {
		return value[:limit], true
	}
	return value, false
}
