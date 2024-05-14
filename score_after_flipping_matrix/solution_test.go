package score_after_flipping_matrix

import "testing"

func TestMatrixScore(t *testing.T) {
	tests := []struct {
		Grid     [][]int
		Expected int
	}{
		{Grid: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, Expected: 39},
		{Grid: [][]int{{0}}, Expected: 1},
	}

	for _, test := range tests {
		result := matrixScore(test.Grid)
		if result != test.Expected {
			t.Errorf("matrixScore(%+v) = %d; expected %d", test.Grid, result, test.Expected)
		}
	}
}
