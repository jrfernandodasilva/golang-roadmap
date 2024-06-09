package functions_test

import (
	"golang-roadmap/functions"
	"testing"
)

func TestSum(t *testing.T) {
	a := 10
	b := 20
	expected := 30

	result := functions.Sum(a, b)

	if result != expected {
		t.Errorf("Sum(%d, %d) = %d, expected %d", a, b, result, expected)
	}
}

func TestSumAll(t *testing.T) {
	values := []int{10, 20, 5, 1, 4}
	expected := 40

	result := functions.SumAll(values...)

	if result != expected {
		t.Errorf("SumAll(%v) = %d, expected %d", values, result, expected)
	}
}

func TestReturnNamedVar(t *testing.T) {
	expected := "Hello"
	result := functions.ReturnNamedVar()

	if result != expected {
		t.Errorf("ReturnNamedVar() = %s, expected %s", result, expected)
	}
}
