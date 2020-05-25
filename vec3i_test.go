package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3i_String(t *testing.T) {
	assert.Equal(t, "Vec3i[5 x -3 x 9]", Vec3i{5, -3, 9}.String())
}

func TestVec3i_Format(t *testing.T) {
	assert.Equal(t, "vec 5, -03, 7", Vec3i{5, -3, 7}.Format("vec %d, %03d, %d"))
}

func TestVec3i_IsOrthogonal(t *testing.T) {
	assert.True(t, Vec3i{2, 0, 8}.IsOrthogonal())
	assert.True(t, Vec3i{-2, 0, 8}.IsOrthogonal())
	assert.True(t, Vec3i{0, 32, 8}.IsOrthogonal())
	assert.True(t, Vec3i{0, -12, -7}.IsOrthogonal())
	assert.True(t, Vec3i{32, 8, 0}.IsOrthogonal())
	assert.True(t, Vec3i{-12, -7, 0}.IsOrthogonal())

	assert.False(t, Vec3i{2, 2, 5}.IsOrthogonal())
	assert.False(t, Vec3i{-1, 1, -3}.IsOrthogonal())
}

func TestVec3i_Abs(t *testing.T) {
	assert.Equal(t, Vec3i{2, 9, 4}, Vec3i{2, 9, 4}.Abs())
	assert.Equal(t, Vec3i{2, 9, 9}, Vec3i{-2, 9, 9}.Abs())
	assert.Equal(t, Vec3i{2, 9, 5}, Vec3i{2, -9, -5}.Abs())
	assert.Equal(t, Vec3i{2, 9, 43}, Vec3i{-2, -9, -43}.Abs())
}

func TestVec3i_Add(t *testing.T) {
	assert.Equal(t, Vec3i{4, 6, 6}, Vec3i{2, 9, 4}.Add(Vec3i{2, -3, 2}))
}

func TestVec3i_IsZero(t *testing.T) {
	assert.True(t, Vec3i{0, 0, 0}.IsZero())
	assert.False(t, Vec3i{1, 0, 0}.IsZero())
	assert.False(t, Vec3i{0, 1, 0}.IsZero())
	assert.False(t, Vec3i{0, 0, 1}.IsZero())
}

func TestVec3i_AddScalar(t *testing.T) {
	assert.Equal(t, Vec3i{4, 11, 3}, Vec3i{2, 9, 1}.AddScalar(2))
	assert.Equal(t, Vec3i{1, -2, -4}, Vec3i{4, 1, -1}.AddScalar(-3))
}

func TestVec3i_AddScalarf(t *testing.T) {
	AssertVec3f(t, Vec3f{4.5, 11.5, 3.5}, Vec3i{2, 9, 1}.AddScalarf(2.5))
	AssertVec3f(t, Vec3f{0.5, -2.5, -4.5}, Vec3i{4, 1, -1}.AddScalarf(-3.5))
}

func TestVec3i_Clamp(t *testing.T) {
	assert.Equal(t, Vec3i{0, 0, 0}, Vec3i{0, 0, 0}.Clamp(-5, 5))
	assert.Equal(t, Vec3i{-5, -5, 5}, Vec3i{-8, -9, 6}.Clamp(-5, 5))
	assert.Equal(t, Vec3i{5, 5, -5}, Vec3i{8, 12, -7}.Clamp(-5, 5))
}

func TestVec3i_Distance(t *testing.T) {
	AssertFloat(t, 2, Vec3i{0, 0, 0}.Distance(Vec3i{2, 0, 0}))
	AssertFloat(t, 2, Vec3i{0, 0, 0}.Distance(Vec3i{0, 0, -2}))
	AssertFloat(t, 0, Vec3i{3, -4, 5}.Distance(Vec3i{3, -4, 5}))
	AssertFloat(t, 20.09975, Vec3i{3, -4, 7}.Distance(Vec3i{1, 8, -9}))
}

func TestVec3i_Div(t *testing.T) {
	assert.Equal(t, Vec3i{3, -8, 7}, Vec3i{12, 56, 42}.Div(Vec3i{4, -7, 6}))
}

func TestVec3i_DivScalar(t *testing.T) {
	assert.Equal(t, Vec3i{2, 7, 10}, Vec3i{20, 56, 82}.DivScalar(8))
}

func TestVec3i_DivScalarf(t *testing.T) {
	AssertVec3f(t, Vec3f{6, 4, 8}, Vec3i{9, 6, 12}.DivScalarf(1.5))
}

func TestVec3i_Dot(t *testing.T) {
	assert.Equal(t, -14, Vec3i{4, 7, 2}.Dot(Vec3i{2, -4, 3}))
}

func TestVec3i_Equal(t *testing.T) {
	assert.True(t, Vec3i{0, 0, 0}.Equal(Vec3i{0, 0, 0}))
	assert.True(t, Vec3i{6, 7, 8}.Equal(Vec3i{6, 7, 8}))
	assert.True(t, Vec3i{-4, -9, -5}.Equal(Vec3i{-4, -9, -5}))

	assert.False(t, Vec3i{6, 7, 8}.Equal(Vec3i{7, 6, 8}))
	assert.False(t, Vec3i{6, 7, 8}.Equal(Vec3i{6, 7, 6}))
}

