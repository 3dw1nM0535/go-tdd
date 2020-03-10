package area

import "math"

// Shape : interface for area
type Shape interface {
	Area() float64
}

// Rectangle : shape reactangle
type Rectangle struct {
	length float64
	width  float64
}

// Triangle : shape triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area : find area of a triangle
func (r Triangle) Area() float64 {
	return 0.5 * r.Base * r.Height
}

// Area : find area of a rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Circle : shape circle
type Circle struct {
	radius float64
}

// Area : find area of a circle
func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}
