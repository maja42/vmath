package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec2i_String(t *testing.T) {
	assert.Equal(t, "Vec2i[5 x -3]", Vec2i{5, -3}.String())
}

func TestVec2i_Format(t *testing.T) {
	assert.Equal(t, "vec 5, -03", Vec2i{5, -3}.Format("vec %d, %03d"))
}

func TestVec2i_IsOrthogonal(t *testing.T) {
	assert.True(t, Vec2i{2, 0}.IsOrthogonal())
	assert.True(t, Vec2i{-2, 0}.IsOrthogonal())
	assert.True(t, Vec2i{0, 32}.IsOrthogonal())
	assert.True(t, Vec2i{0, -12}.IsOrthogonal())

	assert.False(t, Vec2i{2, 2}.IsOrthogonal())
}

func TestVec2i_Abs(t *testing.T) {
	assert.Equal(t, Vec2i{2, 9}, Vec2i{2, 9}.Abs())
	assert.Equal(t, Vec2i{2, 9}, Vec2i{-2, 9}.Abs())
	assert.Equal(t, Vec2i{2, 9}, Vec2i{2, -9}.Abs())
	assert.Equal(t, Vec2i{2, 9}, Vec2i{-2, -9}.Abs())
}

func TestVec2i_Add(t *testing.T) {
	assert.Equal(t, Vec2i{4, 6}, Vec2i{2, 9}.Add(Vec2i{2, -3}))
}

func TestVec2i_IsZero(t *testing.T) {
	assert.True(t, Vec2i{0, 0}.IsZero())
	assert.False(t, Vec2i{1, 0}.IsZero())
	assert.False(t, Vec2i{0, 1}.IsZero())
}

func TestVec2i_AddScalar(t *testing.T) {
	assert.Equal(t, Vec2i{4, 11}, Vec2i{2, 9}.AddScalar(2))
	assert.Equal(t, Vec2i{1, -2}, Vec2i{4, 1}.AddScalar(-3))
}

func TestVec2i_AddScalarf(t *testing.T) {
	assert.Equal(t, Vec2f{4.5, 11.5}, Vec2i{2, 9}.AddScalarf(2.5))
	assert.Equal(t, Vec2f{0.5, -2.5}, Vec2i{4, 1}.AddScalarf(-3.5))
}

func TestVec2i_Clamp(t *testing.T) {
	assert.Equal(t, Vec2i{0, 0}, Vec2i{0, 0}.Clamp(-5, 5))
	assert.Equal(t, Vec2i{-5, -5}, Vec2i{-8, -9}.Clamp(-5, 5))
	assert.Equal(t, Vec2i{5, 5}, Vec2i{8, 12}.Clamp(-5, 5))
}

func TestVec2i_Distance(t *testing.T) {
	AssertFloat(t, 2, Vec2i{0, 0}.Distance(Vec2i{2, 0}))
	AssertFloat(t, 0, Vec2i{3, -4}.Distance(Vec2i{3, -4}))
	AssertFloat(t, 12.165525, Vec2i{3, -4}.Distance(Vec2i{1, 8}))
}

func TestVec2i_Div(t *testing.T) {
	assert.Equal(t, Vec2i{3, -8}, Vec2i{12, 56}.Div(Vec2i{4, -7}))
	assert.Equal(t, Vec2i{2, 2}, Vec2i{12, 12}.Div(Vec2i{5, 5}))
	assert.Equal(t, Vec2i{2, 2}, Vec2i{-12, -12}.Div(Vec2i{-5, -5}))
}

func TestVec2i_DivScalar(t *testing.T) {
	assert.Equal(t, Vec2i{2, 7}, Vec2i{20, 56}.DivScalar(8))
}

func TestVec2i_DivScalarf(t *testing.T) {
	assert.Equal(t, Vec2f{2.5, -7}, Vec2i{20, -56}.DivScalarf(8))
}

func TestVec2i_Dot(t *testing.T) {
	assert.Equal(t, -20, Vec2i{4, 7}.Dot(Vec2i{2, -4}))
}

func TestVec2i_Equal(t *testing.T) {
	assert.True(t, Vec2i{0, 0}.Equal(Vec2i{0, 0}))
	assert.True(t, Vec2i{6, 7}.Equal(Vec2i{6, 7}))
	assert.True(t, Vec2i{-4, -9}.Equal(Vec2i{-4, -9}))

	assert.False(t, Vec2i{6, 7}.Equal(Vec2i{7, 6}))
}

func TestVec2i_IsParallel(t *testing.T) {
	assert.True(t, Vec2i{1, 1}.IsParallel(Vec2i{5, 5}))
	assert.True(t, Vec2i{1, 1}.IsParallel(Vec2i{-1, -1}))
	assert.True(t, Vec2i{1, 0}.IsParallel(Vec2i{1, 0}))
	assert.False(t, Vec2i{1, 1}.IsParallel(Vec2i{-1, 2}))
}

func TestVec2i_IsCollinear(t *testing.T) {
	assert.True(t, Vec2i{1, 1}.IsCollinear(Vec2i{5, 5}))
	assert.False(t, Vec2i{1, 1}.IsCollinear(Vec2i{-1, -1}))

	assert.True(t, Vec2i{0, 1}.IsCollinear(Vec2i{0, 1}))
	assert.False(t, Vec2i{0, 1}.IsCollinear(Vec2i{0, -1}))

	assert.False(t, Vec2i{1, 1}.IsCollinear(Vec2i{-1, 2}))
}

