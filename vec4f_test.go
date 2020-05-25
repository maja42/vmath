package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec4f_String(t *testing.T) {
	assert.Equal(t, "Vec4f[5.000000 x -3.100000 x 9.000000 x 2.400000]", Vec4f{5, -3.1, 9, 2.4}.String())
}

func TestVec4f_Format(t *testing.T) {
	assert.Equal(t, "vec 5, -3.2, 7, 0", Vec4f{5, -3.2, 7, 0}.Format("vec %.0f, %.1f, %.f, %.f"))
}

func TestVec4f_Abs(t *testing.T) {
	AssertVec4f(t, Vec4f{2, 9, 4, 8}, Vec4f{2, 9, 4, 8}.Abs())
	AssertVec4f(t, Vec4f{2, 9, 9, 8}, Vec4f{-2, 9, 9, -8}.Abs())
	AssertVec4f(t, Vec4f{2, 9, 5, 8}, Vec4f{2, -9, -5, 8}.Abs())
	AssertVec4f(t, Vec4f{2, 9, 43, 65}, Vec4f{-2, -9, -43, -65}.Abs())
}

func TestVec4f_Add(t *testing.T) {
	AssertVec4f(t, Vec4f{4, 6, 6, -1}, Vec4f{2, 9, 4, -5}.Add(Vec4f{2, -3, 2, 4}))
}

func TestVec4f_IsZero(t *testing.T) {
	assert.True(t, Vec4f{0, 0, 0, 0}.IsZero())
	assert.False(t, Vec4f{0.0001, 0, 0, 0}.IsZero())
	assert.False(t, Vec4f{0, 0.0001, 0, 0}.IsZero())
	assert.False(t, Vec4f{0, 0, 0.0001, 0}.IsZero())
	assert.False(t, Vec4f{0, 0, 0, 0.0001}.IsZero())
}

func TestVec4f_AddScalar(t *testing.T) {
	AssertVec4f(t, Vec4f{4, 11, 3, -4}, Vec4f{2, 9, 1, -6}.AddScalar(2))
	AssertVec4f(t, Vec4f{1, -2, -4, -7}, Vec4f{4, 1, -1, -4}.AddScalar(-3))
}

func TestVec4f_Clamp(t *testing.T) {
	AssertVec4f(t, Vec4f{0, 0, 0, 0}, Vec4f{0, 0, 0, 0}.Clamp(-5, 5))
	AssertVec4f(t, Vec4f{-5, -5, 5, 5}, Vec4f{-8, -9, 5.01, 7}.Clamp(-5, 5))
	AssertVec4f(t, Vec4f{5, 5, -5, -5}, Vec4f{8, 12, -7, -12}.Clamp(-5, 5))
}

func TestVec4f_Distance(t *testing.T) {
	AssertFloat(t, 2, Vec4f{0, 0, 0, 0}.Distance(Vec4f{2, 0, 0, 0}))
	AssertFloat(t, 2, Vec4f{0, 0, 0, 0}.Distance(Vec4f{0, 0, -2, 0}))
	AssertFloat(t, 5, Vec4f{0, 0, 0, 7}.Distance(Vec4f{0, 0, 0, 2}))
	AssertFloat(t, 0, Vec4f{3, -4, 5, 9}.Distance(Vec4f{3, -4, 5, 9}))
	AssertFloat(t, 21.633308, Vec4f{3, -4, 7, 2}.Distance(Vec4f{1, 8, -9, -6}))
}

func TestVec4f_Div(t *testing.T) {
	AssertVec4f(t, Vec4f{3, -8, 7, 11}, Vec4f{12, 56, 42, 88}.Div(Vec4f{4, -7, 6, 8}))
}

func TestVec4f_DivScalar(t *testing.T) {
	AssertVec4f(t, Vec4f{2.5, 7, 10.25, 5.5}, Vec4f{20, 56, 82, 44}.DivScalar(8))
}

func TestVec4f_Dot(t *testing.T) {
	AssertFloat(t, 16, Vec4f{4, 7, 2, 5}.Dot(Vec4f{2, -4, 3, 6}))
}

func TestVec4f_Equal(t *testing.T) {
	assert.True(t, Vec4f{0, 0, 0, 0}.Equal(Vec4f{0, 0, 0, 0}))
	assert.True(t, Vec4f{6, 7, 8, 8}.Equal(Vec4f{6, 7, 8, 8}))
	assert.True(t, Vec4f{-4, -9, -5, -7}.Equal(Vec4f{-4, -9, -5, -7}))

	assert.False(t, Vec4f{6, 7, 8, 9}.Equal(Vec4f{7, 6, 8, 9}))
	assert.False(t, Vec4f{6, 7, 8, 9}.Equal(Vec4f{6, 7, 8, 7}))
}

