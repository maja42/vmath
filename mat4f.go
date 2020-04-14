package vmath

import (
	"fmt"
)

// matrices are stored in column major order

// Mat4f is a 4x4 float32 matrix.
// Values are stored in column major order: [<col0>, <col1>, <col2>, <col4>]
//
// 0, 4,  8, 12,
// 1, 5,  9, 13,
// 2, 6, 10, 14,
// 3, 7, 11, 15
type Mat4f [16]float32

func (m Mat4f) String() string {
	return fmt.Sprintf("Mat4f[(%f x %f x %f x %f)/(%f x %f x %f x %f)/(%f x %f x %f x %f)/(%f x %f x %f x %f)]",
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15])
}

// Ident4f returns the 4x4 identity matrix.
func Ident4f() Mat4f {
	return Mat4f{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// Mat4fFromRows builds a new 4x4 matrix from row vectors.
func Mat4fFromRows(row0, row1, row2, row3 Vec4f) Mat4f {
	return Mat4f{
		row0[0], row1[0], row2[0], row3[0],
		row0[1], row1[1], row2[1], row3[1],
		row0[2], row1[2], row2[2], row3[2],
		row0[3], row1[3], row2[3], row3[3]}
}

// Mat4fFromCols builds a new 4x4 matrix from column vectors.
func Mat4fFromCols(col0, col1, col2, col3 Vec4f) Mat4f {
	return Mat4f{
		col0[0], col0[1], col0[2], col0[3],
		col1[0], col1[1], col1[2], col1[3],
		col2[0], col2[1], col2[2], col2[3],
		col3[0], col3[1], col3[2], col3[3]}
}

// Mat2f shrinks the matrix to 2x2.
// The right columns and bottom rows are removed.
func (m Mat4f) Mat2f() Mat2f {
	col0, col1, _, _ := m.Cols()
	return Mat2fFromCols(
		col0.XY(),
		col1.XY(),
	)
}

// Mat3f shrinks the matrix to 3x3.
// The right column and bottom row are removed.
func (m Mat4f) Mat3f() Mat3f {
	col0, col1, col2, _ := m.Cols()
	return Mat3fFromCols(
		col0.XYZ(),
		col1.XYZ(),
		col2.XYZ(),
	)
}

// SetMat3f sets the upper-left 3x3 matrix.
func (m Mat4f) SetMat3f(other Mat3f) Mat4f {
	m[0] = other[0]
	m[4] = other[3]
	m[8] = other[6]

	m[1] = other[1]
	m[5] = other[4]
	m[9] = other[7]

	m[2] = other[2]
	m[6] = other[5]
	m[10] = other[8]
	return m
}

// Index returns the cell index with the given row and column.
func (m Mat4f) Index(row, col int) int {
	return col*4 + row
}

// Cell returns the element at the given row and column.
func (m Mat4f) Cell(row, col int) float32 {
	return m[col*4+row]
}

// Row returns a vector with the requested row.
func (m Mat4f) Row(row int) Vec4f {
	return Vec4f{m[row+0], m[row+4], m[row+8], m[row+12]}
}

// Rows returns vectors representing all rows.
func (m Mat4f) Rows() (row0, row1, row2, row3 Vec4f) {
	return m.Row(0), m.Row(1), m.Row(2), m.Row(3)
}

// Col returns a vector with the requested column.
func (m Mat4f) Col(col int) Vec4f {
	return Vec4f{m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3]}
}

// Cols returns vectors representing all columns.
func (m Mat4f) Cols() (col0, col1, col2, col3 Vec4f) {
	return m.Col(0), m.Col(1), m.Col(2), m.Col(3)
}

// Diag returns the matrix's diagonal values.
func (m Mat4f) Diag() Vec4f {
	return Vec4f{m[0], m[5], m[10], m[15]}
}

// Set sets a cell value.
func (m *Mat4f) Set(row, col int, v float32) {
	m[col*4+row] = v
}

// SetRow sets the values within a specific row.
func (m *Mat4f) SetRow(row int, v Vec4f) {
	m[row+0] = v[0]
	m[row+4] = v[1]
	m[row+8] = v[2]
	m[row+12] = v[3]
}

// SetCol sets the values within a specific column.
func (m *Mat4f) SetCol(col int, v Vec4f) {
	m[col*4+0] = v[0]
	m[col*4+1] = v[1]
	m[col*4+2] = v[2]
	m[col*4+3] = v[3]
}

// Transpose returns the transposed matrix.
// Transposing converts between column-major and row-major order.
func (m Mat4f) Transpose() Mat4f {
	return Mat4f{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15]}
}

// IsAffine checks if this is an affine matrix.
func (m Mat4f) IsAffine() bool {
	return Equalf(m[12], 0) && Equalf(m[13], 0) && Equalf(m[14], 0) && Equalf(m[15], 0)
}

