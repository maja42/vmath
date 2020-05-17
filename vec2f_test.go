package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec2f_String(t *testing.T) {
	assert.Equal(t, "Vec2f[5.000000 x -3.100000]", Vec2f{5, -3.1}.String())
}

func TestVec2f_Format(t *testing.T) {
	assert.Equal(t, "vec 5, -3", Vec2f{5, -3.1}.Format("vec %.0f, %.0f"))
}

func TestVec2f_IsOrthogonal(t *testing.T) {
	assert.True(t, Vec2f{2, 0}.IsOrthogonal())
	assert.True(t, Vec2f{-2, 0}.IsOrthogonal())
	assert.True(t, Vec2f{0, 32}.IsOrthogonal())
	assert.True(t, Vec2f{0, -12}.IsOrthogonal())

	assert.False(t, Vec2f{2, 2}.IsOrthogonal())
	assert.False(t, Vec2f{-0.0001, 1}.IsOrthogonal())
}

func TestVec2f_Abs(t *testing.T) {
	AssertVec2f(t, Vec2f{2, 9}, Vec2f{2, 9}.Abs())
	AssertVec2f(t, Vec2f{2, 9}, Vec2f{-2, 9}.Abs())
	AssertVec2f(t, Vec2f{2, 9}, Vec2f{2, -9}.Abs())
	AssertVec2f(t, Vec2f{2, 9}, Vec2f{-2, -9}.Abs())
}

func TestVec2f_Add(t *testing.T) {
	AssertVec2f(t, Vec2f{4, 6}, Vec2f{2, 9}.Add(Vec2f{2, -3}))
}

func TestVec2f_IsZero(t *testing.T) {
	assert.True(t, Vec2f{0, 0}.IsZero())
	assert.False(t, Vec2f{0.0001, 0}.IsZero())
	assert.False(t, Vec2f{0, 0.0001}.IsZero())
}

func TestVec2f_AddScalar(t *testing.T) {
	AssertVec2f(t, Vec2f{4, 11}, Vec2f{2, 9}.AddScalar(2))
	AssertVec2f(t, Vec2f{1, -2}, Vec2f{4, 1}.AddScalar(-3))
}

func TestVec2f_Angle(t *testing.T) {
	AssertFloat(t, 0, Vec2f{0, 0}.Angle(Vec2f{5, 0}))
	AssertFloat(t, pi/2, Vec2f{1, 0}.Angle(Vec2f{0, 1}))
	AssertFloat(t, pi, Vec2f{8, 0}.Angle(Vec2f{-2, 0}))
	AssertFloat(t, -pi/2, Vec2f{1, 0}.Angle(Vec2f{0, -1}))

	AssertFloat(t, pi, Vec2f{1, 0}.Angle(Vec2f{-4, 0}))
	AssertFloat(t, pi/2, Vec2f{1, 1}.Angle(Vec2f{-4, 4}))
}

func TestVec2f_Clamp(t *testing.T) {
	AssertVec2f(t, Vec2f{0, 0}, Vec2f{0, 0}.Clamp(-5, 5))
	AssertVec2f(t, Vec2f{-5, -5}, Vec2f{-8, -9}.Clamp(-5, 5))
	AssertVec2f(t, Vec2f{5, 5}, Vec2f{8, 12}.Clamp(-5, 5))
}

func TestVec2f_Distance(t *testing.T) {
	AssertFloat(t, 2, Vec2f{0, 0}.Distance(Vec2f{2, 0}))
	AssertFloat(t, 0, Vec2f{3, -4}.Distance(Vec2f{3, -4}))
	AssertFloat(t, 12.165525, Vec2f{3, -4}.Distance(Vec2f{1, 8}))
}

func TestVec2f_Div(t *testing.T) {
	AssertVec2f(t, Vec2f{3, -8}, Vec2f{12, 56}.Div(Vec2f{4, -7}))
}

func TestVec2f_DivScalar(t *testing.T) {
	AssertVec2f(t, Vec2f{2.5, 7}, Vec2f{20, 56}.DivScalar(8))
}

func TestVec2f_Dot(t *testing.T) {
	AssertFloat(t, -20, Vec2f{4, 7}.Dot(Vec2f{2, -4}))
}

func TestVec2f_Equal(t *testing.T) {
	assert.True(t, Vec2f{0, 0}.Equal(Vec2f{0, 0}))
	assert.True(t, Vec2f{6, 7}.Equal(Vec2f{6, 7}))
	assert.True(t, Vec2f{-4, -9}.Equal(Vec2f{-4, -9}))

	assert.False(t, Vec2f{6, 7}.Equal(Vec2f{7, 6}))
}

