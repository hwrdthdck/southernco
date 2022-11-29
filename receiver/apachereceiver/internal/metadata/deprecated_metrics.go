package metadata

import (
	"fmt"
	"strconv"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (m *metricApacheCPULoad) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheCPUTime) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, serverNameAttributeValue string, cpuLevelAttributeValue string, cpuModeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
	dp.Attributes().PutStr("level", cpuLevelAttributeValue)
	dp.Attributes().PutStr("mode", cpuModeAttributeValue)
}

func (m *metricApacheCurrentConnections) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheLoad1) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheLoad15) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheLoad5) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheRequestTime) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheRequests) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheScoreboard) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string, scoreboardStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
	dp.Attributes().PutStr("state", scoreboardStateAttributeValue)
}

func (m *metricApacheTraffic) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheUptime) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
}

func (m *metricApacheWorkers) recordDataPointWithServerName(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string, workersStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("server_name", serverNameAttributeValue)
	dp.Attributes().PutStr("state", workersStateAttributeValue)
}

// RecordApacheCPULoadDataPoint adds a data point to apache.cpu.load metric.
func (mb *MetricsBuilder) RecordApacheCPULoadDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheCPULoad, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCPULoad.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheCPUTimeDataPoint adds a data point to apache.cpu.time metric.
func (mb *MetricsBuilder) RecordApacheCPUTimeDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string, cpuLevelAttributeValue ApacheCPUTimeAttributeLevel, cpuModeAttributeValue ApacheCPUTimeAttributeMode) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheCPUTime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCPUTime.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue, cpuLevelAttributeValue.String(), cpuModeAttributeValue.String())
	return nil
}

// RecordApacheCurrentConnectionsDataPoint adds a data point to apache.current_connections metric.
func (mb *MetricsBuilder) RecordApacheCurrentConnectionsDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheCurrentConnections, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCurrentConnections.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheLoad1DataPoint adds a data point to apache.load.1 metric.
func (mb *MetricsBuilder) RecordApacheLoad1DataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheLoad1, value was %s: %w", inputVal, err)
	}
	mb.metricApacheLoad1.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheLoad15DataPoint adds a data point to apache.load.15 metric.
func (mb *MetricsBuilder) RecordApacheLoad15DataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheLoad15, value was %s: %w", inputVal, err)
	}
	mb.metricApacheLoad15.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheLoad5DataPoint adds a data point to apache.load.5 metric.
func (mb *MetricsBuilder) RecordApacheLoad5DataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheLoad5, value was %s: %w", inputVal, err)
	}
	mb.metricApacheLoad5.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheRequestTimeDataPoint adds a data point to apache.request.time metric.
func (mb *MetricsBuilder) RecordApacheRequestTimeDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheRequestTime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheRequestTime.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheRequestsDataPoint adds a data point to apache.requests metric.
func (mb *MetricsBuilder) RecordApacheRequestsDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheRequests, value was %s: %w", inputVal, err)
	}
	mb.metricApacheRequests.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheScoreboardDataPoint adds a data point to apache.scoreboard metric.
func (mb *MetricsBuilder) RecordApacheScoreboardDataPointWithServerName(ts pcommon.Timestamp, val int64, serverNameAttributeValue string, scoreboardStateAttributeValue ApacheScoreboardAttributeState) {
	mb.metricApacheScoreboard.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue, scoreboardStateAttributeValue.String())
}

// RecordApacheTrafficDataPoint adds a data point to apache.traffic metric.
func (mb *MetricsBuilder) RecordApacheTrafficDataPointWithServerName(ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	mb.metricApacheTraffic.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
}

// RecordApacheUptimeDataPoint adds a data point to apache.uptime metric.
func (mb *MetricsBuilder) RecordApacheUptimeDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheUptime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheUptime.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheWorkersDataPoint adds a data point to apache.workers metric.
func (mb *MetricsBuilder) RecordApacheWorkersDataPointWithServerName(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string, workersStateAttributeValue ApacheWorkersAttributeState) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheWorkers, value was %s: %w", inputVal, err)
	}
	mb.metricApacheWorkers.recordDataPointWithServerName(mb.startTime, ts, val, serverNameAttributeValue, workersStateAttributeValue.String())
	return nil
}
