package blank_identifier

import (
	"errors"
	"fmt"
)

func Run() {
	fmt.Println("=====================================")
	fmt.Println("==========blank_identifier===========")
	fmt.Println("=====================================")

	fmt.Println("❯ Sum(1, -15)")
	result, _ := Sum(1, -15)
	fmt.Printf("%T -> ", result)
	fmt.Printf("%v \n", result)

	fmt.Println("❯ for _, value := range arr {...")
	arr := [5]int{10, 20, 30, 40, 50}
	for _, value := range arr {
		fmt.Printf("%T -> ", value)
		fmt.Printf("%v \n", value)
	}
}

// sum returns the sum of two int values
func Sum(a int, b int) (int, error) {

	result := a + b

	if result < 0 {
		return 0, errors.New("Negative number")
	}

	return a + b, nil
}
