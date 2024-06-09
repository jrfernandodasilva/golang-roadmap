package methods

import "fmt"

type Car struct {
	Name string
}

func (c *Car) PrintName() {
	fmt.Println(c.Name)
}

func (c *Car) Run() {
	fmt.Println("Running")
}

func Run() {
	fmt.Println("=====================================")
	fmt.Println("===============methods===============")
	fmt.Println("=====================================")

	car := Car{Name: "Ferrari"}

	fmt.Println("❯ car.PrintName()")
	car.PrintName()

	fmt.Println("❯ car.Run()")
	car.Run()
}
