package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec4i_String(t *testing.T) {
	assert.Equal(t, "Vec4i[5 x -3 x 9 x 2]", Vec4i{5, -3, 9, 2}.String())
}

func TestVec4i_Format(t *testing.T) {
	assert.Equal(t, "vec 5, -003, 7, 0", Vec4i{5, -3, 7, 0}.Format("vec %d, %04d, %d, %d"))
}

func TestVec4i_Abs(t *testing.T) {
	assert.Equal(t, Vec4i{2, 9, 4, 8}, Vec4i{2, 9, 4, 8}.Abs())
	assert.Equal(t, Vec4i{2, 9, 9, 8}, Vec4i{-2, 9, 9, -8}.Abs())
	assert.Equal(t, Vec4i{2, 9, 5, 8}, Vec4i{2, -9, -5, 8}.Abs())
	assert.Equal(t, Vec4i{2, 9, 43, 65}, Vec4i{-2, -9, -43, -65}.Abs())
}

func TestVec4i_Add(t *testing.T) {
	assert.Equal(t, Vec4i{4, 6, 6, -1}, Vec4i{2, 9, 4, -5}.Add(Vec4i{2, -3, 2, 4}))
}

func TestVec4i_IsZero(t *testing.T) {
	assert.True(t, Vec4i{0, 0, 0, 0}.IsZero())
	assert.False(t, Vec4i{1, 0, 0, 0}.IsZero())
	assert.False(t, Vec4i{0, 1, 0, 0}.IsZero())
	assert.False(t, Vec4i{0, 0, 1, 0}.IsZero())
	assert.False(t, Vec4i{0, 0, 0, 1}.IsZero())
}

func TestVec4i_AddScalar(t *testing.T) {
	assert.Equal(t, Vec4i{4, 11, 3, -4}, Vec4i{2, 9, 1, -6}.AddScalar(2))
	assert.Equal(t, Vec4i{1, -2, -4, -7}, Vec4i{4, 1, -1, -4}.AddScalar(-3))
}

func TestVec4i_Clamp(t *testing.T) {
	assert.Equal(t, Vec4i{0, 0, 0, 0}, Vec4i{0, 0, 0, 0}.Clamp(-5, 5))
	assert.Equal(t, Vec4i{-5, -5, 5, 5}, Vec4i{-8, -9, 6, 7}.Clamp(-5, 5))
	assert.Equal(t, Vec4i{5, 5, -5, -5}, Vec4i{8, 12, -7, -12}.Clamp(-5, 5))
}

func TestVec4i_Distance(t *testing.T) {
	AssertFloat(t, 2, Vec4i{0, 0, 0, 0}.Distance(Vec4i{2, 0, 0, 0}))
	AssertFloat(t, 2, Vec4i{0, 0, 0, 0}.Distance(Vec4i{0, 0, -2, 0}))
	AssertFloat(t, 5, Vec4i{0, 0, 0, 7}.Distance(Vec4i{0, 0, 0, 2}))
	AssertFloat(t, 0, Vec4i{3, -4, 5, 9}.Distance(Vec4i{3, -4, 5, 9}))
	AssertFloat(t, 21.633308, Vec4i{3, -4, 7, 2}.Distance(Vec4i{1, 8, -9, -6}))
}

func TestVec4i_Div(t *testing.T) {
	assert.Equal(t, Vec4i{3, -8, 7, 11}, Vec4i{12, 56, 42, 88}.Div(Vec4i{4, -7, 6, 8}))
}

func TestVec4i_DivScalar(t *testing.T) {
	assert.Equal(t, Vec4i{2, 7, 10, 5}, Vec4i{20, 56, 82, 44}.DivScalar(8))
}

func TestVec4i_Dot(t *testing.T) {
	assert.Equal(t, 16, Vec4i{4, 7, 2, 5}.Dot(Vec4i{2, -4, 3, 6}))
}

func TestVec4i_Equal(t *testing.T) {
	assert.True(t, Vec4i{0, 0, 0, 0}.Equal(Vec4i{0, 0, 0, 0}))
	assert.True(t, Vec4i{6, 7, 8, 8}.Equal(Vec4i{6, 7, 8, 8}))
	assert.True(t, Vec4i{-4, -9, -5, -7}.Equal(Vec4i{-4, -9, -5, -7}))

	assert.False(t, Vec4i{6, 7, 8, 9}.Equal(Vec4i{7, 6, 8, 9}))
	assert.False(t, Vec4i{6, 7, 8, 9}.Equal(Vec4i{6, 7, 8, 7}))
}

