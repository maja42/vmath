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
func Radians(deg float32) float32 {
	return math.Pi * deg / 180.0
}

// Degrees converts radians into degrees.
func Degrees(rad float32) float32 {
	return rad * (180.0 / math.Pi)
}

// CartesianToSpherical converts cartesian coordinates into spherical coordinates.
// Returns the radius, azimuth (angle on XY-plane) and inclination.
func CartesianToSpherical(pos Vec3f) (float32, float32, float32) {
	radius := pos.Length()
	azimuth := Atan2(pos[1], pos[0])
	inclination := Acos(pos[2] / radius)
	return radius, azimuth, inclination
}

// SphericalToCartesian converts spherical coordinates into cartesian coordinates.
func SphericalToCartesian(radius, azimuth, inclination float32) Vec3f {
	sinAz, cosAz := Sincos(azimuth)
	sinInc, cosInc := Sincos(inclination)

	return Vec3f{
		radius * sinInc * cosAz,
		radius * sinInc * sinAz,
		radius * cosInc,
	}
}

// Lerp performs a linear interpolation between a and b.
// The parameter t should be in range [0, 1].
func Lerp(a, b, t float64) float64 {
	return a*(1-t) + b*t
}

// NormalizeRadians returns the angle in radians in the range [0, 2*PI[.
func NormalizeRadians(rad float32) float32 {
	var pi2 float32 = math.Pi * 2
	rad += pi2 * float32(int(rad/-pi2)+1)
	rad -= pi2 * float32(int(rad/pi2))
	return rad
}

// NormalizeDegrees returns the angle in degrees in the range [0, 360[.
func NormalizeDegrees(deg float32) float32 {
	deg += float32(360 * (int(deg/-360) + 1))
	deg -= float32(360 * int(deg/360))
	return deg
}

// AngleToVector returns a 2D vector with the given length and angle to the x-axis.
func AngleToVector(rad float32, length float32) Vec2f {
	sin, cos := Sincos(rad)
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
func PointToLineDistance2D(a, b, point Vec2f) float32 {
	// Source: http://geomalgorithms.com/a02-_lines.html
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
func PointToLineSegmentDistance2D(a, b, point Vec2f) float32 {
	// Source: http://geomalgorithms.com/a02-_lines.html
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

// IsPointOnLine returns true if the give point lies to the line a->b;
// Uses the default Epsilon as relative tolerance.
func IsPointOnLine(a, b Vec2f, point Vec2f) bool {
	return IsPointOnLineEps(a, b, point, Epsilon)
}

// IsPointOnLine returns true if the give point lies to the line a->b;
// Uses the given Epsilon as relative tolerance.
func IsPointOnLineEps(a, b Vec2f, point Vec2f, eps float32) bool {
	lineVec := b.Sub(a)
	pointVec := point.Sub(a)
	crossZ := lineVec[0]*pointVec[1] - lineVec[1]*pointVec[0]
	return EqualEps(crossZ, 0, eps)
}

// PolarToCartesian2D converts length and angle into a 2D position.
func PolarToCartesian2D(distance, rad float32) Vec2f {
	sin, cos := Sincos(rad)
	return Vec2f{
		cos * distance,
		sin * distance,
	}
}

// IsPointOnLeft returns true if the give point lies to the left of line a->b;
// If the point lies directly on the line, false is returned.
func IsPointOnLeft(a, b Vec2f, point Vec2f) bool {
	lineVec := b.Sub(a)
	pointVec := point.Sub(a)
	crossZ := lineVec[0]*pointVec[1] - lineVec[1]*pointVec[0]
	return crossZ > 0
}
