package vmath

import (
	"math"
)

// Equalf compares two floats for equality, using the given epsilon as the relative tolerance.
// Performs a relative difference comparison (see https://floating-point-gui.de/errors/comparison/ and https://stackoverflow.com/q/4915462/2224996)
func Equalf(a, b, epsilon float32) bool {
	if a == b { // shortcut; also handles +-Inf
		return true
	}

	diff := Absf(a - b)
	if a == 0 || b == 0 || diff < minNormal { // If a or b are 0 or both are extremely close to it
		return diff < epsilon*minNormal
	}

	return diff/(Absf(a)+Absf(b)) < epsilon
}

// minNormal is he smallest possible float32 number, provided that there is a 1 in front of the binary (=decimal) point.
// Do not confuse with "math.SmallestNonzeroFloat32", where this restriction is not present
// 1 / 2^(127 - 1)
const minNormal = float32(1.1754943508222875e-38)

// Absf converts a float into an absolute (=positive) value.
func Absf(v float32) float32 {
	return math.Float32frombits(math.Float32bits(v) &^ (1 << 31))
}

// Absi converts an int into an absolute (=positive) value.
func Absi(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// Roundf returns the nearest integer, rounding half away from zero.
func Roundf(v float32) int {
	return int(math.Round(float64(v)))
}

// Minf returns the smaller value.
func Minf(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

// Maxf returns the bigger value.
func Maxf(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

// Mini returns the smaller integer value.
func Mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Mini returns the smaller integer value.
func Maxi(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Clampf returns the value v clamped to the range of [min, max].
func Clampf(v, min, max float32) float32 {
	if v <= min {
		return min
	}
	if v >= max {
		return max
	}
	return v
}

// Clampi returns the value v clamped to the range of [min, max].
func Clampi(v, min, max int) int {
	if v <= min {
		return min
	}
	if v >= max {
		return max
	}
	return v
}

// Radians converts degrees into radians.
func Radians(degrees float32) (radians float32) {
	return math.Pi * degrees / 180.0
}

// Degrees converts radians into degrees.
func Degrees(radians float32) (degrees float32) {
	return radians * (180.0 / math.Pi)
}

// Lerp performs a linear interpolation between a and b.
// The parameter t should be in range [0, 1].
func Lerp(a, b, t float64) float64 {
	return a*(1-t) + b*t
}
