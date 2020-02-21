package area

import "math"

type Figure struct {
	length float64
	width  float64
}

func (r Figure) Area() float64 {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}
