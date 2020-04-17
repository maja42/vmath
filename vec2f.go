package vmath

import (
	"fmt"
)

type Vec2f [2]float32

func (v Vec2f) String() string {
	return fmt.Sprintf("Vec2f[%f x %f]", v[0], v[1])
}

// Vec2i returns an integer representation of the vector.
// Decimals are truncated (rounded down).
func (v Vec2f) Vec2i() Vec2i {
	return Vec2i{int(v[0]), int(v[1])}
}

// Round returns an integer representation of the vector.
// Decimals are rounded.
func (v Vec2f) Round() Vec2i {
	return Vec2i{int(Round(v[0])), int(Round(v[1]))}
}

// Vec3f creates a 3D vector.
func (v Vec2f) Vec3f(z float32) Vec3f {
	return Vec3f{v[0], v[1], z}
}

// Vec4f creates a 4D vector.
func (v Vec2f) Vec4f(z, w float32) Vec4f {
	return Vec4f{v[0], v[1], z, w}
}

// Split returns the vector's components.
func (v Vec2f) Split() (x, y float32) {
	return v[0], v[1]
}

// X returns the vector's first component.
// Performance is equivalent to using v[0].
func (v Vec2f) X() float32 {
	return v[0]
}

// Y returns the vector's second component.
// Performance is equivalent to using v[1].
func (v Vec2f) Y() float32 {
	return v[1]
}

// Add performs component-wise addition.
func (v Vec2f) Add(other Vec2f) Vec2f {
	return Vec2f{v[0] + other[0], v[1] + other[1]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec2f) AddScalar(s float32) Vec2f {
	return Vec2f{v[0] + s, v[1] + s}
}

// Sub performs component-wise subtraction.
func (v Vec2f) Sub(other Vec2f) Vec2f {
	return Vec2f{v[0] - other[0], v[1] - other[1]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec2f) SubScalar(s float32) Vec2f {
	return Vec2f{v[0] - s, v[1] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec2f) Mul(other Vec2f) Vec2f {
	return Vec2f{v[0] * other[0], v[1] * other[1]}
}

// MulScalar performs a scalar multiplication.
func (v Vec2f) MulScalar(s float32) Vec2f {
	return Vec2f{v[0] * s, v[1] * s}
}

// Div performs a component-wise division.
func (v Vec2f) Div(other Vec2f) Vec2f {
	return Vec2f{v[0] / other[0], v[1] / other[1]}
}

// DivScalar performs a scalar division.
func (v Vec2f) DivScalar(s float32) Vec2f {
	return Vec2f{v[0] / s, v[1] / s}
}

// Normalize the vector. Its length will be 1 afterwards.
// If the vector's length is zero, a zero vector will be returned.
func (v Vec2f) Normalize() Vec2f {
	length := v.Length()
	if Equalf(length, 0) {
		return Vec2f{}
	}
	return Vec2f{v[0] / length, v[1] / length}
}

// Length returns the vector's length.
func (v Vec2f) Length() float32 {
	return Hypot(v[0], v[1])
}

// SquareLength returns the vector's squared length.
func (v Vec2f) SquareLength() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

// Equal compares two vectors component-wise.
// Uses the default Epsilon as relative tolerance.
func (v Vec2f) Equal(other Vec2f) bool {
	return v.EqualEps(other, Epsilon)
}

// EqualEps compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec2f) EqualEps(other Vec2f, epsilon float32) bool {
	return EqualEps(v[0], other[0], epsilon) && EqualEps(v[1], other[1], epsilon)
}

// Clamp clamps each component to the range of [min, max].
func (v Vec2f) Clamp(min, max float32) Vec2f {
	return Vec2f{
		Clampf(v[0], min, max),
		Clampf(v[1], min, max),
	}
}

// Negate inverts all components.
func (v Vec2f) Negate() Vec2f {
	return Vec2f{-v[0], -v[1]}
}

// Dot performs a dot product with another vector.
func (v Vec2f) Dot(other Vec2f) float32 {
	return v[0]*other[0] + v[1]*other[1]
}

// Cross performs a cross product with another vector.
func (v Vec2f) Cross(other Vec2f) Vec2f {
	return Vec2f{
		v[1]*other[0] - v[0]*other[1],
		v[0]*other[1] - v[1]*other[0],
	}
}

// NormalVec returns a normal vector on the 2D plane that is either on the left or right hand side.
func (v Vec2f) NormalVec(onLeft bool) Vec2f {
	if onLeft {
		return Vec2f{-v[1], v[0]}
	}
	return Vec2f{v[1], -v[0]}
}

// Project returns a vector representing the projection of vector v onto "other".
func (v Vec2f) Project(other Vec2f) Vec2f {
	return other.MulScalar(v.Dot(other) / other.SquareLength())
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec2f) Lerp(other Vec2f, t float32) Vec2f {
	return v.Mul(other.MulScalar(t))
}

// Angle returns the angle between two vectors in radians.
func (v Vec2f) Angle(other Vec2f) float32 {
	return Atan2(other[1]-v[1], other[0]-v[0])
}

// FlatAngle returns the angle of a vector in radians.
// This is the angle between the vector and the x-axis.
func (v Vec2f) FlatAngle() float32 {
	return Atan2(v[1], v[0])
}

// Rotate rotates the vector on the 2D plane.
func (v Vec2f) Rotate(rad float32) Vec2f {
	angle := v.FlatAngle() + rad
	return AngleToVector(angle, v.Length())
}

// Distance returns the euclidean distance to another position.
func (v Vec2f) Distance(other Vec2f) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec2f) SquareDistance(other Vec2f) float32 {
	return other.Sub(v).SquareLength()
}
