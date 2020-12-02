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

package attributes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/translator/conventions"
)

func TestTagsFromAttributes(t *testing.T) {
	attributeMap := map[string]pdata.AttributeValue{
		conventions.AttributeProcessExecutableName: pdata.NewAttributeValueString("otelcol"),
		conventions.AttributeProcessExecutablePath: pdata.NewAttributeValueString("/usr/bin/cmd/otelcol"),
		conventions.AttributeProcessCommand:        pdata.NewAttributeValueString("cmd/otelcol"),
		conventions.AttributeProcessCommandLine:    pdata.NewAttributeValueString("cmd/otelcol --config=\"/path/to/config.yaml\""),
		conventions.AttributeProcessID:             pdata.NewAttributeValueInt(1),
		conventions.AttributeProcessOwner:          pdata.NewAttributeValueString("root"),
	}
	attrs := pdata.NewAttributeMap().InitFromMap(attributeMap)

	assert.Equal(t, []string{
		fmt.Sprintf("%s:%s", conventions.AttributeProcessExecutableName, "otelcol"),
	}, TagsFromAttributes(attrs))
}

func TestTagsFromAttributesEmpty(t *testing.T) {
	attrs := pdata.NewAttributeMap()

	assert.Equal(t, []string{}, TagsFromAttributes(attrs))
}
