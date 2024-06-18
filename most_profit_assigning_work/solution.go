package most_profit_assigning_work

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	answer := 0
	for _, w := range worker {
		maximum := 0
		for i := 0; i < len(profit); i++ {
			if w >= difficulty[i] && maximum < profit[i] {
				maximum = profit[i]
			}
		}
		answer += maximum
	}
	return answer
}
