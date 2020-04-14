package vmath

import (
	"fmt"
)

// Mat3f is a 3x3 float32 matrix.
// Values are stored in column major order: [<col0>, <col1>, <col2>]
//
// 0, 3, 6
// 1, 4, 7
// 2, 5, 8
type Mat3f [9]float32

func (m Mat3f) String() string {
	return fmt.Sprintf("Mat3f[(%f x %f x %f)/(%f x %f x %f)/(%f x %f x %f)]",
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8])
}

// Ident3f returns the 3x3 identity matrix.
func Ident3f() Mat3f {
	return Mat3f{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1}
}

// Mat3fFromCols builds a new 3x3 matrix from row vectors.
func Mat3fFromRows(row0, row1, row2 Vec3f) Mat3f {
	return Mat3f{
		row0[0], row1[0], row2[0],
		row0[1], row1[1], row2[1],
		row0[2], row1[2], row2[2]}
}

// Mat3fFromCols builds a new 3x3 matrix from column vectors.
func Mat3fFromCols(col0, col1, col2 Vec3f) Mat3f {
	return Mat3f{
		col0[0], col0[1], col0[2],
		col1[0], col1[1], col1[2],
		col2[0], col2[1], col2[2]}
}

// Mat2f shrinks the matrix to 2x2.
// The right column and bottom row are removed.
func (m Mat3f) Mat2f() Mat2f {
	col0, col1, _ := m.Cols()
	return Mat2fFromCols(
		col0.XY(),
		col1.XY(),
	)
}

// Mat4f extends the matrix to 4x4.
// The diagonal cell is set to 1, all other values are 0.
func (m Mat3f) Mat4f() Mat4f {
	col0, col1, col2 := m.Cols()
	return Mat4fFromCols(
		col0.Vec4f(0),
		col1.Vec4f(0),
		col2.Vec4f(0),
		Vec4f{0, 0, 0, 1},
	)
}

// Index returns the cell index with the given row and column.
func (m Mat3f) Index(row, col int) int {
	return col*3 + row
}

// Cell returns the element at the given row and column.
func (m Mat3f) Cell(row, col int) float32 {
	return m[col*3+row]
}

// Row returns a vector with the requested row.
func (m Mat3f) Row(row int) Vec3f {
	return Vec3f{m[row+0], m[row+3], m[row+6]}
}

// Rows returns vectors representing all rows.
func (m Mat3f) Rows() (row0, row1, row2 Vec3f) {
	return m.Row(0), m.Row(1), m.Row(2)
}

// Row returns a vector with the requested column.
func (m Mat3f) Col(col int) Vec3f {
	return Vec3f{m[col*3+0], m[col*3+1], m[col*3+2]}
}

// Cols returns vectors representing all columns.
func (m Mat3f) Cols() (col0, col1, col2 Vec3f) {
	return m.Col(0), m.Col(1), m.Col(2)
}

// Diag returns the matrix's diagonal values.
func (m Mat3f) Diag() Vec3f {
	return Vec3f{m[0], m[4], m[8]}
}

// Set sets a cell value.
func (m *Mat3f) Set(row, col int, v float32) {
	m[col*3+row] = v
}

// SetRow sets the values within a specific row.
func (m *Mat3f) SetRow(row int, v Vec3f) {
	m[row+0] = v[0]
	m[row+3] = v[1]
	m[row+6] = v[2]
}

// SetCol sets the values within a specific column.
func (m *Mat3f) SetCol(col int, v Vec3f) {
	m[col*3+0] = v[0]
	m[col*3+1] = v[1]
	m[col*3+2] = v[2]
}

// Transpose returns the transposed matrix.
// Transposing converts between column-major and row-major order.
func (m Mat3f) Transpose() Mat3f {
	return Mat3f{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8]}
}

// det2x2 returns the determinant of a 2x2 matrix
func det2x2(v00, v01, v10, v11 float32) float32 {
	return v00*v11 - v10*v01
}

// Inverse calculates the inverse matrix.
// If the matrix cannot be inverted (singular), the identity matrix and false is returned.
func (m Mat3f) Inverse() (Mat3f, bool) {
	det := m.Det()
	if Equalf(det, 0) {
		return Ident3f(), false
	}

	invDet := 1.0 / det
	return Mat3f{
		invDet * det2x2(m[4], m[7], m[5], m[8]),
		-invDet * det2x2(m[1], m[7], m[2], m[8]),
		invDet * det2x2(m[1], m[4], m[2], m[5]),
		-invDet * det2x2(m[3], m[6], m[5], m[8]),
		invDet * det2x2(m[0], m[6], m[2], m[8]),
		-invDet * det2x2(m[0], m[3], m[2], m[5]),
		invDet * det2x2(m[3], m[6], m[4], m[7]),
		-invDet * det2x2(m[0], m[6], m[1], m[7]),
		invDet * det2x2(m[0], m[3], m[1], m[4]),
	}, true
}

