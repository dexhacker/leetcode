package minimum_difference_between_largest_and_smallest_value_in_three_moves

import "testing"

func TestMinDifference(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{Nums: []int{5, 3, 2, 4}, Expected: 0},
		{Nums: []int{1, 5, 0, 10, 14}, Expected: 1},
		{Nums: []int{3, 100, 20}, Expected: 0},
	}

	for _, test := range tests {
		result := minDifference(test.Nums)
		if result != test.Expected {
			t.Errorf("minDifference(%v) = %d; expect %d", test.Nums, result, test.Expected)
		}
	}
}
