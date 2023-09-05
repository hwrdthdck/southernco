package sampling

import (
	"io"
	"regexp"
	"strconv"
)

type OTelTraceState struct {
	commonTraceState

	// sampling r and t-values
	rnd Randomness // r value parsed, as unsigned
	r   string     // 14 ASCII hex digits
	tt  Threshold  // t value parsed, as a threshold
	t   string     // 1-14 ASCII hex digits
}

const (
	// hardMaxOTelLength is the maximum encoded size of an OTel
	// tracestate value.
	hardMaxOTelLength = 256

	// chr        = ucalpha / lcalpha / DIGIT / "." / "_" / "-"
	// ucalpha    = %x41-5A ; A-Z
	// lcalpha    = %x61-7A ; a-z
	// key        = lcalpha *(lcalpha / DIGIT )
	// value      = *(chr)
	// list-member = key ":" value
	// list        = list-member *( ";" list-member )
	otelKeyRegexp             = lcAlphaRegexp + lcDigitRegexp + `*`
	otelValueRegexp           = `[a-zA-Z0-9._\-]*`
	otelMemberRegexp          = `(?:` + otelKeyRegexp + `:` + otelValueRegexp + `)`
	otelSemicolonMemberRegexp = `(?:` + `;` + otelMemberRegexp + `)`
	otelTracestateRegexp      = `^` + otelMemberRegexp + otelSemicolonMemberRegexp + `*$`
)

var (
	otelTracestateRe = regexp.MustCompile(otelTracestateRegexp)

	otelSyntax = keyValueScanner{
		maxItems:  -1,
		trim:      false,
		separator: ';',
		equality:  ':',
	}
)

func NewOTelTraceState(input string) (otts OTelTraceState, _ error) {
	if len(input) > hardMaxOTelLength {
		return otts, ErrTraceStateSize
	}

	if !otelTracestateRe.MatchString(input) {
		return OTelTraceState{}, strconv.ErrSyntax
	}

	err := otelSyntax.scanKeyValues(input, func(key, value string) error {
		var err error
		switch key {
		case "r":
			if otts.rnd, err = RValueToRandomness(value); err == nil {
				otts.r = value
			} else {
				otts.rnd = Randomness{} // @@@
			}
		case "t":
			if otts.tt, err = TValueToThreshold(value); err == nil {
				otts.t = value
			} else {
				otts.tt = AlwaysSampleThreshold
			}
		default:
			otts.kvs = append(otts.kvs, KV{
				Key:   key,
				Value: value,
			})
		}
		return err
	})

	return otts, err
}

func (otts *OTelTraceState) HasRValue() bool {
	return otts.r != ""
}

func (otts *OTelTraceState) RValue() string {
	return otts.r
}

func (otts *OTelTraceState) RValueRandomness() Randomness {
	return otts.rnd
}

func (otts *OTelTraceState) HasTValue() bool {
	return otts.t != ""
}

func (otts *OTelTraceState) HasNonZeroTValue() bool {
	return otts.HasTValue() && otts.TValueThreshold() != NeverSampleThreshold
}

func (otts *OTelTraceState) TValue() string {
	return otts.t
}

func (otts *OTelTraceState) TValueThreshold() Threshold {
	return otts.tt
}

func (otts *OTelTraceState) SetTValue(threshold Threshold, encoded string) {
	otts.tt = threshold
	otts.t = encoded
}

func (otts *OTelTraceState) UnsetTValue() {
	otts.t = ""
	otts.tt = Threshold{}
}

func (otts *OTelTraceState) HasAnyValue() bool {
	return otts.HasRValue() || otts.HasTValue() || otts.HasExtraValues()
}

func (otts *OTelTraceState) Serialize(w io.StringWriter) {
	cnt := 0
	sep := func() {
		if cnt != 0 {
			w.WriteString(";")
		}
		cnt++
	}
	if otts.HasRValue() {
		sep()
		w.WriteString("r:")
		w.WriteString(otts.RValue())
	}
	if otts.HasTValue() {
		sep()
		w.WriteString("t:")
		w.WriteString(otts.TValue())
	}
	for _, kv := range otts.ExtraValues() {
		sep()
		w.WriteString(kv.Key)
		w.WriteString(":")
		w.WriteString(kv.Value)
	}
}
