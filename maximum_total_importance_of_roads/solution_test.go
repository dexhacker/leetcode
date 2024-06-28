package maximum_total_importance_of_roads

import "testing"

func TestMaximumImportance(t *testing.T) {
	tests := []struct {
		N        int
		Roads    [][]int
		Expected int64
	}{
		{N: 5, Roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}, Expected: 43},
		{N: 5, Roads: [][]int{{0, 3}, {2, 4}, {1, 3}}, Expected: 20},
	}

	for _, test := range tests {
		result := maximumImportance(test.N, test.Roads)
		if result != test.Expected {
			t.Errorf("Expected: %d, Actual: %d", test.Expected, result)
		}
	}
}
