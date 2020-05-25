package math32

import "math"

// Sin returns the sine of the radian argument x.
//
// Special cases are:
//	Sin(±0) = ±0
//	Sin(±Inf) = NaN
//	Sin(NaN) = NaN
func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

// Asin returns the arcsine, in radians, of x.
//
// Special cases are:
//	Asin(±0) = ±0
//	Asin(x) = NaN if x < -1 or x > 1
func Asin(x float32) float32 {
	return float32(math.Asin(float64(x)))
}

// Cos returns the cosine of the radian argument x.
//
// Special cases are:
//	Cos(±Inf) = NaN
//	Cos(NaN) = NaN
func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

// Acos returns the arccosine, in radians, of x.
//
// Special case is:
//	Acos(x) = NaN if x < -1 or x > 1
func Acos(x float32) float32 {
	return float32(math.Acos(float64(x)))
}

// Sincos returns Sin(x), Cos(x).
//
// Special cases are:
//	Sincos(±0) = ±0, 1
//	Sincos(±Inf) = NaN, NaN
//	Sincos(NaN) = NaN, NaN
func Sincos(x float32) (sin, cos float32) {
	s, c := math.Sincos(float64(x))
	return float32(s), float32(c)
}

// Tan returns the tangent of the radian argument x.
//
// Special cases are:
//	Tan(±0) = ±0
//	Tan(±Inf) = NaN
//	Tan(NaN) = NaN
func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}

// Atan returns the arctangent, in radians, of x.
//
// Special cases are:
//      Atan(±0) = ±0
//      Atan(±Inf) = ±Pi/2
func Atan(x float32) float32 {
	return float32(math.Atan(float64(x)))
}

// Atan2 returns the arc tangent of y/x, using
// the signs of the two to determine the quadrant
// of the return value.
//
// Special cases are (in order):
//	Atan2(y, NaN) = NaN
//	Atan2(NaN, x) = NaN
//	Atan2(+0, x>=0) = +0
//	Atan2(-0, x>=0) = -0
//	Atan2(+0, x<=-0) = +Pi
//	Atan2(-0, x<=-0) = -Pi
//	Atan2(y>0, 0) = +Pi/2
//	Atan2(y<0, 0) = -Pi/2
//	Atan2(+Inf, +Inf) = +Pi/4
//	Atan2(-Inf, +Inf) = -Pi/4
//	Atan2(+Inf, -Inf) = 3Pi/4
//	Atan2(-Inf, -Inf) = -3Pi/4
//	Atan2(y, +Inf) = 0
//	Atan2(y>0, -Inf) = +Pi
//	Atan2(y<0, -Inf) = -Pi
//	Atan2(+Inf, x) = +Pi/2
//	Atan2(-Inf, x) = -Pi/2
func Atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}
