package mathi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 7, Abs(7))
	assert.Equal(t, 7, Abs(-7))
	assert.Equal(t, 0, Abs(0))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, Min(3, 7))
	assert.Equal(t, -7, Min(3, -7))
	assert.Equal(t, -3, Min(-3, 7))
	assert.Equal(t, -7, Min(-3, -7))
	assert.Equal(t, 0, Min(0, 0))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 7, Max(3, 7))
	assert.Equal(t, 3, Max(3, -7))
	assert.Equal(t, 7, Max(-3, 7))
	assert.Equal(t, -3, Max(-3, -7))
	assert.Equal(t, 0, Max(0, 0))
}
