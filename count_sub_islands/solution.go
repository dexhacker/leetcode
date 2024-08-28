package count_sub_islands

var g1 [][]int
var g2 [][]int
var used [][]bool

func island(i int, j int, makeUp bool) bool {
	answer := makeUp
	if i-1 >= 0 && g2[i-1][j] == 1 && !used[i-1][j] {
		used[i-1][j] = true
		answer = island(i-1, j, makeUp && g1[i-1][j] == g2[i-1][j]) && answer
	}
	if i+1 < len(g2) && g2[i+1][j] == 1 && !used[i+1][j] {
		used[i+1][j] = true
		answer = island(i+1, j, makeUp && g1[i+1][j] == g2[i+1][j]) && answer
	}
	if j-1 >= 0 && g2[i][j-1] == 1 && !used[i][j-1] {
		used[i][j-1] = true
		answer = island(i, j-1, makeUp && g1[i][j-1] == g2[i][j-1]) && answer
	}
	if j+1 < len(g2[i]) && g2[i][j+1] == 1 && !used[i][j+1] {
		used[i][j+1] = true
		answer = island(i, j+1, makeUp && g1[i][j+1] == g2[i][j+1]) && answer
	}
	return answer
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	g1 = grid1
	g2 = grid2
	used = make([][]bool, len(grid2))
	for i := range used {
		used[i] = make([]bool, len(grid2[i]))
	}
	answer := 0

	for i := 0; i < len(g2); i++ {
		for j := 0; j < len(g2[i]); j++ {
			if used[i][j] {
				continue
			}

			if g2[i][j] == 1 {
				used[i][j] = true
				if island(i, j, g1[i][j] == g2[i][j]) {
					answer++
				}
			}
		}
	}

	return answer
}
