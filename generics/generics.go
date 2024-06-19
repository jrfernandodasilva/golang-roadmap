package generics

import "fmt"

func Run() {
	fmt.Println("=====================================")
	fmt.Println("===============generics==============")
	fmt.Println("=====================================")

	fmt.Println("❯ generics start")
	fmt.Println("sumInts:", sumInts(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println("sumFloats:", sumFloats(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}))
	fmt.Println("sumGenerics:", sumGenerics(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println("sumGenerics:", sumGenerics(map[string]float64{"a": 1.5, "b": 2.3, "c": 3.6}))
	fmt.Println("sumGenericsWithInterface:", sumGenericsWithInterface(map[string]int64{"a": 1, "b": 2, "c": 4}))
	fmt.Println("sumGenericsWithInterface:", sumGenericsWithInterface(map[string]float64{"a": 1.5, "b": 2.2, "c": 3.6}))

	var x, y, z MyNumber = 1, 3, 5
	fmt.Println("sumGenericsWithCustomType:", sumGenericsWithCustomType(map[string]MyNumber{"a": x, "b": y, "c": z}))

	fmt.Println("comparable.getIfEqual:", getIfEqual(5, 5))
	fmt.Println("comparable.getIfEqual:", getIfEqual(5, 11))
	fmt.Println("comparable.getIfEqual:", getIfEqual("Fernando", "Jr"))

	fmt.Println("lib-constraints.max:", max(5, 12))
	fmt.Println("lib-constraints.max:", max(111, 110))
	fmt.Println("lib-constraints.min:", min(5, 12))
	fmt.Println("lib-constraints.min:", min(111, 110))

	fmt.Println("in-method.concat:", concat([]MyString{"brazil", "-", "2024"}))

	fmt.Println("❯ generics end")
}
