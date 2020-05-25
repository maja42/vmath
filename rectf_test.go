package vmath

import (
	"testing"

	"github.com/maja42/vmath/math32"
	"github.com/stretchr/testify/assert"
)

func TestRectfFromCorners(t *testing.T) {
	expected := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}

	assert.Equal(t, expected, RectfFromCorners(Vec2f{-3, 4}, Vec2f{10, 7}))
	assert.Equal(t, expected, RectfFromCorners(Vec2f{10, 7}, Vec2f{-3, 4}))
	assert.Equal(t, expected, RectfFromCorners(Vec2f{10, 4}, Vec2f{-3, 7}))
	assert.Equal(t, expected, RectfFromCorners(Vec2f{-3, 7}, Vec2f{10, 4}))
}

func TestRectfFromPosSize(t *testing.T) {
	expected := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}

	assert.Equal(t, expected, RectfFromPosSize(Vec2f{-3, 4}, Vec2f{13, 3}))
	assert.Equal(t, expected, RectfFromPosSize(Vec2f{10, 7}, Vec2f{-13, -3}))
	assert.Equal(t, expected, RectfFromPosSize(Vec2f{10, 4}, Vec2f{-13, 3}))
	assert.Equal(t, expected, RectfFromPosSize(Vec2f{-3, 7}, Vec2f{13, -3}))
}

func TestRectfFromEdges(t *testing.T) {
	expected := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}

	assert.Equal(t, expected, RectfFromEdges(-3, 10, 4, 7))
	assert.Equal(t, expected, RectfFromEdges(10, -3, 7, 4))
	assert.Equal(t, expected, RectfFromEdges(10, -3, 4, 7))
	assert.Equal(t, expected, RectfFromEdges(-3, 10, 7, 4))
}

func TestRectf_Normalize(t *testing.T) {
	expected := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}

	assert.Equal(t, expected, Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}.Normalize())

	assert.Equal(t, expected, Rectf{
		Min: Vec2f{10, 7},
		Max: Vec2f{-3, 4},
	}.Normalize())

	assert.Equal(t, expected, Rectf{
		Min: Vec2f{10, 4},
		Max: Vec2f{-3, 7},
	}.Normalize())
}

func TestRectf_String(t *testing.T) {
	str := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}.String()
	assert.Equal(t, `Rectf([-3.000000 x 4.000000]-[10.000000 x 7.000000])`, str)
}

func TestRectf_Recti(t *testing.T) {
	rf := Rectf{
		Min: Vec2f{-3.7, 4.5},
		Max: Vec2f{10.2, -7.5},
	}
	ri := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, -7},
	}
	assert.Equal(t, ri, rf.Recti())
}

func TestRectf_Round(t *testing.T) {
	rf := Rectf{
		Min: Vec2f{-3.7, 4.5},
		Max: Vec2f{10.2, -7.5},
	}
	ri := Recti{
		Min: Vec2i{-4, 5},
		Max: Vec2i{10, -8},
	}
	assert.Equal(t, ri, rf.Round())
}

func TestRectf_Size(t *testing.T) {
	r := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}
	assert.Equal(t, Vec2f{13, 3}, r.Size())
}

func TestRectf_Area(t *testing.T) {
	r := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}
	assert.Equal(t, float32(13*3), r.Area())
}

func TestRectf_Edges(t *testing.T) {
	r := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, 7},
	}
	assert.Equal(t, float32(-3), r.Left())
	assert.Equal(t, float32(10), r.Right())
	assert.Equal(t, float32(4), r.Bottom())
	assert.Equal(t, float32(7), r.Top())
}

func TestRectf_SquarePointDistance(t *testing.T) {
	r := Rectf{
		Min: Vec2f{0, 0},
		Max: Vec2f{10, 10},
	}

	// within the rectangle
	assert.Equal(t, float32(0), r.SquarePointDistance(Vec2f{0, 0}))
	assert.Equal(t, float32(0), r.SquarePointDistance(Vec2f{5, 5}))
	assert.Equal(t, float32(0), r.SquarePointDistance(Vec2f{10, 10}))
	assert.Equal(t, float32(0), r.SquarePointDistance(Vec2f{0, 10}))

	// to the left
	assert.Equal(t, float32(25), r.SquarePointDistance(Vec2f{-5, 0}))
	assert.Equal(t, float32(16), r.SquarePointDistance(Vec2f{-4, 10}))
	assert.Equal(t, float32(4), r.SquarePointDistance(Vec2f{-2, 2}))
	assert.Equal(t, float32(9), r.SquarePointDistance(Vec2f{-3, 5}))
	// to the right
	assert.Equal(t, float32(25), r.SquarePointDistance(Vec2f{15, 7.5}))
	// below
	assert.Equal(t, float32(49), r.SquarePointDistance(Vec2f{3, -7}))
	// above
	assert.Equal(t, float32(16), r.SquarePointDistance(Vec2f{8, 14}))

	// below + left
	assert.Equal(t, float32(50), r.SquarePointDistance(Vec2f{-5, -5}))
	assert.Equal(t, float32(68), r.SquarePointDistance(Vec2f{-2, -8}))
	// above + left
	assert.Equal(t, float32(41), r.SquarePointDistance(Vec2f{-4, 15}))
	assert.Equal(t, float32(13), r.SquarePointDistance(Vec2f{-2, 13}))

	// below + right
	assert.Equal(t, float32(34), r.SquarePointDistance(Vec2f{13, -5}))
	assert.Equal(t, float32(53), r.SquarePointDistance(Vec2f{17, -2}))
	// above + right
	assert.Equal(t, float32(37), r.SquarePointDistance(Vec2f{16, 11}))
	assert.Equal(t, float32(85), r.SquarePointDistance(Vec2f{19, 12}))
}

func TestRectf_PointDistance(t *testing.T) {
	r := Rectf{
		Min: Vec2f{0, 0},
		Max: Vec2f{10, 10},
	}
	// within the rectangle
	AssertFloat(t, float32(0), r.PointDistance(Vec2f{5, 5}))
	// to the left
	AssertFloat(t, float32(3), r.PointDistance(Vec2f{-3, 5}))
	// above
	AssertFloat(t, float32(4), r.PointDistance(Vec2f{8, 14}))
	// below + left
	AssertFloat(t, math32.Sqrt(50), r.PointDistance(Vec2f{-5, -5}))
	// below + right
	AssertFloat(t, math32.Sqrt(34), r.PointDistance(Vec2f{13, -5}))
	// above + right
	AssertFloat(t, math32.Sqrt(37), r.PointDistance(Vec2f{16, 11}))
}
