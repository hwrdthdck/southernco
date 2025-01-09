// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwmetricstream // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.27.0"
	"go.uber.org/zap"

	expmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics"
)

const (
	TypeStr = "cwmetrics"

	attributeAWSCloudWatchMetricStreamName = "aws.cloudwatch.metric_stream_name"
	dimensionInstanceID                    = "InstanceId"
	namespaceDelimiter                     = "/"
)

var errInvalidRecords = errors.New("record format invalid")

// Unmarshaler for the CloudWatch Metric Stream JSON record format.
//
// More details can be found at:
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats-json.html
type Unmarshaler struct {
	logger *zap.Logger
}

var _ pmetric.Unmarshaler = (*Unmarshaler)(nil)

// NewUnmarshaler creates a new instance of the Unmarshaler.
func NewUnmarshaler(logger *zap.Logger) *Unmarshaler {
	return &Unmarshaler{logger}
}

// Unmarshal deserializes the records into cWMetrics and uses the
// resourceMetricsBuilder to group them into a single pmetric.Metrics.
// Skips invalid cWMetrics received in the record and
func (u Unmarshaler) UnmarshalMetrics(record []byte) (pmetric.Metrics, error) {
	metrics := pmetric.NewMetrics()

	// Multiple metrics in each record separated by newline character
	decoder := json.NewDecoder(bytes.NewReader(record))
	for datumIndex := 0; ; datumIndex++ {
		var cwMetric cWMetric
		if err := decoder.Decode(&cwMetric); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			u.logger.Error(
				"Unable to unmarshal input",
				zap.Error(err),
				zap.Int("datum_index", datumIndex),
			)
			continue
		}
		if !u.isValid(cwMetric) {
			u.logger.Error(
				"Invalid metric",
				zap.Int("datum_index", datumIndex),
			)
			continue
		}

		rm := metrics.ResourceMetrics().AppendEmpty()
		setResourceAttributes(cwMetric, rm.Resource())

		metric := rm.ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
		metric.SetName(cwMetric.MetricName)
		metric.SetUnit(cwMetric.Unit)

		dp := metric.SetEmptySummary().DataPoints().AppendEmpty()
		dp.SetTimestamp(pcommon.NewTimestampFromTime(time.UnixMilli(cwMetric.Timestamp)))
		setDataPointAttributes(cwMetric, dp)
		dp.SetCount(uint64(cwMetric.Value.Count))
		dp.SetSum(cwMetric.Value.Sum)
		minQ := dp.QuantileValues().AppendEmpty()
		minQ.SetQuantile(0)
		minQ.SetValue(cwMetric.Value.Min)
		maxQ := dp.QuantileValues().AppendEmpty()
		maxQ.SetQuantile(1)
		maxQ.SetValue(cwMetric.Value.Max)
	}

	if metrics.MetricCount() == 0 {
		return metrics, errInvalidRecords
	}

	metrics = expmetrics.Merge(pmetric.NewMetrics(), metrics)
	return metrics, nil
}

// isValid validates that the cWMetric has been unmarshalled correctly.
func (u Unmarshaler) isValid(metric cWMetric) bool {
	return metric.MetricName != "" && metric.Namespace != "" && metric.Unit != "" && metric.Value != nil
}

// Type of the serialized messages.
func (u Unmarshaler) Type() string {
	return TypeStr
}

// setResourceAttributes sets attributes on a pcommon.Resource from a cwMetric.
func setResourceAttributes(m cWMetric, resource pcommon.Resource) {
	attributes := resource.Attributes()
	attributes.PutStr(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderAWS)
	attributes.PutStr(conventions.AttributeCloudAccountID, m.AccountID)
	attributes.PutStr(conventions.AttributeCloudRegion, m.Region)
	serviceNamespace, serviceName := toServiceAttributes(m.Namespace)
	if serviceNamespace != "" {
		attributes.PutStr(conventions.AttributeServiceNamespace, serviceNamespace)
	}
	attributes.PutStr(conventions.AttributeServiceName, serviceName)
	attributes.PutStr(attributeAWSCloudWatchMetricStreamName, m.MetricStreamName)
}

// toServiceAttributes splits the CloudWatch namespace into service namespace/name
// if prepended by AWS/. Otherwise, it returns the CloudWatch namespace as the
// service name with an empty service namespace
func toServiceAttributes(namespace string) (serviceNamespace, serviceName string) {
	index := strings.Index(namespace, namespaceDelimiter)
	if index != -1 && strings.EqualFold(namespace[:index], conventions.AttributeCloudProviderAWS) {
		return namespace[:index], namespace[index+1:]
	}
	return "", namespace
}

// setResourceAttributes sets attributes on a metric data point from a cwMetric.
func setDataPointAttributes(m cWMetric, dp pmetric.SummaryDataPoint) {
	attrs := dp.Attributes()
	for k, v := range m.Dimensions {
		switch k {
		case dimensionInstanceID:
			attrs.PutStr(conventions.AttributeServiceInstanceID, v)
		default:
			attrs.PutStr(k, v)
		}
	}
}
