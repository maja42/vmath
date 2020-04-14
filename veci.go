package vmath

import (
	"fmt"
	"math"
)

type Vec2i [2]int
type Vec3i [3]int
type Vec4i [4]int

func (v Vec2i) String() string {
	return fmt.Sprintf("Vec2i[%d x %d]", v[0], v[1])
}

func (v Vec3i) String() string {
	return fmt.Sprintf("Vec3i[%d x %d x %d]", v[0], v[1], v[2])
}

func (v Vec4i) String() string {
	return fmt.Sprintf("Vec4i[%d x %d x %d x %d]", v[0], v[1], v[2], v[3])
}

func (v Vec2i) Vec2f() Vec2f {
	return Vec2f{float32(v[0]), float32(v[1])}
}

func (v Vec3i) Vec3f() Vec3f {
	return Vec3f{float32(v[0]), float32(v[1]), float32(v[2])}
}

func (v Vec4i) Vec4f() Vec4f {
	return Vec4f{float32(v[0]), float32(v[1]), float32(v[2]), float32(v[3])}
}

func (v Vec2i) Vec3i(z int) Vec3i {
	return Vec3i{v[0], v[1], z}
}

func (v Vec2i) Vec4i(z, w int) Vec4i {
	return Vec4i{v[0], v[1], z, w}
}

func (v Vec3i) Vec4i(w int) Vec4i {
	return Vec4i{v[0], v[1], v[2], w}
}

// Elem returns the vector's components.
func (v Vec2i) Elem() (x, y int) {
	return v[0], v[1]
}

// Elem returns the vector's components.
func (v Vec3i) Elem() (x, y, z int) {
	return v[0], v[1], v[2]
}

