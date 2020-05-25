package vmath

import (
	"fmt"

	"github.com/maja42/vmath/math32"
)

// Rectf represents a 2D, axis-aligned rectangle.
type Rectf struct {
	Min Vec2f
	Max Vec2f
}

// RectfFromCorners creates a new rectangle given two opposite corners.
// If necessary, coordinates are swapped to create a normalized rectangle.
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
// Negative dimensions are inverted to create a normalized rectangle.
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
// If necessary, edges are swapped to create a normalized rectangle.
func RectfFromEdges(left, right, bottom, top float32) Rectf {
	return RectfFromCorners(Vec2f{left, bottom}, Vec2f{right, top})
}

// Normalize ensures that the Min position is smaller than the Max position in every dimension.
func (r Rectf) Normalize() Rectf {
	if r.Min[0] > r.Max[0] {
		r.Min[0], r.Max[0] = r.Max[0], r.Min[0]
	}
	if r.Min[1] > r.Max[1] {
		r.Min[1], r.Max[1] = r.Max[1], r.Min[1]
	}
	return r
}

func (r Rectf) String() string {
	return fmt.Sprintf("Rectf([%f x %f]-[%f x %f])",
		r.Min[0], r.Min[1],
		r.Max[0], r.Max[1])
}

// Recti returns an integer representation of the rectangle.
// Decimals are truncated.
func (r Rectf) Recti() Recti {
	return Recti{
		r.Min.Vec2i(),
		r.Max.Vec2i(),
	}
}

// Round returns an integer representation of the rectangle.
// Decimals are rounded.
func (r Rectf) Round() Recti {
	return Recti{
		r.Min.Round(),
		r.Max.Round(),
	}
}

// Size returns the rectangle's dimensions.
func (r Rectf) Size() Vec2f {
	return r.Max.Sub(r.Min)
}

// Area returns the rectangle's area.
func (r Rectf) Area() float32 {
	size := r.Max.Sub(r.Min)
	return size[0] * size[1]
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

// SetPos changes the rectangle position by modifying min, but keeps the rectangle's size.
func (r Rectf) SetPos(pos Vec2f) {
	size := r.Size()
	r.Min = pos
	r.Max = r.Min.Add(size)
}

// SetSize changes the rectangle size by keeping the min-position.
func (r Rectf) SetSize(size Vec2f) {
	r.Max = r.Min.Add(size)
}

// Add moves the rectangle with the given vector by adding it to the min- and max- components.
func (r Rectf) Add(v Vec2f) Rectf {
	return Rectf{
		Min: r.Min.Add(v),
		Max: r.Max.Add(v),
	}
}

// Sub moves the rectangle with the given vector by subtracting it to the min- and max- components.
func (r Rectf) Sub(v Vec2f) Rectf {
	return Rectf{
		Min: r.Min.Sub(v),
		Max: r.Max.Sub(v),
	}
}

// Intersects checks if this rectangle intersects another rectangle.
// Touching rectangles where floats are exactly equal are not considered to intersect.
func (r Rectf) Intersects(other Rectf) bool {
	return r.Min[0] <= other.Max[0] &&
		r.Max[0] >= other.Min[0] &&
		r.Max[1] >= other.Min[1] &&
		r.Min[1] <= other.Max[1]
}

// ContainsPoint checks if a given point resides within the rectangle.
// If the point is on an edge, it is also considered to be contained within the rectangle.
func (r Rectf) ContainsPoint(point Vec2f) bool {
	return point[0] >= r.Min[0] && point[0] <= r.Max[0] &&
		point[1] >= r.Min[1] && point[1] <= r.Max[1]
}

// ContainsRectf checks if this rectangle completely contains another rectangle.
func (r Rectf) ContainsRectf(other Rectf) bool {
	return r.Min[0] <= other.Min[0] &&
		r.Max[0] >= other.Max[0] &&
		r.Min[1] <= other.Min[1] &&
		r.Max[1] >= other.Max[1]
}

// Merge returns a rectangle that contains both smaller rectangles.
func (r Rectf) Merge(other Rectf) Rectf {
	min := Vec2f{
		math32.Min(r.Min[0], other.Min[0]),
		math32.Min(r.Min[1], other.Min[1]),
	}
	max := Vec2f{
		math32.Max(r.Max[0], other.Max[0]),
		math32.Max(r.Max[1], other.Max[1]),
	}
	return Rectf{min, max}
}

// SquarePointDistance returns the squared distance between the rectangle and a point.
// If the point is contained within the rectangle, 0 is returned.
// Otherwise, the squared distance between the point and the nearest edge or corner is returned.
func (r Rectf) SquarePointDistance(pos Vec2f) float32 {
	// Source: "Nearest Neighbor Queries" by N. Roussopoulos, S. Kelley and F. Vincent, ACM SIGMOD, pages 71-79, 1995.
	sum := float32(0.0)
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
func (r Rectf) PointDistance(pos Vec2f) float32 {
	return math32.Sqrt(r.SquarePointDistance(pos))
}
