package metrics

import (
	"errors"
)

type metricType string

const (
	metricTypeAll       metricType = "all"
	metricTypeGauge     metricType = "gauge"
	metricTypeHistogram metricType = "histogram"
	metricTypeSum       metricType = "sum"
)

// String is used both by fmt.Print and by Cobra in help text
func (e *metricType) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *metricType) Set(v string) error {
	switch v {
	case "gauge", "sum", "histogram", "all":
		*e = metricType(v)
		return nil
	default:
		return errors.New(`must be one of "gauge", "sum", "histogram" or "all"`)
	}
}

// Type is only used in help text
func (e *metricType) Type() string {
	return "metricType"
}
