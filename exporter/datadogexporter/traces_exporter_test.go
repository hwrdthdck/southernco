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

// +build !windows

package datadogexporter

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DataDog/datadog-agent/pkg/trace/pb"
	"github.com/DataDog/datadog-agent/pkg/trace/stats"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/config"
)

func testTraceExporterHelper(td pdata.Traces, t *testing.T) []string {
	var got []string
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", req.Header.Get("DD-Api-Key"))

		contentType := req.Header.Get("Content-Type")

		data := []string{contentType}
		got = append(got, data...)

		if contentType == "application/x-protobuf" {
			testProtobufTracePayload(t, rw, req)
		} else if contentType == "application/json" {
			testJSONTraceStatsPayload(t, rw, req)
		}
		rw.WriteHeader(http.StatusAccepted)
	}))

	defer server.Close()
	cfg := config.Config{
		API: config.APIConfig{
			Key: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		TagsConfig: config.TagsConfig{
			Hostname: "test_host",
			Env:      "test_env",
			Tags:     []string{"key:val"},
		},
		Traces: config.TracesConfig{
			SampleRate: 1,
			TCPAddr: confignet.TCPAddr{
				Endpoint: server.URL,
			},
		},
	}

	params := component.ExporterCreateParams{Logger: zap.NewNop()}

	exporter, err := createTraceExporter(context.Background(), params, &cfg)

	assert.NoError(t, err)

	defer exporter.Shutdown(context.Background())

	ctx := context.Background()
	errConsume := exporter.ConsumeTraces(ctx, td)
	assert.NoError(t, errConsume)

	return got
}

func testProtobufTracePayload(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	var traceData pb.TracePayload
	b, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		assert.NoError(t, err, "http server received malformed trace payload")
		return
	}

	defer req.Body.Close()

	if marshallErr := proto.Unmarshal(b, &traceData); marshallErr != nil {
		http.Error(rw, marshallErr.Error(), http.StatusInternalServerError)
		assert.NoError(t, marshallErr, "http server received malformed trace payload")
		return
	}

	assert.NotNil(t, traceData.Env)
	assert.NotNil(t, traceData.HostName)
	assert.NotNil(t, traceData.Traces)
}

func testJSONTraceStatsPayload(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	var statsData stats.Payload

	gz, err := gzip.NewReader(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		require.NoError(t, err, "http server received malformed stats payload")
		return
	}

	defer req.Body.Close()
	defer gz.Close()

	statsBytes, err := ioutil.ReadAll(gz)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		require.NoError(t, err, "http server received malformed stats payload")
		return
	}

	if marshallErr := json.Unmarshal(statsBytes, &statsData); marshallErr != nil {
		http.Error(rw, marshallErr.Error(), http.StatusInternalServerError)
		require.NoError(t, marshallErr, "http server received malformed stats payload")
		return
	}

	assert.NotNil(t, statsData.Env)
	assert.NotNil(t, statsData.HostName)
	assert.NotNil(t, statsData.Stats)
}

func TestNewTraceExporter(t *testing.T) {
	cfg := &config.Config{}
	cfg.API.Key = "ddog_32_characters_long_api_key1"
	logger := zap.NewNop()

	// The client should have been created correctly
	exp, err := newTraceExporter(logger, cfg)
	assert.NoError(t, err)
	assert.NotNil(t, exp)
}

func TestPushTraceData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", req.Header.Get("DD-Api-Key"))
		rw.WriteHeader(http.StatusAccepted)
	}))

	defer server.Close()
	cfg := &config.Config{
		API: config.APIConfig{
			Key: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		TagsConfig: config.TagsConfig{
			Hostname: "test_host",
			Env:      "test_env",
			Tags:     []string{"key:val"},
		},
		Traces: config.TracesConfig{
			SampleRate: 1,
			TCPAddr: confignet.TCPAddr{
				Endpoint: server.URL,
			},
		},
	}
	logger := zap.NewNop()

	exp, err := newTraceExporter(logger, cfg)

	assert.NoError(t, err)

	tracesLength, err := exp.pushTraceData(context.Background(), func() pdata.Traces {
		traces := pdata.NewTraces()
		resourceSpans := traces.ResourceSpans()
		resourceSpans.Resize(1)
		resourceSpans.At(0).InitEmpty()
		resourceSpans.At(0).InstrumentationLibrarySpans().Resize(1)
		resourceSpans.At(0).InstrumentationLibrarySpans().At(0).Spans().Resize(1)
		return traces
	}())

	assert.NoError(t, err)
	assert.Equal(t, 1, tracesLength)

}

func TestTraceAndStatsExporter(t *testing.T) {
	// ensure that the protobuf serialized traces payload contains HostName Env and Traces
	// ensure that the json gzipped stats payload contains HostName Env and Stats
	got := testTraceExporterHelper(simpleTraces(), t)

	// ensure a protobuf and json payload are sent
	assert.Equal(t, 2, len(got))
	assert.Equal(t, "application/json", got[1])
	assert.Equal(t, "application/x-protobuf", got[0])
}

func simpleTraces() pdata.Traces {
	return simpleTracesWithID(pdata.NewTraceID([]byte{1, 2, 3, 4}))
}

func simpleTracesWithID(traceID pdata.TraceID) pdata.Traces {
	span := pdata.NewSpan()
	span.InitEmpty()
	span.SetTraceID(traceID)

	ils := pdata.NewInstrumentationLibrarySpans()
	ils.InitEmpty()
	ils.Spans().Append(span)

	rs := pdata.NewResourceSpans()
	rs.InitEmpty()
	rs.InstrumentationLibrarySpans().Append(ils)

	traces := pdata.NewTraces()
	traces.ResourceSpans().Append(rs)

	return traces
}
