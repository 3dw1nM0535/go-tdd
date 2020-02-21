package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	peri := Perimeter(rect)
	expected := 40.0

	if peri != expected {
		t.Errorf("got '%.2f' but expected '%.2f'", rect, expected)
	}
}
