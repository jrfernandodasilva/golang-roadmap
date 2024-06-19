package generics

// Non-Generic Sums - sumInts sums the values of a map of ints
func sumInts(v map[string]int64) int64 {
	var result int64
	for _, v := range v {
		result += v
	}
	return result
}

// Non-Generic Sums - sumFloats sums the values of a map of floats
func sumFloats(v map[string]float64) float64 {
	var result float64
	for _, v := range v {
		result += v
	}
	return result
}

// sumGenerics sums the values of a map of generics
// this is a generic function that can be used with any type
// replace sumInts and sumFloats with this function
func sumGenerics[T int64 | float64](v map[string]T) T {
	var result T
	for _, v := range v {
		result += v
	}
	return result
}
