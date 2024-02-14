// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"hash"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil"
)

type Stream struct {
	metric
	attrs [16]byte
}

func (i Stream) Hash() hash.Hash64 {
	sum := i.metric.Hash()
	sum.Write(i.attrs[:])
	return sum
}

func (i Stream) Metric() Metric {
	return i.metric
}

func OfStream[DataPoint attrPoint](m Metric, dp DataPoint) Stream {
	return Stream{metric: m, attrs: pdatautil.MapHash(dp.Attributes())}
}

type attrPoint interface {
	pmetric.NumberDataPoint | pmetric.HistogramDataPoint | pmetric.ExponentialHistogramDataPoint | pmetric.SummaryDataPoint
	Attributes() pcommon.Map
}
