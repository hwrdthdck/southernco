package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"strings"
	"unicode"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

func ToSnakeCase(metric pmetric.Metric) {

	s := metric.Name()

	if s == "" {
		return
	}
	if len(s) == 1 {
		metric.SetName(strings.ToLower(s))
		return
	}
	source := []rune(s)
	dist := strings.Builder{}
	dist.Grow(len(s) + len(s)/3) // avoid reallocation memory, 33% ~ 50% is recommended
	skipNext := false
	for i := 0; i < len(source); i++ {
		cur := source[i]
		switch cur {
		case '-', '_':
			dist.WriteRune('_')
			skipNext = true
			continue
		}
		if unicode.IsLower(cur) || unicode.IsDigit(cur) {
			dist.WriteRune(cur)
			continue
		}

		if i == 0 {
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}

		last := source[i-1]
		if (!unicode.IsLetter(last)) || unicode.IsLower(last) {
			if skipNext {
				skipNext = false
			} else {
				dist.WriteRune('_')
			}
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}
		// last is upper case
		if i < len(source)-1 {
			next := source[i+1]
			if unicode.IsLower(next) {
				if skipNext {
					skipNext = false
				} else {
					dist.WriteRune('_')
				}
				dist.WriteRune(unicode.ToLower(cur))
				continue
			}
		}
		dist.WriteRune(unicode.ToLower(cur))
	}

	metric.SetName(dist.String())
}
