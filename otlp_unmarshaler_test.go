// Copyright 2020 The OpenTelemetry Authors
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

package kafkareceiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/internal/testdata"
)

func TestUnmarshalOTLPTraces(t *testing.T) {
	td := pdata.NewTraces()
	td.ResourceSpans().AppendEmpty().Resource().Attributes().InsertString("foo", "bar")

	expected, err := td.ToOtlpProtoBytes()
	require.NoError(t, err)

	p := otlpTracesPbUnmarshaler{}
	got, err := p.Unmarshal(expected)
	require.NoError(t, err)

	assert.Equal(t, td, got)
	assert.Equal(t, "otlp_proto", p.Encoding())
}

func TestUnmarshalOTLPTraces_error(t *testing.T) {
	p := otlpTracesPbUnmarshaler{}
	_, err := p.Unmarshal([]byte("+$%"))
	assert.Error(t, err)
}

func TestUnmarshalOTLPLogs(t *testing.T) {
	ld := testdata.GenerateLogsOneLogRecord()

	expected, err := ld.ToOtlpProtoBytes()
	require.NoError(t, err)

	p := otlpLogsPbUnmarshaler{}
	got, err := p.Unmarshal(expected)
	require.NoError(t, err)

	assert.Equal(t, ld, got)
	assert.Equal(t, "otlp_proto", p.Encoding())
}

func TestUnmarshalOTLPLogs_error(t *testing.T) {
	p := otlpLogsPbUnmarshaler{}
	_, err := p.Unmarshal([]byte("+$%"))
	assert.Error(t, err)
}
