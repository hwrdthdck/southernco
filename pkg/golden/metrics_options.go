package golden

// WriteMetricsOption is an option for the WriteMetrics function
type WriteMetricsOption func(*writeMetricsOptions)

type writeMetricsOptions struct {
	normalizeTimestamps bool
}

// SkipMetricTimestampNormalization is an option that skips normalizing timestamps before writing metrics to disk.
func SkipMetricTimestampNormalization() WriteMetricsOption {
	return func(wmo *writeMetricsOptions) {
		wmo.normalizeTimestamps = false
	}
}
