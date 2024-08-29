package most_stones_removed_with_same_row_or_column

import "testing"

func TestRemoveStones(t *testing.T) {
	tests := []struct {
		Stones   [][]int
		Expected int
	}{
		{Stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}, Expected: 5},
		{Stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}, Expected: 3},
		{Stones: [][]int{{0, 0}}, Expected: 0},
	}

	for _, test := range tests {
		result := removeStones(test.Stones)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d", test.Expected, result)
		}
	}
}
