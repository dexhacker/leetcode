package magic_squares_in_grid

var g [][]int

func isUniqueAndLowerTen(i int, j int) bool {
	mp := make(map[int]bool)
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if g[i+k][j+l] > 9 {
				return false
			}
			if g[i+k][j+l] == 0 {
				return false
			}
			if _, ok := mp[g[i+k][j+l]]; ok {
				return false
			}
			mp[g[i+k][j+l]] = true
		}
	}
	return true
}

func validSums(i int, j int) bool {
	currentSum := g[i][j] + g[i+1][j] + g[i+2][j]
	if currentSum != g[i][j+1]+g[i+1][j+1]+g[i+2][j+1] {
		return false
	}
	if currentSum != g[i][j+2]+g[i+1][j+2]+g[i+2][j+2] {
		return false
	}
	if currentSum != g[i][j]+g[i][j+1]+g[i][j+2] {
		return false
	}
	if currentSum != g[i+1][j]+g[i+1][j+1]+g[i+1][j+2] {
		return false
	}
	if currentSum != g[i+2][j]+g[i+2][j+1]+g[i+2][j+2] {
		return false
	}
	if currentSum != g[i][j]+g[i+1][j+1]+g[i+2][j+2] {
		return false
	}
	if currentSum != g[i][j+2]+g[i+1][j+1]+g[i+2][j] {
		return false
	}
	return true
}

func isMagicSquares(i int, j int) bool {
	if !isUniqueAndLowerTen(i, j) {
		return false
	}
	if !validSums(i, j) {
		return false
	}
	return true
}

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	g = grid
	answer := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if isMagicSquares(i, j) {
				answer++
			}
		}
	}
	return answer
}
