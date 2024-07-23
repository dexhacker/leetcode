package sort_array_by_increasing_frequency

import (
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected []int
	}{
		{Nums: []int{1, 1, 2, 2, 2, 3}, Expected: []int{3, 1, 1, 2, 2, 2}},
		{Nums: []int{2, 3, 1, 3, 2}, Expected: []int{1, 3, 3, 2, 2}},
		{Nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, Expected: []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	for _, test := range tests {
		result := frequencySort(test.Nums)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Got %v, expected %v", result, test.Expected)
		}
	}
}
