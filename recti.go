package vmath

import (
	"fmt"

	"github.com/maja42/vmath/math32"
	"github.com/maja42/vmath/mathi"
)

// Recti represents a 2D, axis-aligned rectangle.
type Recti struct {
	Min Vec2i
	Max Vec2i
}

// RectiFromCorners creates a new rectangle given two opposite corners.
func RectiFromCorners(c1, c2 Vec2i) Recti {
	if c1[0] > c2[0] {
		c1[0], c2[0] = c2[0], c1[0]
	}
	if c1[1] > c2[1] {
		c1[1], c2[1] = c2[1], c1[1]
	}
	return Recti{c1, c2}
}

// RectiFromPosSize creates a new rectangle with the given size and position.
// Negative dimensions are correctly inverted.
func RectiFromPosSize(pos, size Vec2i) Recti {
	if size[0] < 0 {
		size[0] = -size[0]
		pos[0] -= size[0]
	}
	if size[1] < 0 {
		size[1] = -size[1]
		pos[1] -= size[1]
	}
	return Recti{
		pos,
		pos.Add(size),
	}
}

// RectfFromEdges creates a new rectangle with the given edge positions.
func RectiFromEdges(left, right, bottom, top int) Recti {
	return RectiFromCorners(Vec2i{left, bottom}, Vec2i{right, top})
}

// Normalize ensures that the Min position is smaller than the Max position in every dimension.
func (r Recti) Normalize() Recti {
	if r.Min[0] > r.Max[0] {
		r.Min[0], r.Max[0] = r.Max[0], r.Min[0]
	}
	if r.Min[1] > r.Max[1] {
		r.Min[1], r.Max[1] = r.Max[1], r.Min[1]
	}
	return r
}

func (r Recti) String() string {
	return fmt.Sprintf("Recti([%d x %d]-[%d x %d])",
		r.Min[0], r.Min[1],
		r.Max[0], r.Max[1])
}

// Rectf returns a float representation of the rectangle.
func (r Recti) Rectf() Rectf {
	return Rectf{
		r.Min.Vec2f(),
		r.Max.Vec2f(),
	}
}

// Size returns the rectangle's dimensions.
func (r *Recti) Size() Vec2i {
	return r.Max.Sub(r.Min)
}

// Area returns the rectangle's area.
func (r Recti) Area() int {
	size := r.Max.Sub(r.Min)
	return size[0] * size[1]
}

// Left returns the rectangle's left position (smaller X).
func (r Recti) Left() int {
	return r.Min[0]
}

// Right returns the rectangle's right position (bigger X).
func (r Recti) Right() int {
	return r.Max[0]
}

// Bottom returns the rectangle's bottom position (smaller Y).
func (r Recti) Bottom() int {
	return r.Min[1]
}

// Top returns the rectangle's top position (bigger Y).
func (r Recti) Top() int {
	return r.Max[1]
}

// SetPos changes the rectangle position by modifying min, but keeps the rectangle's size.
func (r Recti) SetPos(pos Vec2i) {
	size := r.Size()
	r.Min = pos
	r.Max = r.Min.Add(size)
}

// SetSize changes the rectangle size by keeping the min-position.
func (r Recti) SetSize(size Vec2i) {
	r.Max = r.Min.Add(size)
}

// Add moves the rectangle with the given vector by adding it to the min- and max- components.
func (r Recti) Add(v Vec2i) Recti {
	return Recti{
		Min: r.Min.Add(v),
		Max: r.Max.Add(v),
	}
}

// Sub moves the rectangle with the given vector by subtracting it to the min- and max- components.
func (r Recti) Sub(v Vec2i) Recti {
	return Recti{
		Min: r.Min.Sub(v),
		Max: r.Max.Sub(v),
	}
}

// Overlaps checks if this rectangle overlaps another rectangle.
// Touching rectangles where floats are exactly equal are not considered to overlap.
func (r Recti) Overlaps(other Recti) bool {
	return r.Min[0] < other.Max[0] &&
		r.Max[0] > other.Min[0] &&
		r.Max[1] > other.Min[1] &&
		r.Min[1] < other.Max[1]
}

// OverlapsOrTouches checks if this rectangle overlaps or touches another rectangle.
func (r Recti) OverlapsOrTouches(other Recti) bool {
	return r.Min[0] <= other.Max[0] &&
		r.Max[0] >= other.Min[0] &&
		r.Max[1] >= other.Min[1] &&
		r.Min[1] <= other.Max[1]
}

// Contains checks if a given point resides within the rectangle.
// If the point is on an edge, it is also considered to be contained within the rectangle.
func (r Recti) ContainsPoint(point Vec2i) bool {
	return point[0] >= r.Min[0] && point[0] <= r.Max[0] &&
		point[1] >= r.Min[1] && point[1] <= r.Max[1]
}

// ContainsRecti checks if this rectangle completely contains another rectangle.
func (r Recti) ContainsRecti(other Recti) bool {
	return r.Min[0] <= other.Min[0] &&
		r.Max[0] >= other.Max[0] &&
		r.Min[1] <= other.Min[1] &&
		r.Max[1] >= other.Max[1]
}

// Merge returns a rectangle that contains both smaller rectangles.
func (r Recti) Merge(other Recti) Recti {
	min := Vec2i{
		mathi.Min(r.Min[0], other.Min[0]),
		mathi.Min(r.Min[1], other.Min[1]),
	}
	max := Vec2i{
		mathi.Max(r.Max[0], other.Max[0]),
		mathi.Max(r.Max[1], other.Max[1]),
	}
	return Recti{min, max}
}

// SquarePointDistance returns the squared distance between the rectangle and a point.
// If the point is contained within the rectangle, 0 is returned.
// Otherwise, the squared distance between the point and the nearest edge or corner is returned.
func (r Recti) SquarePointDistance(pos Vec2i) int {
	// Source: "Nearest Neighbor Queries" by N. Roussopoulos, S. Kelley and F. Vincent, ACM SIGMOD, pages 71-79, 1995.
	sum := 0
	for dim, val := range pos {
		if val < r.Min[dim] {
			// below/left of edge
			d := val - r.Min[dim]
			sum += d * d
		} else if val > r.Max[dim] {
			// above/right of edge
			d := val - r.Max[dim]
			sum += d * d
		} else {
			sum += 0
		}
	}
	return sum
}

// PointDistance returns the distance between the rectangle and a point.
// If the point is contained within the rectangle, 0 is returned.
// Otherwise, the distance between the point and the nearest edge or corner is returned.
func (r Recti) PointDistance(pos Vec2i) float32 {
	return math32.Sqrt(float32(r.SquarePointDistance(pos)))
}
