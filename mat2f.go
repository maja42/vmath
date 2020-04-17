package vmath

import (
	"fmt"
)

// Mat2f is a 2x2 float32 matrix.
// Values are stored in column major order: [<col0>, <col1>]
type Mat2f [4]float32

func (m Mat2f) String() string {
	return fmt.Sprintf("Mat2f[(%f x %f)/(%f x %f)]",
		m[0], m[2],
		m[1], m[3])
}

// Ident2f returns the 2x2 identity matrix
func Ident2f() Mat2f {
	return Mat2f{
		1, 0,
		0, 1}
}

// Mat2fFromRows creates a new 2x2 matrix from row vectors.
func Mat2fFromRows(row0, row1 Vec2f) Mat2f {
	return Mat2f{
		row0[0], row1[0],
		row0[1], row1[1]}
}

// Mat2fFromCols creates a new 2x2 matrix from column vectors.
func Mat2fFromCols(col0, col1 Vec2f) Mat2f {
	return Mat2f{
		col0[0], col0[1],
		col1[0], col1[1]}
}

// Mat3f extends the matrix to 3x3.
// The diagonal cell is set to 1, all other values are 0.
func (m Mat2f) Mat3f() Mat3f {
	col0, col1 := m.Cols()
	return Mat3fFromCols(
		col0.Vec3f(0),
		col1.Vec3f(0),
		Vec3f{0, 0, 1},
	)
}

// Mat4f extends the matrix to 4x4.
// The diagonal cells are set to 1, all other values are 0.
func (m Mat2f) Mat4f() Mat4f {
	col0, col1 := m.Cols()
	return Mat4fFromCols(
		col0.Vec4f(0, 0),
		col1.Vec4f(0, 0),
		Vec4f{0, 0, 1, 0},
		Vec4f{0, 0, 0, 1},
	)
}

// Index returns the cell index with the given row and column.
func (m Mat2f) Index(row, col int) int {
	return col*2 + row
}

// Cell returns the element at the given row and column.
func (m Mat2f) Cell(row, col int) float32 {
	return m[col*2+row]
}

// Row returns a vector with the requested row.
func (m Mat2f) Row(row int) Vec2f {
	return Vec2f{m[row+0], m[row+2]}
}

// Rows returns vectors representing all rows.
func (m Mat2f) Rows() (row0, row1 Vec2f) {
	return m.Row(0), m.Row(1)
}

// Col returns a vector with the requested column.
func (m Mat2f) Col(col int) Vec2f {
	return Vec2f{m[col*2+0], m[col*2+1]}
}

// Cols returns vectors representing all columns.
func (m Mat2f) Cols() (col0, col1 Vec2f) {
	return m.Col(0), m.Col(1)
}

// Diag returns the matrix's diagonal values.
func (m Mat2f) Diag() Vec2f {
	return Vec2f{m[0], m[3]}
}

// Set sets a cell value.
func (m *Mat2f) Set(row, col int, v float32) {
	m[col*2+row] = v
}

// SetRow sets the values within a specific row.
func (m *Mat2f) SetRow(row int, v Vec2f) {
	m[row+0] = v[0]
	m[row+2] = v[1]
}

// SetCol sets the values within a specific column.
func (m *Mat2f) SetCol(col int, v Vec2f) {
	m[col*2+0] = v[0]
	m[col*2+1] = v[1]
}

// Transpose returns the transposed matrix.
// Transposing converts between column-major and row-major order.
func (m Mat2f) Transpose() Mat2f {
	return Mat2f{
		m[0], m[2],
		m[1], m[3]}
}

// Inverse calculates the inverse matrix.
// If the matrix cannot be inverted (singular), the identity matrix and false is returned.
func (m Mat2f) Inverse() (Mat2f, bool) {
	det := m.Det()
	if Equalf(det, 0) {
		return Ident2f(), false
	}

	invDet := 1.0 / det
	return Mat2f{
		invDet * m[3], -invDet * m[1],
		-invDet * m[2], invDet * m[0],
	}, true
}

// Det returns the determinant.
func (m Mat2f) Det() float32 {
	return m[0]*m[3] - m[1]*m[2]
}

// Add performs a component-wise addition.
func (m Mat2f) Add(other Mat2f) Mat2f {
	return Mat2f{
		m[0] + other[0], m[1] + other[1],
		m[2] + other[2], m[3] + other[3]}
}

// AddScalar performs a component-wise scalar addition.
func (m Mat2f) AddScalar(s float32) Mat2f {
	return Mat2f{
		m[0] + s, m[1] + s,
		m[2] + s, m[3] + s}
}

// Sub performs a component-wise subtraction.
func (m Mat2f) Sub(other Mat2f) Mat2f {
	return Mat2f{
		m[0] - other[0], m[1] - other[1],
		m[2] - other[2], m[3] - other[3]}
}

// SubScalar performs a component-wise scalar subtraction.
func (m Mat2f) SubScalar(s float32) Mat2f {
	return Mat2f{
		m[0] - s, m[1] - s,
		m[2] - s, m[3] - s}
}

// Mul performs a matrix multiplication.
func (m Mat2f) Mul(other Mat2f) Mat2f {
	return Mat2f{
		m[0]*other[0] + m[2]*other[1],
		m[1]*other[0] + m[3]*other[1],

		m[0]*other[2] + m[2]*other[3],
		m[1]*other[2] + m[3]*other[3]}
}

// Mul performs a component-wise scalar multiplication.
func (m Mat2f) MulScalar(s float32) Mat2f {
	return Mat2f{
		m[0] * s, m[1] * s,
		m[2] * s, m[3] * s}
}

// MulVec multiples the matrix with a vector.
func (m Mat2f) MulVec(v Vec2f) Vec2f {
	return Vec2f{
		m[0]*v[0] + m[2]*v[1],
		m[1]*v[0] + m[3]*v[1],
	}
}

// Equal compares two matrices component-wise.
// Uses the default Epsilon as relative tolerance.
func (m Mat2f) Equal(other Mat2f) bool {
	return m.EqualEps(other, Epsilon)
}

// Equal compares two matrices component-wise, using the given epsilon as a relative tolerance.
func (m Mat2f) EqualEps(other Mat2f, epsilon float32) bool {
	return EqualEps(m[0], other[0], epsilon) &&
		EqualEps(m[1], other[1], epsilon) &&
		EqualEps(m[2], other[2], epsilon) &&
		EqualEps(m[3], other[3], epsilon)
}
