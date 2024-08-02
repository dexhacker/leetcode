package minimum_swaps_to_group_all_1s_together_ii

import "testing"

func TestMinSwaps(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{Nums: []int{0, 1, 0, 1, 1, 0, 0}, Expected: 1},
		{Nums: []int{0, 1, 1, 1, 0, 0, 1, 1, 0}, Expected: 2},
		{Nums: []int{1, 1, 0, 0, 1}, Expected: 0},
		{Nums: []int{1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, Expected: 2},
		{Nums: []int{0, 1, 0, 0, 1, 0, 0, 0, 1}, Expected: 1},
	}

	for _, test := range tests {
		result := minSwaps(test.Nums)
		if result != test.Expected {
			t.Errorf("minSwaps(%v) = %d, want %d", test.Nums, result, test.Expected)
		}
	}
}
