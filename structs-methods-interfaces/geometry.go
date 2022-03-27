package structs_methods_n_interfaces

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Lenght float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Lenght * t.Height) / 2
}

type Rectangle struct {
	Lenght float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Lenght * r.Height
}

type Circle struct {
	Radiu float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radiu * c.Radiu
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Lenght + rectangle.Height)
}
