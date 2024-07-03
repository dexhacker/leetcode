package minimum_difference_between_largest_and_smallest_value_in_three_moves

import "sort"

func minDifference(nums []int) int {
	if len(nums) < 4 {
		return 0
	}

	last := len(nums) - 1
	sort.Ints(nums)
	answer := nums[last] - nums[3]
	answer = min(answer, nums[last-1]-nums[2])
	answer = min(answer, nums[last-2]-nums[1])
	answer = min(answer, nums[last-3]-nums[0])
	return answer
}
