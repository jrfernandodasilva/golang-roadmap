package variables

import "fmt"

func Run() {
	fmt.Println("=====================================")
	fmt.Println("==============variables==============")
	fmt.Println("=====================================")

	firstName := "Fernando"
	age := 35
	height := 1.87
	isDev := true
	about := `one men...

	lost in space
	`
	arr := []int{10, 20, 30, 40, 50}   // slice
	arr2 := [5]int{10, 20, 30, 40, 50} // fixed array

	fmt.Printf("%T -> ", firstName)
	fmt.Printf("%v \n", firstName)

	fmt.Printf("%T -> ", age)
	fmt.Printf("%v \n", age)

	fmt.Printf("%T -> ", height)
	fmt.Printf("%v \n", height)

	fmt.Printf("%T -> ", isDev)
	fmt.Printf("%v \n", isDev)

	fmt.Printf("%T -> ", about)
	fmt.Printf("%v \n", about)

	fmt.Printf("%T -> ", arr)
	fmt.Printf("%v \n", arr)

	fmt.Printf("%T -> ", arr2)
	fmt.Printf("%v \n", arr2)
}
