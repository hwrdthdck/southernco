// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attraction // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/model/pdata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterhelper"
)

// Settings specifies the processor settings.
type Settings struct {
	// Actions specifies the list of attributes to act on.
	// The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT, COERCE}.
	// This is a required field.
	Actions []ActionKeyValue `mapstructure:"actions"`
}

// ActionKeyValue specifies the attribute key to act upon.
type ActionKeyValue struct {
	// Key specifies the attribute to act upon.
	// This is a required field.
	Key string `mapstructure:"key"`

	// Value specifies the value to populate for the key.
	// The type of the value is inferred from the configuration.
	Value interface{} `mapstructure:"value"`

	// A regex pattern  must be specified for the action EXTRACT.
	// It uses the attribute specified by `key' to extract values from
	// The target keys are inferred based on the names of the matcher groups
	// provided and the names will be inferred based on the values of the
	// matcher group.
	// Note: All subexpressions must have a name.
	// Note: The value type of the source key must be a string. If it isn't,
	// no extraction will occur.
	RegexPattern string `mapstructure:"pattern"`

	// FromAttribute specifies the attribute to use to populate
	// the value. If the attribute doesn't exist, no action is performed.
	FromAttribute string `mapstructure:"from_attribute"`

	// FromContext specifies the context value to use to populate
	// the value. The values would be searched in client.Info.Metadata.
	// If the key doesn't exist, no action is performed.
	// If the key has multiple values the values will be joined with `;` separator.
	FromContext string `mapstructure:"from_context"`

	// TargetType specifies the target type of an attribute to be coerced
	// If the key doesn't exist, no action is performed.
	// If the value cannot be coerced, the null value will be inserted.
	TargetType string `mapstructure:"target_type"`

	// Action specifies the type of action to perform.
	// The set of values are {INSERT, UPDATE, UPSERT, DELETE, HASH}.
	// Both lower case and upper case are supported.
	// INSERT -  Inserts the key/value to attributes when the key does not exist.
	//           No action is applied to attributes where the key already exists.
	//           Either Value, FromAttribute or FromContext must be set.
	// UPDATE -  Updates an existing key with a value. No action is applied
	//           to attributes where the key does not exist.
	//           Either Value, FromAttribute or FromContext must be set.
	// UPSERT -  Performs insert or update action depending on the attributes
	//           containing the key. The key/value is inserted to attributes
	//           that did not originally have the key. The key/value is updated
	//           for attributes where the key already existed.
	//           Either Value, FromAttribute or FromContext must be set.
	// DELETE  - Deletes the attribute. If the key doesn't exist,
	//           no action is performed.
	// HASH    - Calculates the SHA-1 hash of an existing value and overwrites the
	//           value with it's SHA-1 hash result.
	// EXTRACT - Extracts values using a regular expression rule from the input
	//           'key' to target keys specified in the 'rule'. If a target key
	//           already exists, it will be overridden.
	// COERCE  - coerces the type of an existing attribute, or sets it to the null
	//	         value if the existing value cannot be coerced
	// This is a required field.
	Action Action `mapstructure:"action"`
}

func (a *ActionKeyValue) valueSourceCount() int {
	count := 0
	if a.Value != nil {
		count++
	}

	if a.FromAttribute != "" {
		count++
	}

	if a.FromContext != "" {
		count++
	}
	return count
}

// Action is the enum to capture the four types of actions to perform on an
// attribute.
type Action string

const (
	// INSERT adds the key/value to attributes when the key does not exist.
	// No action is applied to attributes where the key already exists.
	INSERT Action = "insert"

	// UPDATE updates an existing key with a value. No action is applied
	// to attributes where the key does not exist.
	UPDATE Action = "update"

	// UPSERT performs the INSERT or UPDATE action. The key/value is
	// inserted to attributes that did not originally have the key. The key/value is
	// updated for attributes where the key already existed.
	UPSERT Action = "upsert"

	// DELETE deletes the attribute. If the key doesn't exist, no action is performed.
	DELETE Action = "delete"

	// HASH calculates the SHA-1 hash of an existing value and overwrites the
	// value with it's SHA-1 hash result.
	HASH Action = "hash"

	// EXTRACT extracts values using a regular expression rule from the input
	// 'key' to target keys specified in the 'rule'. If a target key already
	// exists, it will be overridden.
	EXTRACT Action = "extract"

	// COERCE coerces the type of an existing attribute, or sets it to the null
	// value if the existing value cannot be coerced
	COERCE Action = "coerce"
)

