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
	"fmt"
	"strings"
	"unicode"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func ConvertCase[K any](target ottl.Getter[K], toCase string) (ottl.ExprFunc[K], error) {
	var caseInvalid = true
	for _, validCase := range []string{"lower", "upper", "snake"} {
		if toCase == validCase {
			caseInvalid = false
			break
		}
	}
	if caseInvalid {
		return nil, fmt.Errorf(`invalid case "%s", allowed cases are "lower", "upper", or "snake"`, toCase)
	}

	return func(ctx context.Context, tCtx K) (interface{}, error) {
		val, err := target.Get(ctx, tCtx)

		if err != nil {
			return nil, err
		}

		if val != nil {

			if valStr, ok := val.(string); ok {

				if valStr == "" {
					return valStr, nil
				}

				switch toCase {
				// Convert string to snake case (someName -> some_name)
				case "snake":
					return convertCaseSnake(valStr), nil

				// Convert string to lowercase (SOME_NAME -> some_name)
				case "lower":
					return convertCaseLower(valStr), nil

				// Convert string to uppercase (some_name -> SOME_NAME)
				case "upper":
					return convertCaseUpper(valStr), nil

				default:
					return nil, fmt.Errorf(`error handling unexpected case "%s"`, toCase)
				}
			}
		}
		return nil, nil
	}, nil
}

func convertCaseLower(input string) string {
	return strings.ToLower(input)
}

func convertCaseUpper(input string) string {
	return strings.ToUpper(input)
}

func convertCaseSnake(input string) string {
	if len(input) == 1 {
		return strings.ToLower(input)
	}
	source := []rune(input)
	dist := strings.Builder{}
	dist.Grow(len(input) + len(input)/3)
	skipNext := false

	for i := 0; i < len(source); i++ {
		cur := source[i]

		switch cur {
		case '-', '_':
			dist.WriteRune('_')
			skipNext = true
			continue
		}
		if unicode.IsLower(cur) || unicode.IsDigit(cur) {
			dist.WriteRune(cur)
			continue
		}

		if i == 0 {
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}

		last := source[i-1]

		if (!unicode.IsLetter(last)) || unicode.IsLower(last) {
			if skipNext {
				skipNext = false
			} else {
				dist.WriteRune('_')
			}
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}

		if i < len(source)-1 {
			next := source[i+1]
			if unicode.IsLower(next) {
				if skipNext {
					skipNext = false
				} else {
					dist.WriteRune('_')
				}
				dist.WriteRune(unicode.ToLower(cur))
				continue
			}
		}

		dist.WriteRune(unicode.ToLower(cur))
	}

	return dist.String()
}
