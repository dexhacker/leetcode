package score_after_flipping_matrix

import (
	"fmt"
	"strconv"
)

func matrixScore(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j] = (grid[i][j] + 1) % 2
			}
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		even := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j]%2 == 0 {
				even++
			}
		}

		if len(grid)-even < even {
			for i := 0; i < len(grid); i++ {
				grid[i][j] = (grid[i][j] + 1) % 2
			}
		}
	}

	answer := 0
	for i := 0; i < len(grid); i++ {
		sum := ""
		for j := 0; j < len(grid[i]); j++ {
			sum += fmt.Sprintf("%v", grid[i][j])
		}
		number, err := strconv.ParseInt(sum, 2, 64)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return 0
		}

		answer += int(number)
	}

	return answer
}
