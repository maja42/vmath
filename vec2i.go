package vmath

import (
	"fmt"
)

type Vec2i [2]int

func (v Vec2i) String() string {
	return fmt.Sprintf("Vec2i[%d x %d]", v[0], v[1])
}

// Vec2f returns a float representation of the vector.
func (v Vec2i) Vec2f() Vec2f {
	return Vec2f{float32(v[0]), float32(v[1])}
}

// Vec3i creates a 3D vector.
func (v Vec2i) Vec3i(z int) Vec3i {
	return Vec3i{v[0], v[1], z}
}

// Vec4i creates a 4D vector.
func (v Vec2i) Vec4i(z, w int) Vec4i {
	return Vec4i{v[0], v[1], z, w}
}

// Split returns the vector's components.
func (v Vec2i) Split() (x, y int) {
	return v[0], v[1]
}

// X returns the vector's first component.
// Performance is equivalent to using v[0].
func (v Vec2i) X() int {
	return v[0]
}

// Y returns the vector's second component.
// Performance is equivalent to using v[1].
func (v Vec2i) Y() int {
	return v[1]
}

// Add performs component-wise addition between two vectors.
func (v Vec2i) Add(other Vec2i) Vec2i {
	return Vec2i{v[0] + other[0], v[1] + other[1]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec2i) AddScalar(s int) Vec2i {
	return Vec2i{v[0] + s, v[1] + s}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec2i) Sub(other Vec2i) Vec2i {
	return Vec2i{v[0] - other[0], v[1] - other[1]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec2i) SubScalar(s int) Vec2i {
	return Vec2i{v[0] - s, v[1] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec2i) Mul(other Vec2i) Vec2i {
	return Vec2i{v[0] * other[0], v[1] * other[1]}
}

// MulScalar performs a scalar multiplication.
func (v Vec2i) MulScalar(s int) Vec2i {
	return Vec2i{v[0] * s, v[1] * s}
}

// Div performs a component-wise division.
func (v Vec2i) Div(other Vec2i) Vec2i {
	return Vec2i{v[0] / other[0], v[1] / other[1]}
}

// DivScalar performs a scalar division.
func (v Vec2i) DivScalar(s int) Vec2i {
	return Vec2i{v[0] / s, v[1] / s}
}

// Length returns the vector's length.
func (v Vec2i) Length() float32 {
	return Hypot(float32(v[0]*v[0]), float32(v[1]*v[1]))
}

// SquareLength returns the vector's squared length.
func (v Vec2i) SquareLength() int {
	return v[0]*v[0] + v[1]*v[1]
}

// Equal compares two vectors component-wise.
func (v Vec2i) Equal(other Vec2i) bool {
	return v[0] == other[0] && v[1] == other[1]
}

// Clamp clamps each component to the range of [min, max].
func (v Vec2i) Clamp(min, max int) Vec2i {
	return Vec2i{
		Clampi(v[0], min, max),
		Clampi(v[1], min, max),
	}
}

// Negate inverts all components.
func (v Vec2i) Negate() Vec2i {
	return Vec2i{-v[0], -v[1]}
}

// Dot performs a dot product with another vector.
func (v Vec2i) Dot(other Vec2i) int {
	return v[0]*other[0] + v[1]*other[1]
}

// Distance returns the euclidean distance to another position.
func (v Vec2i) Distance(other Vec2i) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec2i) SquareDistance(other Vec2i) int {
	return other.Sub(v).SquareLength()
}
