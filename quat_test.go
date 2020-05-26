package vmath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const deg90 = float32(math.Pi / 2)
const deg180 = float32(math.Pi)
const deg270 = float32(math.Pi * 3 / 2)

func TestQuat_String(t *testing.T) {
	quat := QuatFromAxisAngle(Vec3f{1, 0, 0}, deg180)
	assert.Equal(t, "Quat[-0.000000, 1.000000 x 0.000000 x 0.000000]", quat.String())

	quat = QuatFromAxisAngle(Vec3f{0, 1, 1}, deg90)
	assert.Equal(t, "Quat[0.707107, 0.000000 x 0.500000 x 0.500000]", quat.String())
}

func TestIdentQuat(t *testing.T) {
	y, p, r := IdentQuat().ToEuler()
	assert.Equal(t, 0, y)
	assert.Equal(t, 0, p)
	assert.Equal(t, 0, r)

	assert.Equal(t, Vec4f{0, 0, 0, 1}, IdentQuat())
}

func TestQuatFromAxisAngle(t *testing.T) {
	quat := QuatFromAxisAngle(Vec3f{0, 0, 1}, deg180)
	AssertVec4f(t, Vec4f{0, 0, 0, 1}, quat.Vec4f())

	quat = QuatFromAxisAngle(Vec3f{1, 0, 0}, deg90)
	AssertVec4f(t, Vec4f{0.707106, 0.707106, 0, 0}, quat.Vec4f())

	quat = QuatFromAxisAngle(Vec3f{0, 1, 1}, deg90)
	AssertVec4f(t, Vec4f{0.707107, 0, 0.5, 0.5}, quat.Vec4f())
}

func TestQuatFromEuler(t *testing.T) {
	quat := QuatFromEuler(deg90, 0, 0)
	AssertQuat(t, Quat{W: 0.707107, X: 0, Y: 0, Z: 0.707107}, quat)

	quat = QuatFromEuler(0, 0, -deg90)
	AssertQuat(t, Quat{W: 0.707107, X: -0.707107, Y: 0, Z: 0}, quat)

	quat = QuatFromEuler(0, deg180, 0)
	AssertQuat(t, Quat{W: 0, X: 0, Y: 1, Z: 0}, quat)
}

func TestQuat_Equals(t *testing.T) {
	quatA := Quat{W: 1, X: 3, Y: 5, Z: 7}
	quatB := Quat{W: 1, X: 3, Y: 5, Z: 7}
	assert.True(t, quatA.Equals(quatB))
	assert.False(t, quatA.Equals(IdentQuat()))
}

func TestQuat_Add(t *testing.T) {
	quatA := Quat{W: 1, X: 3, Y: 5, Z: 7}
	quatB := Quat{W: 9, X: 8, Y: 7, Z: 6}
	AssertQuat(t, Quat{W: 10, X: 11, Y: 12, Z: 13}, quatA.Add(quatB))
}

func TestQuat_AddScalar(t *testing.T) {
	quat := Quat{W: 1, X: 2, Y: 3, Z: 4}
	AssertQuat(t, Quat{W: 4, X: 5, Y: 6, Z: 7}, quat.AddScalar(3))
}

func TestQuat_Sub(t *testing.T) {
	quatA := Quat{W: 1, X: 2, Y: 3, Z: 4}
	quatB := Quat{W: 9, X: 8, Y: 7, Z: 6}
	AssertQuat(t, Quat{W: -8, X: -8, Y: -4, Z: -2}, quatA.Sub(quatB))
}

func TestQuat_SubScalar(t *testing.T) {
	quat := Quat{W: 1, X: 2, Y: 3, Z: 4}
	AssertQuat(t, Quat{W: -1, X: 0, Y: 1, Z: 2}, quat.SubScalar(2))
}

func TestQuat_Mul(t *testing.T) {
	quatA := Quat{W: 1, X: 3, Y: 5, Z: 7}
	quatB := Quat{W: 9, X: 8, Y: 7, Z: 6}
	AssertQuat(t, Quat{W: 9, X: 24, Y: 35, Z: 42}, quatA.Mul(quatB))
}

func TestQuat_MulScalar(t *testing.T) {
	quat := Quat{W: 1, X: 2, Y: 3, Z: 4}
	AssertQuat(t, Quat{W: 2, X: 4, Y: 6, Z: 8}, quat.MulScalar(2))
}

func TestQuat_Div(t *testing.T) {
	quatA := Quat{W: 81, X: 56, Y: 21, Z: 72}
	quatB := Quat{W: 9, X: 8, Y: 7, Z: 6}
	AssertQuat(t, Quat{W: 9, X: 7, Y: 3, Z: 12}, quatA.Div(quatB))
}

func TestQuat_DivScalar(t *testing.T) {
	quat := Quat{W: 1, X: 2, Y: 3, Z: 4}
	AssertQuat(t, Quat{W: 0.5, X: 1, Y: 1.5, Z: 2}, quat.DivScalar(2))
}

func TestQuat_Rotate(t *testing.T) {
	quatA := QuatFromAxisAngle(Vec3f{1, 0, 0}, deg90)

	quatB := QuatFromAxisAngle(Vec3f{1, 0, 0}, deg90)
	AssertQuat(t, Quat{W: 0, X: 1, Y: 0, Z: 0}, quatA.Rotate(quatB))

	quatC := QuatFromAxisAngle(Vec3f{1, 0, 0}, -deg90)
	AssertQuat(t, Quat{W: 1, X: 0, Y: 0, Z: 0}, quatA.Rotate(quatC))

	quatD := QuatFromAxisAngle(Vec3f{0, 1, 0}, deg90)
	AssertQuat(t, Quat{W: 0.5, X: 0.5, Y: 0.5, Z: -0.5}, quatA.Rotate(quatD))
}

