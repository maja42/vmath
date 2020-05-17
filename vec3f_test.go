package vmath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3f_String(t *testing.T) {
	assert.Equal(t, "Vec3f[5.000000 x -3.100000 x 9.000000]", Vec3f{5, -3.1, 9}.String())
}

func TestVec3f_Format(t *testing.T) {
	assert.Equal(t, "vec 5, -3.2, 7", Vec3f{5, -3.2, 7}.Format("vec %.0f, %.1f, %.f"))
}

func TestVec3f_IsOrthogonal(t *testing.T) {
	assert.True(t, Vec3f{2, 0, 8}.IsOrthogonal())
	assert.True(t, Vec3f{-2, 0, 8}.IsOrthogonal())
	assert.True(t, Vec3f{0, 32, 8}.IsOrthogonal())
	assert.True(t, Vec3f{0, -12, -7}.IsOrthogonal())
	assert.True(t, Vec3f{32, 8, 0}.IsOrthogonal())
	assert.True(t, Vec3f{-12, -7, 0}.IsOrthogonal())

	assert.False(t, Vec3f{2, 2, 5}.IsOrthogonal())
	assert.False(t, Vec3f{-0.0001, 1, -3}.IsOrthogonal())
}

func TestVec3f_Abs(t *testing.T) {
	AssertVec3f(t, Vec3f{2, 9, 4}, Vec3f{2, 9, 4}.Abs())
	AssertVec3f(t, Vec3f{2, 9, 9}, Vec3f{-2, 9, 9}.Abs())
	AssertVec3f(t, Vec3f{2, 9, 5}, Vec3f{2, -9, -5}.Abs())
	AssertVec3f(t, Vec3f{2, 9, 43}, Vec3f{-2, -9, -43}.Abs())
}

func TestVec3f_Add(t *testing.T) {
	AssertVec3f(t, Vec3f{4, 6, 6}, Vec3f{2, 9, 4}.Add(Vec3f{2, -3, 2}))
}

func TestVec3f_IsZero(t *testing.T) {
	assert.True(t, Vec3f{0, 0, 0}.IsZero())
	assert.False(t, Vec3f{0.0001, 0, 0}.IsZero())
	assert.False(t, Vec3f{0, 0.0001, 0}.IsZero())
}

func TestVec3f_AddScalar(t *testing.T) {
	AssertVec3f(t, Vec3f{4, 11, 3}, Vec3f{2, 9, 1}.AddScalar(2))
	AssertVec3f(t, Vec3f{1, -2, -4}, Vec3f{4, 1, -1}.AddScalar(-3))
}

func TestVec3f_Angle(t *testing.T) {
	AssertFloat(t, pi/2, Vec3f{1, 0, 0}.Angle(Vec3f{0, 1, 0}))
	AssertFloat(t, pi, Vec3f{8, 0, 0}.Angle(Vec3f{-2, 0, 0}))
	AssertFloat(t, pi, Vec3f{0, 0, 7}.Angle(Vec3f{0, 0, -4}))
	AssertFloat(t, pi, Vec3f{0, 2, 0}.Angle(Vec3f{0, -9, 0}))

	AssertFloat(t, pi/2, Vec3f{1, 0, 0}.Angle(Vec3f{0, -1, 0}))

	AssertFloat(t, pi, Vec3f{1, 0, 0}.Angle(Vec3f{-4, 0, 0}))
	AssertFloat(t, pi/2, Vec3f{1, 1, 0}.Angle(Vec3f{-4, 4, 0}))
}

func TestVec3f_Clamp(t *testing.T) {
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 0, 0}.Clamp(-5, 5))
	AssertVec3f(t, Vec3f{-5, -5, 5}, Vec3f{-8, -9, 5.01}.Clamp(-5, 5))
	AssertVec3f(t, Vec3f{5, 5, -5}, Vec3f{8, 12, -7}.Clamp(-5, 5))
}

func TestVec3f_Distance(t *testing.T) {
	AssertFloat(t, 2, Vec3f{0, 0, 0}.Distance(Vec3f{2, 0, 0}))
	AssertFloat(t, 2, Vec3f{0, 0, 0}.Distance(Vec3f{0, 0, -2}))
	AssertFloat(t, 0, Vec3f{3, -4, 5}.Distance(Vec3f{3, -4, 5}))
	AssertFloat(t, 20.09975, Vec3f{3, -4, 7}.Distance(Vec3f{1, 8, -9}))
}

