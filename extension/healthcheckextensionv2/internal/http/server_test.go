// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap/confmaptest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/status"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/testhelpers"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/testutil"
)

type componentStatusExpectation struct {
	healthy      bool
	status       component.Status
	err          error
	nestedStatus map[string]*componentStatusExpectation
}

type teststep struct {
	step                    func()
	queryParams             string
	eventually              bool
	expectedStatusCode      int
	expectedComponentStatus *componentStatusExpectation
}

func TestStatus(t *testing.T) {
	var server *Server
	traces := testhelpers.NewPipelineMetadata("traces")
	metrics := testhelpers.NewPipelineMetadata("metrics")

	tests := []struct {
		name                    string
		settings                *Settings
		componentHealthSettings *common.ComponentHealthSettings
		pipelines               map[string]*testhelpers.PipelineMetadata
		teststeps               []teststep
	}{
		{
			name: "exclude recoverable and permanent errors",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
			},
		},
		{
			name: "exclude recoverable and permanent errors	- verbose",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStarting,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStarting,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStarting,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusRecoverableError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusRecoverableError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusRecoverableError,
								err:     assert.AnError,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusPermanentError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusPermanentError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusPermanentError,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
			},
		},
		{
			name: "include recoverable and exclude permanent errors",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			componentHealthSettings: &common.ComponentHealthSettings{
				IncludePermanent:   false,
				IncludeRecoverable: true,
				RecoveryDuration:   2 * time.Millisecond,
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					eventually:         true,
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
			},
		},
		{
			name: "include recoverable and exclude permanent errors	- verbose",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			componentHealthSettings: &common.ComponentHealthSettings{
				IncludePermanent:   false,
				IncludeRecoverable: true,
				RecoveryDuration:   2 * time.Millisecond,
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					eventually:         true,
					queryParams:        "verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: false,
								status:  component.StatusRecoverableError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: false,
										status:  component.StatusRecoverableError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: false,
								status:  component.StatusRecoverableError,
								err:     assert.AnError,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusPermanentError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusPermanentError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusPermanentError,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
			},
		},
		{
			name: "include permanent and exclude recoverable errors",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			componentHealthSettings: &common.ComponentHealthSettings{
				IncludePermanent: true,
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
			},
		},
		{
			name: "include permanent and exclude recoverable errors	- verbose",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			componentHealthSettings: &common.ComponentHealthSettings{
				IncludePermanent: true,
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					eventually:         true,
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusRecoverableError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusRecoverableError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusRecoverableError,
								err:     assert.AnError,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: false,
								status:  component.StatusPermanentError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: false,
										status:  component.StatusPermanentError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: false,
								status:  component.StatusPermanentError,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
			},
		},
		{
			name: "include permanent and recoverable errors",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			componentHealthSettings: &common.ComponentHealthSettings{
				IncludePermanent:   true,
				IncludeRecoverable: true,
				RecoveryDuration:   2 * time.Millisecond,
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					eventually:         true,
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=traces",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
				{
					queryParams:        "pipeline=metrics",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
					},
				},
			},
		},
		{
			name: "include permanent and recoverable errors	- verbose",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			componentHealthSettings: &common.ComponentHealthSettings{
				IncludePermanent:   true,
				IncludeRecoverable: true,
				RecoveryDuration:   2 * time.Millisecond,
			},
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(server.aggregator,
							traces.InstanceIDs(),
							component.StatusStarting,
						)
						testhelpers.SeedAggregator(server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStarting,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStarting,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStarting,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStarting,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStarting,
									},
								},
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusOK,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewRecoverableErrorEvent(assert.AnError),
						)
					},
					eventually:         true,
					queryParams:        "verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: false,
								status:  component.StatusRecoverableError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: false,
										status:  component.StatusRecoverableError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusRecoverableError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: false,
								status:  component.StatusRecoverableError,
								err:     assert.AnError,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewStatusEvent(component.StatusOK),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusOK,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusOK,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusOK,
							},
						},
					},
				},
				{
					step: func() {
						server.aggregator.RecordStatus(
							metrics.ExporterID,
							component.NewPermanentErrorEvent(assert.AnError),
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusOK,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusOK,
									},
								},
							},
							"pipeline:metrics": {
								healthy: false,
								status:  component.StatusPermanentError,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusOK,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusOK,
									},
									"exporter:metrics/out": {
										healthy: false,
										status:  component.StatusPermanentError,
										err:     assert.AnError,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusInternalServerError,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: false,
						status:  component.StatusPermanentError,
						err:     assert.AnError,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusOK,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusOK,
							},
							"exporter:metrics/out": {
								healthy: false,
								status:  component.StatusPermanentError,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopping,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopping,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopping,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopping,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopping,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopping,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopping,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopping,
							},
						},
					},
				},
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusStopped,
						)
						testhelpers.SeedAggregator(
							server.aggregator,
							metrics.InstanceIDs(),
							component.StatusStopped,
						)
					},
					queryParams:        "verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"pipeline:traces": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:traces/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:traces/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
							"pipeline:metrics": {
								healthy: true,
								status:  component.StatusStopped,
								nestedStatus: map[string]*componentStatusExpectation{
									"receiver:metrics/in": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"processor:batch": {
										healthy: true,
										status:  component.StatusStopped,
									},
									"exporter:metrics/out": {
										healthy: true,
										status:  component.StatusStopped,
									},
								},
							},
						},
					},
				},
				{
					queryParams:        "pipeline=traces&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:traces/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:traces/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
				{
					queryParams:        "pipeline=metrics&verbose",
					expectedStatusCode: http.StatusServiceUnavailable,
					expectedComponentStatus: &componentStatusExpectation{
						healthy: true,
						status:  component.StatusStopped,
						nestedStatus: map[string]*componentStatusExpectation{
							"receiver:metrics/in": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"processor:batch": {
								healthy: true,
								status:  component.StatusStopped,
							},
							"exporter:metrics/out": {
								healthy: true,
								status:  component.StatusStopped,
							},
						},
					},
				},
			},
		},
		{
			name: "pipeline non-existent",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: true,
					Path:    "/status",
				},
			},
			pipelines: testhelpers.NewPipelines("traces"),
			teststeps: []teststep{
				{
					step: func() {
						testhelpers.SeedAggregator(
							server.aggregator,
							traces.InstanceIDs(),
							component.StatusOK,
						)
					},
					queryParams:        "pipeline=nonexistent",
					expectedStatusCode: http.StatusNotFound,
				},
			},
		},
		{
			name: "status disabled",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{Enabled: false},
				Status: PathSettings{
					Enabled: false,
				},
			},
			teststeps: []teststep{
				{
					expectedStatusCode: http.StatusNotFound,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server = NewServer(
				tc.settings,
				tc.componentHealthSettings,
				componenttest.NewNopTelemetrySettings(),
				status.NewAggregator(testhelpers.ErrPriority(tc.componentHealthSettings)),
			)

			require.NoError(t, server.Start(context.Background(), componenttest.NewNopHost()))
			defer func() { require.NoError(t, server.Shutdown(context.Background())) }()

			client := &http.Client{}
			url := fmt.Sprintf("http://%s%s", tc.settings.Endpoint, tc.settings.Status.Path)

			for _, ts := range tc.teststeps {
				if ts.step != nil {
					ts.step()
				}

				stepURL := url
				if ts.queryParams != "" {
					stepURL = fmt.Sprintf("%s?%s", stepURL, ts.queryParams)
				}

				var err error
				var resp *http.Response

				if ts.eventually {
					assert.Eventually(t, func() bool {
						resp, err = client.Get(stepURL)
						require.NoError(t, err)
						return ts.expectedStatusCode == resp.StatusCode
					}, time.Second, 10*time.Millisecond)
				} else {
					resp, err = client.Get(stepURL)
					require.NoError(t, err)
					assert.Equal(t, ts.expectedStatusCode, resp.StatusCode)
				}

				if ts.expectedComponentStatus != nil {
					body, err := io.ReadAll(resp.Body)
					require.NoError(t, err)

					st := &serializableStatus{}
					require.NoError(t, json.Unmarshal(body, st))
					fmt.Println(string(body))
					if strings.Contains(ts.queryParams, "verbose") {
						assertStatusDetailed(t, ts.expectedComponentStatus, st)
						continue
					}

					assertStatusSimple(t, ts.expectedComponentStatus, st)
				}
			}
		})
	}
}

