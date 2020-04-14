package vmath

import (
	"fmt"
	"math"
)

type Vec2f [2]float32
type Vec3f [3]float32
type Vec4f [4]float32

func (v Vec2f) String() string {
	return fmt.Sprintf("Vec2f[%f x %f]", v[0], v[1])
}

func (v Vec3f) String() string {
	return fmt.Sprintf("Vec3f[%f x %f x %f]", v[0], v[1], v[2])
}

func (v Vec4f) String() string {
	return fmt.Sprintf("Vec4f[%f x %f x %f x %f]", v[0], v[1], v[2], v[3])
}

func (v Vec2f) Vec2i() Vec2i {
	return Vec2i{int(v[0]), int(v[1])}
}

func (v Vec3f) Vec3i() Vec3i {
	return Vec3i{int(v[0]), int(v[1]), int(v[2])}
}

func (v Vec4f) Vec4fi() Vec4i {
	return Vec4i{int(v[0]), int(v[1]), int(v[2]), int(v[3])}
}

func (v Vec2f) Vec3f(z float32) Vec3f {
	return Vec3f{v[0], v[1], z}
}

func (v Vec2f) Vec4f(z, w float32) Vec4f {
	return Vec4f{v[0], v[1], z, w}
}

func (v Vec3f) Vec4f(w float32) Vec4f {
	return Vec4f{v[0], v[1], v[2], w}
}

// Elem returns the vector's components.
func (v Vec2f) Elem() (x, y float32) {
	return v[0], v[1]
}

// Elem returns the vector's components.
func (v Vec3f) Elem() (x, y, z float32) {
	return v[0], v[1], v[2]
}

// Elem returns the vector's components.
func (v Vec4f) Elem() (x, y, z, w float32) {
	return v[0], v[1], v[2], v[3]
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

// Add performs component-wise addition between two vectors.
func (v Vec2f) Add(other Vec2f) Vec2f {
	return Vec2f{v[0] + other[0], v[1] + other[1]}
}

// Add performs component-wise addition between two vectors.
func (v Vec3f) Add(other Vec3f) Vec3f {
	return Vec3f{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

// Add performs component-wise addition between two vectors.
func (v Vec4f) Add(other Vec4f) Vec4f {
	return Vec4f{v[0] + other[0], v[1] + other[1], v[2] + other[2], v[3] + other[3]}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec2f) AddScalar(s float32) Vec2f {
	return Vec2f{v[0] + s, v[1] + s}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec3f) AddScalar(s float32) Vec3f {
	return Vec3f{v[0] + s, v[1] + s, v[2] + s}
}

// AddScalar performs a component-wise scalar addition.
func (v Vec4f) AddScalar(s float32) Vec4f {
	return Vec4f{v[0] + s, v[1] + s, v[2] + s, v[3] + s}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec2f) Sub(other Vec2f) Vec2f {
	return Vec2f{v[0] - other[0], v[1] - other[1]}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec3f) Sub(other Vec3f) Vec3f {
	return Vec3f{v[0] - other[0], v[1] - other[1], v[2] - other[2]}
}

// Sub performs component-wise subtraction between two vectors.
func (v Vec4f) Sub(other Vec4f) Vec4f {
	return Vec4f{v[0] - other[0], v[1] - other[1], v[2] - other[2], v[3] - other[3]}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec2f) SubScalar(s float32) Vec2f {
	return Vec2f{v[0] - s, v[1] - s}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec3f) SubScalar(s float32) Vec3f {
	return Vec3f{v[0] - s, v[1] - s, v[2] - s}
}

// SubScalar performs a component-wise scalar subtraction.
func (v Vec4f) SubScalar(s float32) Vec4f {
	return Vec4f{v[0] - s, v[1] - s, v[2] - s, v[3] - s}
}

// Mul performs a component-wise multiplication.
func (v Vec2f) Mul(other Vec2f) Vec2f {
	return Vec2f{v[0] * other[0], v[1] * other[1]}
}

// Mul performs a component-wise multiplication.
func (v Vec3f) Mul(other Vec3f) Vec3f {
	return Vec3f{v[0] * other[0], v[1] * other[1], v[2] * other[2]}
}

// Mul performs a component-wise multiplication.
func (v Vec4f) Mul(other Vec4f) Vec4f {
	return Vec4f{v[0] * other[0], v[1] * other[1], v[2] * other[2], v[3] * other[3]}
}

// MulScalar performs a scalar multiplication.
func (v Vec2f) MulScalar(s float32) Vec2f {
	return Vec2f{v[0] * s, v[1] * s}
}

// MulScalar performs a scalar multiplication.
func (v Vec3f) MulScalar(s float32) Vec3f {
	return Vec3f{v[0] * s, v[1] * s, v[2] * s}
}

// MulScalar performs a scalar multiplication.
func (v Vec4f) MulScalar(s float32) Vec4f {
	return Vec4f{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// Div performs a component-wise division.
func (v Vec2f) Div(other Vec2f) Vec2f {
	return Vec2f{v[0] / other[0], v[1] / other[1]}
}

// Div performs a component-wise division.
func (v Vec3f) Div(other Vec3f) Vec3f {
	return Vec3f{v[0] / other[0], v[1] / other[1], v[2] / other[2]}
}

// Div performs a component-wise division.
func (v Vec4f) Div(other Vec4f) Vec4f {
	return Vec4f{v[0] / other[0], v[1] / other[1], v[2] / other[2], v[3] / other[3]}
}

// DivScalar performs a scalar division.
func (v Vec2f) DivScalar(s float32) Vec2f {
	return Vec2f{v[0] / s, v[1] / s}
}

// DivScalar performs a scalar division.
func (v Vec3f) DivScalar(s float32) Vec3f {
	return Vec3f{v[0] / s, v[1] / s, v[2] / s}
}

// DivScalar performs a scalar division.
func (v Vec4f) DivScalar(s float32) Vec4f {
	return Vec4f{v[0] / s, v[1] / s, v[2] / s, v[3] / s}
}

// Normalize the vector. Its length will be 1 afterwards.
// If the vector is zero, all components will be infinite afterwards.
func (v Vec2f) Normalize() Vec2f {
	l := 1.0 / v.Len()
	return Vec2f{v[0] * l, v[1] * l}
}

// Normalize the vector. Its length will be 1 afterwards.
// If the vector is zero, all components will be infinite afterwards.
func (v Vec3f) Normalize() Vec3f {
	l := 1.0 / v.Len()
	return Vec3f{v[0] * l, v[1] * l, v[2] * l}
}

// Normalize the vector. Its length will be 1 afterwards.
// If the vector is zero, all components will be infinite afterwards.
func (v Vec4f) Normalize() Vec4f {
	l := 1.0 / v.Len()
	return Vec4f{v[0] * l, v[1] * l, v[2] * l, v[3] * l}
}

// Len returns the vector's length.
func (v Vec2f) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
}

// Len returns the vector's length.
func (v Vec3f) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
}

// Len returns the vector's length.
func (v Vec4f) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
}

// SquareLen returns the vector's squared length.
func (v Vec2f) SquareLen() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

// SquareLen returns the vector's squared length.
func (v Vec3f) SquareLen() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// SquareLen returns the vector's squared length.
func (v Vec4f) SquareLen() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// Equal compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec2f) Equal(other Vec2f, epsilon float32) bool {
	return Equalf(v[0], other[0], epsilon) && Equalf(v[1], other[1], epsilon)
}

// Equal compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec3f) Equal(other Vec2f, epsilon float32) bool {
	return Equalf(v[0], other[0], epsilon) &&
		Equalf(v[1], other[1], epsilon) &&
		Equalf(v[2], other[2], epsilon)
}

