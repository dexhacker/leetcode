package sum_of_all_subset_xor_totals

import "testing"

func TestSubsetXORSum(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{Nums: []int{1, 3}, Expected: 6},
		{Nums: []int{5, 1, 6}, Expected: 28},
		{Nums: []int{3, 4, 5, 6, 7, 8}, Expected: 480},
		{Nums: []int{3, 4, 5, 6, 7, 8, 4, 5}, Expected: 1920},
	}
	for _, test := range tests {
		result := subsetXORSum(test.Nums)
		if result != test.Expected {
			t.Errorf("subsetXORSum(%+v) = %d; expected %d", test.Nums, test.Expected, result)
		}
	}
}
