// Copyright The OpenTelemetry Authors
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

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"bytes"
	"fmt"
	"time"
	"strings"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
	jsoniter "github.com/json-iterator/go"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.13.0"

	"go.uber.org/zap"
)

type azureResourceMetricsUnmarshaler struct {
	buildInfo component.BuildInfo
	logger    *zap.Logger
}

// azureMetricRecords represents an array of Azure metric records
// as exported via an Azure Event Hub
type azureMetricRecords struct {
	Records []azureMetricRecord `json:"records"`
}

// azureMetricRecord represents a single Azure Metric following
// the common schema does not exist (yet):
type azureMetricRecord struct {
	Time       string  `json:"time"`
	ResourceID string  `json:"resourceId"`
	MetricName string  `json:"metricName"`
	TimeGrain  string  `json:"timeGrain"`
	Total      float64 `json:"total"`
	Count      float64 `json:"count"`
	Minimum    float64 `json:"minimum"`
	Maximum    float64 `json:"maximum"`
	Average    float64 `json:"average"`
}

func newAzureResourceMetricsUnmarshaler(buildInfo component.BuildInfo, logger *zap.Logger) eventMetricsUnmarshaler {

	return azureResourceMetricsUnmarshaler{
		buildInfo: buildInfo,
		logger:    logger,
	}
}

// UnmarshalMetrics takes a byte array containing a JSON-encoded
// payload with Azure metric records and transforms it into
// an OpenTelemetry pmetric.Metrics object. The data in the Azure
// metric record appears as fields and attributes in the
// OpenTelemetry representation;
func (r azureResourceMetricsUnmarshaler) UnmarshalMetrics(event *eventhub.Event) (pmetric.Metrics, error) {

	md := pmetric.NewMetrics()

	var azureMetrics azureMetricRecords
	decoder := jsoniter.NewDecoder(bytes.NewReader(event.Data))
	err := decoder.Decode(&azureMetrics)
	if err != nil {
		return md, err
	}

	resourceMetrics := md.ResourceMetrics().AppendEmpty()
	scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
	scopeMetrics.Scope().SetName(receiverScopeName)
	scopeMetrics.Scope().SetVersion(r.buildInfo.Version)
	scopeMetrics.Scope().Attributes().PutStr(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderAzure)

	metrics := scopeMetrics.Metrics()
	metrics.EnsureCapacity(len(azureMetrics.Records))

	for _, azureMetric := range azureMetrics.Records {
		resourceID := azureMetric.ResourceID
		if resourceID == "" {
			r.logger.Warn("No ResourceID Set on Metric")
			continue
		}

		nanos, err := asTimestamp(azureMetric.Time)
		if err != nil {
			r.logger.Warn("Invalid Timestamp", zap.String("time", azureMetric.Time))
			continue
		}

		var startTimestamp pcommon.Timestamp
		if azureMetric.TimeGrain == "PT1M" {
			startTimestamp = pcommon.NewTimestampFromTime(nanos.AsTime().Add(-time.Minute))
		} else {
			r.logger.Warn("Unhandled Time Grain", zap.String("timegrain", azureMetric.TimeGrain))
			continue
		}

		metricTotal := metrics.AppendEmpty()
		metricTotal.SetName(strings.ToLower(fmt.Sprintf("%s_%s", strings.ReplaceAll(azureMetric.MetricName, " ", "_"), "Total")))
		dpTotal := metricTotal.Gauge().DataPoints().AppendEmpty()
		dpTotal.SetStartTimestamp(startTimestamp)
		dpTotal.SetTimestamp(nanos)
		dpTotal.SetDoubleValue(azureMetric.Total)
		dpTotal.Attributes().PutStr(azureResourceID, resourceID)

		metricCount := metrics.AppendEmpty()
		metricCount.SetName(strings.ToLower(fmt.Sprintf("%s_%s", strings.ReplaceAll(azureMetric.MetricName, " ", "_"), "Count")))
		dpCount := metricCount.Gauge().DataPoints().AppendEmpty()
		dpCount.SetStartTimestamp(startTimestamp)
		dpCount.SetTimestamp(nanos)
		dpCount.SetDoubleValue(azureMetric.Count)
		dpCount.Attributes().PutStr(azureResourceID, resourceID)
		
		metricMin := metrics.AppendEmpty()
		metricMin.SetName(strings.ToLower(fmt.Sprintf("%s_%s", strings.ReplaceAll(azureMetric.MetricName, " ", "_"), "Minimum")))
		dpMin := metricMin.Gauge().DataPoints().AppendEmpty()
		dpMin.SetStartTimestamp(startTimestamp)
		dpMin.SetTimestamp(nanos)
		dpMin.SetDoubleValue(azureMetric.Minimum)
		dpMin.Attributes().PutStr(azureResourceID, resourceID)

		metricMax := metrics.AppendEmpty()
		metricMax.SetName(strings.ToLower(fmt.Sprintf("%s_%s", strings.ReplaceAll(azureMetric.MetricName, " ", "_"), "Maximum")))
		dpMax := metricMax.Gauge().DataPoints().AppendEmpty()
		dpMax.SetStartTimestamp(startTimestamp)
		dpMax.SetTimestamp(nanos)
		dpMax.SetDoubleValue(azureMetric.Maximum)
		dpMax.Attributes().PutStr(azureResourceID, resourceID)

		metricAverage := metrics.AppendEmpty()
		metricAverage.SetName(strings.ToLower(fmt.Sprintf("%s_%s", strings.ReplaceAll(azureMetric.MetricName, " ", "_"), "Average")))
		dpAverage := metricAverage.Gauge().DataPoints().AppendEmpty()
		dpAverage.SetStartTimestamp(startTimestamp)
		dpAverage.SetTimestamp(nanos)
		dpAverage.SetDoubleValue(azureMetric.Average)
		dpAverage.Attributes().PutStr(azureResourceID, resourceID)
	}

	return md, nil
}
