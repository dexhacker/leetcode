package count_sub_islands

import "testing"

func TestCountSubIslands(t *testing.T) {
	tests := []struct {
		Grid1    [][]int
		Grid2    [][]int
		Expected int
	}{
		{
			Grid1:    [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
			Grid2:    [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}},
			Expected: 3,
		},
		{
			Grid1:    [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
			Grid2:    [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}},
			Expected: 2,
		},
	}

	for _, test := range tests {
		result := countSubIslands(test.Grid1, test.Grid2)
		if result != test.Expected {
			t.Errorf("expect %d, got %d", test.Expected, result)
		}
	}
}