// InverseAffine calculates the inverse of an affine matrix.
// If the matrix cannot be inverted (singular), the identity matrix and false is returned.
func (m Mat4f) InverseAffine() (Mat4f, bool) {
	inv3, ok := m.Mat3f().Inverse()
	if !ok {
		return Ident4f(), false
	}

	res := inv3.Mat4f()
	res[3] = -(m[3]*res[0] + m[7]*res[1] + m[11]*res[2])
	res[7] = -(m[3]*res[4] + m[7]*res[5] + m[11]*res[6])
	res[11] = -(m[3]*res[8] + m[7]*res[9] + m[11]*res[10])
	return res, true
}

// Inverse calculates the inverse matrix.
// If the matrix cannot be inverted (singular), the identity matrix and false is returned.
func (m Mat4f) Inverse() (Mat4f, bool) {
	if m.IsAffine() {
		return m.InverseAffine()
	}

	det := m.Det()
	if Equalf(det, 0) {
		return Ident4f(), false
	}
	ret := Mat4f{
		-m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] - m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15],
		m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] + m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15],
		-m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] - m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15],
		m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] + m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11],

		m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] + m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15],
		-m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] - m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15],
		m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] + m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15],
		-m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] - m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11],

		-m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] - m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15],
		m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] + m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15],
		-m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] - m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15],
		m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] + m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11],

		m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] + m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14],
		-m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] - m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14],
		m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] + m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14],
		-m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] - m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10],
	}
	return ret.MulScalar(1 / det), true
}

// Det returns the determinant.
func (m Mat4f) Det() float32 {
	return m[3]*m[6]*m[9]*m[12] - m[2]*m[7]*m[9]*m[12] - m[3]*m[5]*m[10]*m[12] + m[1]*m[7]*m[10]*m[12] +
		m[2]*m[5]*m[11]*m[12] - m[1]*m[6]*m[11]*m[12] - m[3]*m[6]*m[8]*m[13] + m[2]*m[7]*m[8]*m[13] +
		m[3]*m[4]*m[10]*m[13] - m[0]*m[7]*m[10]*m[13] - m[2]*m[4]*m[11]*m[13] + m[0]*m[6]*m[11]*m[13] +
		m[3]*m[5]*m[8]*m[14] - m[1]*m[7]*m[8]*m[14] - m[3]*m[4]*m[9]*m[14] + m[0]*m[7]*m[9]*m[14] +
		m[1]*m[4]*m[11]*m[14] - m[0]*m[5]*m[11]*m[14] - m[2]*m[5]*m[8]*m[15] + m[1]*m[6]*m[8]*m[15] +
		m[2]*m[4]*m[9]*m[15] - m[0]*m[6]*m[9]*m[15] - m[1]*m[4]*m[10]*m[15] + m[0]*m[5]*m[10]*m[15]
}

// Add performs a component-wise addition.
func (m Mat4f) Add(other Mat4f) Mat4f {
	return Mat4f{
		m[0] + other[0], m[1] + other[1], m[2] + other[2], m[3] + other[3],
		m[4] + other[4], m[5] + other[5], m[6] + other[6], m[7] + other[7],
		m[8] + other[8], m[9] + other[9], m[10] + other[10], m[11] + other[11],
		m[12] + other[12], m[13] + other[13], m[14] + other[14], m[15] + other[15]}
}

// AddScalar performs a component-wise scalar addition.
func (m Mat4f) AddScalar(s float32) Mat4f {
	return Mat4f{
		m[0] + s, m[1] + s, m[2] + s, m[3] + s,
		m[4] + s, m[5] + s, m[6] + s, m[7] + s,
		m[8] + s, m[9] + s, m[10] + s, m[11] + s,
		m[12] + s, m[13] + s, m[14] + s, m[15] + s}
}

// SubScalar performs a component-wise scalar subtraction.
func (m Mat4f) SubScalar(s float32) Mat4f {
	return Mat4f{
		m[0] - s, m[1] - s, m[2] - s, m[3] - s,
		m[4] - s, m[5] - s, m[6] - s, m[7] - s,
		m[8] - s, m[9] - s, m[10] - s, m[11] - s,
		m[12] - s, m[13] - s, m[14] - s, m[15] - s}
}

// Sub performs a component-wise subtraction.
func (m Mat4f) Sub(other Mat4f) Mat4f {
	return Mat4f{
		m[0] - other[0], m[1] - other[1], m[2] - other[2], m[3] - other[3],
		m[4] - other[4], m[5] - other[5], m[6] - other[6], m[7] - other[7],
		m[8] - other[8], m[9] - other[9], m[10] - other[10], m[11] - other[11],
		m[12] - other[12], m[13] - other[13], m[14] - other[14], m[15] - other[15]}
}

