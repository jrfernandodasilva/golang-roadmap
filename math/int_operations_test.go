package math_test

import (
	"golang-roadmap/math"
	"testing"
)

func TestMultiplyInt(t *testing.T) {
	// Define test values
	a := 5
	b := 3
	expected := 15

	// Call function to test
	result := math.MultiplyInt(a, b)

	// Check result if expected
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestAddInt(t *testing.T) {
	// Define test values
	a := 5
	b := 3
	expected := 8

	// Call function to test
	result := math.AddInt(a, b)

	// Check result if expected
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSubInt(t *testing.T) {
	// Define test values
	a := 5
	b := 3
	expected := 2

	// Call function to test
	result := math.SubInt(a, b)

	// Check result if expected
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Define test values to new test
	a = 3
	b = 6
	expected = -3

	// Call function to test
	result = math.SubInt(a, b)

	// Check result if expected
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