func TestQuat_RotateX(t *testing.T) {
	quat := IdentQuat().RotateX(deg90)

	result := quat.RotateVec(Vec3f{1, 0, 0})
	AssertVec3f(t, Vec3f{1, 0, 0}, result)

	result = quat.RotateVec(Vec3f{0, 0, -1})
	AssertVec3f(t, Vec3f{0, 1, 0}, result)
}

func TestQuat_RotateY(t *testing.T) {
	quat := IdentQuat().RotateY(deg90)

	result := quat.RotateVec(Vec3f{0, 1, 0})
	AssertVec3f(t, Vec3f{0, 1, 0}, result)

	result = quat.RotateVec(Vec3f{0, 0, -1})
	AssertVec3f(t, Vec3f{-1, 0, 0}, result)
}

func TestQuat_RotateZ(t *testing.T) {
	quat := IdentQuat().RotateZ(deg90)

	result := quat.RotateVec(Vec3f{0, 0, 1})
	AssertVec3f(t, Vec3f{0, 0, 1}, result)

	result = quat.RotateVec(Vec3f{0, 1, 0})
	AssertVec3f(t, Vec3f{-1, 0, 0}, result)
}

func TestQuat_Dot(t *testing.T) {
	quatA := Quat{W: 1, X: 3, Y: 5, Z: 7}
	quatB := Quat{W: 9, X: 8, Y: 7, Z: 6}
	AssertFloat(t, 9+24+35+42, quatA.Dot(quatB))
}

func TestQuat_Inverse(t *testing.T) {
	quat := QuatFromAxisAngle(Vec3f{1, 5, 2}, deg90)
	expected := QuatFromAxisAngle(Vec3f{1, 5, 2}, deg270)
	AssertQuat(t, expected, quat.Inverse())

	quat = QuatFromAxisAngle(Vec3f{1, 1, 6}, deg180)
	expected = QuatFromAxisAngle(Vec3f{1, 1, 6}, deg180)
	AssertQuat(t, expected, quat.Inverse())
}

func TestQuat_Conjugate(t *testing.T) {
	quat := QuatFromAxisAngle(Vec3f{1, 2, -4}, deg90)
	expected := QuatFromAxisAngle(Vec3f{-1, -2, 4}, deg90)

	AssertQuat(t, expected, quat.Conjugate())
}

func TestQuat_Length(t *testing.T) {
	quat := Quat{W: 2, X: 0, Y: 0, Z: 0}
	AssertFloat(t, 2, quat.Length())

	quat = Quat{W: 1, X: 2, Y: 3, Z: 4}
	AssertFloat(t, 5.477226, quat.Length())
}

func TestQuat_SquareLength(t *testing.T) {
	quat := Quat{W: 2, X: 0, Y: 0, Z: 0}
	AssertFloat(t, 4, quat.SquareLength())

	quat = Quat{W: 1, X: 2, Y: 3, Z: 4}
	AssertFloat(t, 30, quat.SquareLength())
}

func TestQuat_Normalize(t *testing.T) {
	quat := Quat{X: 10, Y: 0, Z: 0, W: 0}
	AssertQuat(t, Quat{X: 1, Y: 0, Z: 0, W: 0}, quat.Normalize())

	quat = QuatFromAxisAngle(Vec3f{100, 0, 0}, 2*deg180+deg90)
	AssertQuat(t, Quat{X: -0.707106, Y: 0, Z: 0, W: -0.707106}, quat.Normalize())
}

func TestQuat_Directions(t *testing.T) {
	quat := IdentQuat()
	AssertVec3f(t, Vec3f{0, 1, 0}, quat.Up())
	AssertVec3f(t, Vec3f{0, 0, -1}, quat.Forward())
	AssertVec3f(t, Vec3f{1, 0, 0}, quat.Right())

	quat = quat.RotateY(deg90)
	AssertVec3f(t, Vec3f{0, 1, 0}, quat.Up())
	AssertVec3f(t, Vec3f{-1, 0, 0}, quat.Forward())
	AssertVec3f(t, Vec3f{0, 0, -1}, quat.Right())
}

func TestQuat_Axis_Angle(t *testing.T) {
	quat := IdentQuat()
	AssertVec3f(t, Vec3f{0, 0, 0}, quat.Axis())
	AssertFloat(t, 0, quat.Angle())

	quat = quat.RotateY(deg90)
	AssertVec3f(t, Vec3f{0, 1, 0}, quat.Axis().Normalize())
	AssertFloat(t, deg90, quat.Angle())
}

func TestQuat_AxisRotation(t *testing.T) {
	quat := IdentQuat()
	axis, rad := quat.AxisRotation()
	AssertVec3f(t, Vec3f{1, 0, 0}, axis)
	AssertFloat(t, 0, rad)

	quat = quat.RotateY(deg90)
	axis, rad = quat.AxisRotation()
	AssertVec3f(t, Vec3f{0, 1, 0}, axis)
	AssertFloat(t, deg90, rad)
}

func TestQuat_AngleTo(t *testing.T) {
	quatA := QuatFromAxisAngle(Vec3f{1, 0, 0}, deg90)

	quatB := QuatFromAxisAngle(Vec3f{1, 0, 0}, deg270)
	AssertFloat(t, deg180, quatA.AngleTo(quatB))

	quatC := QuatFromAxisAngle(Vec3f{0, 1, 0}, deg90)
	AssertFloat(t, deg90, quatA.AngleTo(quatC))
}
