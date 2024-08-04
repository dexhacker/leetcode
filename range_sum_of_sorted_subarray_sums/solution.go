package range_sum_of_sorted_subarray_sums

import "sort"

func rangeSum(nums []int, _ int, left int, right int) int {
	subArray := make([]int, 0)
	subArray = append(subArray, 0)
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			subArray = append(subArray, sum)
		}
	}
	sort.Ints(subArray)
	answer := 0
	for i := left; i <= right; i++ {
		answer += subArray[i]
		answer %= 1000000007
	}
	return answer
}
