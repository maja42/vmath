package vmath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const pi = math.Pi
const eps = Epsilon * 5

func TestAngleDiff(t *testing.T) {
	AssertFloat(t, pi/3, AngleDiff(pi/3, 2*pi/3))
	AssertFloat(t, -pi/3, AngleDiff(2*pi/3, pi/3))
	AssertFloat(t, 0, AngleDiff(pi, 5*pi))
	AssertFloat(t, 0, AngleDiff(-3*pi, -1*pi))
	AssertFloat(t, pi, AngleDiff(-2*pi, -1*pi))
}

func TestAngleToVector(t *testing.T) {
	AssertVec2f(t, Vec2f{3.5, 0}, AngleToVector(0, 3.5))

	AssertVec2f(t, Vec2f{0, 1}, AngleToVector(pi/2, 1))
	AssertVec2f(t, Vec2f{-1, 0}, AngleToVector(pi, 1))
	AssertVec2f(t, Vec2f{0, -1}, AngleToVector(3*pi/2, 1))

	AssertVec2f(t, Vec2f{0, 42}, AngleToVector(pi/2, 42))
	AssertVec2f(t, Vec2f{0, 0}, AngleToVector(pi, 0))
}

func TestClampf(t *testing.T) {
	AssertFloat(t, 3, Clampf(2, 3, 5))
	AssertFloat(t, 4, Clampf(4, 3, 5))
	AssertFloat(t, 5, Clampf(6, 3, 5))

	AssertFloat(t, -3, Clampf(-4, -3, -1))
	AssertFloat(t, -2, Clampf(-2, -3, -1))
	AssertFloat(t, -1, Clampf(0, -3, -1))
}

func TestClampi(t *testing.T) {
	assert.Equal(t, 3, Clampi(2, 3, 5))
	assert.Equal(t, 4, Clampi(4, 3, 5))
	assert.Equal(t, 5, Clampi(6, 3, 5))

	assert.Equal(t, -3, Clampi(-4, -3, -1))
	assert.Equal(t, -2, Clampi(-2, -3, -1))
	assert.Equal(t, -1, Clampi(0, -3, -1))
}

func TestDegrees(t *testing.T) {
	AssertFloat(t, 0, Degrees(0))
	AssertFloat(t, 90, Degrees(pi/2))
	AssertFloat(t, 180, Degrees(pi))
	AssertFloat(t, 270, Degrees(3*pi/2))
	AssertFloat(t, 360, Degrees(2*pi))
	AssertFloat(t, 720, Degrees(4*pi))

	AssertFloat(t, -180, Degrees(-pi))
	AssertFloat(t, -360, Degrees(-2*pi))
}

func TestEqualEps(t *testing.T) {
	assert.True(t, EqualEps(0, 0, 0))
	assert.True(t, EqualEps(-4, -4, 0))
	assert.True(t, EqualEps(-4, -4.001, 0.001))
	assert.True(t, EqualEps(-4, -3.999, 0.001))
	assert.False(t, EqualEps(-4, -4.001, 0))

	assert.True(t, EqualEps(1e6, 1e6+1, 1e-6))
	assert.False(t, EqualEps(1e6, 1e6+1, 1e-7))
}

func TestEqualf(t *testing.T) {
	assert.True(t, Equalf(0, 0))
	assert.True(t, Equalf(-4, -4))

	assert.False(t, Equalf(-4, -4.001))
	assert.False(t, Equalf(-4, -3.999))

	assert.True(t, Equalf(-4, -4.000001))
	assert.True(t, Equalf(-4, -3.999999))

	assert.True(t, Equalf(1e6, 1e6*(1+Epsilon)))
	assert.False(t, Equalf(1e6, 1e6*(1+10*Epsilon)))
}

