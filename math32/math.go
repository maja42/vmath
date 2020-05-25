package math32

import "math"

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(v float32) float32 {
	return math.Float32frombits(math.Float32bits(v) &^ (1 << 31))
}

// Cbrt returns the cube root of x.
//
// Special cases are:
//	Cbrt(±0) = ±0
//	Cbrt(±Inf) = ±Inf
//	Cbrt(NaN) = NaN
func Cbrt(x float32) float32 {
	return float32(math.Cbrt(float64(x)))
}

// FMA returns x * y + z, computed with only one rounding.
// (That is, FMA returns the fused multiply-add of x, y, and z.)
func FMA(x, y, z float32) float32 {
	return float32(math.FMA(float64(x), float64(y), float64(z)))
}

// Frexp breaks f into a normalized fraction
// and an integral power of two.
// It returns frac and exp satisfying f == frac × 2**exp,
// with the absolute value of frac in the interval [½, 1).
//
// Special cases are:
//	Frexp(±0) = ±0, 0
//	Frexp(±Inf) = ±Inf, 0
//	Frexp(NaN) = NaN, 0
func Frexp(f float32) (float32, int) {
	frac, exp := math.Frexp(float64(f))
	return float32(frac), exp
}

// Ldexp is the inverse of Frexp.
// It returns frac × 2**exp.
//
// Special cases are:
//	Ldexp(±0, exp) = ±0
//	Ldexp(±Inf, exp) = ±Inf
//	Ldexp(NaN, exp) = NaN
func Ldexp(frac float32, exp int) float32 {
	return float32(math.Ldexp(float64(frac), exp))
}

// Gamma returns the Gamma function of x.
//
// Special cases are:
//	Gamma(+Inf) = +Inf
//	Gamma(+0) = +Inf
//	Gamma(-0) = -Inf
//	Gamma(x) = NaN for integer x < 0
//	Gamma(-Inf) = NaN
//	Gamma(NaN) = NaN
func Gamma(x float32) float32 {
	return float32(math.Gamma(float64(x)))
}

// Min returns the smaller of x or y.
//
// Special cases are:
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
func Min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

// Max returns the larger of x or y.
//
// Special cases are:
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
func Max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

// Hypot returns Sqrt(p*p + q*q), taking care to avoid
// unnecessary overflow and underflow.
//
// Special cases are:
//	Hypot(±Inf, q) = +Inf
//	Hypot(p, ±Inf) = +Inf
//	Hypot(NaN, q) = NaN
//	Hypot(p, NaN) = NaN
func Hypot(p, q float32) float32 {
	return float32(math.Hypot(float64(p), float64(q)))
}

// Ceil returns the least integer value greater than or equal to x.
//
// Special cases are:
//	Ceil(±0) = ±0
//	Ceil(±Inf) = ±Inf
//	Ceil(NaN) = NaN
func Ceil(x float32) float32 {
	return float32(math.Ceil(float64(x)))
}

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

// Round returns the nearest integer, rounding half away from zero.
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round(v float32) float32 {
	return float32(math.Round(float64(v)))
}

// RoundToEven returns the nearest integer, rounding ties to even.
//
// Special cases are:
//	RoundToEven(±0) = ±0
//	RoundToEven(±Inf) = ±Inf
//	RoundToEven(NaN) = NaN
func RoundToEven(x float32) float32 {
	return float32(math.RoundToEven(float64(x)))
}

// Sqrt returns the square root of x.
//
// Special cases are:
//	Sqrt(+Inf) = +Inf
//	Sqrt(±0) = ±0
//	Sqrt(x < 0) = NaN
//	Sqrt(NaN) = NaN
func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

// Mod returns the floating-point remainder of x/y.
// The magnitude of the result is less than y and its
// sign agrees with that of x.
//
// Special cases are:
//	Mod(±Inf, y) = NaN
//	Mod(NaN, y) = NaN
//	Mod(x, 0) = NaN
//	Mod(x, ±Inf) = x
//	Mod(x, NaN) = NaN
func Mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

// Modf returns integer and fractional floating-point numbers
// that sum to f. Both values have the same sign as f.
//
// Special cases are:
//	Modf(±Inf) = ±Inf, NaN
//	Modf(NaN) = NaN, NaN
func Modf(f float32) (int float32, frac float32) {
	i, fr := math.Modf(float64(f))
	return float32(i), float32(fr)
}

// Remainder returns the IEEE 754 floating-point remainder of x/y.
//
// Special cases are:
//	Remainder(±Inf, y) = NaN
//	Remainder(NaN, y) = NaN
//	Remainder(x, 0) = NaN
//	Remainder(x, ±Inf) = x
//	Remainder(x, NaN) = NaN
func Remainder(x, y float32) float32 {
	return float32(math.Remainder(float64(x), float64(y)))
}

// Signbit reports whether x is negative or negative zero.
func Signbit(x float32) bool {
	return math.Float32bits(x)&(1<<31) != 0
}

// Copysign returns a value with the magnitude
// of x and the sign of y.
func Copysign(x, y float32) float32 {
	const sign = 1 << 31
	return math.Float32frombits(math.Float32bits(x)&^sign | math.Float32bits(y)&sign)
}

// Trunc returns the integer value of x.
//
// Special cases are:
//	Trunc(±0) = ±0
//	Trunc(±Inf) = ±Inf
//	Trunc(NaN) = NaN
func Trunc(x float32) float32 {
	return float32(math.Trunc(float64(x)))
}
