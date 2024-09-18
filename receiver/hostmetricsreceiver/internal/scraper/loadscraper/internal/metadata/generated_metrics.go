// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

type metricSystemCPULoadAverage15m struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.15m metric with initial data.
func (m *metricSystemCPULoadAverage15m) init() {
	m.data.SetName("system.cpu.load_average.15m")
	m.data.SetDescription("Average CPU Load over 15 minutes.")
	m.data.SetUnit("{thread}")
	m.data.SetEmptyGauge()
}

func (m *metricSystemCPULoadAverage15m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage15m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage15m) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage15m(cfg MetricConfig) metricSystemCPULoadAverage15m {
	m := metricSystemCPULoadAverage15m{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULoadAverage1m struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.1m metric with initial data.
func (m *metricSystemCPULoadAverage1m) init() {
	m.data.SetName("system.cpu.load_average.1m")
	m.data.SetDescription("Average CPU Load over 1 minute.")
	m.data.SetUnit("{thread}")
	m.data.SetEmptyGauge()
}

func (m *metricSystemCPULoadAverage1m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage1m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage1m) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage1m(cfg MetricConfig) metricSystemCPULoadAverage1m {
	m := metricSystemCPULoadAverage1m{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULoadAverage5m struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.5m metric with initial data.
func (m *metricSystemCPULoadAverage5m) init() {
	m.data.SetName("system.cpu.load_average.5m")
	m.data.SetDescription("Average CPU Load over 5 minutes.")
	m.data.SetUnit("{thread}")
	m.data.SetEmptyGauge()
}

func (m *metricSystemCPULoadAverage5m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage5m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage5m) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage5m(cfg MetricConfig) metricSystemCPULoadAverage5m {
	m := metricSystemCPULoadAverage5m{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                        MetricsBuilderConfig // config of the metrics builder.
	startTime                     pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity               int                  // maximum observed number of metrics per resource.
	metricsBuffer                 pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                     component.BuildInfo  // contains version information.
	metricSystemCPULoadAverage15m metricSystemCPULoadAverage15m
	metricSystemCPULoadAverage1m  metricSystemCPULoadAverage1m
	metricSystemCPULoadAverage5m  metricSystemCPULoadAverage5m
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                        mbc,
		startTime:                     pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                 pmetric.NewMetrics(),
		buildInfo:                     settings.BuildInfo,
		metricSystemCPULoadAverage15m: newMetricSystemCPULoadAverage15m(mbc.Metrics.SystemCPULoadAverage15m),
		metricSystemCPULoadAverage1m:  newMetricSystemCPULoadAverage1m(mbc.Metrics.SystemCPULoadAverage1m),
		metricSystemCPULoadAverage5m:  newMetricSystemCPULoadAverage5m(mbc.Metrics.SystemCPULoadAverage5m),
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemCPULoadAverage15m.emit(ils.Metrics())
	mb.metricSystemCPULoadAverage1m.emit(ils.Metrics())
	mb.metricSystemCPULoadAverage5m.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSystemCPULoadAverage15mDataPoint adds a data point to system.cpu.load_average.15m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage15mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage15m.recordDataPoint(mb.startTime, ts, val)
}

// RecordSystemCPULoadAverage1mDataPoint adds a data point to system.cpu.load_average.1m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage1mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage1m.recordDataPoint(mb.startTime, ts, val)
}

// RecordSystemCPULoadAverage5mDataPoint adds a data point to system.cpu.load_average.5m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage5mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage5m.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
