// Copyright  OpenTelemetry Authors
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

package encoding

import (
	"errors"

	"go.opentelemetry.io/collector/consumer/pdata"
)

var (
	// ErrUnsupportedEncodedType is used when the encoder type does not the type of encoding
	ErrUnsupportedEncodedType = errors.New("unsupported type to encode")
)

// Encoder allows for the internal types to be converted to an consumable
// exported type which is written to the kinesis stream
type Encoder interface {
	EncodeMetrics(md pdata.Metrics) error

	EncodeTraces(td pdata.Traces) error

	EncodeLogs(ld pdata.Logs) error
}
