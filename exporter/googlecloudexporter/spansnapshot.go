// Copyright 2020, OpenTelemetry Authors
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

package googlecloudexporter

import (
	"time"

	cloudtrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	apitrace "go.opentelemetry.io/otel/trace"
)

type spanSnapshot struct {
	spanContext, parent                                                   apitrace.SpanContext
	spanKind                                                              apitrace.SpanKind
	startTime, endTime                                                    time.Time
	name                                                                  string
	attributes                                                            []attribute.KeyValue
	links                                                                 []apitrace.Link
	events                                                                []sdktrace.Event
	droppedAttributes, droppedMessageEvents, droppedLinks, childSpanCount int
	resource                                                              *sdkresource.Resource
	instrumentationLibrary                                                instrumentation.Library
	status                                                                sdktrace.Status
}

var _ cloudtrace.ReadOnlySpan = spanSnapshot{}

func (s spanSnapshot) Name() string                     { return s.name }
func (s spanSnapshot) SpanContext() trace.SpanContext   { return s.spanContext }
func (s spanSnapshot) Parent() trace.SpanContext        { return s.parent }
func (s spanSnapshot) SpanKind() trace.SpanKind         { return s.spanKind }
func (s spanSnapshot) StartTime() time.Time             { return s.startTime }
func (s spanSnapshot) EndTime() time.Time               { return s.endTime }
func (s spanSnapshot) Attributes() []attribute.KeyValue { return s.attributes }
func (s spanSnapshot) Links() []trace.Link              { return s.links }
func (s spanSnapshot) Events() []sdktrace.Event         { return s.events }
func (s spanSnapshot) Status() sdktrace.Status          { return s.status }
func (s spanSnapshot) Resource() *sdkresource.Resource  { return s.resource }
func (s spanSnapshot) DroppedAttributes() int           { return s.droppedAttributes }
func (s spanSnapshot) DroppedLinks() int                { return s.droppedLinks }
func (s spanSnapshot) DroppedEvents() int               { return s.droppedLinks }
func (s spanSnapshot) ChildSpanCount() int              { return s.childSpanCount }
func (s spanSnapshot) InstrumentationLibrary() instrumentation.Library {
	return s.instrumentationLibrary
}
