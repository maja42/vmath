package vmath

// Absi returns the absolute value of v.
func Absi(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// Mini returns the smaller integer value.
func Mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Mini returns the smaller integer value.
func Maxi(a, b int) int {
	if a < b {
		return b
	}
	return a
}
