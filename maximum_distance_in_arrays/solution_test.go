package maximum_distance_in_arrays

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		Arrays   [][]int
		Expected int
	}{
		{Arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, Expected: 4},
		{Arrays: [][]int{{1}, {1}}, Expected: 0},
	}

	for _, test := range tests {
		result := maxDistance(test.Arrays)
		if result != test.Expected {
			t.Errorf("maxDistance(%v) = %d; expected %d", test.Arrays, result, test.Expected)
		}
	}
}
