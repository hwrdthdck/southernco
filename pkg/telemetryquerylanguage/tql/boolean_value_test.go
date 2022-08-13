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

package tql

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/telemetryquerylanguage/tql/tqltest"
)

func valueFor(x any) Value {
	val := Value{}
	switch v := x.(type) {
	case []byte:
		var b Bytes = v
		val.Bytes = &b
	case string:
		if v == "NAME" {
			// if the string is NAME construct a path of "name".
			val.Path = &Path{
				Fields: []Field{
					{
						Name: "name",
					},
				},
			}
		} else if strings.Contains(v, "ENUM") {
			// if the string contains ENUM construct an EnumSymbol from it.
			val.Enum = (*EnumSymbol)(tqltest.Strp(v))
		} else {
			val.String = tqltest.Strp(v)
		}
	case float64:
		val.Float = tqltest.Floatp(v)
	case int:
		val.Int = tqltest.Intp(int64(v))
	case bool:
		val.Bool = Booleanp(Boolean(v))
	case nil:
		var n IsNil = true
		val.IsNil = &n
	default:
		panic("test error!")
	}
	return val
}

func comparison(left any, right any, op string) *Comparison {
	return &Comparison{
		Left:  valueFor(left),
		Right: valueFor(right),
		Op:    op,
	}
}

func Test_newComparisonEvaluator(t *testing.T) {
	tests := []struct {
		name       string
		comparison *Comparison
		item       interface{}
	}{
		{
			name: "literals match",
			comparison: &Comparison{
				Left: Value{
					String: tqltest.Strp("hello"),
				},
				Right: Value{
					String: tqltest.Strp("hello"),
				},
				Op: "==",
			},
		},
		{
			name: "literals don't match",
			comparison: &Comparison{
				Left: Value{
					String: tqltest.Strp("hello"),
				},
				Right: Value{
					String: tqltest.Strp("goodbye"),
				},
				Op: "!=",
			},
		},
		{
			name: "path expression matches",
			comparison: &Comparison{
				Left: Value{
					Path: &Path{
						Fields: []Field{
							{
								Name: "name",
							},
						},
					},
				},
				Right: Value{
					String: tqltest.Strp("bear"),
				},
				Op: "==",
			},
			item: "bear",
		},
		{
			name: "path expression not matches",
			comparison: &Comparison{
				Left: Value{
					Path: &Path{
						Fields: []Field{
							{
								Name: "name",
							},
						},
					},
				},
				Right: Value{
					String: tqltest.Strp("cat"),
				},
				Op: "!=",
			},
			item: "bear",
		},
		{
			name:       "no condition",
			comparison: nil,
		},

		{
			name: "compare Enum to int",
			comparison: &Comparison{
				Left: Value{
					Enum: (*EnumSymbol)(tqltest.Strp("TEST_ENUM")),
				},
				Right: Value{
					Int: tqltest.Intp(0),
				},
				Op: "==",
			},
		},
		{
			name: "compare int to Enum",
			comparison: &Comparison{
				Left: Value{
					Int: tqltest.Intp(2),
				},
				Op: "==",
				Right: Value{
					Enum: (*EnumSymbol)(tqltest.Strp("TEST_ENUM_TWO")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluate, err := newComparisonEvaluator(tt.comparison, DefaultFunctionsForTests(), testParsePath, testParseEnum)
			assert.NoError(t, err)
			assert.True(t, evaluate(tqltest.TestTransformContext{
				Item: tt.item,
			}))
		})
	}
}

func Test_newComparisonEvaluatorExtended(t *testing.T) {
	tests := []struct {
		name   string
		l      any
		r      any
		op     string
		item   interface{}
		expect bool
	}{
		{"literals match", "hello", "hello", "==", nil, true},
		{"literals don't match", "hello", "goodbye", "!=", nil, true},
		{"path expression matches", "NAME", "bear", "==", "bear", true},
		{"path expression not matches", "NAME", "cat", "!=", "bear", true},
		{"compare Enum to int", "TEST_ENUM", int(0), "==", nil, true},
		{"compare int to Enum", int(2), "TEST_ENUM_TWO", "==", nil, true},
		{"2 > Enum 0", int(2), "TEST_ENUM", ">", nil, true},
		{"2 not < Enum 0", int(2), "TEST_ENUM", "<", nil, false},
		{"6 == 3.14", 6, 3.14, "==", nil, false},
		{"6 != 3.14", 6, 3.14, "!=", nil, true},
		{"6 > 3.14", 6, 3.14, ">", nil, true},
		{"6 >= 3.14", 6, 3.14, ">=", nil, true},
		{"6 < 3.14", 6, 3.14, "<", nil, false},
		{"6 <= 3.14", 6, 3.14, "<=", nil, false},
		{"'foo' > 'bar'", "foo", "bar", ">", nil, true},
		{"'foo' > bear", "foo", "NAME", ">", "bear", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp := comparison(tt.l, tt.r, tt.op)
			evaluate, err := newComparisonEvaluator(comp, DefaultFunctionsForTests(), testParsePath, testParseEnum)
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, evaluate(tqltest.TestTransformContext{
				Item: tt.item,
			}))
		})
	}
}

