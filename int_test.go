package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsi(t *testing.T) {
	assert.Equal(t, 4, Absi(4))
	assert.Equal(t, 9, Absi(-9))
	assert.Equal(t, 0, Absi(0))
}

func TestMaxi(t *testing.T) {
	assert.Equal(t, 42, Maxi(12, 42))
	assert.Equal(t, 42, Maxi(42, -62))
	assert.Equal(t, -5, Maxi(-5, -9))
	assert.Equal(t, -5, Maxi(-5, -5))
}

func TestMini(t *testing.T) {
	assert.Equal(t, 12, Mini(12, 42))
	assert.Equal(t, -62, Mini(97, -62))
	assert.Equal(t, -9, Mini(-5, -9))
	assert.Equal(t, -5, Mini(-5, -5))
}
