package spiral_matrix_iii

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	used := make([][]bool, rows)
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, cols)
	}

	answer := make([][]int, 0)
	used[rStart][cStart] = true
	answer = append(answer, []int{rStart, cStart})
	direction := "right"
	length := 1
	lastUserDirection := "up"
	repeatCount := 0
	for {
		if lastUserDirection == direction {
			if repeatCount == 2 {
				break
			}
			repeatCount++
		}
		for i := 0; i < length; i++ {
			if direction == "right" {
				cStart++
			}
			if direction == "down" {
				rStart++
			}
			if direction == "left" {
				cStart--
			}
			if direction == "up" {
				rStart--
			}
			if rStart > -1 && cStart > -1 && rStart < rows && cStart < cols && !used[rStart][cStart] {
				used[rStart][cStart] = true
				lastUserDirection = direction
				repeatCount = 0
				answer = append(answer, []int{rStart, cStart})
			}
		}
		if direction == "right" {
			direction = "down"
		} else if direction == "down" {
			direction = "left"
			length++
		} else if direction == "left" {
			direction = "up"
		} else if direction == "up" {
			direction = "right"
			length++
		}
	}
	return answer
}
