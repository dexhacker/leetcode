package count_number_of_teams

import "testing"

func TestNumTeams(t *testing.T) {
	tests := []struct {
		Rating []int
		Expect int
	}{
		{Rating: []int{2, 5, 3, 4, 1}, Expect: 3},
		{Rating: []int{2, 1, 3}, Expect: 0},
		{Rating: []int{1, 2, 3, 4}, Expect: 4},
	}

	for _, test := range tests {
		result := numTeams(test.Rating)
		if result != test.Expect {
			t.Errorf("numTeams(%v) = %v; expect %v", test.Rating, result, test.Expect)
		}
	}
}
