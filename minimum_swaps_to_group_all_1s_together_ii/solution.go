package minimum_swaps_to_group_all_1s_together_ii

func minimum(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minSwaps(nums []int) int {
	slidingWindowSize := 0
	for _, num := range nums {
		if num == 1 {
			slidingWindowSize++
		}
	}

	onesCount := 0
	for i := 0; i < slidingWindowSize; i++ {
		if nums[i] == 1 {
			onesCount++
		}
	}

	answer := slidingWindowSize - onesCount
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == 1 {
			onesCount--
		}
		if nums[(i+slidingWindowSize-1)%len(nums)] == 1 {
			onesCount++
		}
		answer = minimum(answer, slidingWindowSize-onesCount)
	}
	return answer
}