func TestVec4i_Length(t *testing.T) {
	AssertFloat(t, 5, Vec4i{5, 0, 0, 0}.Length())
	AssertFloat(t, 5, Vec4i{0, -5, 0, 0}.Length())
	AssertFloat(t, 8, Vec4i{0, 0, 8}.Length())
	AssertFloat(t, 4, Vec4i{0, 0, 0, -4}.Length())
	AssertFloat(t, 6.244998, Vec4i{2, 1, 3, 5}.Length())
}

func TestVec4i_Mul(t *testing.T) {
	assert.Equal(t, Vec4i{6, -27, 16, 28}, Vec4i{3, 9, 2, 7}.Mul(Vec4i{2, -3, 8, 4}))
}

func TestVec4i_MulScalar(t *testing.T) {
	assert.Equal(t, Vec4i{8, 18, -2, 6}, Vec4i{4, 9, -1, 3}.MulScalar(2))
	assert.Equal(t, Vec4i{-12, -3, 6, -24}, Vec4i{4, 1, -2, 8}.MulScalar(-3))
}

func TestVec4i_Negate(t *testing.T) {
	assert.Equal(t, Vec4i{6, -4, 9, 1}, Vec4i{-6, 4, -9, -1}.Negate())
}

func TestVec4i_Split(t *testing.T) {
	x, y, z, w := Vec4i{6, -9, 12, -2}.Split()
	assert.Equal(t, 6, x)
	assert.Equal(t, -9, y)
	assert.Equal(t, 12, z)
	assert.Equal(t, -2, w)
}

func TestVec4i_SquareDistance(t *testing.T) {
	assert.Equal(t, 0, Vec4i{3, 4, 5, 6}.SquareDistance(Vec4i{3, 4, 5, 6}))
	assert.Equal(t, 16, Vec4i{3, 4, 5, 6}.SquareDistance(Vec4i{1, 2, 3, 4}))
}

func TestVec4i_SquareLength(t *testing.T) {
	assert.Equal(t, 16, Vec4i{4, 0, 0, 0}.SquareLength())
	assert.Equal(t, 16, Vec4i{0, 4, 0, 0}.SquareLength())
	assert.Equal(t, 16, Vec4i{0, 0, 4, 0}.SquareLength())
	assert.Equal(t, 16, Vec4i{0, 0, 0, 4}.SquareLength())

	assert.Equal(t, 54, Vec4i{2, -3, 4, 5}.SquareLength())
}

func TestVec4i_Sub(t *testing.T) {
	assert.Equal(t, Vec4i{0, 12, 2, -2}, Vec4i{2, 9, 7, -4}.Sub(Vec4i{2, -3, 5, -2}))
}

func TestVec4i_SubScalar(t *testing.T) {
	assert.Equal(t, Vec4i{-1, 7, -7, 4}, Vec4i{1, 9, -5, 6}.SubScalar(2))
	assert.Equal(t, Vec4i{7, 4, -1, 2}, Vec4i{4, 1, -4, -1}.SubScalar(-3))
}

func TestVec4i_Vec3i(t *testing.T) {
	assert.Equal(t, Vec4f{-6, 7, 9, -2}, Vec4i{-6, 7, 9, -2}.Vec4f())
	assert.Equal(t, Vec4f{-5, 7, -9, 12}, Vec4i{-5, 7, -9, 12}.Vec4f())
}

func TestVec4i_X(t *testing.T) {
	assert.Equal(t, -45, Vec4i{-45, 12, 7, -3}.X())
}

func TestVec4i_Y(t *testing.T) {
	assert.Equal(t, 12, Vec4i{-45, 12, 7, -3}.Y())
}

func TestVec4i_Z(t *testing.T) {
	assert.Equal(t, 7, Vec4i{-45, 12, 7, -3}.Z())
}

func TestVec4i_W(t *testing.T) {
	assert.Equal(t, -3, Vec4i{-45, 12, 7, -3}.W())
}

func TestVec4i_XY(t *testing.T) {
	assert.Equal(t, Vec2i{-45, 12}, Vec4i{-45, 12, 7, -3}.XY())
}

func TestVec4i_XYZ(t *testing.T) {
	assert.Equal(t, Vec3i{-45, 12, 7}, Vec4i{-45, 12, 7, -3}.XYZ())
}
