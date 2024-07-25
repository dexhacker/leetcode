package sort_an_array

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivotIndex := len(nums) / 2

	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	sortArray(nums[:left])
	sortArray(nums[left+1:])

	return nums
}
