package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 1)
	expected := "a"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
