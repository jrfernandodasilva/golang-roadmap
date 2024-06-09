package uuid_generator

import (
	"fmt"

	"github.com/google/uuid"
)

func Run() {
	fmt.Println("=====================================")
	fmt.Println("================uuid=================")
	fmt.Println("=====================================")
	newUUID := uuid.NewString()
	fmt.Println("newUUID -> " + newUUID)

	// Parse a UUID string
	parsedUUID, err := uuid.Parse("73b1812a-b9e4-4662-a7c7-5960242b191f")
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}

	fmt.Println("❯ parsedUUID")
	fmt.Println(parsedUUID)

	fmt.Println("❯ uuid.Validate(\"73b1812a-b9e4-4662-a7c7-5960242b1911\")")
	check := uuid.Validate("73b1812a-b9e4-4662-a7c7-5960242b1911")
	fmt.Println(check)
}
