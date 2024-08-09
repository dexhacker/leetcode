package magic_squares_in_grid

import "testing"

func TestNumMagicSquaresInside(t *testing.T) {
	tests := []struct {
		Grid     [][]int
		Expected int
	}{
		{Grid: [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}, Expected: 1},
		{Grid: [][]int{{8}}, Expected: 0},
		{Grid: [][]int{{7, 0, 5}, {2, 4, 6}, {3, 8, 1}}, Expected: 0},
	}

	for _, test := range tests {
		result := numMagicSquaresInside(test.Grid)
		if result != test.Expected {
			t.Errorf("input: %v, expect: %v, got: %v", test.Grid, test.Expected, result)
		}
	}
}
