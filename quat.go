package vmath

import (
	"fmt"
	"math"

	"github.com/maja42/vmath/math32"
)

// Quat represents a Quaternion.
type Quat struct {
	W       float32
	X, Y, Z float32
}

func (q Quat) String() string {
	return fmt.Sprintf("Quat[%f, %f x %f x %f]", q.W, q.X, q.Y, q.Z)
}

// IdentQuat returns the identity quaternion.
func IdentQuat() Quat {
	return Quat{1, 0, 0, 0}
}

// QuatFromAxisAngle returns a quaternion representing a rotation around a given axis.
func QuatFromAxisAngle(axis Vec3f, angle float32) Quat {
	axis = axis.Normalize()
	sinAngle, cosAngle := math32.Sincos(angle * 0.5)
	return Quat{
		cosAngle,
		axis[0] * sinAngle,
		axis[1] * sinAngle,
		axis[2] * sinAngle,
	}
}

// Equals compares two quaternions.
// Uses the default Epsilon as relative tolerance.
func (q Quat) Equals(other Quat) bool {
	return q.EqualsEps(other, Epsilon)
}

// EqualsEps compares two quaternions, using the given epsilon as a relative tolerance.
func (q Quat) EqualsEps(other Quat, epsilon float32) bool {
	return EqualEps(q.W, other.W, epsilon) &&
		EqualEps(q.X, other.X, epsilon) && EqualEps(q.Y, other.Y, epsilon) && EqualEps(q.Z, other.Z, epsilon)
}

// Vec4f returns the quaternion as a vector representation.
func (q Quat) Vec4f() Vec4f {
	return Vec4f{q.W, q.X, q.Y, q.Z}
}

// Add performs component-wise addition.
func (q Quat) Add(other Quat) Quat {
	return Quat{q.W + other.W, q.X + other.X, q.Y + other.Y, q.Z + other.Z}
}

// AddScalar performs component-wise scalar addition.
func (q Quat) AddScalar(s float32) Quat {
	return Quat{q.W + s, q.X + s, q.Y + s, q.Z + s}
}

// Sub performs component-wise subtraction.
func (q Quat) Sub(other Quat) Quat {
	return Quat{q.W - other.W, q.X - other.X, q.Y - other.Y, q.Z - other.Z}
}

// SubScalar performs component-wise scalar subtraction.
func (q Quat) SubScalar(s float32) Quat {
	return Quat{q.W - s, q.X - s, q.Y - s, q.Z - s}
}

// Mul multiplies two quaternions, performing a rotation.
func (q Quat) Mul(other Quat) Quat {
	return Quat{
		(other.W * q.W) - (other.X * q.X) - (other.Y * q.Y) - (other.Z * q.Z),
		(other.X * q.W) + (other.W * q.X) - (other.Z * q.Y) + (other.Y * q.Z),
		(other.Y * q.W) + (other.Z * q.X) + (other.W * q.Y) - (other.X * q.Z),
		(other.Z * q.W) - (other.Y * q.X) + (other.X * q.Y) + (other.W * q.Z),
	}
}

// MulScalar performs component-wise scalar multiplication.
func (q Quat) MulScalar(s float32) Quat {
	return Quat{q.W * s, q.X * s, q.Y * s, q.Z * s}
}

// Div performs component-wise division.
func (q Quat) Div(other Quat) Quat {
	return Quat{q.W / other.W, q.X / other.X, q.Y / other.Y, q.Z / other.Z}
}

// DivScalar performs component-wise scalar division.
func (q Quat) DivScalar(s float32) Quat {
	return Quat{q.W / s, q.X / s, q.Y / s, q.Z / s}
}

// RotateX rotates the quaternion with a given angle round its X axis.
func (q Quat) RotateX(rad float32) Quat {
	// Source: http://glmatrix.net/docs/module-quat.html
	sinR, cosR := math32.Sincos(rad * 0.5)
	return Quat{
		q.W*cosR - q.X*sinR,
		q.X*cosR + q.W*sinR,
		q.Y*cosR + q.Z*sinR,
		q.Z*cosR - q.Y*sinR,
	}
}

