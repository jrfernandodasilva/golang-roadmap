package reflection

import (
	"fmt"
	"reflect"
)

// Validator struct to hold validation criteria
type Validator struct {
	ExpectedType string
	IsRequired   bool
	MinSize      int
	MaxSize      int
}

// Validate value against the provided Validator criteria
func validate(value any, validator Validator) bool {
	val := reflect.ValueOf(value)

	// Check if value is required and empty
	if validator.IsRequired && (val.Kind() == reflect.Ptr && val.IsNil() || val.Kind() == reflect.String && val.Len() == 0 || !val.IsValid()) {
		return false
	}

	// Type verification
	if validator.ExpectedType != "" {
		currentType := val.Kind().String()
		if currentType != validator.ExpectedType {
			return false
		}
	}

	// Minimum and maximum size verification
	switch validator.ExpectedType {
	case "string":
		// Minimum size verification for string
		if validator.MinSize > 0 && val.Len() < validator.MinSize {
			return false
		}

		// Maximum size verification for string
		if validator.MaxSize > 0 && val.Len() > validator.MaxSize {
			return false
		}
	case "int":
		intValue := val.Int()

		// Minimum size (value) verification for int
		if validator.MinSize > 0 && intValue < int64(validator.MinSize) {
			return false
		}

		// Maximum size (value) verification for int
		if validator.MaxSize > 0 && intValue > int64(validator.MaxSize) {
			return false
		}
	}

	return true
}

func dataValidation() {
	fmt.Println("❯ reflection.dataValidation()")

	name := "Fernando"
	age := 35

	// Validating if name is string, required, with minimum 3 and maximum 50 characters
	nameValidator := Validator{
		ExpectedType: "string",
		IsRequired:   true,
		MinSize:      3,
		MaxSize:      30,
	}

	if !validate(name, nameValidator) {
		fmt.Println("Invalid name (must be a string with ", nameValidator.MinSize, " to ", nameValidator.MaxSize, " characters)")
		return
	}

	// Validating if age is int and greater than 0
	ageValidator := Validator{
		ExpectedType: "int",
		IsRequired:   true,
		MinSize:      18,
	}

	if !validate(age, ageValidator) {
		fmt.Println("Invalid age (must be an integer and greater than", ageValidator.MinSize, ")")
		return
	}

	fmt.Println("Valid data:", name, age)
}