func TestVec3f_Div(t *testing.T) {
	AssertVec3f(t, Vec3f{3, -8, 7}, Vec3f{12, 56, 42}.Div(Vec3f{4, -7, 6}))
}

func TestVec3f_DivScalar(t *testing.T) {
	AssertVec3f(t, Vec3f{2.5, 7, 10.25}, Vec3f{20, 56, 82}.DivScalar(8))
}

func TestVec3f_Dot(t *testing.T) {
	AssertFloat(t, -14, Vec3f{4, 7, 2}.Dot(Vec3f{2, -4, 3}))
}

func TestVec3f_Equal(t *testing.T) {
	assert.True(t, Vec3f{0, 0, 0}.Equal(Vec3f{0, 0, 0}))
	assert.True(t, Vec3f{6, 7, 8}.Equal(Vec3f{6, 7, 8}))
	assert.True(t, Vec3f{-4, -9, -5}.Equal(Vec3f{-4, -9, -5}))

	assert.False(t, Vec3f{6, 7}.Equal(Vec3f{7, 6}))
}

func TestVec3f_IsParallel(t *testing.T) {
	assert.True(t, Vec3f{1, 1, 1}.IsParallel(Vec3f{5, 5, 5}))
	assert.True(t, Vec3f{1, 1, 1}.IsParallel(Vec3f{-1, -1, -1}))
	assert.True(t, Vec3f{2, 1, 0}.IsParallel(Vec3f{2, 1, 0}))
	assert.False(t, Vec3f{1, 1, 1}.IsParallel(Vec3f{-1, 2, 1}))
}

func TestVec3f_IsCollinear(t *testing.T) {
	assert.True(t, Vec3f{1, 1, 1}.IsCollinear(Vec3f{5, 5, 5}))
	assert.False(t, Vec3f{1, 1, 1}.IsCollinear(Vec3f{-1, -1, -1}))

	assert.True(t, Vec3f{0, 1, 0}.IsCollinear(Vec3f{0, 1, 0}))
	assert.False(t, Vec3f{0, 1, 0}.IsCollinear(Vec3f{0, -1, 0}))

	assert.False(t, Vec3f{1, 1, 1}.IsCollinear(Vec3f{-1, 2, 1}))
}

func TestVec3f_Length(t *testing.T) {
	AssertFloat(t, 5, Vec3f{5, 0, 0}.Length())
	AssertFloat(t, 5, Vec3f{0, -5, 0}.Length())
	AssertFloat(t, 8, Vec3f{0, 0, 8}.Length())
	AssertFloat(t, 3.7416575, Vec3f{2, 1, 3}.Length())
}

func TestVec3f_Lerp(t *testing.T) {
	AssertVec3f(t, Vec3f{-1, 0, 0}, Vec3f{-1, 0, 0}.Lerp(Vec3f{1, 0, 0}, 0))
	AssertVec3f(t, Vec3f{-0.5, 0, 0}, Vec3f{-1, 0, 0}.Lerp(Vec3f{1, 0, 0}, 0.25))
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{-1, 0, 0}.Lerp(Vec3f{1, 0, 0}, 0.5))
	AssertVec3f(t, Vec3f{0.5, 0, 0}, Vec3f{-1, 0, 0}.Lerp(Vec3f{1, 0, 0}, 0.75))
	AssertVec3f(t, Vec3f{1, 0, 0}, Vec3f{-1, 0, 0}.Lerp(Vec3f{1, 0, 0}, 1))

	AssertVec3f(t, Vec3f{0, 0.5, 0.5}, Vec3f{0, 0, 1}.Lerp(Vec3f{0, 1, 0}, 0.5))
}

func TestVec3f_Mul(t *testing.T) {
	AssertVec3f(t, Vec3f{6, -27, 16}, Vec3f{3, 9, 2}.Mul(Vec3f{2, -3, 8}))
}

func TestVec3f_MulScalar(t *testing.T) {
	AssertVec3f(t, Vec3f{8, 18, -2}, Vec3f{4, 9, -1}.MulScalar(2))
	AssertVec3f(t, Vec3f{-12, -3, 6}, Vec3f{4, 1, -2}.MulScalar(-3))
}

func TestVec3f_Negate(t *testing.T) {
	AssertVec3f(t, Vec3f{6, -4, 9}, Vec3f{-6, 4, -9}.Negate())
}

