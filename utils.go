package vmath

import (
	"math"
)

// Epsilon is the default epsilon value for float comparisons.
const Epsilon = 1.0E-8

// Equalf compares two floats for equality.
// Uses the default Epsilon as relative tolerance.
func Equalf(a, b float32) bool {
	return EqualEps(a, b, Epsilon)
}

// Equalf compares two floats for equality, using the given epsilon as the relative tolerance.
// Performs a relative difference comparison (see https://floating-point-gui.de/errors/comparison/ and https://stackoverflow.com/q/4915462/2224996)
func EqualEps(a, b, epsilon float32) bool {
	if a == b { // shortcut; also handles +-Inf
		return true
	}

	diff := Abs(a - b)
	if a == 0 || b == 0 || diff < minNormal {
		return diff < epsilon*minNormal
	}

	return diff/(Abs(a)+Abs(b)) < epsilon
}

// minNormal is he smallest possible float32 number, provided that there is a 1 in front of the binary (=decimal) point.
// Do not confuse with "math.SmallestNonzeroFloat32", where this restriction is not present
// 1 / 2^(127 - 1)
const minNormal = float32(1.1754943508222875e-38)

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

// Wrapf returns the value v in the range of [min, max] by wrapping it around.
func Wrapf(v, min, max float32) float32 {
	diff := max - min
	v -= min
	return min + v - diff*Floor(v/diff)
}

// Wrapi returns the value v in the range of [min, max] by wrapping it around.
func Wrapi(v, min, max int) int {
	return int(Wrapf(float32(v), float32(min), float32(max)))
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

// NormalizeRadians returns the angle in radians in the range [0, 2*PI[.
func NormalizeRadians(radians float32) float32 {
	var pi2 float32 = math.Pi * 2
	radians += pi2 * float32(int(radians/-pi2)+1)
	radians -= pi2 * float32(int(radians/pi2))
	return radians
}

// NormalizeDegrees returns the angle in degrees in the range [0, 360[.
func NormalizeDegrees(degrees float32) float32 {
	degrees += float32(360 * (int(degrees/-360) + 1))
	degrees -= float32(360 * int(degrees/360))
	return degrees
}

// AngleToVector returns a 2D vector with the given length and angle to the x-axis.
func AngleToVector(radians float32, length float32) Vec2f {
	sin, cos := Sincos(radians)
	vec := Vec2f{cos, sin}
	return vec.Normalize().MulScalar(length)
}

// AngleDiff compares to angles and returns their distance in the range ]-PI, PI].
func AngleDiff(fromRad, toRad float32) float32 {
	angle := NormalizeRadians(toRad - fromRad)
	if angle > math.Pi {
		angle -= 2 * math.Pi
	}
	return angle
}

// PointToLineDistance2D returns the distance between a point and an infinitely long line passing through a and b.
// Source: http://geomalgorithms.com/a02-_lines.html
func PointToLineDistance2D(a, b, point Vec2f) float32 {
	// 1) project the a->point vector onto the a->b vector
	// 2) calculate the intersection point
	// 3) return the distance between the point and the intersection

	lineVec := b.Sub(a)
	pointVec := point.Sub(a)
	// calc perpendicular base
	pb := a.Add(pointVec.Project(lineVec))

	return point.Sub(pb).Length()
}

// PointToLineSegmentDistance2D returns the distance between a point and a line segment between a and b.
// Source: http://geomalgorithms.com/a02-_lines.html
func PointToLineSegmentDistance2D(a, b, point Vec2f) float32 {
	// 1) determine if the point is before (a) by comparing the angle between the a->b and a->point vector
	//	  if the point is before, return the distance between the point and point a
	// 2) determine if the point is after (b) by comparing the angle between the a->b and b->point vector
	//	  if the point is afterwards, return the distance between the point and point b
	// 3) otherwise, proceed like `PointToLineDistance2D`
	lineVec := b.Sub(a)
	pointVec := point.Sub(a)

	c1 := pointVec.Dot(lineVec)
	if c1 <= 0 { // angle >= 90° --> point is before (a)
		return point.Sub(a).Length()
	}

	c2 := lineVec.Dot(lineVec)
	if c2 <= c1 { // point is after (b)
		return point.Sub(b).Length()
	}

	// calc perpendicular base
	ratio := c1 / c2
	pb := a.Add(lineVec.MulScalar(ratio))
	return point.Sub(pb).Length()
}

// PolarToCartesian2D converts length and angle into a 2D position.
func PolarToCartesian2D(distance, radians float32) Vec2f {
	sin, cos := Sincos(radians)
	return Vec2f{
		cos * distance,
		sin * distance,
	}
}
