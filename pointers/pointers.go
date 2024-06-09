package pointers

import "fmt"

type Car struct {
	Name string
}

func (c Car) running() {
	c.Name = "Oroch"
	fmt.Println(c.Name, "Running")
}

func (c *Car) runningp() {
	c.Name = "Oroch"
	fmt.Println(c.Name, "Running")
}

func Run() {
	fmt.Println("=====================================")
	fmt.Println("===============pointers==============")
	fmt.Println("=====================================")

	// Declaring a pointer to string
	var name string = "Fernando"
	var namep *string = &name // namep points to name memory address

	age := 34
	var agep *int = &age
	*agep = 35 // update age value using pointer dereferencing

	height := 1.86
	heightp := &height
	*heightp = 1.87

	isDev := false

	fmt.Println("❯ vars")
	fmt.Println(name, age, height, isDev)

	becomeAnDev(&isDev)

	fmt.Println("❯ dereferencing pointers")
	fmt.Println(*namep, *agep, *heightp, isDev)

	c := Car{Name: "BMW"}
	fmt.Println("❯ struct method WITHOUT pointer")
	c.running()
	fmt.Println("Car Name:", c.Name)
	fmt.Println("❯ struct method WITH pointer")
	c.runningp()
	fmt.Println("Car Name:", c.Name)
}

func becomeAnDev(v *bool) {
	*v = true
}
