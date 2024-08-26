package structs_methods_interfaces

import "math"

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}

type Shape interface {
	Area() float64
}