func TestVec4f_Length(t *testing.T) {
	AssertFloat(t, 5, Vec4f{5, 0, 0, 0}.Length())
	AssertFloat(t, 5, Vec4f{0, -5, 0, 0}.Length())
	AssertFloat(t, 8, Vec4f{0, 0, 8}.Length())
	AssertFloat(t, 4, Vec4f{0, 0, 0, -4}.Length())
	AssertFloat(t, 6.244998, Vec4f{2, 1, 3, 5}.Length())
}

func TestVec4f_Lerp(t *testing.T) {
	AssertVec4f(t, Vec4f{-1, 0, 0, 0}, Vec4f{-1, 0, 0, 0}.Lerp(Vec4f{1, 0, 0, 0}, 0))
	AssertVec4f(t, Vec4f{-0.5, 0, 0, 0}, Vec4f{-1, 0, 0, 0}.Lerp(Vec4f{1, 0, 0, 0}, 0.25))
	AssertVec4f(t, Vec4f{0, 0, 0, 0}, Vec4f{-1, 0, 0, 0}.Lerp(Vec4f{1, 0, 0, 0}, 0.5))
	AssertVec4f(t, Vec4f{0.5, 0, 0, 0}, Vec4f{-1, 0, 0, 0}.Lerp(Vec4f{1, 0, 0, 0}, 0.75))
	AssertVec4f(t, Vec4f{1, 0, 0, 0}, Vec4f{-1, 0, 0, 0}.Lerp(Vec4f{1, 0, 0, 0}, 1))

	AssertVec4f(t, Vec4f{0, 0.3, 0.7, 0}, Vec4f{0, 0, 1, 0}.Lerp(Vec4f{0, 1, 0, 0}, 0.3))
	AssertVec4f(t, Vec4f{0.3, 0, 0, 0.7}, Vec4f{0, 0, 0, 1}.Lerp(Vec4f{1, 0, 0, 0}, 0.3))
}

func TestVec4f_Mul(t *testing.T) {
	AssertVec4f(t, Vec4f{6, -27, 16, 28}, Vec4f{3, 9, 2, 7}.Mul(Vec4f{2, -3, 8, 4}))
}

func TestVec4f_MulScalar(t *testing.T) {
	AssertVec4f(t, Vec4f{8, 18, -2, 6}, Vec4f{4, 9, -1, 3}.MulScalar(2))
	AssertVec4f(t, Vec4f{-12, -3, 6, -24}, Vec4f{4, 1, -2, 8}.MulScalar(-3))
}

func TestVec4f_Negate(t *testing.T) {
	AssertVec4f(t, Vec4f{6, -4, 9, 1}, Vec4f{-6, 4, -9, -1}.Negate())
}

func TestVec4f_Normalize(t *testing.T) {
	AssertVec4f(t, Vec4f{1, 0, 0, 0}, Vec4f{8, 0, 0, 0}.Normalize())
	AssertVec4f(t, Vec4f{0, 1, 0, 0}, Vec4f{0, 23, 0, 0}.Normalize())
	AssertVec4f(t, Vec4f{0, 0, 1, 0}, Vec4f{0, 0, 7, 0}.Normalize())
	AssertVec4f(t, Vec4f{0, 0, 0, 1}, Vec4f{0, 0, 0, 17}.Normalize())

	AssertVec4f(t, Vec4f{-1, 0, 0, 0}, Vec4f{-15, 0, 0, 0}.Normalize())
	AssertVec4f(t, Vec4f{0, -1, 0, 0}, Vec4f{0, -54, 0, 0}.Normalize())
	AssertVec4f(t, Vec4f{0, 0, -1, 0}, Vec4f{0, 0, -7, 0}.Normalize())
	AssertVec4f(t, Vec4f{0, 0, 0, -1}, Vec4f{0, 0, 0, -12}.Normalize())

	AssertVec4f(t, Vec4f{0.730297, 0.365148, 0.182574, -0.547723}, Vec4f{12, 6, 3, -9}.Normalize())
	AssertVec4f(t, Vec4f{0, 0, 0, 0}, Vec4f{0, 0, 0, 0}.Normalize())
}

