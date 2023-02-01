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

package datadogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver"

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/DataDog/datadog-agent/pkg/trace/pb"
	"github.com/stretchr/testify/assert"
	vmsgp "github.com/vmihailenco/msgpack/v4"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var data = [2]interface{}{
	0: []string{
		0:  "baggage",
		1:  "item",
		2:  "elasticsearch.version",
		3:  "7.0",
		4:  "my-name",
		5:  "X",
		6:  "my-service",
		7:  "my-resource",
		8:  "_dd.sampling_rate_whatever",
		9:  "value whatever",
		10: "sql",
	},
	1: [][][12]interface{}{
		{
			{
				6,
				4,
				7,
				uint64(12345678901234561234),
				uint64(2),
				uint64(3),
				int64(123),
				int64(456),
				1,
				map[interface{}]interface{}{
					8: 9,
					0: 1,
					2: 3,
				},
				map[interface{}]float64{
					5: 1.2,
				},
				10,
			},
		},
	},
}
var output ptrace.Traces

func TestTracePayloadV05Unmarshalling(t *testing.T) {
	var traces pb.Traces
	payload, err := vmsgp.Marshal(&data)
	assert.NoError(t, err)
	if err := traces.UnmarshalMsgDictionary(payload); err != nil {
		t.Fatal(err)
	}
	req := &http.Request{RequestURI: "/v0.5/traces", Body: io.NopCloser(bytes.NewReader(payload))}

	translated := toTraces(&pb.TracerPayload{
		LanguageName:    req.Header.Get("Datadog-Meta-Lang"),
		LanguageVersion: req.Header.Get("Datadog-Meta-Lang-Version"),
		Chunks:          traceChunksFromTraces(traces),
		TracerVersion:   req.Header.Get("Datadog-Meta-Tracer-Version"),
	}, req)
	assert.Equal(t, 1, translated.SpanCount(), "Span Count wrong")
	span := translated.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
	assert.NotNil(t, span)
	assert.Equal(t, 4, span.Attributes().Len(), "missing tags")
	value, exists := span.Attributes().Get("service.name")
	assert.True(t, exists, "service.name missing")
	assert.Equal(t, "my-service", value.AsString(), "service.name tag value incorrect")
	assert.Equal(t, span.Name(), "my-resource")
}

func BenchmarkTranslator(b *testing.B) {

	payload, err := vmsgp.Marshal(&data)
	assert.NoError(b, err)

	var traces pb.Traces
	if err := traces.UnmarshalMsgDictionary(payload); err != nil {
		b.Fatal(err)
	}

	req := &http.Request{RequestURI: "/v0.5/traces", Body: io.NopCloser(bytes.NewReader(payload))}
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		output = toTraces(&pb.TracerPayload{
			LanguageName:    req.Header.Get("Datadog-Meta-Lang"),
			LanguageVersion: req.Header.Get("Datadog-Meta-Lang-Version"),
			Chunks:          traceChunksFromTraces(traces),
			TracerVersion:   req.Header.Get("Datadog-Meta-Tracer-Version"),
		}, req)
	}
	b.StopTimer()
}
