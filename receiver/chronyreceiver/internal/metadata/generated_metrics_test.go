// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), componenttest.NewNopReceiverCreateSettings(), WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	mb.RecordNtpFrequencyOffsetDataPoint(ts, 1, AttributeLeapStatus(1))

	enabledMetrics["ntp.skew"] = true
	mb.RecordNtpSkewDataPoint(ts, 1)

	mb.RecordNtpStratumDataPoint(ts, 1)

	enabledMetrics["ntp.time.correction"] = true
	mb.RecordNtpTimeCorrectionDataPoint(ts, 1, AttributeLeapStatus(1))

	enabledMetrics["ntp.time.last_offset"] = true
	mb.RecordNtpTimeLastOffsetDataPoint(ts, 1, AttributeLeapStatus(1))

	mb.RecordNtpTimeRmsOffsetDataPoint(ts, 1, AttributeLeapStatus(1))

	mb.RecordNtpTimeRootDelayDataPoint(ts, 1, AttributeLeapStatus(1))

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		NtpFrequencyOffset: MetricSettings{Enabled: true},
		NtpSkew:            MetricSettings{Enabled: true},
		NtpStratum:         MetricSettings{Enabled: true},
		NtpTimeCorrection:  MetricSettings{Enabled: true},
		NtpTimeLastOffset:  MetricSettings{Enabled: true},
		NtpTimeRmsOffset:   MetricSettings{Enabled: true},
		NtpTimeRootDelay:   MetricSettings{Enabled: true},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordNtpFrequencyOffsetDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpSkewDataPoint(ts, 1)
	mb.RecordNtpStratumDataPoint(ts, 1)
	mb.RecordNtpTimeCorrectionDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpTimeLastOffsetDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpTimeRmsOffsetDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpTimeRootDelayDataPoint(ts, 1, AttributeLeapStatus(1))

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "ntp.frequency.offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The frequency is the rate by which the system s clock would be wrong if chronyd was not correcting it.", ms.At(i).Description())
			assert.Equal(t, "ppm", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("leap.status")
			assert.True(t, ok)
			assert.Equal(t, "normal", attrVal.Str())
			validatedMetrics["ntp.frequency.offset"] = struct{}{}
		case "ntp.skew":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "This is the estimated error bound on the frequency.", ms.At(i).Description())
			assert.Equal(t, "ppm", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["ntp.skew"] = struct{}{}
		case "ntp.stratum":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The number of hops away from the reference system keeping the reference time", ms.At(i).Description())
			assert.Equal(t, "{count}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["ntp.stratum"] = struct{}{}
		case "ntp.time.correction":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The number of seconds difference between the system's clock and the reference clock", ms.At(i).Description())
			assert.Equal(t, "seconds", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("leap.status")
			assert.True(t, ok)
			assert.Equal(t, "normal", attrVal.Str())
			validatedMetrics["ntp.time.correction"] = struct{}{}
		case "ntp.time.last_offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The estimated local offset on the last clock update", ms.At(i).Description())
			assert.Equal(t, "seconds", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("leap.status")
			assert.True(t, ok)
			assert.Equal(t, "normal", attrVal.Str())
			validatedMetrics["ntp.time.last_offset"] = struct{}{}
		case "ntp.time.rms_offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "the long term average of the offset value", ms.At(i).Description())
			assert.Equal(t, "seconds", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("leap.status")
			assert.True(t, ok)
			assert.Equal(t, "normal", attrVal.Str())
			validatedMetrics["ntp.time.rms_offset"] = struct{}{}
		case "ntp.time.root_delay":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "This is the total of the network path delays to the stratum-1 system from which the system is ultimately synchronised.", ms.At(i).Description())
			assert.Equal(t, "seconds", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("leap.status")
			assert.True(t, ok)
			assert.Equal(t, "normal", attrVal.Str())
			validatedMetrics["ntp.time.root_delay"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		NtpFrequencyOffset: MetricSettings{Enabled: false},
		NtpSkew:            MetricSettings{Enabled: false},
		NtpStratum:         MetricSettings{Enabled: false},
		NtpTimeCorrection:  MetricSettings{Enabled: false},
		NtpTimeLastOffset:  MetricSettings{Enabled: false},
		NtpTimeRmsOffset:   MetricSettings{Enabled: false},
		NtpTimeRootDelay:   MetricSettings{Enabled: false},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())
	mb.RecordNtpFrequencyOffsetDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpSkewDataPoint(ts, 1)
	mb.RecordNtpStratumDataPoint(ts, 1)
	mb.RecordNtpTimeCorrectionDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpTimeLastOffsetDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpTimeRmsOffsetDataPoint(ts, 1, AttributeLeapStatus(1))
	mb.RecordNtpTimeRootDelayDataPoint(ts, 1, AttributeLeapStatus(1))

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