func Test_newConditionEvaluator_invalid(t *testing.T) {
	tests := []struct {
		name       string
		comparison *Comparison
	}{
		{
			name: "unknown operation",
			comparison: &Comparison{
				Left: Value{
					String: tqltest.Strp("bear"),
				},
				Op: "<>",
				Right: Value{
					String: tqltest.Strp("cat"),
				},
			},
		},
		{
			name: "unknown Path",
			comparison: &Comparison{
				Left: Value{
					Enum: (*EnumSymbol)(tqltest.Strp("SYMBOL_NOT_FOUND")),
				},
				Op: "==",
				Right: Value{
					String: tqltest.Strp("trash"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newComparisonEvaluator(tt.comparison, DefaultFunctionsForTests(), testParsePath, testParseEnum)
			assert.Error(t, err)
		})
	}
}

func Test_newBooleanExpressionEvaluator(t *testing.T) {
	tests := []struct {
		name string
		want bool
		expr *BooleanExpression
	}{
		{"a", false,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(true),
					},
					Right: []*OpAndBooleanValue{
						{
							Operator: "and",
							Value: &BooleanValue{
								ConstExpr: Booleanp(false),
							},
						},
					},
				},
			},
		},
		{"b", true,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(true),
					},
					Right: []*OpAndBooleanValue{
						{
							Operator: "and",
							Value: &BooleanValue{
								ConstExpr: Booleanp(true),
							},
						},
					},
				},
			},
		},
		{"c", false,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(true),
					},
					Right: []*OpAndBooleanValue{
						{
							Operator: "and",
							Value: &BooleanValue{
								ConstExpr: Booleanp(true),
							},
						},
						{
							Operator: "and",
							Value: &BooleanValue{
								ConstExpr: Booleanp(false),
							},
						},
					},
				},
			},
		},
		{"d", true,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(true),
					},
				},
				Right: []*OpOrTerm{
					{
						Operator: "or",
						Term: &Term{
							Left: &BooleanValue{
								ConstExpr: Booleanp(false),
							},
						},
					},
				},
			},
		},
		{"e", true,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(false),
					},
				},
				Right: []*OpOrTerm{
					{
						Operator: "or",
						Term: &Term{
							Left: &BooleanValue{
								ConstExpr: Booleanp(true),
							},
						},
					},
				},
			},
		},
		{"f", false,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(false),
					},
				},
				Right: []*OpOrTerm{
					{
						Operator: "or",
						Term: &Term{
							Left: &BooleanValue{
								ConstExpr: Booleanp(false),
							},
						},
					},
				},
			},
		},
		{"g", true,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(false),
					},
					Right: []*OpAndBooleanValue{
						{
							Operator: "and",
							Value: &BooleanValue{
								ConstExpr: Booleanp(false),
							},
						},
					},
				},
				Right: []*OpOrTerm{
					{
						Operator: "or",
						Term: &Term{
							Left: &BooleanValue{
								ConstExpr: Booleanp(true),
							},
						},
					},
				},
			},
		},
		{"h", true,
			&BooleanExpression{
				Left: &Term{
					Left: &BooleanValue{
						ConstExpr: Booleanp(true),
					},
					Right: []*OpAndBooleanValue{
						{
							Operator: "and",
							Value: &BooleanValue{
								SubExpr: &BooleanExpression{
									Left: &Term{
										Left: &BooleanValue{
											ConstExpr: Booleanp(true),
										},
									},
									Right: []*OpOrTerm{
										{
											Operator: "or",
											Term: &Term{
												Left: &BooleanValue{
													ConstExpr: Booleanp(false),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluate, err := newBooleanExpressionEvaluator(tt.expr, DefaultFunctionsForTests(), testParsePath, testParseEnum)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, evaluate(tqltest.TestTransformContext{
				Item: nil,
			}))
		})
	}
}