// Mul performs a matrix multiplication.
func (m Mat4f) Mul(other Mat4f) Mat4f {
	return Mat4f{
		m[0]*other[0] + m[4]*other[1] + m[8]*other[2] + m[12]*other[3],
		m[1]*other[0] + m[5]*other[1] + m[9]*other[2] + m[13]*other[3],
		m[2]*other[0] + m[6]*other[1] + m[10]*other[2] + m[14]*other[3],
		m[3]*other[0] + m[7]*other[1] + m[11]*other[2] + m[15]*other[3],

		m[0]*other[4] + m[4]*other[5] + m[8]*other[6] + m[12]*other[7],
		m[1]*other[4] + m[5]*other[5] + m[9]*other[6] + m[13]*other[7],
		m[2]*other[4] + m[6]*other[5] + m[10]*other[6] + m[14]*other[7],
		m[3]*other[4] + m[7]*other[5] + m[11]*other[6] + m[15]*other[7],

		m[0]*other[8] + m[4]*other[9] + m[8]*other[10] + m[12]*other[11],
		m[1]*other[8] + m[5]*other[9] + m[9]*other[10] + m[13]*other[11],
		m[2]*other[8] + m[6]*other[9] + m[10]*other[10] + m[14]*other[11],
		m[3]*other[8] + m[7]*other[9] + m[11]*other[10] + m[15]*other[11],

		m[0]*other[12] + m[4]*other[13] + m[8]*other[14] + m[12]*other[15],
		m[1]*other[12] + m[5]*other[13] + m[9]*other[14] + m[13]*other[15],
		m[2]*other[12] + m[6]*other[13] + m[10]*other[14] + m[14]*other[15],
		m[3]*other[12] + m[7]*other[13] + m[11]*other[14] + m[15]*other[15]}
}

// MulScalar performs a component-wise scalar multiplication.
func (m Mat4f) MulScalar(s float32) Mat4f {
	return Mat4f{
		m[0] * s, m[1] * s, m[2] * s, m[3] * s,
		m[4] * s, m[5] * s, m[6] * s, m[7] * s,
		m[8] * s, m[9] * s, m[10] * s, m[11] * s,
		m[12] * s, m[13] * s, m[14] * s, m[15] * s}
}

// MulVec multiples the matrix with a vector.
func (m Mat4f) MulVec(v Vec4f) Vec4f {
	return Vec4f{
		m[0]*v[0] + m[4]*v[1] + m[8]*v[2] + m[12]*v[3],
		m[1]*v[0] + m[5]*v[1] + m[9]*v[2] + m[13]*v[3],
		m[2]*v[0] + m[6]*v[1] + m[10]*v[2] + m[14]*v[3],
		m[3]*v[0] + m[7]*v[1] + m[11]*v[2] + m[15]*v[3],
	}
}

// Equal compares two matrices component-wise.
// Uses the default Epsilon as relative tolerance.
func (m Mat4f) Equal(other Mat4f) bool {
	return m.EqualEps(other, Epsilon)
}

// Equal compares two matrices component-wise, using the given epsilon as a relative tolerance.
func (m Mat4f) EqualEps(other Mat4f, epsilon float32) bool {
	for i := range m {
		if !EqualEps(m[i], other[i], epsilon) {
			return false
		}
	}
	return true
}

// Translation returns the translation vector of the matrix.
func (m Mat4f) Translation() Vec3f {
	return Vec3f{m[3], m[7], m[11]}
}

// SetTranslation sets the translation vector of the matrix.
func (m Mat4f) SetTranslation(translation Vec3f) Mat4f {
	m[3] = translation[0]
	m[7] = translation[1]
	m[11] = translation[2]
	return m
}

// Mat4fFromTranslation returns the 4x4 matrix with the given translation vector.
func Mat4fFromTranslation(translation Vec3f) Mat4f {
	return Mat4f{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		translation[0], translation[1], translation[2], 1}
}

// Scaling returns the scaling of the matrix.
func (m Mat4f) Scaling() Vec3f {
	return Vec3f{m[0], m[5], m[10]}
}

// SetScaling sets the scaling of the matrix.
func (m Mat4f) SetScaling(scaling Vec3f) Mat4f {
	m[0] = scaling[0]
	m[5] = scaling[1]
	m[10] = scaling[2]
	return m
}

// Mat4fFromScaling returns a 4x4 matrix with the given scaling.
func Mat4fFromScaling(scaling Vec3f) Mat4f {
	return Mat4f{
		scaling[0], 0, 0, 0,
		0, scaling[1], 0, 0,
		0, 0, scaling[2], 0,
		0, 0, 0, 1}
}