func TestVec2f_FlatAngle(t *testing.T) {
	AssertFloat(t, 0, Vec2f{0, 0}.FlatAngle())
	AssertFloat(t, 0, Vec2f{6, 0}.FlatAngle())
	AssertFloat(t, pi/4, Vec2f{1, 1}.FlatAngle())
	AssertFloat(t, pi/2, Vec2f{0, 1}.FlatAngle())
	AssertFloat(t, pi, Vec2f{-5, 0}.FlatAngle())
	AssertFloat(t, 3*pi/2, NormalizeRadians(Vec2f{0, -2}.FlatAngle()))
}

func TestVec2f_IsParallel(t *testing.T) {
	assert.True(t, Vec2f{1, 1}.IsParallel(Vec2f{5, 5}))
	assert.True(t, Vec2f{1, 1}.IsParallel(Vec2f{-1, -1}))
	assert.True(t, Vec2f{1, 0}.IsParallel(Vec2f{1, 0}))
	assert.False(t, Vec2f{1, 1}.IsParallel(Vec2f{-1, 2}))
}

func TestVec2f_IsCollinear(t *testing.T) {
	assert.True(t, Vec2f{1, 1}.IsCollinear(Vec2f{5, 5}))
	assert.False(t, Vec2f{1, 1}.IsCollinear(Vec2f{-1, -1}))

	assert.True(t, Vec2f{0, 1}.IsCollinear(Vec2f{0, 1}))
	assert.False(t, Vec2f{0, 1}.IsCollinear(Vec2f{0, -1}))

	assert.False(t, Vec2f{1, 1}.IsCollinear(Vec2f{-1, 2}))
}

func TestVec2f_Length(t *testing.T) {
	AssertFloat(t, 5, Vec2f{5, 0}.Length())
	AssertFloat(t, 5, Vec2f{0, -5}.Length())
	AssertFloat(t, 2.236068, Vec2f{2, 1}.Length())
}

func TestVec2f_Lerp(t *testing.T) {
	AssertVec2f(t, Vec2f{-1, 0}, Vec2f{-1, 0}.Lerp(Vec2f{1, 0}, 0))
	AssertVec2f(t, Vec2f{-0.5, 0}, Vec2f{-1, 0}.Lerp(Vec2f{1, 0}, 0.25))
	AssertVec2f(t, Vec2f{0, 0}, Vec2f{-1, 0}.Lerp(Vec2f{1, 0}, 0.5))
	AssertVec2f(t, Vec2f{0.5, 0}, Vec2f{-1, 0}.Lerp(Vec2f{1, 0}, 0.75))
	AssertVec2f(t, Vec2f{1, 0}, Vec2f{-1, 0}.Lerp(Vec2f{1, 0}, 1))

	AssertVec2f(t, Vec2f{0.5, 0.5}, Vec2f{0, 1}.Lerp(Vec2f{1, 0}, 0.5))
}

func TestVec2f_MagCross(t *testing.T) {
	AssertFloat(t, -1, Vec2f{0, 1}.MagCross(Vec2f{1, 0}))
	AssertFloat(t, -6, Vec2f{0, 2}.MagCross(Vec2f{3, 0}))
	AssertFloat(t, 47, Vec2f{5, -3}.MagCross(Vec2f{4, 7}))
}

func TestVec2f_Mul(t *testing.T) {
	AssertVec2f(t, Vec2f{6, -27}, Vec2f{3, 9}.Mul(Vec2f{2, -3}))
}

func TestVec2f_MulScalar(t *testing.T) {
	AssertVec2f(t, Vec2f{8, 18}, Vec2f{4, 9}.MulScalar(2))
	AssertVec2f(t, Vec2f{-12, -3}, Vec2f{4, 1}.MulScalar(-3))
}

func TestVec2f_Negate(t *testing.T) {
	AssertVec2f(t, Vec2f{6, -4}, Vec2f{-6, 4}.Negate())
}

func TestVec2f_NormalVec(t *testing.T) {
	AssertVec2f(t, Vec2f{-1, 1}, Vec2f{1, 1}.NormalVec(true))
	AssertVec2f(t, Vec2f{1, -1}, Vec2f{1, 1}.NormalVec(false))

	AssertVec2f(t, Vec2f{0, -1}, Vec2f{-1, 0}.NormalVec(true))
	AssertVec2f(t, Vec2f{0, 1}, Vec2f{-1, 0}.NormalVec(false))
}

