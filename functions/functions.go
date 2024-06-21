package functions

import (
	"fmt"

	"github.com/jrfernandodasilva/golang-roadmap/math"
)

func Run() {
	fmt.Println("=====================================")
	fmt.Println("==============functions==============")
	fmt.Println("=====================================")

	fmt.Println("❯ math.MultiplyInt(5, math.MultiplyBase)")
	resultMultiplyInt := math.MultiplyInt(5, math.MultiplyBase)
	fmt.Printf("%T -> ", resultMultiplyInt)
	fmt.Printf("%v \n", resultMultiplyInt)

	fmt.Println("❯ math.AddInt(2, 4)")
	resultAddInt := math.AddInt(2, 4)
	fmt.Printf("%T -> ", resultAddInt)
	fmt.Printf("%v \n", resultAddInt)

	fmt.Println("❯ math.SubInt(2, 4)")
	resultSubInt := math.SubInt(2, 4)
	fmt.Printf("%T -> ", resultSubInt)
	fmt.Printf("%v \n", resultSubInt)

	fmt.Println("❯ math.MultiplyFloat(2.0, 4.0)")
	resultMultiplyFloat := math.MultiplyFloat(2.0, 4.0)
	fmt.Printf("%T -> ", resultMultiplyFloat)
	fmt.Printf("%v \n", resultMultiplyFloat)

	fmt.Println("❯ math.AddFloat(2.0, 4.0)")
	resultAddFloat := math.AddFloat(2.0, 4.0)
	fmt.Printf("%T -> ", resultAddFloat)
	fmt.Printf("%v \n", resultAddFloat)

	fmt.Println("❯ math.SubFloat(7.0, 4.0)")
	resultSubFloat := math.SubFloat(7.0, 4.0)
	fmt.Printf("%T -> ", resultSubFloat)
	fmt.Printf("%v \n", resultSubFloat)

	fmt.Println("❯ Sum(1, 2)")
	resultSum := Sum(1, 2)
	fmt.Printf("%T -> ", resultSum)
	fmt.Printf("%v \n", resultSum)

	fmt.Println("❯ ReturnNamedVar()")
	fmt.Println(ReturnNamedVar())

	fmt.Println("❯ SumAll(10, 20, 30, 40, 50, 5)")
	fmt.Println(SumAll(10, 20, 30, 40, 50, 5))

	fmt.Println("❯ Anonymous Function")
	fmt.Println("❯ result := func(v ...int) func() int {...")
	result := func(v ...int) func() int {
		res := 0

		for _, value := range v {
			res += value
		}

		return func() int {
			return res * res
		}
	}

	fmt.Printf("%T -> ", result)
	fmt.Printf("%v \n", result()())
	fmt.Printf("%T -> ", result(1, 2, 3, 4, 5)())
	fmt.Printf("%v \n", result(1, 2, 3, 4, 5)())
}

// Sum returns the sum of two int values
func Sum(a int, b int) int {
	return a + b
}

// Sum all int values
func SumAll(v ...int) int {
	var sum int
	for _, value := range v {
		sum += value
	}
	return sum
}

// Return a named variable
func ReturnNamedVar() (result string) {
	result = "Hello"
	return
}
