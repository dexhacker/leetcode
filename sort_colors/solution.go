package sort_colors

func sortColors(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
	}

	for j := i; j < len(nums); j++ {
		if nums[j] == 1 {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
	}
}
