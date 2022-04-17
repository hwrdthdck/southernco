// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"fmt"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/model/semconv/v1.9.0"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/load metrics.
type MetricsSettings struct {
	SystemCPULoadAverage15m MetricSettings `mapstructure:"system.cpu.load_average.15m"`
	SystemCPULoadAverage1m  MetricSettings `mapstructure:"system.cpu.load_average.1m"`
	SystemCPULoadAverage5m  MetricSettings `mapstructure:"system.cpu.load_average.5m"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemCPULoadAverage15m: MetricSettings{
			Enabled: true,
		},
		SystemCPULoadAverage1m: MetricSettings{
			Enabled: true,
		},
		SystemCPULoadAverage5m: MetricSettings{
			Enabled: true,
		},
	}
}

type metricSystemCPULoadAverage15m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.15m metric with initial data.
func (m *metricSystemCPULoadAverage15m) init() {
	m.data.SetName("system.cpu.load_average.15m")
	m.data.SetDescription("Average CPU Load over 15 minutes.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
}

func (m *metricSystemCPULoadAverage15m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage15m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage15m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage15m(settings MetricSettings) metricSystemCPULoadAverage15m {
	m := metricSystemCPULoadAverage15m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULoadAverage1m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.1m metric with initial data.
func (m *metricSystemCPULoadAverage1m) init() {
	m.data.SetName("system.cpu.load_average.1m")
	m.data.SetDescription("Average CPU Load over 1 minute.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
}

func (m *metricSystemCPULoadAverage1m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage1m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage1m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage1m(settings MetricSettings) metricSystemCPULoadAverage1m {
	m := metricSystemCPULoadAverage1m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULoadAverage5m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.5m metric with initial data.
func (m *metricSystemCPULoadAverage5m) init() {
	m.data.SetName("system.cpu.load_average.5m")
	m.data.SetDescription("Average CPU Load over 5 minutes.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
}

func (m *metricSystemCPULoadAverage5m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage5m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage5m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage5m(settings MetricSettings) metricSystemCPULoadAverage5m {
	m := metricSystemCPULoadAverage5m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                     pcommon.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity               int               // maximum observed number of metrics per resource.
	resourceCapacity              int               // maximum observed number of resource attributes.
	metricsBuffer                 pmetric.Metrics   // accumulates metrics data before emitting.
	metricSystemCPULoadAverage15m metricSystemCPULoadAverage15m
	metricSystemCPULoadAverage1m  metricSystemCPULoadAverage1m
	metricSystemCPULoadAverage5m  metricSystemCPULoadAverage5m
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                     pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                 pmetric.NewMetrics(),
		metricSystemCPULoadAverage15m: newMetricSystemCPULoadAverage15m(settings.SystemCPULoadAverage15m),
		metricSystemCPULoadAverage1m:  newMetricSystemCPULoadAverage1m(settings.SystemCPULoadAverage1m),
		metricSystemCPULoadAverage5m:  newMetricSystemCPULoadAverage5m(settings.SystemCPULoadAverage5m),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceOption applies changes to provided resource.
type ResourceOption func(pcommon.Resource)

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/load")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemCPULoadAverage15m.emit(ils.Metrics())
	mb.metricSystemCPULoadAverage1m.emit(ils.Metrics())
	mb.metricSystemCPULoadAverage5m.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pmetric.Metrics {
	mb.EmitForResource(ro...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSystemCPULoadAverage15mDataPoint adds a data point to system.cpu.load_average.15m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage15mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage15m.recordDataPoint(mb.startTime, ts, val)
}

// ParseSystemCPULoadAverage15mDataPoint attempts to parse and add a data point to system.cpu.load_average.15m metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseSystemCPULoadAverage15mDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors) {
	if f, err := strconv.ParseFloat(val, 64); err != nil {
		errors.AddPartial(1, fmt.Errorf("failed to parse float for SystemCPULoadAverage15m, value was %s: %w", val, err))
	} else {
		mb.metricSystemCPULoadAverage15m.recordDataPoint(mb.startTime, ts, f)
	}
}

// RecordSystemCPULoadAverage1mDataPoint adds a data point to system.cpu.load_average.1m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage1mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage1m.recordDataPoint(mb.startTime, ts, val)
}

// ParseSystemCPULoadAverage1mDataPoint attempts to parse and add a data point to system.cpu.load_average.1m metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseSystemCPULoadAverage1mDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors) {
	if f, err := strconv.ParseFloat(val, 64); err != nil {
		errors.AddPartial(1, fmt.Errorf("failed to parse float for SystemCPULoadAverage1m, value was %s: %w", val, err))
	} else {
		mb.metricSystemCPULoadAverage1m.recordDataPoint(mb.startTime, ts, f)
	}
}

// RecordSystemCPULoadAverage5mDataPoint adds a data point to system.cpu.load_average.5m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage5mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage5m.recordDataPoint(mb.startTime, ts, val)
}

// ParseSystemCPULoadAverage5mDataPoint attempts to parse and add a data point to system.cpu.load_average.5m metric.
// Function returns whether or not a data point was successfully recorded
func (mb *MetricsBuilder) ParseSystemCPULoadAverage5mDataPoint(ts pdata.Timestamp, val string, errors scrapererror.ScrapeErrors) {
	if f, err := strconv.ParseFloat(val, 64); err != nil {
		errors.AddPartial(1, fmt.Errorf("failed to parse float for SystemCPULoadAverage5m, value was %s: %w", val, err))
	} else {
		mb.metricSystemCPULoadAverage5m.recordDataPoint(mb.startTime, ts, f)
	}
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
}{}

// A is an alias for Attributes.
var A = Attributes
