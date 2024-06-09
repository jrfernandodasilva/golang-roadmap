package main

import (
	"fmt"
	"golang-roadmap/blank_identifier"
	"golang-roadmap/functions"
	"golang-roadmap/methods"
	"golang-roadmap/pointers"
	"golang-roadmap/uuid_generator"
	"golang-roadmap/variables"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("================main=================")
	fmt.Println("=====================================")
	fmt.Println("call another packages")

	variables.Run()
	uuid_generator.Run()
	blank_identifier.Run()
	methods.Run()
	functions.Run()
	pointers.Run()
}