type attributeAction struct {
	Key           string
	FromAttribute string
	FromContext   string
	TargetType    string
	// Compiled regex if provided
	Regex *regexp.Regexp
	// Attribute names extracted from the regexp's subexpressions.
	AttrNames []string
	// Number of non empty strings in above array

	// TODO https://go.opentelemetry.io/collector/issues/296
	// Do benchmark testing between having action be of type string vs integer.
	// The reason is attributes processor will most likely be commonly used
	// and could impact performance.
	Action         Action
	AttributeValue *pdata.AttributeValue
}

// AttrProc is an attribute processor.
type AttrProc struct {
	actions []attributeAction
}

// NewAttrProc validates that the input configuration has all of the required fields for the processor
// and returns a AttrProc to be used to process attributes.
// An error is returned if there are any invalid inputs.
func NewAttrProc(settings *Settings) (*AttrProc, error) {
	var attributeActions []attributeAction
	for i, a := range settings.Actions {
		// `key` is a required field
		if a.Key == "" {
			return nil, fmt.Errorf("error creating AttrProc due to missing required field \"key\" at the %d-th actions", i)
		}

		// Convert `action` to lowercase for comparison.
		a.Action = Action(strings.ToLower(string(a.Action)))
		action := attributeAction{
			Key:    a.Key,
			Action: a.Action,
		}

		valueSourceCount := a.valueSourceCount()

		switch a.Action {
		case INSERT, UPDATE, UPSERT:
			if valueSourceCount == 0 {
				return nil, fmt.Errorf("error creating AttrProc. Either field \"value\", \"from_attribute\" or \"from_context\" setting must be specified for %d-th action", i)
			}

			if valueSourceCount > 1 {
				return nil, fmt.Errorf("error creating AttrProc due to multiple value sources being set at the %d-th actions", i)
			}
			if a.RegexPattern != "" {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use the \"pattern\" field. This must not be specified for %d-th action", a.Action, i)
			}
			if a.TargetType != "" {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use the \"target_type\" field. This must not be specified for %d-th action", a.Action, i)
			}
			// Convert the raw value from the configuration to the internal trace representation of the value.
			if a.Value != nil {
				val, err := filterhelper.NewAttributeValueRaw(a.Value)
				if err != nil {
					return nil, err
				}
				action.AttributeValue = &val
			} else {
				action.FromAttribute = a.FromAttribute
				action.FromContext = a.FromContext
			}
		case HASH, DELETE:
			if valueSourceCount > 0 || a.RegexPattern != "" {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use value sources or \"pattern\" field. These must not be specified for %d-th action", a.Action, i)
			}
			if a.TargetType != "" {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use the \"target_type\" field. This must not be specified for %d-th action", a.Action, i)
			}
		case EXTRACT:
			if valueSourceCount > 0 {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use a value source field. These must not be specified for %d-th action", a.Action, i)
			}
			if a.RegexPattern == "" {
				return nil, fmt.Errorf("error creating AttrProc due to missing required field \"pattern\" for action \"%s\" at the %d-th action", a.Action, i)
			}
			if a.TargetType != "" {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use the \"target_type\" field. This must not be specified for %d-th action", a.Action, i)
			}
			re, err := regexp.Compile(a.RegexPattern)
			if err != nil {
				return nil, fmt.Errorf("error creating AttrProc. Field \"pattern\" has invalid pattern: \"%s\" to be set at the %d-th actions", a.RegexPattern, i)
			}
			attrNames := re.SubexpNames()
			if len(attrNames) <= 1 {
				return nil, fmt.Errorf("error creating AttrProc. Field \"pattern\" contains no named matcher groups at the %d-th actions", i)
			}

			for subExpIndex := 1; subExpIndex < len(attrNames); subExpIndex++ {
				if attrNames[subExpIndex] == "" {
					return nil, fmt.Errorf("error creating AttrProc. Field \"pattern\" contains at least one unnamed matcher group at the %d-th actions", i)
				}
			}
			action.Regex = re
			action.AttrNames = attrNames
		case COERCE:
			if valueSourceCount > 0 || a.RegexPattern != "" {
				return nil, fmt.Errorf("error creating AttrProc. Action \"%s\" does not use value sources or \"pattern\" field. These must not be specified for %d-th action", a.Action, i)
			}
			switch a.TargetType {
			case "string":
			case "int":
			case "double":
			case "":
				return nil, fmt.Errorf("error creating AttrProc due to missing required field \"target_type\" for action \"%s\" at the %d-th action", a.Action, i)
			default:
				return nil, fmt.Errorf("error creating AttrProc due to invalid value \"%s\" in field \"target_type\" for action \"%s\" at the %d-th action", a.TargetType, a.Action, i)
			}
			action.TargetType = a.TargetType
		default:
			return nil, fmt.Errorf("error creating AttrProc due to unsupported action %q at the %d-th actions", a.Action, i)
		}

		attributeActions = append(attributeActions, action)
	}
	return &AttrProc{actions: attributeActions}, nil
}

