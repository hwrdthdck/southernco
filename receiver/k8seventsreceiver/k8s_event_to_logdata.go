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

package k8seventsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	semconv "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
)

const (
	// Number of log attributes to add to the plog.LogRecordSlice.
	totalLogAttributes = 7

	// Number of resource attributes to add to the plog.ResourceLogs.
	totalResourceAttributes = 6
)

// Only two types of events are created as of now.
// For more info: https://docs.openshift.com/container-platform/4.9/rest_api/metadata_apis/event-core-v1.html
var severityMap = map[string]plog.SeverityNumber{
	"normal":  plog.SeverityNumberInfo,
	"warning": plog.SeverityNumberWarn,
}

// k8sEventToLogRecord converts Kubernetes event to plog.LogRecordSlice and adds the resource attributes.
func k8sEventToLogData(logger *zap.Logger, ev *corev1.Event) plog.Logs {
	ld := plog.NewLogs()
	rl := ld.ResourceLogs().AppendEmpty()
	sl := rl.ScopeLogs().AppendEmpty()
	lr := sl.LogRecords().AppendEmpty()

	resourceAttrs := rl.Resource().Attributes()
	resourceAttrs.EnsureCapacity(totalResourceAttributes)

	resourceAttrs.PutString(semconv.AttributeK8SNodeName, ev.Source.Host)

	// Attributes related to the object causing the event.
	resourceAttrs.PutString("k8s.object.kind", ev.InvolvedObject.Kind)
	resourceAttrs.PutString("k8s.object.name", ev.InvolvedObject.Name)
	resourceAttrs.PutString("k8s.object.uid", string(ev.InvolvedObject.UID))
	resourceAttrs.PutString("k8s.object.fieldpath", ev.InvolvedObject.FieldPath)
	resourceAttrs.PutString("k8s.object.api_version", ev.InvolvedObject.APIVersion)
	resourceAttrs.PutString("k8s.object.resource_version", ev.InvolvedObject.ResourceVersion)

	lr.SetTimestamp(pcommon.NewTimestampFromTime(getEventTimestamp(ev)))

	// The Message field contains description about the event,
	// which is best suited for the "Body" of the LogRecordSlice.
	lr.Body().SetStringVal(ev.Message)

	// Set the "SeverityNumber" and "SeverityText" if a known type of
	// severity is found.
	if severityNumber, ok := severityMap[strings.ToLower(ev.Type)]; ok {
		lr.SetSeverityNumber(severityNumber)
		lr.SetSeverityText(ev.Type)
	} else {
		logger.Debug("unknown severity type", zap.String("type", ev.Type))
	}

	attrs := lr.Attributes()
	attrs.EnsureCapacity(totalLogAttributes)

	attrs.PutString("k8s.event.reason", ev.Reason)
	attrs.PutString("k8s.event.action", ev.Action)
	attrs.PutString("k8s.event.start_time", ev.ObjectMeta.CreationTimestamp.String())
	attrs.PutString("k8s.event.name", ev.ObjectMeta.Name)
	attrs.PutString("k8s.event.uid", string(ev.ObjectMeta.UID))
	attrs.PutString(semconv.AttributeK8SNamespaceName, ev.InvolvedObject.Namespace)

	// "Count" field of k8s event will be '0' in case it is
	// not present in the collected event from k8s.
	if ev.Count != 0 {
		attrs.PutInt("k8s.event.count", int64(ev.Count))
	}

	return ld
}