// InverseTranspose inverts and transposes the matrix in a single step.
// If the matrix cannot be inverted (singular), the identity matrix and false is returned.
func (m Mat3f) InverseTranspose() (Mat3f, bool) {
	det := m.Det()

	if Equalf(det, 0) {
		return Ident3f(), false
	}

	invDet := 1.0 / det
	return Mat3f{
		invDet * det2x2(m[4], m[7], m[5], m[8]),
		-invDet * det2x2(m[3], m[6], m[5], m[8]),
		invDet * det2x2(m[3], m[6], m[4], m[7]),
		-invDet * det2x2(m[1], m[7], m[2], m[8]),
		invDet * det2x2(m[0], m[6], m[2], m[8]),
		-invDet * det2x2(m[0], m[6], m[1], m[7]),
		invDet * det2x2(m[1], m[4], m[2], m[5]),
		-invDet * det2x2(m[0], m[3], m[2], m[5]),
		invDet * det2x2(m[0], m[3], m[1], m[4]),
	}, true
}

// Det returns the determinant.
func (m Mat3f) Det() float32 {
	return m[0]*m[4]*m[8] + m[2]*m[3]*m[7] + m[1]*m[5]*m[6] -
		m[0]*m[5]*m[7] - m[1]*m[3]*m[8] - m[2]*m[4]*m[6]
}

// Add performs a component-wise addition.
func (m Mat3f) Add(other Mat3f) Mat3f {
	return Mat3f{
		m[0] + other[0], m[1] + other[1], m[2] + other[2],
		m[3] + other[3], m[4] + other[4], m[5] + other[5],
		m[6] + other[6], m[7] + other[7], m[8] + other[8]}
}

// AddScalar performs a component-wise scalar addition.
func (m Mat3f) AddScalar(s float32) Mat3f {
	return Mat3f{
		m[0] + s, m[1] + s, m[2] + s,
		m[3] + s, m[4] + s, m[5] + s,
		m[6] + s, m[7] + s, m[8] + s}
}

// Sub performs a component-wise subtraction.
func (m Mat3f) Sub(other Mat3f) Mat3f {
	return Mat3f{
		m[0] - other[0], m[1] - other[1], m[2] - other[2],
		m[3] - other[3], m[4] - other[4], m[5] - other[5],
		m[6] - other[6], m[7] - other[7], m[8] - other[8]}
}

// SubScalar performs a component-wise scalar subtraction.
func (m Mat3f) SubScalar(s float32) Mat3f {
	return Mat3f{
		m[0] - s, m[1] - s, m[2] - s,
		m[3] - s, m[4] - s, m[5] - s,
		m[6] - s, m[7] - s, m[8] - s}
}

// Mul performs a matrix multiplication.
func (m Mat3f) Mul(other Mat3f) Mat3f {
	return Mat3f{
		m[0]*other[0] + m[3]*other[1] + m[6]*other[2],
		m[1]*other[0] + m[4]*other[1] + m[7]*other[2],
		m[2]*other[0] + m[5]*other[1] + m[8]*other[2],

		m[0]*other[3] + m[3]*other[4] + m[6]*other[5],
		m[1]*other[3] + m[4]*other[4] + m[7]*other[5],
		m[2]*other[3] + m[5]*other[4] + m[8]*other[5],

		m[0]*other[6] + m[3]*other[7] + m[6]*other[8],
		m[1]*other[6] + m[4]*other[7] + m[7]*other[8],
		m[2]*other[6] + m[5]*other[7] + m[8]*other[8]}
}

// MulScalar performs a component-wise scalar multiplication.
func (m Mat3f) MulScalar(s float32) Mat3f {
	return Mat3f{
		m[0] * s, m[1] * s, m[2] * s,
		m[3] * s, m[4] * s, m[5] * s,
		m[6] * s, m[7] * s, m[8] * s}
}

// MulVec multiples the matrix with a vector.
func (m Mat3f) MulVec(v Vec3f) Vec3f {
	return Vec3f{
		m[0]*v[0] + m[3]*v[1] + m[6]*v[2],
		m[1]*v[0] + m[4]*v[1] + m[7]*v[2],
		m[2]*v[0] + m[5]*v[1] + m[8]*v[2],
	}
}

// Equal compares two matrices component-wise.
// Uses the default Epsilon as relative tolerance.
func (m Mat3f) Equal(other Mat3f) bool {
	return m.EqualEps(other, Epsilon)
}

// Equal compares two matrices component-wise, using the given epsilon as a relative tolerance.
func (m Mat3f) EqualEps(other Mat3f, epsilon float32) bool {
	for i := range m {
		if !EqualEps(m[i], other[i], epsilon) {
			return false
		}
	}
	return true
}
