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
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	datadogpb "github.com/DataDog/datadog-agent/pkg/trace/exportable/pb"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/obsreport"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/atomic"
)

type datadogReceiver struct {
	config       *Config
	params       receiver.CreateSettings
	nextConsumer consumer.Traces
	server       *http.Server
	shutdownWG   sync.WaitGroup
	tReceiver    *obsreport.Receiver

	startOnce    sync.Once
	stopOnce     sync.Once
	errorCounter atomic.Int64

	mb             *metadata.MetricsBuilder
	startTimestamp pcommon.Timestamp
}

func newDataDogReceiver(config *Config, nextConsumer consumer.Traces, params receiver.CreateSettings) (receiver.Traces, error) {
	if nextConsumer == nil {
		return nil, component.ErrNilNextConsumer
	}

	instance, err := obsreport.NewReceiver(obsreport.ReceiverSettings{LongLivedCtx: false, ReceiverID: params.ID, Transport: "http", ReceiverCreateSettings: params})
	if err != nil {
		return nil, err
	}
	return &datadogReceiver{
		params:       params,
		config:       config,
		nextConsumer: nextConsumer,
		mb:           metadata.NewMetricsBuilder(metadata.DefaultMetricsSettings(), params),
		server: &http.Server{
			ReadTimeout: config.ReadTimeout,
			Addr:        config.HTTPServerSettings.Endpoint,
		},
		tReceiver:      instance,
		startTimestamp: pcommon.NewTimestampFromTime(time.Now()),
	}, nil
}

func (ddr *datadogReceiver) Start(_ context.Context, host component.Host) error {
	ddr.startOnce.Do(func() {
		ddr.shutdownWG.Add(1)
		go func() {
			defer ddr.shutdownWG.Done()

			ddmux := http.NewServeMux()
			ddmux.HandleFunc("/v0.3/traces", ddr.handleTraces)
			ddmux.HandleFunc("/v0.4/traces", ddr.handleTraces)
			ddmux.HandleFunc("/v0.5/traces", ddr.handleTraces)
			ddr.server.Handler = ddmux
			if err := ddr.server.ListenAndServe(); err != http.ErrServerClosed {
				host.ReportFatalError(fmt.Errorf("error starting datadog receiver: %w", err))
			}
		}()
	})
	return nil
}

func (ddr *datadogReceiver) Shutdown(ctx context.Context) (err error) {
	ddr.stopOnce.Do(func() {
		err = ddr.server.Shutdown(ctx)
	})
	ddr.shutdownWG.Wait()
	return err
}

func (ddr *datadogReceiver) handleTraces(w http.ResponseWriter, req *http.Request) {
	now := pcommon.NewTimestampFromTime(time.Now())
	obsCtx := ddr.tReceiver.StartTracesOp(context.Background())
	var err error
	var spanCount int
	defer func(spanCount *int) {
		ddr.tReceiver.EndTracesOp(obsCtx, "datadog", *spanCount, err)
	}(&spanCount)
	var ddTraces datadogpb.Traces
	hostname, _ := os.Hostname()

	defer ddr.mb.Emit()
	err = decodeRequest(req, &ddTraces)
	if err != nil {
		defer ddr.mb.Emit()
		ddr.errorCounter.Inc()
		http.Error(w, "Unable to unmarshal reqs", http.StatusInternalServerError)
		hostname, _ := os.Hostname()
		ddr.mb.RecordOtelReceiverDatadogErrorsDataPoint(now, ddr.errorCounter.Load(), hostname)
		ddr.params.Logger.Error(fmt.Sprintf("Error %d: Unable to unmarshal request. %s", ddr.errorCounter.Load(), err))
	}

	otelTraces := toTraces(ddTraces, req)
	spanCount = otelTraces.SpanCount()

	err = ddr.nextConsumer.ConsumeTraces(obsCtx, otelTraces)
	if err != nil {
		ddr.errorCounter.Inc()
		http.Error(w, "Trace consumer errored out", http.StatusInternalServerError)
		ddr.mb.RecordOtelReceiverDatadogErrorsDataPoint(now, ddr.errorCounter.Load(), hostname)
		ddr.params.Logger.Error(fmt.Sprintf("Error %d: Trace consumer errored out. %s", ddr.errorCounter.Load(), err))
	} else {
		w.Write([]byte("OK"))
		ddr.mb.RecordOtelReceiverDatadogSpansProcessedDataPoint(now, int64(spanCount), hostname)
		ddr.mb.RecordOtelReceiverDatadogSpansRateDataPoint(now, float64(time.Since(now.AsTime()).Milliseconds())/float64(spanCount), hostname)
	}

}