func TestVec3f_Normalize(t *testing.T) {
	AssertVec3f(t, Vec3f{1, 0, 0}, Vec3f{8, 0, 0}.Normalize())
	AssertVec3f(t, Vec3f{0, 1, 0}, Vec3f{0, 23, 0}.Normalize())
	AssertVec3f(t, Vec3f{0, 0, 1}, Vec3f{0, 0, 7}.Normalize())

	AssertVec3f(t, Vec3f{-1, 0, 0}, Vec3f{-15, 0, 0}.Normalize())
	AssertVec3f(t, Vec3f{0, -1, 0}, Vec3f{0, -54, 0}.Normalize())
	AssertVec3f(t, Vec3f{0, 0, -1}, Vec3f{0, 0, -7}.Normalize())
	AssertVec3f(t, Vec3f{0.872872, 0.436436, 0.218218}, Vec3f{12, 6, 3}.Normalize())
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 0, 0}.Normalize())
}

func TestVec3f_Project(t *testing.T) {
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 1, 0}.Project(Vec3f{1, 0, 0}))
	AssertVec3f(t, Vec3f{1, 0, 0}, Vec3f{1, 1, 0}.Project(Vec3f{1, 0, 0}))
	AssertVec3f(t, Vec3f{1, 0, 0}, Vec3f{1, 2, 3}.Project(Vec3f{2, 0, 0}))
	AssertVec3f(t, Vec3f{0, 0, 9}, Vec3f{1, 2, 9}.Project(Vec3f{0, 0, 4.5}))

	AssertVec3f(t, Vec3f{0.5, 0.5, 0}, Vec3f{0, 1, 0}.Project(Vec3f{1, 1, 0}))
	AssertVec3f(t, Vec3f{0.5, 0.5, 0}, Vec3f{1, 0, 0}.Project(Vec3f{1, 1, 0}))
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{1, 0, 0}.Project(Vec3f{0, 1, 1}))
	AssertVec3f(t, Vec3f{0, 0.5, 0.5}, Vec3f{0, 1, 0}.Project(Vec3f{0, 1, 1}))
	AssertVec3f(t, Vec3f{0, 0.5, 0.5}, Vec3f{2, 0, 1}.Project(Vec3f{0, 1, 1}))
}

func TestVec3f_RotateX(t *testing.T) {
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 0, 0}.RotateX(Vec3f{}, pi))
	AssertVec3f(t, Vec3f{7, -1, 0}, Vec3f{7, 0, 1}.RotateX(Vec3f{}, pi/2))
	AssertVec3f(t, Vec3f{-5, 0, 1}, Vec3f{-5, -1, 0}.RotateX(Vec3f{}, -pi/2))

	AssertVec3f(t, Vec3f{0, -5, 0}, Vec3f{0, 5, 0}.RotateX(Vec3f{}, pi))
	AssertVec3f(t, Vec3f{0, 5, 0}, Vec3f{0, 5, 0}.RotateX(Vec3f{0, 5, 0}, pi))
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 5, 0}.RotateX(Vec3f{0, 2.5, 0}, pi))
}

func TestVec3f_RotateY(t *testing.T) {
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 0, 0}.RotateY(Vec3f{}, pi))
	AssertVec3f(t, Vec3f{1, 7, 0}, Vec3f{0, 7, 1}.RotateY(Vec3f{}, pi/2))
	AssertVec3f(t, Vec3f{0, -5, -1}, Vec3f{-1, -5, 0}.RotateY(Vec3f{}, -pi/2))

	AssertVec3f(t, Vec3f{-5, 0, 0}, Vec3f{5, 0, 0}.RotateY(Vec3f{}, pi))
	AssertVec3f(t, Vec3f{5, 0, 0}, Vec3f{5, 0, 0}.RotateY(Vec3f{5, 0, 0}, pi))
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{5, 0, 0}.RotateY(Vec3f{2.5, 0, 0}, pi))
}

func TestVec3f_RotateZ(t *testing.T) {
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 0, 0}.RotateZ(Vec3f{}, pi))
	AssertVec3f(t, Vec3f{0, 1, 7}, Vec3f{1, 0, 7}.RotateZ(Vec3f{}, pi/2))
	AssertVec3f(t, Vec3f{-1, 0, -5}, Vec3f{0, -1, -5}.RotateZ(Vec3f{}, -pi/2))

	AssertVec3f(t, Vec3f{0, -5, 0}, Vec3f{0, 5, 0}.RotateZ(Vec3f{}, pi))
	AssertVec3f(t, Vec3f{0, 5, 0}, Vec3f{0, 5, 0}.RotateZ(Vec3f{0, 5, 0}, pi))
	AssertVec3f(t, Vec3f{0, 0, 0}, Vec3f{0, 5, 0}.RotateZ(Vec3f{0, 2.5, 0}, pi))
}

