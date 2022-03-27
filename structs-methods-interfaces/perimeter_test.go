package structs_methods_n_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	lenght := 10.0
	height := 10.0

	rectangle := Rectangle{lenght, height}
	response := Perimeter(rectangle)

	expected_response := 40.0

	if response != expected_response {
		t.Errorf("response %.2f expected %.2f", response, expected_response)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, expected_response float64) {
		t.Helper()

		response := shape.Area()

		if response != expected_response {
			t.Errorf("response %.10f expected %.10f", response, expected_response)
		}
	}

	t.Run("When a rectangle is given, return a rectangle area", func(t *testing.T) {
		lenght := 10.0
		height := 10.0
		rectangle := Rectangle{lenght, height}

		expected_response := 100.0

		verifyArea(t, rectangle, expected_response)
	})

	t.Run("When a circle is given, return a circle area", func(t *testing.T) {
		radiu := 10.0
		circle := Circle{radiu}

		expected_response := 314.1592653589793

		verifyArea(t, circle, expected_response)
	})
}
