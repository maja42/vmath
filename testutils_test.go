package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertFloat(t *testing.T, expected, actual float32) {
	t.Helper()
	assert.InDelta(t, expected, actual, float64(eps))
}

func AssertVec2f(t *testing.T, expected, actual Vec2f) {
	t.Helper()
	delta := float32(eps)
	diff := expected.Sub(actual)
	if diff[0] < -delta || diff[0] > delta || diff[1] < -delta || diff[1] > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, diff)
	}
}

func AssertVec3f(t *testing.T, expected, actual Vec3f) {
	t.Helper()
	delta := float32(eps)
	diff := expected.Sub(actual)
	if diff[0] < -delta || diff[0] > delta ||
		diff[1] < -delta || diff[1] > delta ||
		diff[2] < -delta || diff[2] > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, diff)
	}
}

func AssertVec4f(t *testing.T, expected, actual Vec4f) {
	t.Helper()
	delta := float32(eps)
	diff := expected.Sub(actual)
	if diff[0] < -delta || diff[0] > delta ||
		diff[1] < -delta || diff[1] > delta ||
		diff[2] < -delta || diff[2] > delta ||
		diff[3] < -delta || diff[3] > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, diff)
	}
}