func TestVec2f_Normalize(t *testing.T) {
	AssertVec2f(t, Vec2f{1, 0}, Vec2f{8, 0}.Normalize())
	AssertVec2f(t, Vec2f{0, 1}, Vec2f{0, 23}.Normalize())
	AssertVec2f(t, Vec2f{-1, 0}, Vec2f{-15, 0}.Normalize())
	AssertVec2f(t, Vec2f{0, -1}, Vec2f{0, -54}.Normalize())
	AssertVec2f(t, Vec2f{0.8, 0.6}, Vec2f{4, 3}.Normalize())
	AssertVec2f(t, Vec2f{0, 0}, Vec2f{0, 0}.Normalize())
}

func TestVec2f_Project(t *testing.T) {
	AssertVec2f(t, Vec2f{0, 0}, Vec2f{0, 1}.Project(Vec2f{1, 0}))
	AssertVec2f(t, Vec2f{1, 0}, Vec2f{1, 1}.Project(Vec2f{1, 0}))
	AssertVec2f(t, Vec2f{1, 0}, Vec2f{1, 2}.Project(Vec2f{2, 0}))
	AssertVec2f(t, Vec2f{0.5, 0.5}, Vec2f{0, 1}.Project(Vec2f{1, 1}))
	AssertVec2f(t, Vec2f{0.5, 0.5}, Vec2f{1, 0}.Project(Vec2f{1, 1}))

}

func TestVec2f_Rotate(t *testing.T) {
	AssertVec2f(t, Vec2f{0, 0}, Vec2f{0, 0}.Rotate(pi))
	AssertVec2f(t, Vec2f{0, 1}, Vec2f{1, 0}.Rotate(pi/2))
	AssertVec2f(t, Vec2f{-1, 0}, Vec2f{0, -1}.Rotate(-pi/2))
}

func TestVec2f_Round(t *testing.T) {
	assert.Equal(t, Vec2i{-6, 7}, Vec2f{-6, 7}.Round())
	assert.Equal(t, Vec2i{-5, 7}, Vec2f{-5.2, 7.2}.Round())
	assert.Equal(t, Vec2i{-6, 8}, Vec2f{-5.8, 7.8}.Round())
	assert.Equal(t, Vec2i{-6, 8}, Vec2f{-5.5, 7.5}.Round())
}

func TestVec2f_Split(t *testing.T) {
	x, y := Vec2f{6, -9}.Split()
	AssertFloat(t, 6, x)
	AssertFloat(t, -9, y)
}

func TestVec2f_SquareDistance(t *testing.T) {
	AssertFloat(t, 0, Vec2f{3, 4}.SquareDistance(Vec2f{3, 4}))
	AssertFloat(t, 8, Vec2f{3, 4}.SquareDistance(Vec2f{1, 2}))
}

func TestVec2f_SquareLength(t *testing.T) {
	AssertFloat(t, 13, Vec2f{2, -3}.SquareLength())
}

func TestVec2f_Sub(t *testing.T) {
	AssertVec2f(t, Vec2f{0, 12}, Vec2f{2, 9}.Sub(Vec2f{2, -3}))
}

func TestVec2f_SubScalar(t *testing.T) {
	AssertVec2f(t, Vec2f{-1, 7}, Vec2f{1, 9}.SubScalar(2))
	AssertVec2f(t, Vec2f{7, 4}, Vec2f{4, 1}.SubScalar(-3))
}

func TestVec2f_Vec2i(t *testing.T) {
	assert.Equal(t, Vec2i{-6, 7}, Vec2f{-6, 7}.Vec2i())
	assert.Equal(t, Vec2i{-5, 7}, Vec2f{-5.8, 7.5}.Vec2i())
}

func TestVec2f_Vec3f(t *testing.T) {
	AssertVec3f(t, Vec3f{1, 2, 3}, Vec2f{1, 2}.Vec3f(3))
}

func TestVec2f_Vec4f(t *testing.T) {
	AssertVec4f(t, Vec4f{1, 2, 3, 4}, Vec2f{1, 2}.Vec4f(3, 4))
}

func TestVec2f_X(t *testing.T) {
	AssertFloat(t, -45, Vec2f{-45, 12}.X())
}

func TestVec2f_Y(t *testing.T) {
	AssertFloat(t, 12, Vec2f{-45, 12}.Y())
}
