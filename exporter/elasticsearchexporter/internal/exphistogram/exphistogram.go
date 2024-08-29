package exphistogram

import (
	"math"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

// LowerBoundary calculates the lower boundary given index and scale.
// Adopted from https://opentelemetry.io/docs/specs/otel/metrics/data-model/#producer-expectations
func LowerBoundary(index, scale int) float64 {
	if scale <= 0 {
		return LowerBoundaryNegativeScale(index, scale)
	}
	// Use this form in case the equation above computes +Inf
	// as the lower boundary of a valid bucket.
	inverseFactor := math.Ldexp(math.Ln2, -scale)
	return 2.0 * math.Exp(float64(index-(1<<scale))*inverseFactor)
}

// LowerBoundaryNegativeScale calculates the lower boundary for scale <= 0.
func LowerBoundaryNegativeScale(index, scale int) float64 {
	return math.Ldexp(1, index<<-scale)
}

// MapToIndex gets bucket index from value and scale.
// Adopted from https://opentelemetry.io/docs/specs/otel/metrics/data-model/#producer-expectations
func MapToIndex(value float64, scale int) int {
	// Special case for power-of-two values.
	if frac, exp := math.Frexp(value); frac == 0.5 {
		return ((exp - 1) << scale) - 1
	}
	scaleFactor := math.Ldexp(math.Log2E, scale)
	// Note: math.Floor(value) equals math.Ceil(value)-1 when value
	// is not a power of two, which is checked above.
	return int(math.Floor(math.Log(value) * scaleFactor))
}

func ToTDigest(dp pmetric.ExponentialHistogramDataPoint) (counts []int64, values []float64) {
	scale := int(dp.Scale())

	offset := int(dp.Negative().Offset())
	bucketCounts := dp.Negative().BucketCounts()
	for i := bucketCounts.Len() - 1; i >= 0; i-- {
		count := bucketCounts.At(i)
		if count == 0 {
			continue
		}
		lb := -LowerBoundary(offset+i+1, scale)
		ub := -LowerBoundary(offset+i, scale)
		counts = append(counts, int64(count))
		values = append(values, lb+(ub-lb)/2)
	}

	if zeroCount := dp.ZeroCount(); zeroCount != 0 {
		counts = append(counts, int64(zeroCount))
		// The midpoint is only non-zero when positive offset and negative offset are not the same,
		// but the midpoint between negative and positive boundaries closest to zero will not be very meaningful anyway.
		// Using a zero here instead.
		values = append(values, 0)
	}

	offset = int(dp.Positive().Offset())
	bucketCounts = dp.Positive().BucketCounts()
	for i := 0; i < bucketCounts.Len(); i++ {
		count := bucketCounts.At(i)
		if count == 0 {
			continue
		}
		lb := LowerBoundary(offset+i, scale)
		ub := LowerBoundary(offset+i+1, scale)
		counts = append(counts, int64(count))
		values = append(values, lb+(ub-lb)/2)
	}
	return
}
