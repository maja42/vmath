package vmath

import (
	"fmt"
)

type Vec2i [2]int

func (v Vec2i) String() string {
	return fmt.Sprintf("Vec2i[%d x %d]", v[0], v[1])
}

// Format the vector to a string.
func (v Vec2i) Format(format string) string {
	return fmt.Sprintf(format, v[0], v[1])
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

// IsOrthogonal returns true if the vector is horizontal or vertical (one of its components is zero).
func (v Vec2i) IsOrthogonal() bool {
	return v[0] == 0 || v[1] == 0
}

// Abs returns a vector with the components turned into absolute values.
func (v Vec2i) Abs() Vec2i {
	return Vec2i{Absi(v[0]), Absi(v[1])}
}

// Add performs component-wise addition between two vectors.
func (v Vec2i) Add(other Vec2i) Vec2i {
	return Vec2i{v[0] + other[0], v[1] + other[1]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec2i) AddScalar(s int) Vec2i {
	return Vec2i{v[0] + s, v[1] + s}
}

// AddScalarf performs a scalar addition.
func (v Vec2i) AddScalarf(s float32) Vec2f {
	return Vec2f{float32(v[0]) + s, float32(v[1]) + s}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec2i) Sub(other Vec2i) Vec2i {
	return Vec2i{v[0] - other[0], v[1] - other[1]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec2i) SubScalar(s int) Vec2i {
	return Vec2i{v[0] - s, v[1] - s}
}

// SubScalarf performs a scalar subtraction.
func (v Vec2i) SubScalarf(s float32) Vec2f {
	return Vec2f{float32(v[0]) - s, float32(v[1]) - s}
}

// Mul performs a component-wise multiplication.
func (v Vec2i) Mul(other Vec2i) Vec2i {
	return Vec2i{v[0] * other[0], v[1] * other[1]}
}

// MulScalar performs a scalar multiplication.
func (v Vec2i) MulScalar(s int) Vec2i {
	return Vec2i{v[0] * s, v[1] * s}
}

// MulScalar performs a scalar multiplication.
func (v Vec2i) MulScalarf(s float32) Vec2f {
	return Vec2f{float32(v[0]) * s, float32(v[1]) * s}
}

// Div performs a component-wise division.
// Decimals are truncated.
func (v Vec2i) Div(other Vec2i) Vec2i {
	return Vec2i{v[0] / other[0], v[1] / other[1]}
}

// DivScalar performs a scalar division.
// Decimals are truncated.
func (v Vec2i) DivScalar(s int) Vec2i {
	return Vec2i{v[0] / s, v[1] / s}
}

// DivScalarf performs a scalar division.
func (v Vec2i) DivScalarf(s float32) Vec2f {
	return Vec2f{float32(v[0]) / s, float32(v[1]) / s}
}

// Length returns the vector's length.
func (v Vec2i) Length() float32 {
	return Hypot(float32(v[0]), float32(v[1]))
}

// SquareLength returns the vector's squared length.
func (v Vec2i) SquareLength() int {
	return v[0]*v[0] + v[1]*v[1]
}

// IsZero returns true if all components are zero.
func (v Vec2i) IsZero() bool {
	return v[0] == 0 && v[1] == 0
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

// MagCross returns the length of the cross product vector.
// This is equal to the magnitude of a 3D cross product vector, with the Z position implicitly set to zero.
// It represents twice the signed area between the two vectors.
func (v Vec2i) MagCross(other Vec2i) int {
	return v[0]*other[1] - v[1]*other[0]
}

// IsParallel returns true if the given vector is parallel.
// Vectors that point in opposite directions are also parallel (but not collinear).
func (v Vec2i) IsParallel(other Vec2i) bool {
	return v.MagCross(other) == 0
}

// IsCollinear returns true if the given vector is collinear (pointing in the same direction).
func (v Vec2i) IsCollinear(other Vec2i) bool {
	return v.MagCross(other) == 0 && // parallel
		(v[0] >= 0) == (other[0] >= 0) && // same x direction
		(v[1] >= 0) == (other[1] >= 0) // same y direction
}

// NormalVec returns a normal vector on the 2D plane that is either on the left or right hand side.
func (v Vec2i) NormalVec(onLeft bool) Vec2i {
	if onLeft {
		return Vec2i{-v[1], v[0]}
	}
	return Vec2i{v[1], -v[0]}
}

// Project returns a vector representing the projection of vector v onto "other".
func (v Vec2i) Project(other Vec2i) Vec2f {
	return v.Vec2f().Project(other.Vec2f())
}

// Distance returns the euclidean distance to another position.
func (v Vec2i) Distance(other Vec2i) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec2i) SquareDistance(other Vec2i) int {
	return other.Sub(v).SquareLength()
}
