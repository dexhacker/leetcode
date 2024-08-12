package kth_largest_element_in_a_stream

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		Add      []int
		Expected []int
	}{
		{Add: []int{3, 5, 10, 9, 4}, Expected: []int{4, 5, 5, 8, 8}},
	}

	for _, test := range tests {
		kLargest := Constructor(3, []int{4, 5, 8, 2})
		for i := range test.Add {
			result := kLargest.Add(test.Add[i])
			if result != test.Expected[i] {
				t.Errorf("Expected: %v, Actual: %v", test.Expected[i], result)
			}
		}
	}
}