// RotateY rotates the quaternion with a given angle round its Y axis.
func (q Quat) RotateY(rad float32) Quat {
	// Source: http://glmatrix.net/docs/module-quat.html
	sinR, cosR := math32.Sincos(rad * 0.5)
	return Quat{
		q.W*cosR - q.Y*sinR,
		q.X*cosR - q.Z*sinR,
		q.Y*cosR + q.W*sinR,
		q.Z*cosR + q.X*sinR,
	}
}

// RotateZ rotates the quaternion with a given angle round its Y axis.
func (q Quat) RotateZ(rad float32) Quat {
	// Source: http://glmatrix.net/docs/module-quat.html
	sinR, cosR := math32.Sincos(rad * 0.5)
	return Quat{
		q.W*cosR - q.Z*sinR,
		q.X*cosR + q.Y*sinR,
		q.Y*cosR - q.X*sinR,
		q.Z*cosR + q.W*sinR,
	}
}

// Dot performs a dot product with another quaternion.
func (q Quat) Dot(other Quat) float32 {
	return q.W*other.W + q.X*other.X + q.Y*other.Y + q.Z*other.Z
}

// Inverse returns the inverse quaternion.
func (q Quat) Inverse() Quat {
	return Quat{-q.W, q.X, q.Y, q.Z}
}

// Conjugate returns the conjugated quaternion.
func (q Quat) Conjugate() Quat {
	return Quat{q.W, -q.X, -q.Y, -q.Z}
}

// Length returns the quaternion's length.
func (q Quat) Length() float32 {
	return math32.Sqrt(q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z)
}

// SquareLength returns the quaternion's squared length.
func (q Quat) SquareLength() float32 {
	return q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z
}

// Normalize the quaternion.
// The quaternion must be non-zero.
func (q Quat) Normalize() Quat {
	length := q.Length()
	if Equalf(length, 1) { // shortcut
		return q
	}
	return Quat{q.W / length, q.X / length, q.Y / length, q.Z / length}
}

// Right returns the up-vector of the quaternion's coordinate system.
func (q Quat) Up() Vec3f {
	return q.RotateVec(Vec3f{0, 1, 0})
}

// Forward returns the forward-vector of the quaternion's coordinate system.
func (q Quat) Forward() Vec3f {
	return q.RotateVec(Vec3f{0, 0, -1})
}

// Right returns the right-vector of the quaternion's coordinate system.
func (q Quat) Right() Vec3f {
	return q.RotateVec(Vec3f{1, 0, 0})
}

// Axis returns the quaternion's rotation axis.
// The returned axis is not normalized.
func (q Quat) Axis() Vec3f {
	return Vec3f{q.X, q.Y, q.Z}
}

// Angle returns the quaternion's rotation angle around its axis.
func (q Quat) Angle() float32 {
	q = q.Normalize()
	return math32.Acos(q.W) * 2
}

// AxisRotation returns the quaternion's rotation angle and axis.
func (q Quat) AxisRotation() (Vec3f, float32) {
	// Based on: http://glmatrix.net/docs/module-quat.html
	rad := q.Angle()
	s := math32.Sin(rad * 0.5)
	if s < Epsilon { // no rotation
		return Vec3f{1, 0, 0}, rad
	}
	return Vec3f{q.X / s, q.Y / s, q.Z / s}, rad
}

// QuatFromEuler returns a quaternion based on the given euler rotations.
// Axis: yaw: Z, pitch: Y, roll: X
func QuatFromEuler(yaw, pitch, roll float32) Quat {
	// Source: https://en.wikipedia.org/wiki/Conversion_between_quaternions_and_Euler_angles
	sinY, cosY := math32.Sincos(yaw * 0.5)
	sinP, cosP := math32.Sincos(pitch * 0.5)
	sinR, cosR := math32.Sincos(roll * 0.5)
	return Quat{
		W: cosR*cosP*cosY + sinR*sinP*sinY,
		X: sinR*cosP*cosY - cosR*sinP*sinY,
		Y: cosR*sinP*cosY + sinR*cosP*sinY,
		Z: cosR*cosP*sinY - sinR*sinP*cosY,
	}
}

