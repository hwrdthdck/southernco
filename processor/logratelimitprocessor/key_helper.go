// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logratelimitprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logratelimitprocessor"

import (
	"fmt"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	// fieldDelimiter is the delimiter used to split a field key into its parts.
	fieldDelimiter = "."

	// fieldEscapeKeyReplacement is the string used to temporarily replace escaped delimters while splitting a field key.
	fieldEscapeKeyReplacement = "{TEMP_REPLACE}"
)

// keyHelper helps in generating a key from the values of the rate_limit_fields
type keyHelper struct {
	fields []*field
}

// field represents a field and it's compound key to match on
type field struct {
	keyParts []string
}

// newKeyHelper creates a new keyHelper based on the keys passed in rateLimiterFields
func newKeyHelper(rateLimiterFields []string) *keyHelper {
	fe := &keyHelper{
		fields: make([]*field, 0, len(rateLimiterFields)),
	}

	for _, f := range rateLimiterFields {
		fe.fields = append(fe.fields, &field{
			keyParts: splitField(f),
		})
	}

	return fe
}

// GenerateKey removes any body or attribute fields that match in the log record
func (fe *keyHelper) GenerateKey(logRecord plog.LogRecord, resource pcommon.Resource) (uint64, string) {
	fieldValueArray := make([]pdatautil.HashOption, 0, len(fe.fields))
	stringBuilder := strings.Builder{}
	for i, field := range fe.fields {
		value := field.getFieldValue(logRecord, resource)
		stringBuilder.WriteString(value.AsString())
		fieldValueArray[i] = pdatautil.WithValue(value)
	}
	return pdatautil.Hash64(fieldValueArray...), stringBuilder.String()
}

// getFieldValue gets the field value for the given log
func (f *field) getFieldValue(logRecord plog.LogRecord, resource pcommon.Resource) pcommon.Value {
	firstPart, remainingParts := f.keyParts[0], f.keyParts[1:]

	res := pcommon.NewValueEmpty()

	switch firstPart {
	case bodyField:
		if len(remainingParts) == 0 || logRecord.Body().Type() != pcommon.ValueTypeMap {
			err := res.FromRaw(logRecord.Body().AsRaw())
			if err != nil {
				return pcommon.NewValueEmpty()
			}
			return res
		}

		// If body is a map then recurse through to remove the field
		return getFieldValueFromMap(logRecord.Body().Map(), remainingParts)

	case attributeField:
		if len(remainingParts) == 0 {
			err := res.FromRaw(logRecord.Attributes().AsRaw())
			if err != nil {
				return pcommon.NewValueEmpty()
			}
			return res
		}

		// Recurse through map and remove fields
		return getFieldValueFromMap(logRecord.Attributes(), remainingParts)

	case resourceField:
		if len(remainingParts) == 0 {
			err := res.FromRaw(resource.Attributes().AsRaw())
			if err != nil {
				return pcommon.NewValueEmpty()
			}
			return res
		}

		// Recurse through the map and remove fields
		return getFieldValueFromMap(resource.Attributes(), remainingParts)
	}

	return res
}

// getFieldValueFromMap gets the field value from map by recursion
func getFieldValueFromMap(valueMap pcommon.Map, keyParts []string) pcommon.Value {
	nextKeyPart, remainingParts := keyParts[0], keyParts[1:]

	// Look for the value associated with the next key part.
	// If we don't find it then return
	value, ok := valueMap.Get(nextKeyPart)
	if !ok {
		return pcommon.NewValueEmpty()
	}

	// No more key parts that means we have found the value and remove it
	if len(remainingParts) == 0 {
		return value
	}

	// If the value is a map then recurse through with the remaining parts
	if value.Type() == pcommon.ValueTypeMap {
		return getFieldValueFromMap(value.Map(), remainingParts)
	}

	return pcommon.NewValueEmpty()
}

// splitField splits a field key into its parts.
// It replaces escaped delimiters with the full delimiter after splitting.
func splitField(fieldKey string) []string {
	escapedKey := strings.ReplaceAll(fieldKey, fmt.Sprintf("\\%s", fieldDelimiter), fieldEscapeKeyReplacement)
	keyParts := strings.Split(escapedKey, fieldDelimiter)

	// Replace the temporarily escaped delimiters with the actual delimiter.
	for i := range keyParts {
		keyParts[i] = strings.ReplaceAll(keyParts[i], fieldEscapeKeyReplacement, fieldDelimiter)
	}

	return keyParts
}
