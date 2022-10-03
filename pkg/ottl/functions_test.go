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

package ottl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottltest"
)

func Test_NewFunctionCall_invalid(t *testing.T) {
	functions := make(map[string]interface{})
	functions["testing_error"] = functionThatHasAnError
	functions["testing_getsetter"] = functionWithGetSetter
	functions["testing_getter"] = functionWithGetter
	functions["testing_multiple_args"] = functionWithMultipleArgs
	functions["testing_string"] = functionWithString
	functions["testing_string_slice"] = functionWithStringSlice
	functions["testing_byte_slice"] = functionWithByteSlice
	functions["testing_string_variadic"] = functionWithStringVariadic
	functions["testing_enum"] = functionWithEnum
	functions["testing_telemetry_settings_first"] = functionWithTelemetrySettingsFirst

	p := NewParser(
		functions,
		testParsePath,
		testParseEnum,
		component.TelemetrySettings{},
	)

	tests := []struct {
		name string
		inv  invocation
	}{
		{
			name: "unknown function",
			inv: invocation{
				Function:  "unknownfunc",
				Arguments: []value{},
			},
		},
		{
			name: "not accessor",
			inv: invocation{
				Function: "testing_getsetter",
				Arguments: []value{
					{
						String: ottltest.Strp("not path"),
					},
				},
			},
		},
		{
			name: "not reader (invalid function)",
			inv: invocation{
				Function: "testing_getter",
				Arguments: []value{
					{
						Invocation: &invocation{
							Function: "unknownfunc",
						},
					},
				},
			},
		},
		{
			name: "not enough args",
			inv: invocation{
				Function: "testing_multiple_args",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
					{
						String: ottltest.Strp("test"),
					},
				},
			},
		},
		{
			name: "too many args",
			inv: invocation{
				Function: "testing_multiple_args",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
				},
			},
		},
		{
			name: "not enough args with telemetrySettings",
			inv: invocation{
				Function: "testing_telemetry_settings_first",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
				},
			},
		},
		{
			name: "too many args with telemetrySettings",
			inv: invocation{
				Function: "testing_telemetry_settings_first",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						Int: ottltest.Intp(10),
					},
					{
						Int: ottltest.Intp(10),
					},
				},
			},
		},
		{
			name: "not matching arg type",
			inv: invocation{
				Function: "testing_string",
				Arguments: []value{
					{
						Int: ottltest.Intp(10),
					},
				},
			},
		},
		{
			name: "not matching arg type when byte slice",
			inv: invocation{
				Function: "testing_byte_slice",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
				},
			},
		},
		{
			name: "mismatching slice element type",
			inv: invocation{
				Function: "testing_string_slice",
				Arguments: []value{
					{
						List: &list{
							Values: []value{
								{
									String: ottltest.Strp("test"),
								},
								{
									Int: ottltest.Intp(10),
								},
							},
						},
					},
				},
			},
		},
		{
			name: "mismatching slice argument type",
			inv: invocation{
				Function: "testing_string_slice",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
				},
			},
		},
		{
			name: "mismatching variadic argument type",
			inv: invocation{
				Function: "testing_string_variadic",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
					{
						Int: ottltest.Intp(10),
					},
				},
			},
		},
		{
			name: "function call returns error",
			inv: invocation{
				Function: "testing_error",
			},
		},
		{
			name: "Enum not found",
			inv: invocation{
				Function: "testing_enum",
				Arguments: []value{
					{
						Enum: (*EnumSymbol)(ottltest.Strp("SYMBOL_NOT_FOUND")),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.newFunctionCall(tt.inv)
			assert.Error(t, err)
		})
	}
}

