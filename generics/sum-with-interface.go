package generics

// interfaces can be used to define generic functions
// using constraints can be used to define interfaces that are more restrictive
// ~int do approximate int types
type Number interface {
	~int | int32 | int64 | float32 | float64
}

// sumGenericsWithInterface sums the values of a map of generics
func sumGenericsWithInterface[T Number](v map[string]T) T {
	var result T
	for _, v := range v {
		result += v
	}
	return result
}
