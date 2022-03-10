// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for couchbasereceiver metrics.
type MetricsSettings struct {
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{}
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	data      *pdata.Metrics // data buffer for generated metric.
	startTime pdata.Timestamp
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime: pdata.NewTimestampFromTime(time.Now()),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit() pdata.Metrics {
	if mb.data == nil {
		return pdata.NewMetrics()
	}
	return *mb.data
}

// EnsureCapacity TODO
func (mb *MetricsBuilder) EnsureCapacity(length int) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().EnsureCapacity(length)
}

// IncreaseCapacity TODO
func (mb *MetricsBuilder) IncreaseCapacity(length int) {
	if mb.data == nil {
		mb.data = mb.newMetricData()
	}
	mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().EnsureCapacity(length + mb.data.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().Len())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// newMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) newMetricData() *pdata.Metrics {
	md := pdata.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/couchbasereceiver")
	return &md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
}{}

// A is an alias for Attributes.
var A = Attributes
