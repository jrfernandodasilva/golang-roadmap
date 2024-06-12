package main

import (
	"fmt"
	"golang-roadmap/api"
	"golang-roadmap/blank_identifier"
	"golang-roadmap/channels"
	"golang-roadmap/functions"
	"golang-roadmap/goroutines"
	"golang-roadmap/methods"
	"golang-roadmap/pointers"
	"golang-roadmap/structs"
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
	structs.Run()
	goroutines.Run()
	channels.Run()
	api.Run()
}
