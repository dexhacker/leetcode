package the_number_of_beautiful_subsets

import "testing"

func TestBeautifulSubsets(t *testing.T) {
	tests := []struct {
		Nums     []int
		K        int
		Expected int
	}{
		{Nums: []int{2, 4, 6}, K: 2, Expected: 4},
		{Nums: []int{1}, K: 1, Expected: 1},
		{Nums: []int{2, 4, 6, 9, 12, 14, 20, 30, 80, 82, 100, 150}, K: 2, Expected: 1439},
		{Nums: []int{2, 4, 6, 9, 12, 14, 20, 30, 80, 82, 100, 150, 156, 160, 170, 180, 19, 20, 21}, K: 2, Expected: 138239},
		{Nums: []int{14, 87, 37, 94, 43, 25, 77, 54, 44, 71, 39, 16, 61, 29, 61, 9, 96}, K: 88, Expected: 131071},
	}

	for _, test := range tests {
		result := beautifulSubsets(test.Nums, test.K)
		if result != test.Expected {
			t.Errorf("input: %v, expect: %v, got: %v", test.Nums, test.Expected, result)
		}
	}
}
