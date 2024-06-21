package main

import (
	"fmt"

	"github.com/jrfernandodasilva/golang-roadmap/api"
	"github.com/jrfernandodasilva/golang-roadmap/blank_identifier"
	"github.com/jrfernandodasilva/golang-roadmap/channels"
	"github.com/jrfernandodasilva/golang-roadmap/functions"
	"github.com/jrfernandodasilva/golang-roadmap/generics"
	"github.com/jrfernandodasilva/golang-roadmap/goroutines"
	"github.com/jrfernandodasilva/golang-roadmap/methods"
	"github.com/jrfernandodasilva/golang-roadmap/pointers"
	"github.com/jrfernandodasilva/golang-roadmap/process"
	"github.com/jrfernandodasilva/golang-roadmap/reflection"
	"github.com/jrfernandodasilva/golang-roadmap/structs"
	"github.com/jrfernandodasilva/golang-roadmap/uuid_generator"
	"github.com/jrfernandodasilva/golang-roadmap/variables"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("================main=================")
	fmt.Println("=====================================")
	fmt.Println("call another packages")

	process.Run()
	variables.Run()
	uuid_generator.Run()
	blank_identifier.Run()
	methods.Run()
	functions.Run()
	pointers.Run()
	structs.Run()
	reflection.Run()
	generics.Run()
	goroutines.Run()
	channels.Run()
	api.Run()
}
