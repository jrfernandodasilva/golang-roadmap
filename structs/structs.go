package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func (p Person) Show() {
	fmt.Printf("FirstName: %v, LastName: %v, Email: %v, Age: %v\n", p.FirstName, p.LastName, p.Email, p.Age)
}

type PersonInternational struct {
	Person
	Country string `json:"country"`
}

func (pi PersonInternational) Show() {
	fmt.Printf("FirstName: %v, LastName: %v, Email: %v, Age: %v, Country: %v\n", pi.FirstName, pi.LastName, pi.Email, pi.Age, pi.Country)
}

func Run() {
	fmt.Println("=====================================")
	fmt.Println("===============structs===============")
	fmt.Println("=====================================")

	person := Person{
		FirstName: "Fernando",
		LastName:  "Jr",
		Email:     "fernandojr@github.com",
		Age:       35,
	}

	personJson, err := json.Marshal(person)
	if err != nil {
		log.Fatal("Error marshalling person:", err)
		return
	}

	person2 := Person{"Levi", "Jr", "levijr@github.com", 1}

	person2Json, err := json.Marshal(person2)
	if err != nil {
		log.Fatal("Error marshalling person2:", err)
		return
	}

	fmt.Printf("%T -> ", person)
	fmt.Printf("%v \n", person)
	person.Show()
	fmt.Println(string(personJson) + "\n")

	fmt.Printf("%T -> ", person2)
	fmt.Printf("%v \n", person2)
	person2.Show()
	fmt.Println(string(person2Json) + "\n")

	personInternational := PersonInternational{
		Person: Person{
			FirstName: "Rafael",
			LastName:  "Jr",
			Email:     "rafaeljr@github.com",
			Age:       13,
		},
		Country: "United States",
	}

	personInternationalJson, err := json.Marshal(personInternational)
	if err != nil {
		log.Fatal("Error marshalling personInternational:", err)
		return
	}

	fmt.Printf("%T -> ", personInternational)
	fmt.Printf("%v \n", personInternational)
	personInternational.Show()
	fmt.Println(string(personInternationalJson) + "\n")

	personInternational2Json := `{"first_name":"Baruc","last_name":"Jr","email":"barucjr@github.com","age":18,"country":"Uruguay"}`
	personInternational2 := PersonInternational{}

	err = json.Unmarshal([]byte(personInternational2Json), &personInternational2)
	if err != nil {
		log.Fatal("Error unmarshalling personInternational2:", err)
		return
	}

	fmt.Printf("%T -> ", personInternational2)
	fmt.Printf("%v \n", personInternational2)
	personInternational2.Show()
	fmt.Println(personInternational2Json)
}
