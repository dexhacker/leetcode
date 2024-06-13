package minimum_number_of_moves_to_seat_everyone

import "sort"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}

func minMovesToSeat(seats []int, students []int) int {
	answer := 0

	sort.Ints(seats)
	sort.Ints(students)

	for i := 0; i < len(seats); i++ {
		answer += abs(seats[i] - students[i])
	}

	return answer
}
