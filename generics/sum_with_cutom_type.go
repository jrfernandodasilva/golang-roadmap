package generics

type MyNumber int

// sumGenericsWithInterface sums the values of a map of generics
func sumGenericsWithCustomType[T MyNumber](v map[string]MyNumber) MyNumber {
	return sumGenericsWithInterface(v)
}
