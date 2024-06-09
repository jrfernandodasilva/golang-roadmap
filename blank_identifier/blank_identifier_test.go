package blank_identifier_test

import (
	blankid "golang-roadmap/blank_identifier"
	"testing"
)

func TestSum(t *testing.T) {
	a := 10
	b := 15
	expected := 25

	result, err := blankid.Sum(a, b)

	if err != nil {
		t.Errorf("Sum(%d, %d) = %v, expected %d", a, b, err, expected)
	}

	if result != expected {
		t.Errorf("Sum(%d, %d) = %d, expected %d", a, b, result, expected)
	}
}

func TestSum_(t *testing.T) {
	a := 10
	b := -15
	expected := 0

	result, _ := blankid.Sum(a, b)

	if result != expected {
		t.Errorf("Sum(%d, %d) = %d, expected %d", a, b, result, expected)
	}
}