func TestVec3f_Round(t *testing.T) {
	assert.Equal(t, Vec3i{-6, 7, 5}, Vec3f{-6, 7, 5}.Round())
	assert.Equal(t, Vec3i{-5, 7, 5}, Vec3f{-5.2, 7.2, 5.2}.Round())
	assert.Equal(t, Vec3i{-6, 8, 10}, Vec3f{-5.8, 7.8, 9.7}.Round())
	assert.Equal(t, Vec3i{-6, 8, 4}, Vec3f{-5.5, 7.5, 3.5}.Round())
}

func TestVec3f_Split(t *testing.T) {
	x, y, z := Vec3f{6, -9, 12}.Split()
	AssertFloat(t, 6, x)
	AssertFloat(t, -9, y)
	AssertFloat(t, 12, z)
}

func TestVec3f_SquareDistance(t *testing.T) {
	AssertFloat(t, 0, Vec3f{3, 4, 5}.SquareDistance(Vec3f{3, 4, 5}))
	AssertFloat(t, 12, Vec3f{3, 4, 5}.SquareDistance(Vec3f{1, 2, 3}))
}

func TestVec3f_SquareLength(t *testing.T) {
	AssertFloat(t, 13, Vec3f{2, -3}.SquareLength())
}

func TestVec3f_Sub(t *testing.T) {
	AssertVec3f(t, Vec3f{0, 12, 2}, Vec3f{2, 9, 7}.Sub(Vec3f{2, -3, 5}))
}

func TestVec3f_SubScalar(t *testing.T) {
	AssertVec3f(t, Vec3f{-1, 7, -7}, Vec3f{1, 9, -5}.SubScalar(2))
	AssertVec3f(t, Vec3f{7, 4, -1}, Vec3f{4, 1, -4}.SubScalar(-3))
}

func TestVec3f_Vec3i(t *testing.T) {
	assert.Equal(t, Vec3i{-6, 7, 9}, Vec3f{-6, 7, 9}.Vec3i())
	assert.Equal(t, Vec3i{-5, 7, -9}, Vec3f{-5.8, 7.5, -9.5}.Vec3i())
}

func TestVec3f_Vec4f(t *testing.T) {
	AssertVec4f(t, Vec4f{1, 2, 3, 4}, Vec3f{1, 2, 3}.Vec4f(4))
}

func TestVec3f_X(t *testing.T) {
	AssertFloat(t, -45, Vec3f{-45, 12, 7}.X())
}

func TestVec3f_Y(t *testing.T) {
	AssertFloat(t, 12, Vec3f{-45, 12, 7}.Y())
}

func TestVec3f_Z(t *testing.T) {
	AssertFloat(t, 7, Vec3f{-45, 12, 7}.Z())
}

func TestVec3f_XY(t *testing.T) {
	AssertVec2f(t, Vec2f{-45, 12}, Vec3f{-45, 12, 7}.XY())
}

func TestVec3f_RotationTo(t *testing.T) {
	// around x
	rot := Vec3f{0, 1, 0}.RotationTo(Vec3f{0, 0, 1})
	expected := QuatFromAxisAngle(Vec3f{1, 0, 0}, math.Pi/2)
	assert.True(t, expected.Equals(rot))
	// around y
	rot = Vec3f{0, 0, 1}.RotationTo(Vec3f{1, 0, 0})
	expected = QuatFromAxisAngle(Vec3f{0, 1, 0}, math.Pi/2)
	assert.True(t, expected.Equals(rot))
	// around z
	rot = Vec3f{1, 0, 0}.RotationTo(Vec3f{0, 1, 0})
	expected = QuatFromAxisAngle(Vec3f{0, 0, 1}, math.Pi/2)
	assert.True(t, expected.Equals(rot))

	// no rotation
	rot = Vec3f{0, 1, 0}.RotationTo(Vec3f{0, 1, 0})
	expected = QuatFromAxisAngle(Vec3f{1, 0, 0}, 0)
	assert.True(t, expected.Equals(rot))
	// 180°
	rot = Vec3f{1, 0, 0}.RotationTo(Vec3f{-1, 0, 0})
	expected = QuatFromAxisAngle(Vec3f{0, 0, 1}, -math.Pi)
	assert.True(t, expected.Equals(rot))
}
