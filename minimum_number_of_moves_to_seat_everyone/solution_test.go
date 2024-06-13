package minimum_number_of_moves_to_seat_everyone

import "testing"

func TestMinMovesToSeat(t *testing.T) {
	tests := []struct {
		Seats    []int
		Students []int
		Expected int
	}{
		{Students: []int{3, 1, 5}, Seats: []int{2, 7, 4}, Expected: 4},
		{Students: []int{4, 1, 5, 9}, Seats: []int{1, 3, 2, 6}, Expected: 7},
		{Students: []int{2, 2, 6, 6}, Seats: []int{1, 3, 2, 6}, Expected: 4},
	}

	for _, test := range tests {
		result := minMovesToSeat(test.Seats, test.Students)
		if result != test.Expected {
			t.Errorf("minimum number of moves to seat should be %d, got %d", test.Expected, result)
		}
	}
}