// Process applies the AttrProc to an attribute map.
func (ap *AttrProc) Process(ctx context.Context, attrs pdata.AttributeMap) {
	for _, action := range ap.actions {
		// TODO https://go.opentelemetry.io/collector/issues/296
		// Do benchmark testing between having action be of type string vs integer.
		// The reason is attributes processor will most likely be commonly used
		// and could impact performance.
		switch action.Action {
		case DELETE:
			attrs.Delete(action.Key)
		case INSERT:
			av, found := getSourceAttributeValue(ctx, action, attrs)
			if !found {
				continue
			}
			attrs.Insert(action.Key, av)
		case UPDATE:
			av, found := getSourceAttributeValue(ctx, action, attrs)
			if !found {
				continue
			}
			attrs.Update(action.Key, av)
		case UPSERT:
			av, found := getSourceAttributeValue(ctx, action, attrs)
			if !found {
				continue
			}
			attrs.Upsert(action.Key, av)
		case HASH:
			hashAttribute(action, attrs)
		case EXTRACT:
			extractAttributes(action, attrs)
		case COERCE:
			coerceAttribute(action, attrs)
		}
	}
}

func getAttributeValueFromContext(ctx context.Context, key string) (pdata.AttributeValue, bool) {
	ci := client.FromContext(ctx)
	vals := ci.Metadata.Get(key)

	if len(vals) == 0 {
		return pdata.AttributeValue{}, false
	}

	return pdata.NewAttributeValueString(strings.Join(vals, ";")), true
}

func getSourceAttributeValue(ctx context.Context, action attributeAction, attrs pdata.AttributeMap) (pdata.AttributeValue, bool) {
	// Set the key with a value from the configuration.
	if action.AttributeValue != nil {
		return *action.AttributeValue, true
	}

	if action.FromContext != "" {
		return getAttributeValueFromContext(ctx, action.FromContext)
	}

	return attrs.Get(action.FromAttribute)
}

func hashAttribute(action attributeAction, attrs pdata.AttributeMap) {
	if value, exists := attrs.Get(action.Key); exists {
		sha1Hasher(value)
	}
}

func coerceAttribute(action attributeAction, attrs pdata.AttributeMap) {
	if value, exists := attrs.Get(action.Key); exists {
		coerceValue(action.TargetType, value)
	}
}

func extractAttributes(action attributeAction, attrs pdata.AttributeMap) {
	value, found := attrs.Get(action.Key)

	// Extracting values only functions on strings.
	if !found || value.Type() != pdata.AttributeValueTypeString {
		return
	}

	// Note: The number of matches will always be equal to number of
	// subexpressions.
	matches := action.Regex.FindStringSubmatch(value.StringVal())
	if matches == nil {
		return
	}

	// Start from index 1, which is the first submatch (index 0 is the entire
	// match).
	for i := 1; i < len(matches); i++ {
		attrs.UpsertString(action.AttrNames[i], matches[i])
	}
}
