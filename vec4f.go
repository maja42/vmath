package vmath

import (
	"fmt"

	"github.com/maja42/vmath/math32"
)

type Vec4f [4]float32

func (v Vec4f) String() string {
	return fmt.Sprintf("Vec4f[%f x %f x %f x %f]", v[0], v[1], v[2], v[3])
}

// Format the vector to a string.
func (v Vec4f) Format(format string) string {
	return fmt.Sprintf(format, v[0], v[1], v[2], v[3])
}

// Vec4i returns an integer representation of the vector.
// Decimals are truncated.
func (v Vec4f) Vec4i() Vec4i {
	return Vec4i{int(v[0]), int(v[1]), int(v[2]), int(v[3])}
}

// Round returns an integer representation of the vector.
// Decimals are rounded.
func (v Vec4f) Round() Vec4i {
	return Vec4i{
		int(math32.Round(v[0])),
		int(math32.Round(v[1])),
		int(math32.Round(v[2])),
		int(math32.Round(v[3]))}
}

// Split returns the vector's components.
func (v Vec4f) Split() (x, y, z, w float32) {
	return v[0], v[1], v[2], v[3]
}

// X returns the vector's first component.
// Performance is equivalent to using v[0].
func (v Vec4f) X() float32 {
	return v[0]
}

// Y returns the vector's second component.
// Performance is equivalent to using v[1].
func (v Vec4f) Y() float32 {
	return v[1]
}

// Z returns the vector's third component.
// Performance is equivalent to using v[2].
func (v Vec4f) Z() float32 {
	return v[2]
}

// W returns the vector's fourth component.
// Performance is equivalent to using v[3].
func (v Vec4f) W() float32 {
	return v[3]
}

// XY returns a 2D vector with the X and Y components.
func (v Vec4f) XY() Vec2f {
	return Vec2f{v[0], v[1]}
}

// XYZ returns a 3D vector with the X, Y and Z components.
func (v Vec4f) XYZ() Vec3f {
	return Vec3f{v[0], v[1], v[2]}
}

// Abs returns a vector with the components turned into absolute values.
func (v Vec4f) Abs() Vec4f {
	return Vec4f{math32.Abs(v[0]), math32.Abs(v[1]), math32.Abs(v[2]), math32.Abs(v[3])}
}

// Add performs component-wise addition.
func (v Vec4f) Add(other Vec4f) Vec4f {
	return Vec4f{v[0] + other[0], v[1] + other[1], v[2] + other[2], v[3] + other[3]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec4f) AddScalar(s float32) Vec4f {
	return Vec4f{v[0] + s, v[1] + s, v[2] + s, v[3] + s}
}

// Sub performs component-wise subtraction.
func (v Vec4f) Sub(other Vec4f) Vec4f {
	return Vec4f{v[0] - other[0], v[1] - other[1], v[2] - other[2], v[3] - other[3]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec4f) SubScalar(s float32) Vec4f {
	return Vec4f{v[0] - s, v[1] - s, v[2] - s, v[3] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec4f) Mul(other Vec4f) Vec4f {
	return Vec4f{v[0] * other[0], v[1] * other[1], v[2] * other[2], v[3] * other[3]}
}

// MulScalar performs a scalar multiplication.
func (v Vec4f) MulScalar(s float32) Vec4f {
	return Vec4f{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// Div performs a component-wise division.
func (v Vec4f) Div(other Vec4f) Vec4f {
	return Vec4f{v[0] / other[0], v[1] / other[1], v[2] / other[2], v[3] / other[3]}
}

// DivScalar performs a scalar division.
func (v Vec4f) DivScalar(s float32) Vec4f {
	return Vec4f{v[0] / s, v[1] / s, v[2] / s, v[3] / s}
}

// Normalize the vector. Its length will be 1 afterwards.
// If the vector's length is zero, a zero vector will be returned.
func (v Vec4f) Normalize() Vec4f {
	length := v.Length()
	if Equalf(length, 0) {
		return Vec4f{}
	}
	return Vec4f{v[0] / length, v[1] / length, v[2] / length, v[3] / length}
}

// Length returns the vector's length.
func (v Vec4f) Length() float32 {
	return math32.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])
}

// SquareLength returns the vector's squared length.
func (v Vec4f) SquareLength() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// IsZero returns true if all components are zero.
// Uses the default Epsilon as relative tolerance.
func (v Vec4f) IsZero() bool {
	return v.EqualEps(Vec4f{}, Epsilon)
}

// Equal compares two vectors component-wise.
// Uses the default Epsilon as relative tolerance.
func (v Vec4f) Equal(other Vec4f) bool {
	return v.EqualEps(other, Epsilon)
}

// EqualEps compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec4f) EqualEps(other Vec4f, epsilon float32) bool {
	return EqualEps(v[0], other[0], epsilon) &&
		EqualEps(v[1], other[1], epsilon) &&
		EqualEps(v[2], other[2], epsilon) &&
		EqualEps(v[3], other[3], epsilon)
}

// Clamp clamps each component to the range of [min, max].
func (v Vec4f) Clamp(min, max float32) Vec4f {
	return Vec4f{
		Clampf(v[0], min, max),
		Clampf(v[1], min, max),
		Clampf(v[2], min, max),
		Clampf(v[3], min, max),
	}
}

// Negate inverts all components.
func (v Vec4f) Negate() Vec4f {
	return Vec4f{-v[0], -v[1], -v[2], -v[3]}
}

// Dot performs a dot product with another vector.
func (v Vec4f) Dot(other Vec4f) float32 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2] + v[3]*other[3]
}

// Project returns a vector representing the projection of vector v onto "other".
func (v Vec4f) Project(other Vec4f) Vec4f {
	return other.MulScalar(v.Dot(other) / other.SquareLength())
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec4f) Lerp(other Vec4f, t float32) Vec4f {
	return other.Sub(v).MulScalar(t).Add(v)
}

// Distance returns the euclidean distance to another position.
func (v Vec4f) Distance(other Vec4f) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec4f) SquareDistance(other Vec4f) float32 {
	return other.Sub(v).SquareLength()
}
