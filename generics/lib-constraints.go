package generics

import "golang.org/x/exp/constraints"

func max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}

	return b
}

func min[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	}

	return b
}
