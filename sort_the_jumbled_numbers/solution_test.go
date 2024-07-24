package sort_the_jumbled_numbers

import (
	"reflect"
	"testing"
)

func TestSortJumbled(t *testing.T) {
	tests := []struct {
		Mapping  []int
		Nums     []int
		Expected []int
	}{
		{Mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, Nums: []int{991, 338, 38}, Expected: []int{338, 38, 991}},
		{Mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, Nums: []int{789, 456, 123}, Expected: []int{123, 456, 789}},
	}

	for _, test := range tests {
		result := sortJumbled(test.Mapping, test.Nums)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("expected: %v, got: %v", test.Expected, result)
		}
	}
}