func TestIsPointOnLeft(t *testing.T) {
	a := Vec2f{2, 2}
	b := Vec2f{4, 4}

	assert.True(t, IsPointOnLeft(a, b, Vec2f{3, 7}))
	assert.False(t, IsPointOnLeft(b, a, Vec2f{3, 7}))

	assert.False(t, IsPointOnLeft(a, b, Vec2f{3, 2}))
	assert.True(t, IsPointOnLeft(b, a, Vec2f{3, 2}))

	assert.False(t, IsPointOnLeft(a, b, Vec2f{2, 2}))
}

func TestIsPointOnLine(t *testing.T) {
	a := Vec2f{2, 2}
	b := Vec2f{4, 4}

	assert.True(t, IsPointOnLine(a, b, Vec2f{8, 8}))
	assert.True(t, IsPointOnLine(b, a, Vec2f{8, 8}))

	assert.False(t, IsPointOnLine(a, b, Vec2f{1, 2}))
	assert.False(t, IsPointOnLine(a, b, Vec2f{1, 1.0001}))
}

func TestIsPointOnLineEps(t *testing.T) {
	a := Vec2f{2, 2}
	b := Vec2f{4, 4}

	assert.True(t, IsPointOnLineEps(a, b, Vec2f{8, 8}, 0))
	assert.True(t, IsPointOnLineEps(b, a, Vec2f{8, 8}, 0))

	assert.False(t, IsPointOnLineEps(a, b, Vec2f{1e5, 1e5 + 1}, 0))
	assert.True(t, IsPointOnLineEps(a, b, Vec2f{1e5, 1e5 + 1}, 1))
}

func TestLerp(t *testing.T) {
	AssertFloat(t, 0, Lerp(0, 10, 0))
	AssertFloat(t, 5, Lerp(0, 10, 0.5))
	AssertFloat(t, 10, Lerp(0, 10, 1))
	AssertFloat(t, 20, Lerp(0, 10, 2))
	AssertFloat(t, -20, Lerp(0, 10, -2))

	AssertFloat(t, -5, Lerp(-10, -5, 1))
	AssertFloat(t, 2.5, Lerp(-5, 5, 0.75))
}

func TestNormalizeDegrees(t *testing.T) {
	AssertFloat(t, 0, NormalizeDegrees(0))
	AssertFloat(t, 0, NormalizeDegrees(360))
	AssertFloat(t, 0, NormalizeDegrees(720))
	AssertFloat(t, 0, NormalizeDegrees(-360))

	AssertFloat(t, 45, NormalizeDegrees(45))
	AssertFloat(t, 315, NormalizeDegrees(-45))
	AssertFloat(t, 270, NormalizeDegrees(-90))
}

func TestNormalizeRadians(t *testing.T) {
	AssertFloat(t, 0, Radians(0))
	AssertFloat(t, pi/2, Radians(90))
	AssertFloat(t, pi, Radians(180))
	AssertFloat(t, 3*pi/2, Radians(270))
	AssertFloat(t, 2*pi, Radians(360))
	AssertFloat(t, 4*pi, Radians(720))

	AssertFloat(t, -pi, Radians(-180))
	AssertFloat(t, -2*pi, Radians(-360))
}

func TestPointToLineDistance2D(t *testing.T) {
	a := Vec2f{0, 0}
	b := Vec2f{1, 0}

	AssertFloat(t, 0, PointToLineDistance2D(a, b, Vec2f{6, 0}))
	AssertFloat(t, 0, PointToLineDistance2D(a, b, Vec2f{-8, 0}))
	AssertFloat(t, 6, PointToLineDistance2D(a, b, Vec2f{4, 6}))
	AssertFloat(t, 2, PointToLineDistance2D(a, b, Vec2f{0, -2}))

	c := Vec2f{1, 1}
	AssertFloat(t, 5, PointToLineDistance2D(b, c, Vec2f{6, 12}))
	AssertFloat(t, 1, PointToLineDistance2D(c, b, Vec2f{0, 0}))
}

