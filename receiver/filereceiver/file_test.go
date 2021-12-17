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

package filereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filereceiver"

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configtest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/model/otlp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/testdata"
)

func TestDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	require.NotNil(t, cfg, "failed to create default config")
	require.NoError(t, configtest.CheckConfigStruct(cfg))
}

func TestFileTracesReceiver(t *testing.T) {
	tempFolder, err := os.MkdirTemp("", "file")
	assert.NoError(t, err)
	factory := NewFactory()
	cfg := &cfg{Path: tempFolder}
	sink := new(consumertest.TracesSink)
	receiver, err := factory.CreateTracesReceiver(context.Background(), componenttest.NewNopReceiverCreateSettings(), cfg, sink)
	assert.NoError(t, err)
	err = receiver.Start(context.Background(), nil)
	assert.NoError(t, err)

	td := testdata.GenerateTracesTwoSpansSameResource()
	marshaler := otlp.NewJSONTracesMarshaler()
	b, err := marshaler.MarshalTraces(td)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempFolder, "traces.json"), b, 0600)
	assert.NoError(t, err)
	time.Sleep(1 * time.Second)

	assert.EqualValues(t, td, sink.AllTraces()[0])
	err = receiver.Shutdown(context.Background())
	assert.NoError(t, err)
}

func TestFileMetricsReceiver(t *testing.T) {
	tempFolder, err := os.MkdirTemp("", "file")
	assert.NoError(t, err)
	factory := NewFactory()
	cfg := &cfg{Path: tempFolder}
	sink := new(consumertest.MetricsSink)
	receiver, err := factory.CreateMetricsReceiver(context.Background(), componenttest.NewNopReceiverCreateSettings(), cfg, sink)
	assert.NoError(t, err)
	err = receiver.Start(context.Background(), nil)
	assert.NoError(t, err)

	md := testdata.GenerateMetricsManyMetricsSameResource(5)
	marshaler := otlp.NewJSONMetricsMarshaler()
	b, err := marshaler.MarshalMetrics(md)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempFolder, "metrics.json"), b, 0600)
	assert.NoError(t, err)
	time.Sleep(1 * time.Second)

	assert.EqualValues(t, md, sink.AllMetrics()[0])
	err = receiver.Shutdown(context.Background())
	assert.NoError(t, err)
}

func TestFileLogsReceiver(t *testing.T) {
	tempFolder, err := os.MkdirTemp("", "file")
	assert.NoError(t, err)
	factory := NewFactory()
	cfg := &cfg{Path: tempFolder}
	sink := new(consumertest.LogsSink)
	receiver, err := factory.CreateLogsReceiver(context.Background(), componenttest.NewNopReceiverCreateSettings(), cfg, sink)
	assert.NoError(t, err)
	err = receiver.Start(context.Background(), nil)
	assert.NoError(t, err)

	ld := testdata.GenerateLogsManyLogRecordsSameResource(5)
	marshaler := otlp.NewJSONLogsMarshaler()
	b, err := marshaler.MarshalLogs(ld)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempFolder, "logs.json"), b, 0600)
	assert.NoError(t, err)
	time.Sleep(1 * time.Second)

	assert.EqualValues(t, ld, sink.AllLogs()[0])
	err = receiver.Shutdown(context.Background())
	assert.NoError(t, err)
}