// ToEuler converts the quaternion into euler rotations.
// Axis: yaw: Z, pitch: Y, roll: X
func (q Quat) ToEuler() (yaw, pitch, roll float32) {
	// Source: https://en.wikipedia.org/wiki/Conversion_between_quaternions_and_Euler_angles

	// roll (x-axis rotation)
	srcp := 2 * (q.W*q.X + q.Y*q.Z)
	crcp := 1 - 2*(q.X*q.X+q.Y*q.Y)
	roll = math32.Atan2(srcp, crcp)

	// pitch (y-axis rotation)
	sp := 2 * (q.W*q.Y - q.Z*q.X)
	if math32.Abs(sp) >= 1 {
		pitch = math32.Copysign(math.Pi/2, sp) // use 90° if out of range
	} else {
		pitch = math32.Asin(sp)
	}

	// yaw (z-axis rotation)
	sycp := 2 * (q.W*q.Z + q.X*q.Y)
	cycp := 1 - 2*(q.Y*q.Y+q.Z*q.Z)
	yaw = math32.Atan2(sycp, cycp)

	return
}

// AngleTo returns the angle between two quaternions by comparing one of their axis.
func (q Quat) AngleTo(other Quat) float32 {
	return q.Forward().Angle(other.Forward())
}

// Mat4f returns a homogeneous 3D rotation matrix based on the quaternion.
func (q Quat) Mat4f() Mat4f {
	return Mat4f{
		1 - 2*q.Y*q.Y - 2*q.Z*q.Z, 2*q.X*q.Y + 2*q.W*q.Z, 2*q.X*q.Z - 2*q.W*q.Y, 0,
		2*q.X*q.Y - 2*q.W*q.Z, 1 - 2*q.X*q.X - 2*q.Z*q.Z, 2*q.Y*q.Z + 2*q.W*q.X, 0,
		2*q.X*q.Z + 2*q.W*q.Y, 2*q.Y*q.Z - 2*q.W*q.X, 1 - 2*q.X*q.X - 2*q.Y*q.Y, 0,
		0, 0, 0, 1,
	}
}

// RotateVec rotates a vector.
func (q Quat) RotateVec(v Vec3f) Vec3f {
	// Source: https://gamedev.stackexchange.com/a/50545/39091
	s := q.W
	u := Vec3f{q.X, q.Y, q.Z}

	a := u.MulScalar(2 * u.Dot(v))
	b := v.MulScalar(s*s - u.Dot(u))
	c := u.Cross(v).MulScalar(2 * s)
	return a.Add(b).Add(c)
}

// Lerp performs a linear interpolation to another quaternion.
// The parameter t should be in range [0, 1].
func (q Quat) Lerp(other Quat, t float32) Quat {
	return q.Mul(other.MulScalar(t))
}

// Slerp performs a spherical linear interpolation to another quaternion.
// The parameter t should be in range [0, 1].
func (q Quat) Slerp(other Quat, t float32) Quat {
	// Source: http://glmatrix.net/docs/module-quat.html
	dot := q.Dot(other)
	if dot > 0.9999 { // quaternions are close together, perform lerp
		return q.Lerp(other, t)
	}

	if dot < 0.0 { // adjust signs
		dot = -dot
		other.W = -other.W
		other.X = -other.X
		other.Y = -other.Y
		other.Z = -other.Z
	}

	return Quat{
		(1-t)*q.W + 1*other.W,
		(1-t)*q.X + 1*other.X,
		(1-t)*q.Y + 1*other.Y,
		(1-t)*q.Z + 1*other.Z,
	}
}
