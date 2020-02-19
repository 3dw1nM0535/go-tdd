package array

import "testing"

func TestArrSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		sum := Sum(nums)
		expected := 15

		if sum != expected {
			t.Errorf("got '%d', given '%v', expected '%d'", sum, nums, expected)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		sum := Sum(nums)
		expected := 55

		if sum != expected {
			t.Errorf("got '%d', given '%v', expected '%d'", sum, nums, expected)
		}
	})
}