func Test_NewFunctionCall(t *testing.T) {
	p := NewParser(
		defaultFunctionsForTests(),
		testParsePath,
		testParseEnum,
		component.TelemetrySettings{},
	)

	tests := []struct {
		name string
		inv  invocation
		want any
	}{
		{
			name: "empty variadic arg",
			inv: invocation{
				Function:  "testing_string_variadic",
				Arguments: []value{},
			},
			want: 0,
		},
		{
			name: "string variadic arg",
			inv: invocation{
				Function: "testing_string_variadic",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
				},
			},
			want: 3,
		},
		{
			name: "float variadic arg",
			inv: invocation{
				Function: "testing_float_variadic",
				Arguments: []value{
					{
						Float: ottltest.Floatp(1.1),
					},
					{
						Float: ottltest.Floatp(1.2),
					},
					{
						Float: ottltest.Floatp(1.3),
					},
				},
			},
			want: 3,
		},
		{
			name: "int variadic arg",
			inv: invocation{
				Function: "testing_int_variadic",
				Arguments: []value{
					{
						Int: ottltest.Intp(1),
					},
					{
						Int: ottltest.Intp(2),
					},
					{
						Int: ottltest.Intp(3),
					},
				},
			},
			want: 3,
		},
		{
			name: "getter variadic arg",
			inv: invocation{
				Function: "testing_getter_variadic",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						Int: ottltest.Intp(1),
					},
					{
						Float: ottltest.Floatp(1.1),
					},
					{
						Bool: (*boolean)(ottltest.Boolp(true)),
					},
					{
						Enum: (*EnumSymbol)(ottltest.Strp("TEST_ENUM")),
					},
					{
						Invocation: &invocation{
							Function: "testing_getter",
							Arguments: []value{
								{
									Path: &Path{
										Fields: []Field{
											{
												Name: "name",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: 7,
		},
		{
			name: "empty slice arg",
			inv: invocation{
				Function: "testing_string_slice",
				Arguments: []value{
					{
						List: &list{
							Values: []value{},
						},
					},
				},
			},
			want: 0,
		},
		{
			name: "string slice arg",
			inv: invocation{
				Function: "testing_string_slice",
				Arguments: []value{
					{
						List: &list{
							Values: []value{
								{
									String: ottltest.Strp("test"),
								},
								{
									String: ottltest.Strp("test"),
								},
								{
									String: ottltest.Strp("test"),
								},
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "float slice arg",
			inv: invocation{
				Function: "testing_float_slice",
				Arguments: []value{
					{
						List: &list{
							Values: []value{
								{
									Float: ottltest.Floatp(1.1),
								},
								{
									Float: ottltest.Floatp(1.2),
								},
								{
									Float: ottltest.Floatp(1.3),
								},
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "int slice arg",
			inv: invocation{
				Function: "testing_int_slice",
				Arguments: []value{
					{
						List: &list{
							Values: []value{
								{
									Int: ottltest.Intp(1),
								},
								{
									Int: ottltest.Intp(1),
								},
								{
									Int: ottltest.Intp(1),
								},
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "getter slice arg",
			inv: invocation{
				Function: "testing_getter_slice",
				Arguments: []value{
					{
						List: &list{
							Values: []value{
								{
									Path: &Path{
										Fields: []Field{
											{
												Name: "name",
											},
										},
									},
								},
								{
									String: ottltest.Strp("test"),
								},
								{
									Int: ottltest.Intp(1),
								},
								{
									Float: ottltest.Floatp(1.1),
								},
								{
									Bool: (*boolean)(ottltest.Boolp(true)),
								},
								{
									Enum: (*EnumSymbol)(ottltest.Strp("TEST_ENUM")),
								},
								{
									Invocation: &invocation{
										Function: "testing_getter",
										Arguments: []value{
											{
												Path: &Path{
													Fields: []Field{
														{
															Name: "name",
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
			},
			want: 7,
		}, {
			name: "setter arg",
			inv: invocation{
				Function: "testing_setter",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "getsetter arg",
			inv: invocation{
				Function: "testing_getsetter",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "getter arg",
			inv: invocation{
				Function: "testing_getter",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "getter arg with nil literal",
			inv: invocation{
				Function: "testing_getter",
				Arguments: []value{
					{
						IsNil: (*isNil)(ottltest.Boolp(true)),
					},
				},
			},
			want: nil,
		},
		{
			name: "string arg",
			inv: invocation{
				Function: "testing_string",
				Arguments: []value{
					{
						String: ottltest.Strp("test"),
					},
				},
			},
			want: nil,
		},
		{
			name: "float arg",
			inv: invocation{
				Function: "testing_float",
				Arguments: []value{
					{
						Float: ottltest.Floatp(1.1),
					},
				},
			},
			want: nil,
		},
		{
			name: "int arg",
			inv: invocation{
				Function: "testing_int",
				Arguments: []value{
					{
						Int: ottltest.Intp(1),
					},
				},
			},
			want: nil,
		},
		{
			name: "bool arg",
			inv: invocation{
				Function: "testing_bool",
				Arguments: []value{
					{
						Bool: (*boolean)(ottltest.Boolp(true)),
					},
				},
			},
			want: nil,
		},
		{
			name: "byteSlice arg",
			inv: invocation{
				Function: "testing_byte_slice",
				Arguments: []value{
					{
						Bytes: (*byteSlice)(&[]byte{1, 2, 3, 4, 5, 6, 7, 8}),
					},
				},
			},
			want: nil,
		},
		{
			name: "multiple args",
			inv: invocation{
				Function: "testing_multiple_args",
				Arguments: []value{
					{
						Path: &Path{
							Fields: []Field{
								{
									Name: "name",
								},
							},
						},
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						Float: ottltest.Floatp(1.1),
					},
					{
						Int: ottltest.Intp(1),
					},
					{
						String: ottltest.Strp("test"),
					},
					{
						String: ottltest.Strp("test"),
					},
				},
			},
			want: nil,
		},
		{
			name: "Enum arg",
			inv: invocation{
				Function: "testing_enum",
				Arguments: []value{
					{
						Enum: (*EnumSymbol)(ottltest.Strp("TEST_ENUM")),
					},
				},
			},
			want: nil,
		},
		{
			name: "telemetrySettings first",
			inv: invocation{
				Function: "testing_telemetry_settings_first",
				Arguments: []value{
					{
						String: ottltest.Strp("test0"),
					},
					{
						String: ottltest.Strp("test1"),
					},
					{
						Int: ottltest.Intp(1),
					},
				},
			},
			want: nil,
		},
		{
			name: "telemetrySettings middle",
			inv: invocation{
				Function: "testing_telemetry_settings_middle",
				Arguments: []value{
					{
						String: ottltest.Strp("test0"),
					},
					{
						String: ottltest.Strp("test1"),
					},
					{
						Int: ottltest.Intp(1),
					},
				},
			},
			want: nil,
		},
		{
			name: "telemetrySettings last",
			inv: invocation{
				Function: "testing_telemetry_settings_last",
				Arguments: []value{
					{
						String: ottltest.Strp("test0"),
					},
					{
						String: ottltest.Strp("test1"),
					},
					{
						Int: ottltest.Intp(1),
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := p.newFunctionCall(tt.inv)
			assert.NoError(t, err)

			if tt.want != nil {
				assert.Equal(t, tt.want, fn(nil))
			}
		})
	}
}

func functionWithStringSlice(strs []string) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(strs)
	}, nil
}

func functionWithFloatSlice(floats []float64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(floats)
	}, nil
}

func functionWithIntSlice(ints []int64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(ints)
	}, nil
}

func functionWithByteSlice(bytes []byte) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(bytes)
	}, nil
}

func functionWithGetterSlice(getters []Getter[interface{}]) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(getters)
	}, nil
}

func functionWithStringVariadic(strs ...string) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(strs)
	}, nil
}

func functionWithFloatVariadic(floats ...float64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(floats)
	}, nil
}

func functionWithIntVariadic(ints ...int64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(ints)
	}, nil
}

func functionWithGetterVariadic(getters ...Getter[interface{}]) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return len(getters)
	}, nil
}

func functionWithSetter(Setter[interface{}]) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithGetSetter(GetSetter[interface{}]) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithGetter(Getter[interface{}]) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithString(string) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithFloat(float64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithInt(int64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithBool(bool) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithMultipleArgs(GetSetter[interface{}], string, float64, int64, ...string) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionThatHasAnError() (ExprFunc[interface{}], error) {
	err := errors.New("testing")
	return func(interface{}) interface{} {
		return "anything"
	}, err
}

func functionWithEnum(Enum) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithTelemetrySettingsFirst(component.TelemetrySettings, string, string, int64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithTelemetrySettingsMiddle(string, string, component.TelemetrySettings, int64) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func functionWithTelemetrySettingsLast(string, string, int64, component.TelemetrySettings) (ExprFunc[interface{}], error) {
	return func(interface{}) interface{} {
		return "anything"
	}, nil
}

func defaultFunctionsForTests() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["testing_string_slice"] = functionWithStringSlice
	functions["testing_float_slice"] = functionWithFloatSlice
	functions["testing_int_slice"] = functionWithIntSlice
	functions["testing_byte_slice"] = functionWithByteSlice
	functions["testing_getter_slice"] = functionWithGetterSlice
	functions["testing_string_variadic"] = functionWithStringVariadic
	functions["testing_float_variadic"] = functionWithFloatVariadic
	functions["testing_int_variadic"] = functionWithIntVariadic
	functions["testing_getter_variadic"] = functionWithGetterVariadic
	functions["testing_setter"] = functionWithSetter
	functions["testing_getsetter"] = functionWithGetSetter
	functions["testing_getter"] = functionWithGetter
	functions["testing_string"] = functionWithString
	functions["testing_float"] = functionWithFloat
	functions["testing_int"] = functionWithInt
	functions["testing_bool"] = functionWithBool
	functions["testing_multiple_args"] = functionWithMultipleArgs
	functions["testing_enum"] = functionWithEnum
	functions["testing_telemetry_settings_first"] = functionWithTelemetrySettingsFirst
	functions["testing_telemetry_settings_middle"] = functionWithTelemetrySettingsMiddle
	functions["testing_telemetry_settings_last"] = functionWithTelemetrySettingsLast
	return functions
}
