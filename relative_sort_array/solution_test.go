package relative_sort_array

import "testing"

func TestRelativeSortArray(t *testing.T) {
	tests := []struct {
		Arr1     []int
		Arr2     []int
		Expected []int
	}{
		{
			Arr1:     []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			Arr2:     []int{2, 1, 4, 3, 9, 6},
			Expected: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			Arr1:     []int{28, 6, 22, 8, 44, 17},
			Arr2:     []int{22, 28, 8, 6},
			Expected: []int{22, 28, 8, 6, 17, 44},
		},
	}

	for _, test := range tests {
		result := relativeSortArray(test.Arr1, test.Arr2)
		for i := 0; i < len(result); i++ {
			if result[i] != test.Expected[i] {
				t.Errorf("expected: %v, got: %v", test.Expected[i], result[i])
			}
		}
	}
}
