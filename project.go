package vmath

// Ortho returns an orthographic projection matrix.
func Ortho(left, right, bottom, top float32, near, far float32) Mat4f {
	return Mat4f{
		2 / (right - left), 0, 0, 0,
		0, 2 / (top - bottom), 0, 0,
		0, 0, -2 / (far - near), 0,

		-(right + left) / (right - left),
		-(top + bottom) / (top - bottom),
		-(far + near) / (far - near),
		1,
	}
}

// UnOrtho returns an orthographic unprojection matrix.
func UnOrtho(left, right, bottom, top float32, near, far float32) Mat4f {
	return Mat4f{
		(right - left) / 2, 0, 0, 0,
		0, (top - bottom) / 2, 0, 0,
		0, 0, (far - near) / -2, 0,

		(left + right) / 2,
		(top + bottom) / 2,
		(far + near) / -2,
		1,
	}
}

// Perspective returns a perspective projection matrix.
// The field-of-view is in radians.
func Perspective(fovY, aspectRatio, near, far float32) Mat4f {
	xHalf := Tan(fovY/2) * near
	yHalf := xHalf * aspectRatio
	return Frustum(-yHalf, yHalf, -xHalf, xHalf, near, far)
}

// Frustum returns a frustum matrix.
func Frustum(left, right, bottom, top float32, near, far float32) Mat4f {
	invX := 1 / (right - left)
	invY := 1 / (top - bottom)
	invZ := 1 / (near - far)

	return Mat4f{
		near * 2 * invX, 0, 0, 0,
		0, near * 2 * invY, 0, 0,

		(right + left) * invX,
		(bottom + top) * invY,
		(near + far) * invZ,
		-1,

		0, 0, far * near * 2 * invZ, 0,
	}
}

// LookAt returns a transformation matrix for a viewer, looking towards the target, with a defined upwards vector.
func LookAt(eye, target, up Vec3f) Mat4f {
	// source: https://stackoverflow.com/a/352957/2224996

	forward := target.Sub(eye).Normalize()
	right := forward.Cross(up).Normalize()
	up = right.Cross(forward)

	res := Mat4f{
		right[0], up[0], -forward[0], 0,
		right[1], up[1], -forward[1], 0,
		right[2], up[2], -forward[2], 0,
		0, 0, 0, 1,
	}
	transMat := Mat4fFromTranslation(eye.Invert())
	return res.Mul(transMat)
}
