package auto

import (
	"bytes"
	"encoding/json"
	"errors"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog/compression"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
)

const (
	TypeStr         = "auto"
	recordDelimiter = "\n"
)

var (
	errInvalidRecords         = errors.New("record format invalid")
	errUnsupportedContentType = errors.New("content type not supported")
)

// Unmarshaler for the CloudWatch Log JSON record format.
type Unmarshaler struct {
	logger *zap.Logger
}

var _ unmarshaler.Unmarshaler = (*Unmarshaler)(nil)

// NewUnmarshaler creates a new instance of the Unmarshaler.
func NewUnmarshaler(logger *zap.Logger) *Unmarshaler {
	return &Unmarshaler{logger}
}

// isCloudWatchLog checks if the data has the entries needed to be considered a cloudwatch log
func isCloudWatchLog(data []byte) bool {
	if !bytes.Contains(data, []byte(`"owner":`)) {
		return false
	}
	if !bytes.Contains(data, []byte(`"logGroup":`)) {
		return false
	}
	if !bytes.Contains(data, []byte(`"logStream":`)) {
		return false
	}
	return true
}

// isCloudwatchMetrics checks if the data has the entries needed to be considered a cloudwatch metric
func isCloudwatchMetrics(data []byte) bool {
	if !bytes.Contains(data, []byte(`"metric_name":`)) {
		return false
	}
	if !bytes.Contains(data, []byte(`"namespace":`)) {
		return false
	}
	if !bytes.Contains(data, []byte(`"unit":`)) {
		return false
	}
	if !bytes.Contains(data, []byte(`"value":`)) {
		return false
	}
	return true
}

func (u *Unmarshaler) addCloudwatchLog(
	record []byte,
	resourceLogs map[cwlog.ResourceAttributes]*cwlog.ResourceLogsBuilder,
	ld plog.Logs,
) error {
	var log cwlog.CWLog
	if err := json.Unmarshal(record, &log); err != nil {
		return err
	}
	attrs := cwlog.ResourceAttributes{
		Owner:     log.Owner,
		LogGroup:  log.LogGroup,
		LogStream: log.LogStream,
	}
	lb, exists := resourceLogs[attrs]
	if !exists {
		lb = cwlog.NewResourceLogsBuilder(ld, attrs)
		resourceLogs[attrs] = lb
	}
	lb.AddLog(log)
	return nil
}

func (u *Unmarshaler) addCloudwatchMetric(
	record []byte,
	resourceMetrics map[cwmetricstream.ResourceAttributes]*cwmetricstream.ResourceMetricsBuilder,
	md pmetric.Metrics,
) error {
	var metric cwmetricstream.CWMetric
	if err := json.Unmarshal(record, &metric); err != nil {
		return err
	}
	attrs := cwmetricstream.ResourceAttributes{
		MetricStreamName: metric.MetricStreamName,
		Namespace:        metric.Namespace,
		AccountID:        metric.AccountID,
		Region:           metric.Region,
	}
	mb, exists := resourceMetrics[attrs]
	if !exists {
		mb = cwmetricstream.NewResourceMetricsBuilder(md, attrs)
		resourceMetrics[attrs] = mb
	}
	mb.AddMetric(metric)
	return nil
}

func (u *Unmarshaler) unmarshalJSON(md pmetric.Metrics, ld plog.Logs, records [][]byte) error {
	resourceLogs := make(map[cwlog.ResourceAttributes]*cwlog.ResourceLogsBuilder)
	resourceMetrics := make(map[cwmetricstream.ResourceAttributes]*cwmetricstream.ResourceMetricsBuilder)
	for i, compressed := range records {
		record, err := compression.Unzip(compressed)
		if err != nil {
			u.logger.Error("Failed to unzip record",
				zap.Error(err),
				zap.Int("record_index", i),
			)
			continue
		}
		// Multiple metrics/logs in each record separated by newline character
		for j, single := range bytes.Split(record, []byte(recordDelimiter)) {
			if len(single) == 0 {
				continue
			}

			if isCloudWatchLog(single) {
				if err = u.addCloudwatchLog(single, resourceLogs, ld); err != nil {
					u.logger.Error(
						"Unable to unmarshal record to cloudwatch log",
						zap.Error(err),
						zap.Int("record_index", i),
						zap.Int("single_index", j),
					)
					continue
				}
			}

			if isCloudwatchMetrics(single) {
				if err = u.addCloudwatchMetric(single, resourceMetrics, md); err != nil {
					u.logger.Error(
						"Unable to unmarshal input",
						zap.Error(err),
						zap.Int("single_index", j),
						zap.Int("record_index", i),
					)
					continue
				}
			}
		}
	}

	if len(resourceLogs) == 0 && len(resourceMetrics) == 0 {
		return errInvalidRecords
	}
	return nil
}

func (u *Unmarshaler) Unmarshal(contentType string, records [][]byte) (pmetric.Metrics, plog.Logs, error) {
	ld := plog.NewLogs()
	md := pmetric.NewMetrics()
	switch contentType {
	case "application/json":
		return md, ld, u.unmarshalJSON(md, ld, records)
	case "application/x-protobuf":
		// TODO implement this
		return md, ld, nil
	default:
		return md, ld, errUnsupportedContentType
	}

}

func (u *Unmarshaler) Type() string {
	return TypeStr
}
