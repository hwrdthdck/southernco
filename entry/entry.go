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

package entry

import (
	"fmt"
	"os"
	"time"
)

const defaultTimestampEnv = "STANZA_DEFAULT_TIMESTAMP"

func getNow() func() time.Time {
	env := os.Getenv(defaultTimestampEnv)
	if env == "" {
		return time.Now
	}

	parsed, err := time.Parse(time.RFC3339, env)
	if err != nil {
		panic(fmt.Sprintf("failed parsing default timestamp: %s", err))
	}

	return func() time.Time {
		return parsed
	}
}

var now = getNow()

// Entry is a flexible representation of log data associated with a timestamp.
type Entry struct {
	Timestamp    time.Time         `json:"timestamp"               yaml:"timestamp"`
	Severity     Severity          `json:"severity"                yaml:"severity"`
	SeverityText string            `json:"severity_text,omitempty" yaml:"severity_text,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"        yaml:"labels,omitempty"`
	Resource     map[string]string `json:"resource,omitempty"      yaml:"resource,omitempty"`
	Record       interface{}       `json:"record"                  yaml:"record"`
}

// New will create a new log entry with current timestamp and an empty record.
func New() *Entry {
	return &Entry{
		Timestamp: now(),
	}
}

// AddLabel will add a key/value pair to the entry's labels.
func (entry *Entry) AddLabel(key, value string) {
	if entry.Labels == nil {
		entry.Labels = make(map[string]string)
	}
	entry.Labels[key] = value
}

// AddResourceKey wil add a key/value pair to the entry's resource.
func (entry *Entry) AddResourceKey(key, value string) {
	if entry.Resource == nil {
		entry.Resource = make(map[string]string)
	}
	entry.Resource[key] = value
}

// Get will return the value of a field on the entry, including a boolean indicating if the field exists.
func (entry *Entry) Get(field FieldInterface) (interface{}, bool) {
	return field.Get(entry)
}

// Set will set the value of a field on the entry.
func (entry *Entry) Set(field FieldInterface, val interface{}) error {
	return field.Set(entry, val)
}

// Delete will delete a field from the entry.
func (entry *Entry) Delete(field FieldInterface) (interface{}, bool) {
	return field.Delete(entry)
}

// Read will read the value of a field into a designated interface.
func (entry *Entry) Read(field FieldInterface, dest interface{}) error {
	switch dest := dest.(type) {
	case *string:
		return entry.readToString(field, dest)
	case *map[string]interface{}:
		return entry.readToInterfaceMap(field, dest)
	case *map[string]string:
		return entry.readToStringMap(field, dest)
	case *interface{}:
		return entry.readToInterface(field, dest)
	default:
		return fmt.Errorf("can not read to unsupported type '%T'", dest)
	}
}

// readToInterface reads a field to a designated interface pointer.
func (entry *Entry) readToInterface(field FieldInterface, dest *interface{}) error {
	val, ok := entry.Get(field)
	if !ok {
		return fmt.Errorf("field '%s' is missing and can not be read as a interface{}", field)
	}

	*dest = val
	return nil
}

// readToString reads a field to a designated string pointer.
func (entry *Entry) readToString(field FieldInterface, dest *string) error {
	val, ok := entry.Get(field)
	if !ok {
		return fmt.Errorf("field '%s' is missing and can not be read as a string", field)
	}

	switch typed := val.(type) {
	case string:
		*dest = typed
	case []byte:
		*dest = string(typed)
	default:
		return fmt.Errorf("field '%s' of type '%T' can not be cast to a string", field, val)
	}

	return nil
}

// readToInterfaceMap reads a field to a designated map interface pointer.
func (entry *Entry) readToInterfaceMap(field FieldInterface, dest *map[string]interface{}) error {
	val, ok := entry.Get(field)
	if !ok {
		return fmt.Errorf("field '%s' is missing and can not be read as a map[string]interface{}", field)
	}

	if m, ok := val.(map[string]interface{}); ok {
		*dest = m
	} else {
		return fmt.Errorf("field '%s' of type '%T' can not be cast to a map[string]interface{}", field, val)
	}

	return nil
}

// readToStringMap reads a field to a designated map string pointer.
func (entry *Entry) readToStringMap(field FieldInterface, dest *map[string]string) error {
	val, ok := entry.Get(field)
	if !ok {
		return fmt.Errorf("field '%s' is missing and can not be read as a map[string]string{}", field)
	}

	switch m := val.(type) {
	case map[string]interface{}:
		newDest := make(map[string]string)
		for k, v := range m {
			if vStr, ok := v.(string); ok {
				newDest[k] = vStr
			} else {
				return fmt.Errorf("can not cast map members '%s' of type '%s' to string", k, v)
			}
		}
		*dest = newDest
	case map[interface{}]interface{}:
		newDest := make(map[string]string)
		for k, v := range m {
			keyStr, ok := k.(string)
			if !ok {
				return fmt.Errorf("can not cast map key of type '%T' to string", k)
			}
			vStr, ok := v.(string)
			if !ok {
				return fmt.Errorf("can not cast map value of type '%T' to string", v)
			}
			newDest[keyStr] = vStr
		}
		*dest = newDest
	}

	return nil
}

// Copy will return a deep copy of the entry.
func (entry *Entry) Copy() *Entry {
	return &Entry{
		Timestamp:    entry.Timestamp,
		Severity:     entry.Severity,
		SeverityText: entry.SeverityText,
		Labels:       copyStringMap(entry.Labels),
		Resource:     copyStringMap(entry.Resource),
		Record:       copyValue(entry.Record),
	}
}