func assertStatusDetailed(
	t *testing.T,
	expected *componentStatusExpectation,
	actual *serializableStatus,
) {
	assert.Equal(t, expected.healthy, actual.Healthy)
	assert.Equal(t, expected.status, actual.Status(),
		"want: %s, got: %s", expected.status, actual.Status())
	if expected.err != nil {
		assert.Equal(t, expected.err.Error(), actual.Error)
	}
	assertNestedStatus(t, expected.nestedStatus, actual.ComponentStatuses)
}

func assertNestedStatus(
	t *testing.T,
	expected map[string]*componentStatusExpectation,
	actual map[string]*serializableStatus,
) {
	for k, expectation := range expected {
		st, ok := actual[k]
		require.True(t, ok, "status for key: %s not found", k)
		assert.Equal(t, expectation.healthy, st.Healthy)
		assert.Equal(t, expectation.status, st.Status(),
			"want: %s, got: %s", expectation.status, st.Status())
		if expectation.err != nil {
			assert.Equal(t, expectation.err.Error(), st.Error)
		}
		assertNestedStatus(t, expectation.nestedStatus, st.ComponentStatuses)
	}
}

func assertStatusSimple(
	t *testing.T,
	expected *componentStatusExpectation,
	actual *serializableStatus,
) {
	assert.Equal(t, expected.status, actual.Status())
	assert.Equal(t, expected.healthy, actual.Healthy)
	if expected.err != nil {
		assert.Equal(t, expected.err.Error(), actual.Error)
	}
	assert.Nil(t, actual.ComponentStatuses)
}

