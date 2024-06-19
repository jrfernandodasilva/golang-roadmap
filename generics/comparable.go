package generics

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var result V
	for _, v := range m {
		result += v
	}
	return result
}

// getIfEqual returns the first value if it is equal to the second value.
func getIfEqual[T comparable](a T, b T) T {
	if a == b {
		return a
	}

	return b
}
