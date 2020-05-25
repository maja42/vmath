package math32

import "math"

// Ilogb returns the binary exponent of x as an integer.
//
// Special cases are:
//	Ilogb(±Inf) = MaxInt32
//	Ilogb(0) = MinInt32
//	Ilogb(NaN) = MaxInt32
func Ilogb(x float32) int {
	return math.Ilogb(float64(x))
}

// Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).
//
// Special cases are:
//	Lgamma(+Inf) = +Inf
//	Lgamma(0) = +Inf
//	Lgamma(-integer) = +Inf
//	Lgamma(-Inf) = -Inf
//	Lgamma(NaN) = NaN
func Lgamma(x float32) (lgamma float32, sign int) {
	g, s := math.Lgamma(float64(x))
	return float32(g), s
}

// Log returns the natural logarithm of x.
//
// Special cases are:
//	Log(+Inf) = +Inf
//	Log(0) = -Inf
//	Log(x < 0) = NaN
//	Log(NaN) = NaN
func Log(x float32) float32 {
	return float32(math.Log(float64(x)))
}

// Log10 returns the decimal logarithm of x.
// The special cases are the same as for Log.
func Log10(x float32) float32 {
	return float32(math.Log10(float64(x)))
}

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
//
// Special cases are:
//	Log1p(+Inf) = +Inf
//	Log1p(±0) = ±0
//	Log1p(-1) = -Inf
//	Log1p(x < -1) = NaN
//	Log1p(NaN) = NaN
func Log1p(x float32) float32 {
	return float32(math.Log1p(float64(x)))
}

// Log2 returns the binary logarithm of x.
// The special cases are the same as for Log.
func Log2(x float32) float32 {
	return float32(math.Log2(float64(x)))
}

// Logb returns the binary exponent of x.
//
// Special cases are:
//	Logb(±Inf) = +Inf
//	Logb(0) = -Inf
//	Logb(NaN) = NaN
func Logb(x float32) float32 {
	return float32(math.Logb(float64(x)))
}
