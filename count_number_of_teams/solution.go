package count_number_of_teams

func numTeams(rating []int) int {
	upper := make([]int, len(rating))
	lower := make([]int, len(rating))
	for i := range rating {
		upperThenJ := 0
		lowerThenJ := 0
		for j := i + 1; j < len(rating); j++ {
			if rating[i] > rating[j] {
				upperThenJ++
			}
			if rating[i] < rating[j] {
				lowerThenJ++
			}
		}
		upper[i] = upperThenJ
		lower[i] = lowerThenJ
	}

	answer := 0
	for i := 0; i < len(rating); i++ {
		for j := i + 1; j < len(rating); j++ {
			if rating[i] < rating[j] {
				answer += lower[j]
			}
			if rating[i] > rating[j] {
				answer += upper[j]
			}
		}
	}
	return answer
}
