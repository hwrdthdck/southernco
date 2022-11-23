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

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func MergeMaps[K any](target ottl.Getter[K], source ottl.Getter[K]) (ottl.ExprFunc[K], error) {
	return func(ctx context.Context, tCtx K) (interface{}, error) {
		targetVal, err := target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}
		if targetMap, ok := targetVal.(pcommon.Map); ok {
			val, err := source.Get(ctx, tCtx)
			if err != nil {
				return nil, err
			}
			if valueMap, ok := val.(pcommon.Map); ok {
				valueMap.Range(func(k string, v pcommon.Value) bool {
					if tv, ok := targetMap.Get(k); ok {
						v.CopyTo(tv)
					} else {
						tv := targetMap.PutEmpty(k)
						v.CopyTo(tv)
					}
					return true
				})
			}
		}
		return nil, nil
	}, nil
}
