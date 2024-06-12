package sort_colors

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected []int
	}{
		{Nums: []int{2, 0, 2, 1, 1, 0}, Expected: []int{0, 0, 1, 1, 2, 2}},
		{Nums: []int{2, 0, 1}, Expected: []int{0, 1, 2}},
		{Nums: []int{0, 0, 1, 1, 2, 2}, Expected: []int{0, 0, 1, 1, 2, 2}},
	}

	for _, test := range tests {
		sortColors(test.Nums)
		for i := 0; i < len(test.Nums); i++ {
			if test.Nums[i] != test.Expected[i] {
				fmt.Printf("Nums: %v, Expected: %v\n", test.Nums[i], test.Expected[i])
			}
		}
	}
}
