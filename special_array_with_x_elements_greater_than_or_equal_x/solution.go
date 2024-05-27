package special_array_with_x_elements_greater_than_or_equal_x

import "slices"

func specialArray(nums []int) int {
	slices.Sort(nums)
	for i := range nums {
		if len(nums)-i <= nums[i] {
			if i > 0 && nums[i-1] >= len(nums)-i {
				return -1
			}

			return len(nums) - i
		}
	}

	return -1
}
