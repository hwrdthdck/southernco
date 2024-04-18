package expo

import (
	"math"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Scale int32

// Idx gives the bucket index v belongs into
func (scale Scale) Idx(v float64) int {
	// from: https://opentelemetry.io/docs/specs/otel/metrics/data-model/#all-scales-use-the-logarithm-function

	// Special case for power-of-two values.
	if frac, exp := math.Frexp(v); frac == 0.5 {
		return ((exp - 1) << scale) - 1
	}

	scaleFactor := math.Ldexp(math.Log2E, int(scale))
	// Note: math.Floor(value) equals math.Ceil(value)-1 when value
	// is not a power of two, which is checked above.
	return int(math.Floor(math.Log(v) * scaleFactor))
}

// Bounds returns the half-open interval (min,max] of the bucket at index.
// This means a value min < v <= max belongs to this bucket.
//
// NOTE: this is different from Go slice intervals, which are [a,b)
func (scale Scale) Bounds(index int) (min, max float64) {
	// from: https://opentelemetry.io/docs/specs/otel/metrics/data-model/#all-scales-use-the-logarithm-function
	lower := func(index int) float64 {
		inverseFactor := math.Ldexp(math.Ln2, int(-scale))
		return math.Exp(float64(index) * inverseFactor)
	}

	return lower(index), lower(index + 1)
}

// Downscale collapses the buckets of bs until scale 'to' is reached
func Downscale(bs pmetric.ExponentialHistogramDataPointBuckets, from, to Scale) {
	switch {
	case from == to:
		return
	case from < to:
		panic("cannot upscale without introducing error")
	}

	for at := from; at > to; at-- {
		Collapse(bs)
	}
}

// Collapse merges adjacent buckets and zeros the remaining area:
//
//	before:	1 1 1 1 1 1 1 1 1 1 1 1
//	after:	 2   2   2   2   2   2   0   0   0   0   0   0
func Collapse(bs pmetric.ExponentialHistogramDataPointBuckets) {
	counts := bs.BucketCounts()
	size := counts.Len() / 2
	if counts.Len()%2 != 0 {
		size++
	}

	shift := 0
	if bs.Offset()%2 != 0 {
		bs.SetOffset(bs.Offset() - 1)
		shift--
	}
	bs.SetOffset(bs.Offset() / 2)

	for i := 0; i < size; i++ {
		k := i*2 + shift

		if i == 0 && k == -1 {
			counts.SetAt(i, counts.At(k+1))
			continue
		}

		counts.SetAt(i, counts.At(k))
		if k+1 < counts.Len() {
			counts.SetAt(i, counts.At(k)+counts.At(k+1))
		}
	}

	for i := size; i < counts.Len(); i++ {
		counts.SetAt(i, 0)
	}
}
