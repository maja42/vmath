package vmath

import (
	"fmt"
	"math"
)

type Vec3f [3]float32

func (v Vec3f) String() string {
	return fmt.Sprintf("Vec3f[%f x %f x %f]", v[0], v[1], v[2])
}

// Vec3i returns an integer representation of the vector.
// Decimals are truncated (rounded down).
func (v Vec3f) Vec3i() Vec3i {
	return Vec3i{int(v[0]), int(v[1]), int(v[2])}
}

// Vec4f creates a 4D vector.
func (v Vec3f) Vec4f(w float32) Vec4f {
	return Vec4f{v[0], v[1], v[2], w}
}

// Split returns the vector's components.
func (v Vec3f) Split() (x, y, z float32) {
	return v[0], v[1], v[2]
}

// X returns the vector's first component.
// Performance is equivalent to using v[0].
func (v Vec3f) X() float32 {
	return v[0]
}

// Y returns the vector's second component.
// Performance is equivalent to using v[1].
func (v Vec3f) Y() float32 {
	return v[1]
}

// Z returns the vector's third component.
// Performance is equivalent to using v[2].
func (v Vec3f) Z() float32 {
	return v[2]
}

// XY returns a 2D vector with the X and Y components.
func (v Vec3f) XY() Vec2f {
	return Vec2f{v[0], v[1]}
}

// Add performs component-wise addition.
func (v Vec3f) Add(other Vec3f) Vec3f {
	return Vec3f{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec3f) AddScalar(s float32) Vec3f {
	return Vec3f{v[0] + s, v[1] + s, v[2] + s}
}

// Sub performs component-wise subtraction.
func (v Vec3f) Sub(other Vec3f) Vec3f {
	return Vec3f{v[0] - other[0], v[1] - other[1], v[2] - other[2]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec3f) SubScalar(s float32) Vec3f {
	return Vec3f{v[0] - s, v[1] - s, v[2] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec3f) Mul(other Vec3f) Vec3f {
	return Vec3f{v[0] * other[0], v[1] * other[1], v[2] * other[2]}
}

// MulScalar performs a scalar multiplication.
func (v Vec3f) MulScalar(s float32) Vec3f {
	return Vec3f{v[0] * s, v[1] * s, v[2] * s}
}

// Div performs a component-wise division.
func (v Vec3f) Div(other Vec3f) Vec3f {
	return Vec3f{v[0] / other[0], v[1] / other[1], v[2] / other[2]}
}

// DivScalar performs a scalar division.
func (v Vec3f) DivScalar(s float32) Vec3f {
	return Vec3f{v[0] / s, v[1] / s, v[2] / s}
}

// Normalize the vector. Its length will be 1 afterwards.
// If the vector is zero, all components will be infinite afterwards.
func (v Vec3f) Normalize() Vec3f {
	l := 1.0 / v.Length()
	return Vec3f{v[0] * l, v[1] * l, v[2] * l}
}

// Length returns the vector's length.
func (v Vec3f) Length() float32 {
	return Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

// SquareLen returns the vector's squared length.
func (v Vec3f) SquareLen() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Equal compares two vectors component-wise.
// Uses the default Epsilon as relative tolerance.
func (v Vec3f) Equal(other Vec3f) bool {
	return v.EqualEps(other, Epsilon)
}

// EqualEps compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec3f) EqualEps(other Vec3f, epsilon float32) bool {
	return EqualEps(v[0], other[0], epsilon) &&
		EqualEps(v[1], other[1], epsilon) &&
		EqualEps(v[2], other[2], epsilon)
}

// Clamp clamps each component to the range of [min, max].
func (v Vec3f) Clamp(min, max float32) Vec3f {
	return Vec3f{
		Clampf(v[0], min, max),
		Clampf(v[1], min, max),
		Clampf(v[2], min, max),
	}
}

// Invert inverts (negates) all components.
func (v Vec3f) Invert() Vec3f {
	return Vec3f{-v[0], -v[1], -v[2]}
}

// Dot performs a dot product with another vector.
func (v Vec3f) Dot(other Vec3f) float32 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}

// Cross performs a cross product with another vector.
func (v Vec3f) Cross(other Vec3f) Vec3f {
	return Vec3f{
		v[1]*other[2] - v[2]*other[1],
		v[2]*other[0] - v[0]*other[2],
		v[0]*other[1] - v[1]*other[0],
	}
}

// Project returns a vector representing the projection of vector v onto "other".
func (v Vec3f) Project(other Vec3f) Vec3f {
	return other.MulScalar(v.Dot(other) / other.SquareLen())
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec3f) Lerp(other Vec3f, t float32) Vec3f {
	return v.Mul(other.MulScalar(t))
}

// Angle returns the angle between two vectors in radians.
func (v Vec3f) Angle(other Vec3f) float32 {
	v = v.Normalize()
	other = other.Normalize()

	if v.Dot(other) < 0 {
		n := v.Add(other).Length() / 2
		return math.Pi - 2.0*Asin(Min(n, 1))
	}
	n := v.Sub(other).Length() / 2
	return 2.0 * Asin(Min(n, 1))
}
