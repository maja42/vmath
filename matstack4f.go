package vmath

import (
	"errors"
)

// MatStack4f represents a stack of 4x4 matrices.
type MatStack4f struct {
	stack []Mat4f
}

// NewMatStack4f creates a new matrix stack containing only the identity matrix.
func NewMatStack4f() *MatStack4f {
	mStack := &MatStack4f{
		stack: make([]Mat4f, 1),
	}
	mStack.stack[0] = Ident4f()
	return mStack
}

// Size returns the current size of the matrix stack
func (m MatStack4f) Size() int {
	return len(m.stack)
}

// Push stores the current top on the stack by duplicating it.
func (m *MatStack4f) Push() {
	m.stack = append(m.stack, m.Top())
}

// Pop removes the current top element from the stack.
// Returns an error if the stack contains only one element.
func (m *MatStack4f) Pop() error {
	if len(m.stack) == 1 {
		return errors.New("cannot pop last element from matrix stack")
	}
	m.stack = m.stack[:len(m.stack)-1]
	return nil
}

// Top returns the current top element without modifying the stack.
func (m MatStack4f) Top() Mat4f {
	return m.stack[len(m.stack)-1]
}

// Set overwrites the top element with a new matrix.
func (m *MatStack4f) Set(mat Mat4f) {
	m.stack[len(m.stack)-1] = mat
}

// SetIdent overwrites the top element with the identity matrix.
func (m *MatStack4f) SetIdent() {
	m.Set(Ident4f())
}

// MulRight multiplies the top element with the given matrix.
func (m *MatStack4f) MulRight(mat Mat4f) {
	top := m.Top()
	m.Set(top.Mul(mat))
}

// MulLeft multiplies the given matrix with the top element and overwrites the top element with the result.
func (m *MatStack4f) MulLeft(mat Mat4f) {
	top := m.Top()
	m.Set(mat.Mul(top))
}
