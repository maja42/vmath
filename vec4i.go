package vmath

import (
	"fmt"
	"math"
)

type Vec4i [4]int

func (v Vec4i) String() string {
	return fmt.Sprintf("Vec4i[%d x %d x %d x %d]", v[0], v[1], v[2], v[3])
}

// Vec2f returns a float representation of the vector.
func (v Vec4i) Vec4f() Vec4f {
	return Vec4f{float32(v[0]), float32(v[1]), float32(v[2]), float32(v[3])}
}

// Split returns the vector's components.
func (v Vec4i) Split() (x, y, z, w int) {
	return v[0], v[1], v[2], v[3]
}

// X returns the vector's first component.
// Performance is equivalent to using v[0].
func (v Vec4i) X() int {
	return v[0]
}

// Y returns the vector's second component.
// Performance is equivalent to using v[1].
func (v Vec4i) Y() int {
	return v[1]
}

// Z returns the vector's third component.
// Performance is equivalent to using v[2].
func (v Vec4i) Z() int {
	return v[2]
}

// W returns the vector's fourth component.
// Performance is equivalent to using v[3].
func (v Vec4i) W() int {
	return v[3]
}

// XY returns a 2D vector with the X and Y components.
func (v Vec4i) XY() Vec2i {
	return Vec2i{v[0], v[1]}
}

// XYZ returns a 3D vector with the X, Y and Z components.
func (v Vec4i) XYZ() Vec3i {
	return Vec3i{v[0], v[1], v[2]}
}

// Add performs component-wise addition between two vectors.
func (v Vec4i) Add(other Vec4i) Vec4i {
	return Vec4i{v[0] + other[0], v[1] + other[1], v[2] + other[2], v[3] + other[3]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec4i) AddScalar(s int) Vec4i {
	return Vec4i{v[0] + s, v[1] + s, v[2] + s, v[3] + s}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec4i) Sub(other Vec4i) Vec4i {
	return Vec4i{v[0] - other[0], v[1] - other[1], v[2] - other[2], v[3] - other[3]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec4i) SubScalar(s int) Vec4i {
	return Vec4i{v[0] - s, v[1] - s, v[2] - s, v[3] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec4i) Mul(other Vec4i) Vec4i {
	return Vec4i{v[0] * other[0], v[1] * other[1], v[2] * other[2], v[3] * other[3]}
}

// MulScalar performs a scalar multiplication.
func (v Vec4i) MulScalar(s int) Vec4i {
	return Vec4i{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// Div performs a component-wise division.
func (v Vec4i) Div(other Vec4i) Vec4i {
	return Vec4i{v[0] / other[0], v[1] / other[1], v[2] / other[2], v[3] / other[3]}
}

// DivScalar performs a scalar division.
func (v Vec4i) DivScalar(s int) Vec4i {
	return Vec4i{v[0] / s, v[1] / s, v[2] / s, v[3] / s}
}

// Length returns the vector's length.
func (v Vec4i) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
}

// SquareLength returns the vector's squared length.
func (v Vec4i) SquareLength() int {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// Equal compares two vectors component-wise.
func (v Vec4i) Equal(other Vec4i) bool {
	return v[0] == other[0] && v[1] == other[1] && v[2] == other[2] && v[3] == other[3]
}

// Clamp clamps each component to the range of [min, max].
func (v Vec4i) Clamp(min, max int) Vec4i {
	return Vec4i{
		Clampi(v[0], min, max),
		Clampi(v[1], min, max),
		Clampi(v[2], min, max),
		Clampi(v[3], min, max),
	}
}

// Negate inverts all components.
func (v Vec4i) Negate() Vec4i {
	return Vec4i{-v[0], -v[1], -v[2], -v[3]}
}

// Dot performs a dot product with another vector.
func (v Vec4i) Dot(other Vec4i) int {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2] + v[3]*other[3]
}

// Distance returns the euclidean distance to another position.
func (v Vec4i) Distance(other Vec4i) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec4i) SquareDistance(other Vec4i) int {
	return other.Sub(v).SquareLength()
}
