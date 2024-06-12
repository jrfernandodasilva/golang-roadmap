package goroutines

import (
	"fmt"
	"time"
)

func print(from string) {
	for i := 0; i < 6; i++ {
		fmt.Println(from, ":", i)
	}
}

func Run() {

	fmt.Println("=====================================")
	fmt.Println("==============goroutines=============")
	fmt.Println("=====================================")

	// print a message directly
	print("direct")

	// print a message in a goroutine
	go print("goroutine")

	// print a message in a goroutine as a annonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
