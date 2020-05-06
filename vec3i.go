package vmath

import (
	"fmt"
	"math"
)

type Vec3i [3]int

func (v Vec3i) String() string {
	return fmt.Sprintf("Vec3i[%d x %d x %d]", v[0], v[1], v[2])
}

// Format the vector to a string.
func (v Vec3i) Format(format string) string {
	return fmt.Sprintf(format, v[0], v[1], v[2])
}

// XVec3i returns a 3D vector representing the X-axis.
func XVec3i() Vec3i {
	return Vec3i{1, 0, 0}
}

// YVec3i returns a 3D vector representing the Y-axis.
func YVec3i() Vec3i {
	return Vec3i{0, 1, 0}
}

// ZVec3i returns a 3D vector representing the Z-axis.
func ZVec3i() Vec3i {
	return Vec3i{0, 0, 1}
}

// Abs returns a vector with the components turned into absolute values.
func (v Vec3i) Abs() Vec3i {
	return Vec3i{Absi(v[0]), Absi(v[1]), Absi(v[2])}
}

// Vec2f returns a float representation of the vector.
func (v Vec3i) Vec3f() Vec3f {
	return Vec3f{float32(v[0]), float32(v[1]), float32(v[2])}
}

// Vec4i creates a 4D vector.
func (v Vec3i) Vec4i(w int) Vec4i {
	return Vec4i{v[0], v[1], v[2], w}
}

// Split returns the vector's components.
func (v Vec3i) Split() (x, y, z int) {
	return v[0], v[1], v[2]
}

// X returns the vector's first component.
// Performance is equivalent to using v[0].
func (v Vec3i) X() int {
	return v[0]
}

// Y returns the vector's second component.
// Performance is equivalent to using v[1].
func (v Vec3i) Y() int {
	return v[1]
}

// Z returns the vector's third component.
// Performance is equivalent to using v[2].
func (v Vec3i) Z() int {
	return v[2]
}

// XY returns a 2D vector with the X and Y components.
func (v Vec3i) XY() Vec2i {
	return Vec2i{v[0], v[1]}
}

// Add performs component-wise addition between two vectors.
func (v Vec3i) Add(other Vec3i) Vec3i {
	return Vec3i{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec3i) AddScalar(s int) Vec3i {
	return Vec3i{v[0] + s, v[1] + s, v[2] + s}
}

// AddScalarf performs a scalar addition.
func (v Vec3i) AddScalarf(s float32) Vec3f {
	return Vec3f{float32(v[0]) + s, float32(v[1]) + s, float32(v[2]) + s}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec3i) Sub(other Vec3i) Vec3i {
	return Vec3i{v[0] - other[0], v[1] - other[1], v[2] - other[2]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec3i) SubScalar(s int) Vec3i {
	return Vec3i{v[0] - s, v[1] - s, v[2] - s}
}

// SubScalarf performs a scalar subtraction.
func (v Vec3i) SubScalarf(s float32) Vec3f {
	return Vec3f{float32(v[0]) - s, float32(v[1]) - s, float32(v[2]) - s}
}

// Mul performs a component-wise multiplication.
func (v Vec3i) Mul(other Vec3i) Vec3i {
	return Vec3i{v[0] * other[0], v[1] * other[1], v[2] * other[2]}
}

// MulScalar performs a scalar multiplication.
func (v Vec3i) MulScalar(s int) Vec3i {
	return Vec3i{v[0] * s, v[1] * s, v[2] * s}
}

// MulScalar performs a scalar multiplication.
func (v Vec3i) MulScalarf(s float32) Vec3f {
	return Vec3f{float32(v[0]) * s, float32(v[1]) * s, float32(v[2]) * s}
}

// Div performs a component-wise division.
func (v Vec3i) Div(other Vec3i) Vec3i {
	return Vec3i{v[0] / other[0], v[1] / other[1], v[2] / other[2]}
}

// DivScalar performs a scalar division.
func (v Vec3i) DivScalar(s int) Vec3i {
	return Vec3i{v[0] / s, v[1] / s, v[2] / s}
}

// DivScalarf performs a scalar division.
func (v Vec3i) DivScalarf(s float32) Vec3f {
	return Vec3f{float32(v[0]) / s, float32(v[1]) / s, float32(v[2]) / s}
}

// Length returns the vector's length.
func (v Vec3i) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
}

// SquareLength returns the vector's squared length.
func (v Vec3i) SquareLength() int {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// IsZero returns true if all components are zero.
func (v Vec3i) IsZero() bool {
	return v[0] == 0 && v[1] == 0 && v[2] == 0
}

// Equal compares two vectors component-wise.
func (v Vec3i) Equal(other Vec3i) bool {
	return v[0] == other[0] && v[1] == other[1] && v[2] == other[2]
}

// Clamp clamps each component to the range of [min, max].
func (v Vec3i) Clamp(min, max int) Vec3i {
	return Vec3i{
		Clampi(v[0], min, max),
		Clampi(v[1], min, max),
		Clampi(v[2], min, max),
	}
}

// Negate inverts all components.
func (v Vec3i) Negate() Vec3i {
	return Vec3i{-v[0], -v[1], -v[2]}
}

// Dot performs a dot product with another vector.
func (v Vec3i) Dot(other Vec3i) int {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}

// Cross performs a cross product with another vector.
func (v Vec3i) Cross(other Vec3i) Vec3i {
	return Vec3i{
		v[1]*other[2] - v[2]*other[1],
		v[2]*other[0] - v[0]*other[2],
		v[0]*other[1] - v[1]*other[0],
	}
}

// IsParallel returns true if the given vector is parallel.
// Vectors that point in opposite directions are also parallel (but not collinear).
func (v Vec3i) IsParallel(other Vec3i) bool {
	return v.Cross(other).SquareLength() == 0
}

// IsCollinear returns true if the given vector is collinear (pointing in the same direction).
// Uses the given Epsilon as relative tolerance.
func (v Vec3i) IsCollinear(other Vec3i) bool {
	return v.Vec3f().IsCollinear(other.Vec3f())
}

// IsCollinearEps returns true if the given vector is collinear (pointing in the same direction).
// Uses the given Epsilon as relative tolerance.
func (v Vec3i) IsCollinearEps(other Vec3i, eps float32) bool {
	return v.Vec3f().IsCollinearEps(other.Vec3f(), eps)
}

// Distance returns the euclidean distance to another position.
func (v Vec3i) Distance(other Vec3i) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec3i) SquareDistance(other Vec3i) int {
	return other.Sub(v).SquareLength()
}
