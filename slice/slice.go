package slice

import (
	"github.com/3dw1nM0535/go-tdd/array"
)

// SumAll return the sum of slices

// SumAll return the sum of slices
func SumAll(numbersToSum ...[]int) (sum []int) {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, array.Sum(numbers))
	}

	return sums
}
