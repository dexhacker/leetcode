package path_with_maximum_gold

func findPath(grid [][]int, i int, j int, sum int) int {
	grid[i][j] *= -1
	answer := sum

	if i-1 >= 0 && grid[i-1][j] > 0 {
		result := findPath(grid, i-1, j, sum+grid[i-1][j])
		if result > answer {
			answer = result
		}
	}

	if i+1 < len(grid) && grid[i+1][j] > 0 {
		result := findPath(grid, i+1, j, sum+grid[i+1][j])
		if result > answer {
			answer = result
		}
	}

	if j-1 >= 0 && grid[i][j-1] > 0 {
		result := findPath(grid, i, j-1, sum+grid[i][j-1])
		if result > answer {
			answer = result
		}
	}

	if j+1 < len(grid[0]) && grid[i][j+1] > 0 {
		result := findPath(grid, i, j+1, sum+grid[i][j+1])
		if result > answer {
			answer = result
		}
	}

	grid[i][j] *= -1
	return answer
}

func getMaximumGold(grid [][]int) int {
	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				result := findPath(grid, i, j, grid[i][j])
				if result > answer {
					answer = result
				}
			}
		}
	}

	return answer
}
