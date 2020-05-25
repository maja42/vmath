package math32

import "math"

const (
	uvinf = float64(0x7FF0000000000000)
)

const (
	Infinity    = float32(uvinf)
	NegInfinity = -Infinity
)

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float32 {
	return float32(math.Inf(sign))
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float32 {
	return float32(math.NaN())
}

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float32) (is bool) {
	return f != f
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf(f float32, sign int) bool {
	return math.IsInf(float64(f), sign)
}
