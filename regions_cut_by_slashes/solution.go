package regions_cut_by_slashes

var used [][]*Cell
var runeGrid [][]rune

type Cell struct {
	Left  bool
	Right bool
}

func dfs(i int, j int, r rune) {
	if i < 0 || i >= len(used) || j < 0 || j >= len(used) {
		return
	}

	cell := used[i][j]
	if runeGrid[i][j] == ' ' {
		if cell.Left && cell.Right {
			return
		}

		cell.Left = true
		cell.Right = true
		dfs(i-1, j, 'u')
		dfs(i+1, j, 'd')
		dfs(i, j-1, 'l')
		dfs(i, j+1, 'r')
	} else if runeGrid[i][j] == '/' {
		if r == 'd' || r == 'r' {
			if cell.Left {
				return
			}

			cell.Left = true
			dfs(i-1, j, 'u')
			dfs(i, j-1, 'l')
		} else {
			if cell.Right {
				return
			}

			cell.Right = true
			dfs(i, j+1, 'r')
			dfs(i+1, j, 'd')
		}
	} else {
		if r == 'd' || r == 'l' {
			if cell.Right {
				return
			}

			cell.Right = true
			dfs(i-1, j, 'u')
			dfs(i, j+1, 'r')
		} else {
			if cell.Left {
				return
			}

			cell.Left = true
			dfs(i+1, j, 'd')
			dfs(i, j-1, 'l')
		}
	}
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	used = make([][]*Cell, n)
	for i := range used {
		used[i] = make([]*Cell, n)
		for j := range used[i] {
			used[i][j] = &Cell{Left: false, Right: false}
		}
	}

	runeGrid = make([][]rune, n)
	for i := range runeGrid {
		runeGrid[i] = []rune(grid[i])
	}

	answer := 0
	for i := range runeGrid {
		for j := range runeGrid[i] {
			cell := used[i][j]

			if !cell.Left {
				answer++
				cell.Left = true
				dfs(i, j-1, 'l')
				if runeGrid[i][j] == '/' {
					dfs(i-1, j, 'u')
				} else {
					dfs(i+1, j, 'd')
				}
			}

			if !cell.Right {
				answer++
				cell.Right = true
				dfs(i, j+1, 'r')
				if runeGrid[i][j] == '/' {
					dfs(i+1, j, 'd')
				} else {
					dfs(i-1, j, 'u')
				}
			}
		}
	}
	return answer
}
