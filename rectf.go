package vmath

import (
	"fmt"
)

// Rectf represents a 2D, axis-aligned rectangle.
type Rectf struct {
	Min Vec2f
	Max Vec2f
}

// RectfFromCorners creates a new rectangle given two opposite corners.
func RectfFromCorners(c1, c2 Vec2f) Rectf {
	if c1[0] > c2[0] {
		c1[0], c2[0] = c2[0], c1[0]
	}
	if c1[1] > c2[1] {
		c1[1], c2[1] = c2[1], c1[1]
	}
	return Rectf{c1, c2}
}

// RectfFromPosSize creates a new rectangle with the given size and position.
// Negative dimensions are correctly inverted.
func RectfFromPosSize(pos, size Vec2f) Rectf {
	if size[0] < 0 {
		size[0] = -size[0]
		pos[0] -= size[0]
	}
	if size[1] < 0 {
		size[1] = -size[1]
		pos[1] -= size[1]
	}
	return Rectf{
		pos,
		pos.Add(size),
	}
}

// RectfFromEdges creates a new rectangle with the given edge positions.
func RectfFromEdges(left, right, bottom, top float32) Rectf {
	return RectfFromCorners(Vec2f{left, bottom}, Vec2f{right, top})
}

func (r Rectf) String() string {
	return fmt.Sprintf("Rectf([%f x %f]-[%f x %f])",
		r.Min[0], r.Min[1],
		r.Max[0], r.Max[1])
}

// Recti returns an integer representation of the rectangle.
// Decimals are truncated (rounded down).
func (r Rectf) Recti() Recti {
	return Recti{
		r.Min.Vec2i(),
		r.Max.Vec2i(),
	}
}

// Size returns the rectangle's dimensions.
func (r Rectf) Size() Vec2f {
	return r.Max.Sub(r.Min)
}

// Left returns the rectangle's left position (smaller X).
func (r Rectf) Left() float32 {
	return r.Min[0]
}

// Right returns the rectangle's right position (bigger X).
func (r Rectf) Right() float32 {
	return r.Max[0]
}

// Bottom returns the rectangle's bottom position (smaller Y).
func (r Rectf) Bottom() float32 {
	return r.Min[1]
}

// Top returns the rectangle's top position (bigger Y).
func (r Rectf) Top() float32 {
	return r.Max[1]
}

// Overlaps checks if this rectangle overlaps another rectangle.
// Touching rectangles where floats are exactly equal are not considered to overlap.
func (r Rectf) Overlaps(other Rectf) bool {
	return r.Min[0] < other.Max[0] &&
		r.Max[0] > other.Min[0] &&
		r.Max[1] > other.Min[1] &&
		r.Min[1] < other.Max[1]
}
