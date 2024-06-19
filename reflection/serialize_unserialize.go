package reflection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Person represents a person with basic information
type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func serializeAndUnserialize() {

	fmt.Println("❯ reflection.serializeAndUnserialize()")

	person := Person{
		Name:       "Fernando",
		Age:        35,
		Occupation: "Developer",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error on JSON convert:", err)
		return
	}

	value := reflect.ValueOf(person)
	typeOf := value.Type()

	fmt.Printf("Type of p: %s and kind: %s\n", typeOf, typeOf.Kind())
	fmt.Println("Json:", string(jsonData))

	var deserializedPerson Person
	err = json.Unmarshal(jsonData, &deserializedPerson)
	if err != nil {
		fmt.Println("Error on JSON unserialize:", err)
		return
	}

	fmt.Printf("Unserialized people type: %s and kind: %s\n", reflect.TypeOf(deserializedPerson), reflect.TypeOf(deserializedPerson).Kind())
	fmt.Println("Unserialized person:", deserializedPerson)
}
