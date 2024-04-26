package minimum_falling_path_sum_ii

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		Grid     [][]int
		Expected int
	}{
		{Grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, Expected: 13},
		{Grid: [][]int{{7}}, Expected: 7},
		{
			Grid: [][]int{
				{-37, 51, -36, 34, -22},
				{82, 4, 30, 14, 38},
				{-68, -52, -92, 65, -85},
				{-49, -3, -77, 8, -19},
				{-60, -71, -21, -62, -73},
			},
			Expected: -268,
		},
	}

	for _, test := range tests {
		result := minFallingPathSum(test.Grid)
		if result != test.Expected {
			t.Errorf("minFallingPathSum(%+v) = %v; expected %v", test.Grid, result, test.Expected)
		}
	}
}
