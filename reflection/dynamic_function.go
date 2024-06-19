package reflection

import (
	"fmt"
	"reflect"
)

func createFunc(returnType reflect.Type, params []reflect.Type) reflect.Value {
	// Example function to sum two int numbers
	fn := func(args []reflect.Value) []reflect.Value {
		a := args[0].Int()
		b := args[1].Int()
		sum := a + b
		return []reflect.Value{reflect.ValueOf(int(sum))}
	}

	// Create function type dynamically
	funcType := reflect.FuncOf(params, []reflect.Type{returnType}, false)

	// Create a reflect value for the function
	return reflect.MakeFunc(funcType, fn)
}

func dynamicFunction() {
	fmt.Println("❯ reflection.dynamicFunction()")

	// Creating a dinamic function to sum two int numbers
	sum := createFunc(reflect.TypeOf(1), []reflect.Type{reflect.TypeOf(1), reflect.TypeOf(1)})

	num1 := reflect.ValueOf(10)
	num2 := reflect.ValueOf(20)

	// Calling the dinamic function
	result := sum.Call([]reflect.Value{num1, num2})

	fmt.Println("Sum:", result[0].Int())
}
