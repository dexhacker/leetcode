package sort_an_array

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected []int
	}{
		{Nums: []int{5, 2, 3, 1}, Expected: []int{1, 2, 3, 5}},
		{Nums: []int{5, 1, 1, 2, 0, 0}, Expected: []int{0, 0, 1, 1, 2, 5}},
		{
			Nums:     []int{4, 7, 8, 9, 2, 1, 3, 7, 9, 8, 6, 2, 2, 7, 3, 6, 4, 8, 7, 2, 8, 2, 1, 3, 9, 8, 9},
			Expected: []int{1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 4, 4, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 8, 9, 9, 9, 9},
		},
	}

	for _, test := range tests {
		result := sortArray(test.Nums)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("sort array expect %v got %v", test.Expected, result)
		}
	}
}
