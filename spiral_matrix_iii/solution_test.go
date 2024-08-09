package spiral_matrix_iii

import "testing"

func TestSpiralMatrixIII(t *testing.T) {
	tests := []struct {
		Rows     int
		Cols     int
		RStart   int
		CStart   int
		Expected [][]int
	}{
		{Rows: 1, Cols: 2, RStart: 0, CStart: 0, Expected: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
		{
			Rows:   5,
			Cols:   6,
			RStart: 1,
			CStart: 4,
			Expected: [][]int{
				{1, 4},
				{1, 5},
				{2, 5},
				{2, 4},
				{2, 3},
				{1, 3},
				{0, 3},
				{0, 4},
				{0, 5},
				{3, 5},
				{3, 4},
				{3, 3},
				{3, 2},
				{2, 2},
				{1, 2},
				{0, 2},
				{4, 5},
				{4, 4},
				{4, 3},
				{4, 2},
				{4, 1},
				{3, 1},
				{2, 1},
				{1, 1},
				{0, 1},
				{4, 0},
				{3, 0},
				{2, 0},
				{1, 0},
				{0, 0},
			},
		},
	}

	for _, test := range tests {
		result := spiralMatrixIII(test.Rows, test.Cols, test.RStart, test.CStart)
		for i := range result {
			for j := range result[i] {
				if result[i][j] != test.Expected[i][j] {
					t.Errorf(
						"i: %d, j: %d, result[i][j]: %d != test.Expected[i][j]: %d",
						i, j, result[i][j], test.Expected[i][j],
					)
				}
			}
		}
	}
}