// Equal compares two vectors component-wise, using the given epsilon as a relative tolerance.
func (v Vec4f) Equal(other Vec2f, epsilon float32) bool {
	return Equalf(v[0], other[0], epsilon) &&
		Equalf(v[1], other[1], epsilon) &&
		Equalf(v[2], other[2], epsilon) &&
		Equalf(v[3], other[3], epsilon)
}

// Clamp clamps each component to the range of [min, max].
func (v Vec2f) Clamp(min, max float32) Vec2f {
	return Vec2f{
		Clampf(v[0], min, max),
		Clampf(v[1], min, max),
	}
}

// Clamp clamps each component to the range of [min, max].
func (v Vec3f) Clamp(min, max float32) Vec3f {
	return Vec3f{
		Clampf(v[0], min, max),
		Clampf(v[1], min, max),
		Clampf(v[2], min, max),
	}
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

// Invert inverts (negates) all components.
func (v Vec2f) Invert() Vec2f {
	return Vec2f{-v[0], -v[1]}
}

// Invert inverts (negates) all components.
func (v Vec3f) Invert() Vec3f {
	return Vec3f{-v[0], -v[1], -v[2]}
}

// Invert inverts (negates) all components.
func (v Vec4f) Invert() Vec4f {
	return Vec4f{-v[0], -v[1], -v[2], -v[3]}
}

// Dot performs a dot product with another vector.
func (v Vec2f) Dot(other Vec2f) float32 {
	return v[0]*other[0] + v[1]*other[1]
}

// Dot performs a dot product with another vector.
func (v Vec3f) Dot(other Vec3f) float32 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}

// Dot performs a dot product with another vector.
func (v Vec4f) Dot(other Vec4f) float32 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2] + v[3]*other[3]
}

// Cross performs a cross product with another vector.
func (v Vec3f) Cross(other Vec3f) Vec3f {
	return Vec3f{
		v[1]*other[2] - v[2]*other[1],
		v[2]*other[0] - v[0]*other[2],
		v[0]*other[1] - v[1]*other[0],
	}
}

// Project returns a vector representing the projection of other onto v.
func (v Vec2f) Project(other Vec2f) Vec2f {
	return v.MulScalar(other.Dot(v) / v.SquareLen())
}

// Project returns a vector representing the projection of other onto v.
func (v Vec3f) Project(other Vec3f) Vec3f {
	return v.MulScalar(other.Dot(v) / v.SquareLen())
}

// Project returns a vector representing the projection of other onto v.
func (v Vec4f) Project(other Vec4f) Vec4f {
	return v.MulScalar(other.Dot(v) / v.SquareLen())
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec2f) Lerp(other Vec2f, t float32) Vec2f {
	return v.Mul(other.MulScalar(t))
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec3f) Lerp(other Vec3f, t float32) Vec3f {
	return v.Mul(other.MulScalar(t))
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec4f) Lerp(other Vec4f, t float32) Vec4f {
	return v.Mul(other.MulScalar(t))
}

// Angle returns the angle between two vectors in radians.
func (v Vec2f) Angle(other Vec2f) float32 {
	return float32(math.Atan2(float64(other[1]-v[1]), float64(other[0]-v[0])))
}

// Angle returns the angle between two vectors in radians.
func (v Vec3f) Angle(other Vec3f) float64 {
	v = v.Normalize()
	other = other.Normalize()

	if v.Dot(other) < 0 {
		n := v.Add(other).Len() / 2
		return math.Pi - 2.0*math.Asin(float64(Minf(n, 1)))
	}
	n := v.Sub(other).Len() / 2
	return 2.0 * math.Asin(float64(Minf(n, 1)))
}
