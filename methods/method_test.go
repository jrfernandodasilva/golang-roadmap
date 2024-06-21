package methods_test

import (
	"testing"

	"github.com/jrfernandodasilva/golang-roadmap/methods"
)

func TestCar(t *testing.T) {
	car := methods.Car{Name: "Oroch"}
	carName := car.Name
	expected := "Oroch"

	if carName != expected {
		t.Errorf("Expected %s, got %s", expected, carName)
	}
}
