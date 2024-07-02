package intersection_of_two_arrays_ii

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		Nums1    []int
		Nums2    []int
		Expected []int
	}{
		{Nums1: []int{1, 2, 2, 1}, Nums2: []int{2, 2}, Expected: []int{2, 2}},
		{Nums1: []int{4, 9, 5}, Nums2: []int{9, 4, 9, 8, 4}, Expected: []int{4, 9}},
	}

	for _, test := range tests {
		result := intersect(test.Nums1, test.Nums2)
		for i := range result {
			fmt.Printf("Expected: %v, Actual: %v\n", test.Expected[i], result[i])
		}
	}
}
