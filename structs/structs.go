package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return 3.141592653589793 * c.Radius * c.Radius
}
