package height_checker

import "sort"

func heightChecker(heights []int) int {
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)
	sort.Ints(sortedHeights)
	answer := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			answer++
		}
	}

	return answer
}