func TestVec2i_Length(t *testing.T) {
	AssertFloat(t, 5, Vec2i{5, 0}.Length())
	AssertFloat(t, 5, Vec2i{0, -5}.Length())
	AssertFloat(t, 2.236068, Vec2i{2, 1}.Length())
}

func TestVec2i_MagCross(t *testing.T) {
	assert.Equal(t, -1, Vec2i{0, 1}.MagCross(Vec2i{1, 0}))
	assert.Equal(t, -6, Vec2i{0, 2}.MagCross(Vec2i{3, 0}))
	assert.Equal(t, 47, Vec2i{5, -3}.MagCross(Vec2i{4, 7}))
}

func TestVec2i_Mul(t *testing.T) {
	assert.Equal(t, Vec2i{6, -27}, Vec2i{3, 9}.Mul(Vec2i{2, -3}))
}

func TestVec2i_MulScalar(t *testing.T) {
	assert.Equal(t, Vec2i{8, 18}, Vec2i{4, 9}.MulScalar(2))
	assert.Equal(t, Vec2i{-12, -3}, Vec2i{4, 1}.MulScalar(-3))
}

func TestVec2i_MulScalarf(t *testing.T) {
	assert.Equal(t, Vec2f{10, 22.5}, Vec2i{4, 9}.MulScalarf(2.5))
	assert.Equal(t, Vec2f{-14, -3.5}, Vec2i{4, 1}.MulScalarf(-3.5))
}

func TestVec2i_Negate(t *testing.T) {
	assert.Equal(t, Vec2i{6, -4}, Vec2i{-6, 4}.Negate())
}

func TestVec2i_NormalVec(t *testing.T) {
	assert.Equal(t, Vec2i{-1, 1}, Vec2i{1, 1}.NormalVec(true))
	assert.Equal(t, Vec2i{1, -1}, Vec2i{1, 1}.NormalVec(false))

	assert.Equal(t, Vec2i{0, -1}, Vec2i{-1, 0}.NormalVec(true))
	assert.Equal(t, Vec2i{0, 1}, Vec2i{-1, 0}.NormalVec(false))
}

func TestVec2i_Project(t *testing.T) {
	assert.Equal(t, Vec2f{0, 0}, Vec2i{0, 1}.Project(Vec2i{1, 0}))
	assert.Equal(t, Vec2f{1, 0}, Vec2i{1, 1}.Project(Vec2i{1, 0}))
	assert.Equal(t, Vec2f{1, 0}, Vec2i{1, 2}.Project(Vec2i{2, 0}))
	assert.Equal(t, Vec2f{0.5, 0.5}, Vec2i{0, 1}.Project(Vec2i{1, 1}))
	assert.Equal(t, Vec2f{0.5, 0.5}, Vec2i{1, 0}.Project(Vec2i{1, 1}))

}

func TestVec2i_Split(t *testing.T) {
	x, y := Vec2i{6, -9}.Split()
	assert.Equal(t, 6, x)
	assert.Equal(t, -9, y)
}

func TestVec2i_SquareDistance(t *testing.T) {
	assert.Equal(t, 0, Vec2i{3, 4}.SquareDistance(Vec2i{3, 4}))
	assert.Equal(t, 8, Vec2i{3, 4}.SquareDistance(Vec2i{1, 2}))
}

func TestVec2i_SquareLength(t *testing.T) {
	assert.Equal(t, 13, Vec2i{2, -3}.SquareLength())
}

func TestVec2i_Sub(t *testing.T) {
	assert.Equal(t, Vec2i{0, 12}, Vec2i{2, 9}.Sub(Vec2i{2, -3}))
}

func TestVec2i_SubScalar(t *testing.T) {
	assert.Equal(t, Vec2i{-1, 7}, Vec2i{1, 9}.SubScalar(2))
	assert.Equal(t, Vec2i{7, 4}, Vec2i{4, 1}.SubScalar(-3))
}

func TestVec2i_SubScalarf(t *testing.T) {
	assert.Equal(t, Vec2f{-1.5, 6.5}, Vec2i{1, 9}.SubScalarf(2.5))
	assert.Equal(t, Vec2f{7.5, 4.5}, Vec2i{4, 1}.SubScalarf(-3.5))
}

func TestVec2i_Vec2f(t *testing.T) {
	assert.Equal(t, Vec2f{-6, 7}, Vec2i{-6, 7}.Vec2f())
	assert.Equal(t, Vec2f{5, -42}, Vec2i{5, -42}.Vec2f())
}

func TestVec2i_Vec3i(t *testing.T) {
	assert.Equal(t, Vec3i{1, 2, 3}, Vec2i{1, 2}.Vec3i(3))
}

func TestVec2i_Vec4f(t *testing.T) {
	assert.Equal(t, Vec4i{1, 2, 3, 4}, Vec2i{1, 2}.Vec4i(3, 4))
}

func TestVec2i_X(t *testing.T) {
	assert.Equal(t, -45, Vec2i{-45, 12}.X())
}

func TestVec2i_Y(t *testing.T) {
	assert.Equal(t, 12, Vec2i{-45, 12}.Y())
}
