package path_with_maximum_gold

import "testing"

func TestGetMaximumGold(t *testing.T) {
	tests := []struct {
		Grid     [][]int
		Expected int
	}{
		{Grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}, Expected: 24},
		{Grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}, Expected: 28},
	}

	for _, test := range tests {
		result := getMaximumGold(test.Grid)
		if result != test.Expected {
			t.Errorf("getMaximumGold(%+v) = %v; expected %v", test.Grid, result, test.Expected)
		}
	}
}
