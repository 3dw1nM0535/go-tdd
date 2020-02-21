package perimeter

type Rectangle struct {
	length float64
	width  float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.length + rect.width)
}
