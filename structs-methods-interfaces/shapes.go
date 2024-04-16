package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter calculates and returns the perimeter of a Rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

// Area calculates the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area calculates the area of a Triangle.
func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}
