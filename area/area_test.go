package area

import "testing"

func TestArea(t *testing.T) {
	t.Run("Area for Rectangle", func(t *testing.T) {
		sq := Figure{12.0, 6.0}
		area := sq.Area()
		expected := 72.0

		if area != expected {
			t.Errorf("got '%.2f' but expected '%.2f'", area, expected)
		}
	})
	t.Run("Area for Circle", func(t *testing.T) {
		circle := Circle{10}
		circleArea := circle.Area()
		expected := 314.1592653589793

		if circleArea != expected {
			t.Errorf("got %g but expected %g", circleArea, expected)
		}
	})
}
