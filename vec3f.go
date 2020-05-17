package vmath

import (
	"fmt"
	"math"
)

type Vec3f [3]float32

func (v Vec3f) String() string {
	return fmt.Sprintf("Vec3f[%f x %f x %f]", v[0], v[1], v[2])
}

// Format the vector to a string.
func (v Vec3f) Format(format string) string {
	return fmt.Sprintf(format, v[0], v[1], v[2])
}

// Vec3i returns an integer representation of the vector.
// Decimals are truncated.
func (v Vec3f) Vec3i() Vec3i {
	return Vec3i{int(v[0]), int(v[1]), int(v[2])}
}

// Round returns an integer representation of the vector.
// Decimals are rounded.
func (v Vec3f) Round() Vec3i {
	return Vec3i{int(Round(v[0])), int(Round(v[1])), int(Round(v[2]))}
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

// IsOrthogonal returns true if the vector is parallel to the X, Y or Z axis (one of its components is zero).
func (v Vec3f) IsOrthogonal() bool {
	return v[0] == 0 || v[1] == 0 || v[2] == 0
}

// Abs returns a vector with the components turned into absolute values.
func (v Vec3f) Abs() Vec3f {
	return Vec3f{Abs(v[0]), Abs(v[1]), Abs(v[2])}
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
// If the vector's length is zero, a zero vector will be returned.
func (v Vec3f) Normalize() Vec3f {
	length := v.Length()
	if Equalf(length, 0) {
		return Vec3f{}
	}
	return Vec3f{v[0] / length, v[1] / length, v[2] / length}
}

// Length returns the vector's length.
func (v Vec3f) Length() float32 {
	return Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

// SquareLength returns the vector's squared length.
func (v Vec3f) SquareLength() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// IsZero returns true if all components are zero.
// Uses the default Epsilon as relative tolerance.
func (v Vec3f) IsZero() bool {
	return v.EqualEps(Vec3f{}, Epsilon)
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

// Negate inverts all components.
func (v Vec3f) Negate() Vec3f {
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

// IsParallel returns true if the given vector is parallel.
// Vectors that point in opposite directions are also parallel.
// Uses the default Epsilon as relative tolerance.
func (v Vec3f) IsParallel(other Vec3f) bool {
	return Equalf(v.Cross(other).SquareLength(), 0)
}

// IsParallel returns true if the given vector is parallel.
// Vectors that point in opposite directions are also parallel.
// Uses the given Epsilon as relative tolerance.
func (v Vec3f) IsParallelEps(other Vec3f, eps float32) bool {
	return EqualEps(v.Cross(other).SquareLength(), 0, eps)
}

// IsCollinear returns true if the given vector is collinear (pointing in the same direction).
// Uses the given Epsilon as relative tolerance.
func (v Vec3f) IsCollinear(other Vec3f) bool {
	return v.IsCollinearEps(other, Epsilon)
}

// IsCollinearEps returns true if the given vector is collinear (pointing in the same direction).
// Uses the given Epsilon as relative tolerance.
func (v Vec3f) IsCollinearEps(other Vec3f, eps float32) bool {
	// Note: Vectors that are nearly zero will not be reported as collinear if they are facing
	// in different directions, even if their size falls within epsilon.
	return v.IsParallelEps(other, eps) &&
		Signbit(v[0]) == Signbit(other[0]) && // same x direction
		Signbit(v[1]) == Signbit(other[1]) && // same y direction
		Signbit(v[2]) == Signbit(other[2]) // same y direction
}

// Project returns a vector representing the projection of vector v onto "other".
func (v Vec3f) Project(other Vec3f) Vec3f {
	return other.MulScalar(v.Dot(other) / other.SquareLength())
}

// Lerp performs a linear interpolation between two vectors.
// The parameter t should be in range [0, 1].
func (v Vec3f) Lerp(other Vec3f, t float32) Vec3f {
	return v.MulScalar(1 - t).Add(other.MulScalar(t))
}

// Angle returns the angle between two vectors in radians.
func (v Vec3f) Angle(other Vec3f) float32 {
	v = v.Normalize()
	other = other.Normalize()
	return Acos(v.Dot(other))
}

// RotationTo returns the shortest rotation to the destination vector.
func (v Vec3f) RotationTo(dest Vec3f) Quat {
	// Source: http://glmatrix.net/docs/module-quat.html

	v = v.Normalize()
	dest = dest.Normalize()
	dot := v.Dot(dest)

	if dot < -1+Epsilon {
		t := Vec3f{1, 0, 0}.Cross(v)
		if t.Length() < Epsilon {
			t = Vec3f{0, 1, 0}.Cross(v)
		}
		return QuatFromAxisAngle(t.Normalize(), math.Pi)
	}
	if dot > 1-Epsilon {
		return Quat{1, 0, 0, 0}
	}
	t := v.Cross(dest)
	return Quat{1 + dot, t[0], t[1], t[2]}.Normalize()
}

// RotateX rotates a point around the X-axis.
func (v Vec3f) RotateX(origin Vec3f, rad float32) Vec3f {
	v = v.Sub(origin) // translate to origin

	sin, cos := Sincos(rad)
	p := Vec3f{
		v[0],
		v[1]*cos - v[2]*sin,
		v[1]*sin + v[2]*cos}
	return p.Add(origin)
}

// RotateY rotates a point around the Y-axis.
func (v Vec3f) RotateY(origin Vec3f, rad float32) Vec3f {
	v = v.Sub(origin) // translate to origin

	sin, cos := Sincos(rad)
	p := Vec3f{
		v[2]*sin + v[0]*cos,
		v[1],
		v[2]*cos - v[0]*sin}
	return p.Add(origin)
}

// RotateZ rotates a point around the Z-axis.
func (v Vec3f) RotateZ(origin Vec3f, rad float32) Vec3f {
	v = v.Sub(origin) // translate to origin

	sin, cos := Sincos(rad)
	p := Vec3f{
		v[0]*cos - v[1]*sin,
		v[0]*sin + v[1]*cos,
		v[2]}
	return p.Add(origin)
}

// Distance returns the euclidean distance to another position.
func (v Vec3f) Distance(other Vec3f) float32 {
	return other.Sub(v).Length()
}

// SquareDistance returns the squared euclidean distance to another position.
func (v Vec3f) SquareDistance(other Vec3f) float32 {
	return other.Sub(v).SquareLength()
}
