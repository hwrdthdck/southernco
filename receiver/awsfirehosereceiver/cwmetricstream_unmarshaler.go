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

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"go.opentelemetry.io/collector/model/pdata"
)

const (
	encodingCWMetricStream = "cwmetrics"
)

type cwMetricStreamUnmarshaler struct {
}

var _ MetricsUnmarshaler = (*cwMetricStreamUnmarshaler)(nil)

func (u cwMetricStreamUnmarshaler) Unmarshal(_ []byte) (pdata.Metrics, error) {
	return pdata.NewMetrics(), nil
}

func (u cwMetricStreamUnmarshaler) Encoding() string {
	return encodingCWMetricStream
}