func TestPointToLineSegmentDistance2D(t *testing.T) {
	a := Vec2f{0, 0}
	b := Vec2f{10, 0}

	AssertFloat(t, 0, PointToLineSegmentDistance2D(a, b, Vec2f{6, 0}))
	AssertFloat(t, 6, PointToLineSegmentDistance2D(a, b, Vec2f{4, 6}))
	AssertFloat(t, 2, PointToLineSegmentDistance2D(a, b, Vec2f{0, -2}))

	AssertFloat(t, 2, PointToLineSegmentDistance2D(a, b, Vec2f{12, 0}))
	AssertFloat(t, 2, PointToLineSegmentDistance2D(a, b, Vec2f{10, -2}))

	AssertFloat(t, 2, PointToLineSegmentDistance2D(a, b, Vec2f{-2, 0}))
	AssertFloat(t, 2, PointToLineSegmentDistance2D(a, b, Vec2f{0, 2}))
}

func TestPolarToCartesian2D(t *testing.T) {
	AssertVec2f(t, Vec2f{0, 0}, PolarToCartesian2D(0, pi/2))
	AssertVec2f(t, Vec2f{0, 1}, PolarToCartesian2D(1, pi/2))
	AssertVec2f(t, Vec2f{-1, 0}, PolarToCartesian2D(1, pi))
	AssertVec2f(t, Vec2f{0, -1}, PolarToCartesian2D(1, 3*pi/2))
	AssertVec2f(t, Vec2f{1, 0}, PolarToCartesian2D(1, 2*pi))
}

func TestRadians(t *testing.T) {
	AssertFloat(t, 0, Radians(0))
	AssertFloat(t, pi/2, Radians(90))
	AssertFloat(t, pi, Radians(180))
	AssertFloat(t, 3*pi/2, Radians(270))
	AssertFloat(t, 2*pi, Radians(360))
	AssertFloat(t, 4*pi, Radians(720))

	AssertFloat(t, -pi, Radians(-180))
	AssertFloat(t, -2*pi, Radians(-360))
}

func TestWrapf(t *testing.T) {
	AssertFloat(t, 4, Wrapf(4, 0, 10))
	AssertFloat(t, 0, Wrapf(0, 0, 10))
	AssertFloat(t, 0, Wrapf(10, 0, 10))

	AssertFloat(t, 6, Wrapf(-4, 0, 10))
	AssertFloat(t, 7, Wrapf(17, 0, 10))
	AssertFloat(t, 3, Wrapf(103, 0, 10))

	AssertFloat(t, 6, Wrapf(6, 5, 10))
	AssertFloat(t, 8, Wrapf(3, 5, 10))
	AssertFloat(t, 7, Wrapf(12, 5, 10))

	AssertFloat(t, -6, Wrapf(-6, -10, -5))
	AssertFloat(t, -8, Wrapf(-3, -10, -5))
	AssertFloat(t, -7, Wrapf(-12, -10, -5))
	AssertFloat(t, -8, Wrapf(7, -10, -5))
}

func TestWrapi(t *testing.T) {
	assert.Equal(t, 4, Wrapi(4, 0, 10))
	assert.Equal(t, 0, Wrapi(0, 0, 10))
	assert.Equal(t, 0, Wrapi(10, 0, 10))

	assert.Equal(t, 6, Wrapi(-4, 0, 10))
	assert.Equal(t, 7, Wrapi(17, 0, 10))
	assert.Equal(t, 3, Wrapi(103, 0, 10))

	assert.Equal(t, 6, Wrapi(6, 5, 10))
	assert.Equal(t, 8, Wrapi(3, 5, 10))
	assert.Equal(t, 7, Wrapi(12, 5, 10))

	assert.Equal(t, -6, Wrapi(-6, -10, -5))
	assert.Equal(t, -8, Wrapi(-3, -10, -5))
	assert.Equal(t, -7, Wrapi(-12, -10, -5))
	assert.Equal(t, -8, Wrapi(7, -10, -5))
}