func TestConfig(t *testing.T) {
	var server *Server
	confMap, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	confJSON, err := os.ReadFile(filepath.Clean(filepath.Join("testdata", "config.json")))
	require.NoError(t, err)

	for _, tc := range []struct {
		name               string
		settings           *Settings
		setup              func()
		expectedStatusCode int
		expectedBody       []byte
	}{
		{
			name: "config not notified",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{
					Enabled: true,
					Path:    "/config",
				},
				Status: PathSettings{
					Enabled: false,
				},
			},
			expectedStatusCode: http.StatusServiceUnavailable,
			expectedBody:       []byte{},
		},
		{
			name: "config notified",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{
					Enabled: true,
					Path:    "/config",
				},
				Status: PathSettings{
					Enabled: false,
				},
			},
			setup: func() {
				require.NoError(t, server.NotifyConfig(context.Background(), confMap))
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       confJSON,
		},
		{
			name: "config disabled",
			settings: &Settings{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: testutil.GetAvailableLocalAddress(t),
				},
				Config: PathSettings{
					Enabled: false,
				},
				Status: PathSettings{
					Enabled: false,
				},
			},
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       []byte("404 page not found\n"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server = NewServer(
				tc.settings,
				&common.ComponentHealthSettings{},
				componenttest.NewNopTelemetrySettings(),
				status.NewAggregator(status.PriorityPermanent),
			)

			require.NoError(t, server.Start(context.Background(), componenttest.NewNopHost()))
			defer func() { require.NoError(t, server.Shutdown(context.Background())) }()

			client := &http.Client{}
			url := fmt.Sprintf("http://%s%s", tc.settings.Endpoint, tc.settings.Config.Path)

			if tc.setup != nil {
				tc.setup()
			}

			resp, err := client.Get(url)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedBody, body)
		})
	}

}
