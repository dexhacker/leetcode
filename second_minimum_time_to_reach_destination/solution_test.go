package second_minimum_time_to_reach_destination

import "testing"

func TestSecondMinimum(t *testing.T) {
	tests := []struct {
		N        int
		Edges    [][]int
		Time     int
		Change   int
		Expected int
	}{
		{
			N:        5,
			Edges:    [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}},
			Time:     3,
			Change:   5,
			Expected: 13,
		},
		{
			N:        2,
			Edges:    [][]int{{1, 2}},
			Time:     3,
			Change:   2,
			Expected: 11,
		},
	}

	for _, test := range tests {
		result := secondMinimum(test.N, test.Edges, test.Time, test.Change)
		if result != test.Expected {
			t.Errorf("test: %v, expect %d, but got %d", test, test.Expected, result)
		}
	}
}
