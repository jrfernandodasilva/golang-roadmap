package reflection

import "fmt"

func Run() {
	fmt.Println("=====================================")
	fmt.Println("==============reflection=============")
	fmt.Println("=====================================")

	serializeAndUnserialize()
	dynamicFunction()
	dataValidation()
}
