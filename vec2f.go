package vmath

import (
	"fmt"

	"github.com/maja42/vmath/math32"
)

type Vec2f [2]float32

func (v Vec2f) String() string {
	return fmt.Sprintf("Vec2f[%f x %f]", v[0], v[1])
}

// Format the vector to a string.
func (v Vec2f) Format(format string) string {
	return fmt.Sprintf(format, v[0], v[1])
}

// Vec2i returns an integer representation of the vector.
// Decimals are truncated.
func (v Vec2f) Vec2i() Vec2i {
	return Vec2i{int(v[0]), int(v[1])}
}

// Round returns an integer representation of the vector.
// Decimals are rounded.
func (v Vec2f) Round() Vec2i {
	return Vec2i{int(math32.Round(v[0])), int(math32.Round(v[1]))}
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

// IsOrthogonal returns true if the vector is horizontal or vertical (one of its components is zero).
func (v Vec2f) IsOrthogonal() bool {
	return v[0] == 0 || v[1] == 0
}

// Abs returns a vector with the components turned into absolute values.
func (v Vec2f) Abs() Vec2f {
	return Vec2f{math32.Abs(v[0]), math32.Abs(v[1])}
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
	return math32.Hypot(v[0], v[1])
}

// SquareLength returns the vector's squared length.
func (v Vec2f) SquareLength() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

// IsZero returns true if all components are zero.
// Uses the default Epsilon as relative tolerance.
func (v Vec2f) IsZero() bool {
	return v.EqualEps(Vec2f{}, Epsilon)
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

// MagCross returns the length of the cross product vector.
// This is equal to the magnitude of a 3D cross product vector, with the Z position implicitly set to zero.
// It represents twice the signed area between the two vectors.
func (v Vec2f) MagCross(other Vec2f) float32 {
	return v[0]*other[1] - v[1]*other[0]
}

// IsParallel returns true if the given vector is parallel.
// Vectors that point in opposite directions are also parallel.
// Uses the default Epsilon as relative tolerance.
func (v Vec2f) IsParallel(other Vec2f) bool {
	return Equalf(v[0]*other[1], v[1]*other[0])
}

// IsParallel returns true if the given vector is parallel.
// Vectors that point in opposite directions are also parallel.
// Uses the given Epsilon as relative tolerance.
func (v Vec2f) IsParallelEps(other Vec2f, eps float32) bool {
	return EqualEps(v[0]*other[1], v[1]*other[0], eps)
}

// IsCollinear returns true if the given vector is collinear (pointing in the same direction).
// Uses the given Epsilon as relative tolerance.
func (v Vec2f) IsCollinear(other Vec2f) bool {
	return v.IsCollinearEps(other, Epsilon)
}

// IsCollinearEps returns true if the given vector is collinear (pointing in the same direction).
// Uses the given Epsilon as relative tolerance.
func (v Vec2f) IsCollinearEps(other Vec2f, eps float32) bool {
	// Note: Vectors that are nearly zero will not be reported as collinear if they are facing
	// in different directions, even if their size falls within epsilon.
	return EqualEps(v[0]*other[1], v[1]*other[0], eps) && // parallel
		math32.Signbit(v[0]) == math32.Signbit(other[0]) && // same x direction
		math32.Signbit(v[1]) == math32.Signbit(other[1]) // same y direction
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
	return other.Sub(v).MulScalar(t).Add(v)
}

// Angle returns the angle relative to another vector.
func (v Vec2f) Angle(other Vec2f) float32 {
	return math32.Atan2(other[1], other[0]) - math32.Atan2(v[1], v[0])
}

// FlatAngle returns the angle of a vector in radians.
// This is the angle between the vector and the x-axis.
func (v Vec2f) FlatAngle() float32 {
	return math32.Atan2(v[1], v[0])
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