func TestVec4f_Project(t *testing.T) {
	AssertVec4f(t, Vec4f{0, 0, 0, 0}, Vec4f{0, 1, 0, -2}.Project(Vec4f{1, 0, 0, 0}))
	AssertVec4f(t, Vec4f{1, 0, 0, 0}, Vec4f{1, 1, 1, 1}.Project(Vec4f{1, 0, 0, 0}))
	AssertVec4f(t, Vec4f{1, 0, 0, 0}, Vec4f{1, 2, 3, 4}.Project(Vec4f{2, 0, 0, 0}))

	AssertVec4f(t, Vec4f{0, 0, 9, 0}, Vec4f{1, 2, 9, 3}.Project(Vec4f{0, 0, 4.5, 0}))
	AssertVec4f(t, Vec4f{0, 0, 0, 4.5}, Vec4f{1, 2, 3, 4.5}.Project(Vec4f{0, 0, 0, 9}))

	AssertVec4f(t, Vec4f{0.5, 0.5, 0, 0}, Vec4f{1, 0, 0, 0}.Project(Vec4f{1, 1, 0, 0}))
	AssertVec4f(t, Vec4f{0.5, 0.5, 0, 0}, Vec4f{0, 1, 0, 0}.Project(Vec4f{1, 1, 0, 0}))

	AssertVec4f(t, Vec4f{0.25, 0.25, 0.25, 0.25}, Vec4f{0, 0, 1, 0}.Project(Vec4f{1, 1, 1, 1}))
	AssertVec4f(t, Vec4f{0.25, 0.25, 0.25, 0.25}, Vec4f{0, 0, 0, 1}.Project(Vec4f{1, 1, 1, 1}))
	AssertVec4f(t, Vec4f{0.5, 0.5, 0.5, 0.5}, Vec4f{1, 1, 0, 0}.Project(Vec4f{1, 1, 1, 1}))

	AssertVec4f(t, Vec4f{0, 0, 0, 0}, Vec4f{1, 0, 0, 0}.Project(Vec4f{0, 1, 1, 0}))
	AssertVec4f(t, Vec4f{0, 0.5, 0.5, 0}, Vec4f{0, 1, 0, 0}.Project(Vec4f{0, 1, 1, 0}))
	AssertVec4f(t, Vec4f{0, 0.5, 0.5, 0}, Vec4f{2, 0, 1, 0}.Project(Vec4f{0, 1, 1, 0}))
}

func TestVec4f_Round(t *testing.T) {
	assert.Equal(t, Vec4i{-6, 7, 5, -3}, Vec4f{-6, 7, 5, -3}.Round())
	assert.Equal(t, Vec4i{-5, 7, 5, -3}, Vec4f{-5.2, 7.2, 5.2, -3.3}.Round())
	assert.Equal(t, Vec4i{-6, 8, 10, 5}, Vec4f{-5.8, 7.8, 9.7, 4.6}.Round())
	assert.Equal(t, Vec4i{-6, 8, 4, 5}, Vec4f{-5.5, 7.5, 3.5, 4.5}.Round())
}

func TestVec4f_Split(t *testing.T) {
	x, y, z, w := Vec4f{6, -9, 12, -2}.Split()
	AssertFloat(t, 6, x)
	AssertFloat(t, -9, y)
	AssertFloat(t, 12, z)
	AssertFloat(t, -2, w)
}

func TestVec4f_SquareDistance(t *testing.T) {
	AssertFloat(t, 0, Vec4f{3, 4, 5, 6}.SquareDistance(Vec4f{3, 4, 5, 6}))
	AssertFloat(t, 16, Vec4f{3, 4, 5, 6}.SquareDistance(Vec4f{1, 2, 3, 4}))
}

func TestVec4f_SquareLength(t *testing.T) {
	AssertFloat(t, 16, Vec4f{4, 0, 0, 0}.SquareLength())
	AssertFloat(t, 16, Vec4f{0, 4, 0, 0}.SquareLength())
	AssertFloat(t, 16, Vec4f{0, 0, 4, 0}.SquareLength())
	AssertFloat(t, 16, Vec4f{0, 0, 0, 4}.SquareLength())

	AssertFloat(t, 54, Vec4f{2, -3, 4, 5}.SquareLength())

}

func TestVec4f_Sub(t *testing.T) {
	AssertVec4f(t, Vec4f{0, 12, 2, -2}, Vec4f{2, 9, 7, -4}.Sub(Vec4f{2, -3, 5, -2}))
}

func TestVec4f_SubScalar(t *testing.T) {
	AssertVec4f(t, Vec4f{-1, 7, -7, 4}, Vec4f{1, 9, -5, 6}.SubScalar(2))
	AssertVec4f(t, Vec4f{7, 4, -1, 2}, Vec4f{4, 1, -4, -1}.SubScalar(-3))
}

func TestVec4f_Vec3i(t *testing.T) {
	assert.Equal(t, Vec4i{-6, 7, 9, -2}, Vec4f{-6, 7, 9, -2}.Vec4i())
	assert.Equal(t, Vec4i{-5, 7, -9, 12}, Vec4f{-5.8, 7.5, -9.5, 12.9}.Vec4i())
}

func TestVec4f_X(t *testing.T) {
	AssertFloat(t, -45, Vec4f{-45, 12, 7, -3}.X())
}

func TestVec4f_Y(t *testing.T) {
	AssertFloat(t, 12, Vec4f{-45, 12, 7, -3}.Y())
}

func TestVec4f_Z(t *testing.T) {
	AssertFloat(t, 7, Vec4f{-45, 12, 7, -3}.Z())
}

func TestVec4f_W(t *testing.T) {
	AssertFloat(t, -3, Vec4f{-45, 12, 7, -3}.W())
}

func TestVec4f_XY(t *testing.T) {
	AssertVec2f(t, Vec2f{-45, 12}, Vec4f{-45, 12, 7, -3}.XY())
}

func TestVec4f_XYZ(t *testing.T) {
	AssertVec3f(t, Vec3f{-45, 12, 7}, Vec4f{-45, 12, 7, -3}.XYZ())
}
