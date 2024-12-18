// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func GetMapValue[K any](ctx context.Context, tCtx K, m pcommon.Map, keys []ottl.Key[K]) (any, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("cannot get map value without keys")
	}

	s, err := keys[0].String(ctx, tCtx)
	if err != nil {
		return nil, err
	}
	if s == nil {
		resString, err := FetchValueFromPath[K, string](ctx, tCtx, keys[0])
		if err != nil {
			return nil, fmt.Errorf("non-string indexing is not supported")
		}
		s = resString
	}

	val, ok := m.Get(*s)
	if !ok {
		return nil, nil
	}

	return getIndexableValue[K](ctx, tCtx, val, keys[1:])
}

func SetMapValue[K any](ctx context.Context, tCtx K, m pcommon.Map, keys []ottl.Key[K], val any) error {
	if len(keys) == 0 {
		return fmt.Errorf("cannot set map value without key")
	}

	s, err := keys[0].String(ctx, tCtx)
	if err != nil {
		return err
	}
	if s == nil {
		resString, err := FetchValueFromPath[K, string](ctx, tCtx, keys[0])
		if err != nil {
			return fmt.Errorf("non-string indexing is not supported")
		}
		s = resString
	}

	currentValue, ok := m.Get(*s)
	if !ok {
		currentValue = m.PutEmpty(*s)
	}
	return setIndexableValue[K](ctx, tCtx, currentValue, val, keys[1:])
}

func FetchValueFromPath[K any, T int64 | string](ctx context.Context, tCtx K, key ottl.Key[K]) (*T, error) {
	p, err := key.PathGetter(ctx, tCtx)
	if err != nil {
		return nil, err
	}
	res, err := p.Get(ctx, tCtx)
	if err != nil {
		return nil, err
	}
	resVal, ok := res.(T)
	if !ok {
		return nil, fmt.Errorf("casting not successful")
	}
	return &resVal, nil
}
