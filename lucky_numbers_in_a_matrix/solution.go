package lucky_numbers_in_a_matrix

func luckyNumbers(matrix [][]int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		minJ := 0
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] < matrix[i][minJ] {
				minJ = j
			}
		}
		if mp[minJ] < matrix[i][minJ] {
			mp[minJ] = matrix[i][minJ]
		}
	}

	for j, minValue := range mp {
		isLuckyNumber := true
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] > minValue {
				isLuckyNumber = false
				break
			}
		}

		if isLuckyNumber {
			return []int{minValue}
		}
	}

	return []int{}
}
