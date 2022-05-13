// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entry // import "github.com/open-telemetry/opentelemetry-log-collection/entry"

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	AttributesPrefix = "attributes"
	ResourcePrefix   = "resource"
	BodyPrefix       = "body"
)

// Field represents a potential field on an entry.
// It is used to get, set, and delete values at this field.
// It is deserialized from JSON dot notation.
type Field struct {
	FieldInterface
}

// FieldInterface is a field on an entry.
type FieldInterface interface {
	Get(*Entry) (interface{}, bool)
	Set(entry *Entry, value interface{}) error
	Delete(entry *Entry) (interface{}, bool)
	String() string
}

// UnmarshalJSON will unmarshal a field from JSON
func (f *Field) UnmarshalJSON(raw []byte) error {
	var s string
	err := json.Unmarshal(raw, &s)
	if err != nil {
		return err
	}
	*f, err = NewField(s)
	return err
}

// UnmarshalYAML will unmarshal a field from YAML
func (f *Field) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	*f, err = NewField(s)
	return err
}

func NewField(s string) (Field, error) {
	keys, err := fromJSONDot(s)
	if err != nil {
		return Field{}, fmt.Errorf("splitting field: %s", err)
	}

	switch keys[0] {
	case AttributesPrefix:
		if len(keys) == 1 {
			return Field{}, fmt.Errorf("attributes cannot be referenced without subfield")
		}
		return NewAttributeField(keys[1:]...), nil
	case ResourcePrefix:
		if len(keys) == 1 {
			return Field{}, fmt.Errorf("resource cannot be referenced without subfield")
		}
		return NewResourceField(keys[1:]...), nil
	case BodyPrefix:
		return NewBodyField(keys[1:]...), nil
	default:
		return Field{}, fmt.Errorf("unrecognized prefix")
	}
}

// MarshalJSON will marshal a field into JSON
func (f Field) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, f.String())), nil
}

// MarshalYAML will marshal a field into YAML
func (f Field) MarshalYAML() (interface{}, error) {
	return f.String(), nil
}

type splitState uint

const (
	// Begin is the beginning state of a field split
	Begin splitState = iota
	// InBracket is the state of a field split inside a bracket
	InBracket
	// InQuote is the state of a field split inside a quote
	InQuote
	// OutQuote is the state of a field split outside a quote
	OutQuote
	// OutBracket is the state of a field split outside a bracket
	OutBracket
	// InUnbracketedToken is the state field split on any token outside brackets
	InUnbracketedToken
)

func fromJSONDot(s string) ([]string, error) {
	fields := make([]string, 0, 1)

	state := Begin
	var quoteChar rune
	var tokenStart int

	for i, c := range s {
		switch state {
		case Begin:
			if c == '[' {
				state = InBracket
				continue
			}
			tokenStart = i
			state = InUnbracketedToken
		case InBracket:
			if !(c == '\'' || c == '"') {
				return nil, fmt.Errorf("strings in brackets must be surrounded by quotes")
			}
			state = InQuote
			quoteChar = c
			tokenStart = i + 1
		case InQuote:
			if c == quoteChar {
				fields = append(fields, s[tokenStart:i])
				state = OutQuote
			}
		case OutQuote:
			if c != ']' {
				return nil, fmt.Errorf("found characters between closed quote and closing bracket")
			}
			state = OutBracket
		case OutBracket:
			switch c {
			case '.':
				state = InUnbracketedToken
				tokenStart = i + 1
			case '[':
				state = InBracket
			default:
				return nil, fmt.Errorf("bracketed access must be followed by a dot or another bracketed access")
			}
		case InUnbracketedToken:
			if c == '.' {
				fields = append(fields, s[tokenStart:i])
				tokenStart = i + 1
			} else if c == '[' {
				fields = append(fields, s[tokenStart:i])
				state = InBracket
			}
		}
	}

	switch state {
	case InBracket, OutQuote:
		return nil, fmt.Errorf("found unclosed left bracket")
	case InQuote:
		if quoteChar == '"' {
			return nil, fmt.Errorf("found unclosed double quote")
		}
		return nil, fmt.Errorf("found unclosed single quote")
	case InUnbracketedToken:
		fields = append(fields, s[tokenStart:])
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("fields size is 0")
	}

	return fields, nil
}

// toJSONDot returns the JSON dot notation for a field.
func toJSONDot(prefix string, keys []string) string {
	if len(keys) == 0 {
		return prefix
	}

	containsDots := false
	for _, key := range keys {
		if strings.Contains(key, ".") {
			containsDots = true
		}
	}

	var b strings.Builder
	b.WriteString(prefix)
	if containsDots {
		for _, key := range keys {
			b.WriteString(`['`)
			b.WriteString(key)
			b.WriteString(`']`)
		}
	} else {
		b.WriteString(".")
		for i, key := range keys {
			if i != 0 {
				b.WriteString(".")
			}
			b.WriteString(key)
		}
	}

	return b.String()
}

// getNestedMap will get a nested map assigned to a key.
// If the map does not exist, it will create and return it.
func getNestedMap(currentMap map[string]interface{}, key string) map[string]interface{} {
	currentValue, ok := currentMap[key]
	if !ok {
		currentMap[key] = map[string]interface{}{}
	}

	nextMap, ok := currentValue.(map[string]interface{})
	if !ok {
		nextMap = map[string]interface{}{}
		currentMap[key] = nextMap
	}

	return nextMap
}
