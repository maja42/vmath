package mathi

// Abs returns the absolute value of v.
func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// Min returns the smaller integer value.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Min returns the smaller integer value.
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
