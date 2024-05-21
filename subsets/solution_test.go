package subsets

import "testing"

func TestSubsets(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected [][]int
	}{
		{Nums: []int{1, 2, 3}, Expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		{Nums: []int{0}, Expected: [][]int{{}, {0}}},
	}

	for _, test := range tests {
		result := subsets(test.Nums)
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(result[i]); j++ {
				if result[i][j] != test.Expected[i][j] {
					t.Errorf(
						"subsets failed for %v and %v: expected %v, got %v",
						i, j, test.Expected[i][j], result[i][j],
					)
				}
			}
		}
	}
}