func TestVec3i_IsParallel(t *testing.T) {
	assert.True(t, Vec3i{1, 1, 1}.IsParallel(Vec3i{5, 5, 5}))
	assert.True(t, Vec3i{1, 1, 1}.IsParallel(Vec3i{-1, -1, -1}))
	assert.True(t, Vec3i{2, 1, 0}.IsParallel(Vec3i{2, 1, 0}))
	assert.False(t, Vec3i{1, 1, 1}.IsParallel(Vec3i{-1, 2, 1}))
}

func TestVec3i_IsCollinear(t *testing.T) {
	assert.True(t, Vec3i{1, 1, 1}.IsCollinear(Vec3i{5, 5, 5}))
	assert.False(t, Vec3i{1, 1, 1}.IsCollinear(Vec3i{-1, -1, -1}))

	assert.True(t, Vec3i{0, 1, 0}.IsCollinear(Vec3i{0, 1, 0}))
	assert.False(t, Vec3i{0, 1, 0}.IsCollinear(Vec3i{0, -1, 0}))

	assert.False(t, Vec3i{1, 1, 1}.IsCollinear(Vec3i{-1, 2, 1}))
}

func TestVec3i_Length(t *testing.T) {
	AssertFloat(t, 5, Vec3i{5, 0, 0}.Length())
	AssertFloat(t, 5, Vec3i{0, -5, 0}.Length())
	AssertFloat(t, 8, Vec3i{0, 0, 8}.Length())
	AssertFloat(t, 3.7416575, Vec3i{2, 1, 3}.Length())
}

func TestVec3i_Mul(t *testing.T) {
	assert.Equal(t, Vec3i{6, -27, 16}, Vec3i{3, 9, 2}.Mul(Vec3i{2, -3, 8}))
}

func TestVec3i_MulScalar(t *testing.T) {
	assert.Equal(t, Vec3i{8, 18, -2}, Vec3i{4, 9, -1}.MulScalar(2))
	assert.Equal(t, Vec3i{-12, -3, 6}, Vec3i{4, 1, -2}.MulScalar(-3))
}

func TestVec3i_MulScalarf(t *testing.T) {
	AssertVec3f(t, Vec3f{10, 22.5, -2.5}, Vec3i{4, 9, -1}.MulScalarf(2.5))
	AssertVec3f(t, Vec3f{-14, -3.5, 7}, Vec3i{4, 1, -2}.MulScalarf(-3.5))
}

func TestVec3i_Negate(t *testing.T) {
	assert.Equal(t, Vec3i{6, -4, 9}, Vec3i{-6, 4, -9}.Negate())
}

func TestVec3i_Split(t *testing.T) {
	x, y, z := Vec3i{6, -9, 12}.Split()
	assert.Equal(t, 6, x)
	assert.Equal(t, -9, y)
	assert.Equal(t, 12, z)
}

func TestVec3i_SquareDistance(t *testing.T) {
	assert.Equal(t, 0, Vec3i{3, 4, 5}.SquareDistance(Vec3i{3, 4, 5}))
	assert.Equal(t, 12, Vec3i{3, 4, 5}.SquareDistance(Vec3i{1, 2, 3}))
}

func TestVec3i_SquareLength(t *testing.T) {
	assert.Equal(t, 16, Vec3i{4, 0, 0}.SquareLength())
	assert.Equal(t, 16, Vec3i{0, 4, 0}.SquareLength())
	assert.Equal(t, 16, Vec3i{0, 0, 4}.SquareLength())
	assert.Equal(t, 29, Vec3i{2, -3, 4}.SquareLength())
}

func TestVec3i_Sub(t *testing.T) {
	assert.Equal(t, Vec3i{0, 12, 2}, Vec3i{2, 9, 7}.Sub(Vec3i{2, -3, 5}))
}

func TestVec3i_SubScalar(t *testing.T) {
	assert.Equal(t, Vec3i{-1, 7, -7}, Vec3i{1, 9, -5}.SubScalar(2))
	assert.Equal(t, Vec3i{7, 4, -1}, Vec3i{4, 1, -4}.SubScalar(-3))
}

func TestVec4i_SubScalarf(t *testing.T) {
	AssertVec3f(t, Vec3f{-1.5, 6.5, -7.5}, Vec3i{1, 9, -5}.SubScalarf(2.5))
	AssertVec3f(t, Vec3f{7.5, 4.5, -0.5}, Vec3i{4, 1, -4}.SubScalarf(-3.5))
}

func TestVec3i_Vec3f(t *testing.T) {
	assert.Equal(t, Vec3f{-6, 7, 9}, Vec3i{-6, 7, 9}.Vec3f())
	assert.Equal(t, Vec3f{-5, 7, -9}, Vec3i{-5, 7, -9}.Vec3f())
}

func TestVec3i_Vec4i(t *testing.T) {
	assert.Equal(t, Vec4i{1, 2, 3, 4}, Vec3i{1, 2, 3}.Vec4i(4))
}

func TestVec3i_X(t *testing.T) {
	assert.Equal(t, -45, Vec3i{-45, 12, 7}.X())
}

func TestVec3i_Y(t *testing.T) {
	assert.Equal(t, 12, Vec3i{-45, 12, 7}.Y())
}

func TestVec3i_Z(t *testing.T) {
	assert.Equal(t, 7, Vec3i{-45, 12, 7}.Z())
}

func TestVec3i_XY(t *testing.T) {
	assert.Equal(t, Vec2i{-45, 12}, Vec3i{-45, 12, 7}.XY())
}
