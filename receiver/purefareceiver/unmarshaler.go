// Copyright 2022 The OpenTelemetry Authors
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

package purefareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// LogsUnmarshaler deserializes the message body.
type LogsUnmarshaler interface {
	// Unmarshal deserializes the message body into traces.
	Unmarshal([]byte) (plog.Logs, error)

	// Encoding of the serialized messages.
	Encoding() string
}

func defaultLogsUnmarshalers() map[string]LogsUnmarshaler {
	otlpPb := newPdataLogsUnmarshaler(&plog.ProtoUnmarshaler{}, defaultEncoding)
	raw := newRawLogsUnmarshaler()
	return map[string]LogsUnmarshaler{
		otlpPb.Encoding(): otlpPb,
		raw.Encoding():    raw,
	}
}