// Elem returns the vector's components.
func (v Vec4i) Elem() (x, y, z, w int) {
	return v[0], v[1], v[2], v[3]
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

// Add performs component-wise addition between two vectors.
func (v Vec2i) Add(other Vec2i) Vec2i {
	return Vec2i{v[0] + other[0], v[1] + other[1]}
}

// Add performs component-wise addition between two vectors.
func (v Vec3i) Add(other Vec3i) Vec3i {
	return Vec3i{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

// Add performs component-wise addition between two vectors.
func (v Vec4i) Add(other Vec4i) Vec4i {
	return Vec4i{v[0] + other[0], v[1] + other[1], v[2] + other[2], v[3] + other[3]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec2i) AddScalar(s int) Vec2i {
	return Vec2i{v[0] + s, v[1] + s}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec3i) AddScalar(s int) Vec3i {
	return Vec3i{v[0] + s, v[1] + s, v[2] + s}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec4i) AddScalar(s int) Vec4i {
	return Vec4i{v[0] + s, v[1] + s, v[2] + s, v[3] + s}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec2i) Sub(other Vec2i) Vec2i {
	return Vec2i{v[0] - other[0], v[1] - other[1]}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec3i) Sub(other Vec3i) Vec3i {
	return Vec3i{v[0] - other[0], v[1] - other[1], v[2] - other[2]}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec4i) Sub(other Vec4i) Vec4i {
	return Vec4i{v[0] - other[0], v[1] - other[1], v[2] - other[2], v[3] - other[3]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec2i) SubScalar(s int) Vec2i {
	return Vec2i{v[0] - s, v[1] - s}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec3i) SubScalar(s int) Vec3i {
	return Vec3i{v[0] - s, v[1] - s, v[2] - s}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec4i) SubScalar(s int) Vec4i {
	return Vec4i{v[0] - s, v[1] - s, v[2] - s, v[3] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec2i) Mul(other Vec2i) Vec2i {
	return Vec2i{v[0] * other[0], v[1] * other[1]}
}

// Mul performs a component-wise multiplication.
func (v Vec3i) Mul(other Vec3i) Vec3i {
	return Vec3i{v[0] * other[0], v[1] * other[1], v[2] * other[2]}
}

// Mul performs a component-wise multiplication.
func (v Vec4i) Mul(other Vec4i) Vec4i {
	return Vec4i{v[0] * other[0], v[1] * other[1], v[2] * other[2], v[3] * other[3]}
}

// MulScalar performs a scalar multiplication.
func (v Vec2i) MulScalar(s int) Vec2i {
	return Vec2i{v[0] * s, v[1] * s}
}

// MulScalar performs a scalar multiplication.
func (v Vec3i) MulScalar(s int) Vec3i {
	return Vec3i{v[0] * s, v[1] * s, v[2] * s}
}

// MulScalar performs a scalar multiplication.
func (v Vec4i) MulScalar(s int) Vec4i {
	return Vec4i{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// Div performs a component-wise division.
func (v Vec2i) Div(other Vec2i) Vec2i {
	return Vec2i{v[0] / other[0], v[1] / other[1]}
}

// Div performs a component-wise division.
func (v Vec3i) Div(other Vec3i) Vec3i {
	return Vec3i{v[0] / other[0], v[1] / other[1], v[2] / other[2]}
}

// Div performs a component-wise division.
func (v Vec4i) Div(other Vec4i) Vec4i {
	return Vec4i{v[0] / other[0], v[1] / other[1], v[2] / other[2], v[3] / other[3]}
}

// DivScalar performs a scalar division.
func (v Vec2i) DivScalar(s int) Vec2i {
	return Vec2i{v[0] / s, v[1] / s}
}

// DivScalar performs a scalar division.
func (v Vec3i) DivScalar(s int) Vec3i {
	return Vec3i{v[0] / s, v[1] / s, v[2] / s}
}

// DivScalar performs a scalar division.
func (v Vec4i) DivScalar(s int) Vec4i {
	return Vec4i{v[0] / s, v[1] / s, v[2] / s, v[3] / s}
}

// Len returns the vector's length.
func (v Vec2i) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
}

// Len returns the vector's length.
func (v Vec3i) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
}

// Len returns the vector's length.
func (v Vec4i) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
}

// SquareLen returns the vector's squared length.
func (v Vec2i) SquareLen() int {
	return v[0]*v[0] + v[1]*v[1]
}

// SquareLen returns the vector's squared length.
func (v Vec3i) SquareLen() int {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// SquareLen returns the vector's squared length.
func (v Vec4i) SquareLen() int {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// Equal compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec2i) Equal(other Vec2i) bool {
	return v[0] == other[0] && v[1] == other[1]
}

// Equal compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec3i) Equal(other Vec2i) bool {
	return v[0] == other[0] && v[1] == other[1] && v[2] == other[2]
}

// Equal compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec4i) Equal(other Vec2i) bool {
	return v[0] == other[0] && v[1] == other[1] && v[2] == other[2] && v[3] == other[3]
}

// Clamp clamps each component to the range of [min, max].
func (v Vec2i) Clamp(min, max int) Vec2i {
	return Vec2i{
		Clampi(v[0], min, max),
		Clampi(v[1], min, max),
	}
}

// Clamp clamps each component to the range of [min, max].
func (v Vec3i) Clamp(min, max int) Vec3i {
	return Vec3i{
		Clampi(v[0], min, max),
		Clampi(v[1], min, max),
		Clampi(v[2], min, max),
	}
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

// Invert inverts (negates) all components.
func (v Vec2i) Invert() Vec2i {
	return Vec2i{-v[0], -v[1]}
}

// Invert inverts (negates) all components.
func (v Vec3i) Invert() Vec3i {
	return Vec3i{-v[0], -v[1], -v[2]}
}

// Invert inverts (negates) all components.
func (v Vec4i) Invert() Vec4i {
	return Vec4i{-v[0], -v[1], -v[2], -v[3]}
}

// Dot performs a dot product with another vector.
func (v Vec2i) Dot(other Vec2i) int {
	return v[0]*other[0] + v[1]*other[1]
}

// Dot performs a dot product with another vector.
func (v Vec3i) Dot(other Vec3i) int {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}

// Dot performs a dot product with another vector.
func (v Vec4i) Dot(other Vec4i) int {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2] + v[3]*other[3]
}
