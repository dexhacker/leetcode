package minimum_increment_to_make_array_unique

import "sort"

func minIncrementForUnique(nums []int) int {
	answer := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		answer += nums[i-1] - nums[i] + 1
		nums[i] = nums[i-1] + 1
	}
	return answer
}
