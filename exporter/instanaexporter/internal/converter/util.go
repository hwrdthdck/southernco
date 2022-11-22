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

package converter

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func containsMetricWithPrefix(metricSlice pmetric.MetricSlice, prefix string) bool {
	for i := 0; i < metricSlice.Len(); i++ {
		metric := metricSlice.At(i)

		if strings.HasPrefix(metric.Name(), prefix) {
			return true
		}
	}

	return false
}

func containsAttributes(attributeMap pcommon.Map, attributes ...string) bool {
	for i := 0; i < len(attributes); i++ {
		_, ex := attributeMap.Get(attributes[i])

		if !ex {
			return false
		}
	}

	return true
}
