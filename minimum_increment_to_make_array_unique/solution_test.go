package minimum_increment_to_make_array_unique

import (
	"fmt"
	"testing"
)

func TestMinIncrementForUnique(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{Nums: []int{1, 2, 2}, Expected: 1},
		{Nums: []int{3, 2, 1, 2, 1, 7}, Expected: 6},
		{Nums: []int{1, 1, 2, 3, 4, 5}, Expected: 5},
	}

	for _, test := range tests {
		result := minIncrementForUnique(test.Nums)
		if result != test.Expected {
			fmt.Printf("Expected: %d, got: %d\n", test.Expected, result)
		}
	}
}
