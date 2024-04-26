package minimum_falling_path_sum_ii

func minFallingPathSum(grid [][]int) int {
	if len(grid) == 1 {
		minimum := grid[0][0]
		for _, element := range grid[0] {
			if element < minimum {
				minimum = element
			}
		}
		return minimum
	}

	for i := 1; i < len(grid); i++ {
		min1 := grid[i-1][0]
		min2 := grid[i-1][1]

		if min1 > min2 {
			tmp := min2
			min2 = min1
			min1 = tmp
		}

		for j := 2; j < len(grid[i]); j++ {
			if grid[i-1][j] < min1 {
				min2 = min1
				min1 = grid[i-1][j]
			} else if grid[i-1][j] < min2 {
				min2 = grid[i-1][j]
			}
		}

		for j := 0; j < len(grid[i]); j++ {
			if grid[i-1][j] == min1 {
				grid[i][j] += min2
			} else {
				grid[i][j] += min1
			}
		}
	}

	lastIndex := len(grid) - 1
	minimum := grid[lastIndex][0]
	for _, element := range grid[lastIndex] {
		if minimum > element {
			minimum = element
		}
	}

	return minimum
}
