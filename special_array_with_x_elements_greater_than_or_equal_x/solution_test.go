package special_array_with_x_elements_greater_than_or_equal_x

import "testing"

func TestSpecialArray(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{Nums: []int{3, 5}, Expected: 2},
		{Nums: []int{0, 0}, Expected: -1},
		{Nums: []int{0, 4, 3, 0, 4}, Expected: 3},
		{Nums: []int{3, 5, 4, 6, 7, 8, 9}, Expected: 5},
		{Nums: []int{3, 6, 7, 7, 0}, Expected: -1},
	}

	for _, test := range tests {
		result := specialArray(test.Nums)
		if result != test.Expected {
			t.Errorf("test %v: wanted %v, got %v", test.Nums, test.Expected, result)
		}
	}
}
