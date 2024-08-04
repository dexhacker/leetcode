package range_sum_of_sorted_subarray_sums

import "testing"

func TestRangeSum(t *testing.T) {
	tests := []struct {
		Nums     []int
		N        int
		Left     int
		Right    int
		Expected int
	}{
		{Nums: []int{1, 2, 3, 4}, N: 4, Left: 1, Right: 5, Expected: 13},
		{Nums: []int{1, 2, 3, 4}, N: 4, Left: 3, Right: 4, Expected: 6},
		{Nums: []int{1, 2, 3, 4}, N: 4, Left: 1, Right: 10, Expected: 50},
	}

	for _, test := range tests {
		result := rangeSum(test.Nums, test.N, test.Left, test.Right)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d", test.Expected, result)
		}
	}
}
