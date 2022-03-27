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
	testsArea := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Lenght: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radiu: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Lenght: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			response := tt.shape.Area()

			if response != tt.hasArea {
				t.Errorf("%#v response %.2f expected %.2f", tt.shape, response, tt.hasArea)
			}
		})
	}
}
