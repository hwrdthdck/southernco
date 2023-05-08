// Copyright The OpenTelemetry Authors
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

package translation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/schema/v1.0/ast"
	"go.opentelemetry.io/otel/schema/v1.0/types"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"
)

func TestNewRevision(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name         string
		inVersion    *Version
		inDefinition ast.VersionDef
		expect       *RevisionV1
	}{
		{
			name:         "no definition defined",
			inVersion:    &Version{1, 1, 1},
			inDefinition: ast.VersionDef{},
			expect: &RevisionV1{
				ver:              &Version{1, 1, 1},
				all:              migrate.NewAttributeChangeSetSlice(),
				resource:         migrate.NewAttributeChangeSetSlice(),
				spans:            migrate.NewConditionalAttributeSetSlice(),
				eventNames:       migrate.NewSignalSlice(),
				eventAttrsOnSpan: migrate.NewConditionalAttributeSetSlice(),
				eventAttrsOnName: migrate.NewConditionalAttributeSetSlice(),
				metricsAttrs:     migrate.NewConditionalAttributeSetSlice(),
				metricNames:      migrate.NewSignalSlice(),
			},
		},
		{
			name:      "complete version definition used",
			inVersion: &Version{1, 0, 0},
			inDefinition: ast.VersionDef{
				All: ast.Attributes{
					Changes: []ast.AttributeChange{
						{
							RenameAttributes: &ast.RenameAttributes{
								AttributeMap: ast.AttributeMap{
									"state": "status",
								},
							},
						},
						{
							RenameAttributes: &ast.RenameAttributes{
								AttributeMap: ast.AttributeMap{
									"status": "state",
								},
							},
						},
					},
				},
				Resources: ast.Attributes{
					Changes: []ast.AttributeChange{
						{
							RenameAttributes: &ast.RenameAttributes{
								AttributeMap: ast.AttributeMap{
									"service_name": "service.name",
								},
							},
						},
					},
				},
				Spans: ast.Spans{
					Changes: []ast.SpansChange{
						{
							RenameAttributes: &ast.AttributeMapForSpans{
								ApplyToSpans: []types.SpanName{
									"application start",
								},
								AttributeMap: ast.AttributeMap{
									"service_version": "service.version",
								},
							},
						},
						{
							RenameAttributes: &ast.AttributeMapForSpans{
								AttributeMap: ast.AttributeMap{
									"deployment.environment": "service.deployment.environment",
								},
							},
						},
					},
				},
				SpanEvents: ast.SpanEvents{
					Changes: []ast.SpanEventsChange{
						{
							RenameEvents: &ast.RenameSpanEvents{
								EventNameMap: map[string]string{
									"started": "application started",
								},
							},
							RenameAttributes: &ast.RenameSpanEventAttributes{
								ApplyToSpans: []types.SpanName{
									"service running",
								},
								ApplyToEvents: []types.EventName{
									"service errored",
								},
								AttributeMap: ast.AttributeMap{
									"service.app.name": "service.name",
								},
							},
						},
					},
				},
				Logs: ast.Logs{
					Changes: []ast.LogsChange{
						{
							RenameAttributes: &ast.RenameAttributes{
								AttributeMap: ast.AttributeMap{
									"ERROR": "error",
								},
							},
						},
					},
				},
				Metrics: ast.Metrics{
					Changes: []ast.MetricsChange{
						{
							RenameMetrics: map[types.MetricName]types.MetricName{
								"service.computed.uptime": "service.uptime",
							},
							RenameAttributes: &ast.AttributeMapForMetrics{
								ApplyToMetrics: []types.MetricName{
									"service.runtime",
								},
								AttributeMap: ast.AttributeMap{
									"runtime": "service.language",
								},
							},
						},
					},
				},
			},
			expect: &RevisionV1{
				ver: &Version{1, 0, 0},
				all: migrate.NewAttributeChangeSetSlice(
					migrate.NewAttributes(map[string]string{
						"state": "status",
					}),
					migrate.NewAttributes(map[string]string{
						"status": "state",
					}),
				),
				resource: migrate.NewAttributeChangeSetSlice(
					migrate.NewAttributes(map[string]string{
						"service_name": "service.name",
					}),
				),
				spans: migrate.NewConditionalAttributeSetSlice(
					migrate.NewConditionalAttributeSet(
						map[string]string{"service_version": "service.version"},
						"application start",
					),
					migrate.NewConditionalAttributeSet[string, string, string](
						map[string]string{"deployment.environment": "service.deployment.environment"},
					),
				),
				eventNames: migrate.NewSignalSlice(
					migrate.NewSignal(map[string]string{
						"started": "application started",
					}),
				),
				eventAttrsOnSpan: migrate.NewConditionalAttributeSetSlice(
					migrate.NewConditionalAttributeSet(
						map[string]string{
							"service.app.name": "service.name",
						},
						"service running",
					),
				),
				eventAttrsOnName: migrate.NewConditionalAttributeSetSlice(
					migrate.NewConditionalAttributeSet(
						map[string]string{
							"service.app.name": "service.name",
						},
						"service errored",
					),
				),
				metricsAttrs: migrate.NewConditionalAttributeSetSlice(
					migrate.NewConditionalAttributeSet(
						map[string]string{
							"runtime": "service.language",
						},
						"service.runtime",
					),
				),
				metricNames: migrate.NewSignalSlice(
					migrate.NewSignal(map[string]string{
						"service.computed.uptime": "service.uptime",
					}),
				),
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rev := NewRevision(tc.inVersion, tc.inDefinition)
			assert.EqualValues(t, tc.expect, rev, "Must match the expected values")
		})
	}
}
