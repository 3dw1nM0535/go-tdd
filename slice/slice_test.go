package slice

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	sum := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("got %v sum but expected %v", sum, expected)
	}
}
