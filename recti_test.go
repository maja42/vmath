package vmath

import (
	"testing"

	"github.com/maja42/vmath/math32"
	"github.com/stretchr/testify/assert"
)

func TestRectiFromCorners(t *testing.T) {
	expected := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}

	assert.Equal(t, expected, RectiFromCorners(Vec2i{-3, 4}, Vec2i{10, 7}))
	assert.Equal(t, expected, RectiFromCorners(Vec2i{10, 7}, Vec2i{-3, 4}))
	assert.Equal(t, expected, RectiFromCorners(Vec2i{10, 4}, Vec2i{-3, 7}))
	assert.Equal(t, expected, RectiFromCorners(Vec2i{-3, 7}, Vec2i{10, 4}))
}

func TestRectiFromPosSize(t *testing.T) {
	expected := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}

	assert.Equal(t, expected, RectiFromPosSize(Vec2i{-3, 4}, Vec2i{13, 3}))
	assert.Equal(t, expected, RectiFromPosSize(Vec2i{10, 7}, Vec2i{-13, -3}))
	assert.Equal(t, expected, RectiFromPosSize(Vec2i{10, 4}, Vec2i{-13, 3}))
	assert.Equal(t, expected, RectiFromPosSize(Vec2i{-3, 7}, Vec2i{13, -3}))
}

func TestRectiFromEdges(t *testing.T) {
	expected := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}

	assert.Equal(t, expected, RectiFromEdges(-3, 10, 4, 7))
	assert.Equal(t, expected, RectiFromEdges(10, -3, 7, 4))
	assert.Equal(t, expected, RectiFromEdges(10, -3, 4, 7))
	assert.Equal(t, expected, RectiFromEdges(-3, 10, 7, 4))
}

func TestRecti_Normalize(t *testing.T) {
	expected := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}

	assert.Equal(t, expected, Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}.Normalize())

	assert.Equal(t, expected, Recti{
		Min: Vec2i{10, 7},
		Max: Vec2i{-3, 4},
	}.Normalize())

	assert.Equal(t, expected, Recti{
		Min: Vec2i{10, 4},
		Max: Vec2i{-3, 7},
	}.Normalize())
}

func TestRecti_String(t *testing.T) {
	str := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}.String()
	assert.Equal(t, `Recti([-3 x 4]-[10 x 7])`, str)
}

func TestRecti_Rectf(t *testing.T) {
	ri := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, -7},
	}
	rf := Rectf{
		Min: Vec2f{-3, 4},
		Max: Vec2f{10, -7},
	}
	assert.Equal(t, rf, ri.Rectf())
}

func TestRecti_Size(t *testing.T) {
	r := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}
	assert.Equal(t, Vec2i{13, 3}, r.Size())
}

func TestRecti_Area(t *testing.T) {
	r := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}
	assert.Equal(t, 13*3, r.Area())
}

func TestRecti_Edges(t *testing.T) {
	r := Recti{
		Min: Vec2i{-3, 4},
		Max: Vec2i{10, 7},
	}
	assert.Equal(t, -3, r.Left())
	assert.Equal(t, 10, r.Right())
	assert.Equal(t, 4, r.Bottom())
	assert.Equal(t, 7, r.Top())
}

func TestRecti_SquarePointDistance(t *testing.T) {
	r := Recti{
		Min: Vec2i{0, 0},
		Max: Vec2i{10, 10},
	}

	// within the rectangle
	assert.Equal(t, 0, r.SquarePointDistance(Vec2i{0, 0}))
	assert.Equal(t, 0, r.SquarePointDistance(Vec2i{5, 5}))
	assert.Equal(t, 0, r.SquarePointDistance(Vec2i{10, 10}))
	assert.Equal(t, 0, r.SquarePointDistance(Vec2i{0, 10}))

	// to the left
	assert.Equal(t, 25, r.SquarePointDistance(Vec2i{-5, 0}))
	assert.Equal(t, 16, r.SquarePointDistance(Vec2i{-4, 10}))
	assert.Equal(t, 4, r.SquarePointDistance(Vec2i{-2, 2}))
	assert.Equal(t, 9, r.SquarePointDistance(Vec2i{-3, 5}))
	// to the right
	assert.Equal(t, 25, r.SquarePointDistance(Vec2i{15, 7}))
	// below
	assert.Equal(t, 49, r.SquarePointDistance(Vec2i{3, -7}))
	// above
	assert.Equal(t, 16, r.SquarePointDistance(Vec2i{8, 14}))

	// below + left
	assert.Equal(t, 50, r.SquarePointDistance(Vec2i{-5, -5}))
	assert.Equal(t, 68, r.SquarePointDistance(Vec2i{-2, -8}))
	// above + left
	assert.Equal(t, 41, r.SquarePointDistance(Vec2i{-4, 15}))
	assert.Equal(t, 13, r.SquarePointDistance(Vec2i{-2, 13}))

	// below + right
	assert.Equal(t, 34, r.SquarePointDistance(Vec2i{13, -5}))
	assert.Equal(t, 53, r.SquarePointDistance(Vec2i{17, -2}))
	// above + right
	assert.Equal(t, 37, r.SquarePointDistance(Vec2i{16, 11}))
	assert.Equal(t, 85, r.SquarePointDistance(Vec2i{19, 12}))
}

func TestRecti_PointDistance(t *testing.T) {
	r := Recti{
		Min: Vec2i{0, 0},
		Max: Vec2i{10, 10},
	}
	// within the rectangle
	AssertFloat(t, float32(0), r.PointDistance(Vec2i{5, 5}))
	// to the left
	AssertFloat(t, float32(3), r.PointDistance(Vec2i{-3, 5}))
	// above
	AssertFloat(t, float32(4), r.PointDistance(Vec2i{8, 14}))
	// below + left
	AssertFloat(t, math32.Sqrt(50), r.PointDistance(Vec2i{-5, -5}))
	// below + right
	AssertFloat(t, math32.Sqrt(34), r.PointDistance(Vec2i{13, -5}))
	// above + right
	AssertFloat(t, math32.Sqrt(37), r.PointDistance(Vec2i{16, 11}))
}
