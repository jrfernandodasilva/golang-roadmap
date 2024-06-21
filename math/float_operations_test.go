package math_test

import (
	"math"
	"testing"

	my_math "github.com/jrfernandodasilva/golang-roadmap/math"
)

func TestMultiplyFloat(t *testing.T) {
	// Define test values
	a := 5.1
	b := 3.0
	expected := 15.3

	// Call function to test
	result := my_math.MultiplyFloat(a, b)

	// Check result if expected
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("Expected %f, got %f, diff %f", expected, result, math.Abs(result-expected))
	}
}

func TestAddFloat(t *testing.T) {
	// Define test values
	a := 5.1
	b := 3.3
	expected := 8.4

	// Call function to test
	result := my_math.AddFloat(a, b)

	// Check result if expected
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("Expected %f, got %f, diff %f", expected, result, math.Abs(result-expected))
	}
}

func TestSubFloat(t *testing.T) {
	// Define test values
	a := 5.0
	b := 3.1
	expected := 1.9

	// Call function to test
	result := my_math.SubFloat(a, b)

	// Check result if expected
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("Expected %f, got %f, diff %f", expected, result, math.Abs(result-expected))
	}

	// Define test values to new test
	a = 3.0
	b = 7.2
	expected = -4.2

	// Call function to test
	result = my_math.SubFloat(a, b)

	// Check result if expected
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
