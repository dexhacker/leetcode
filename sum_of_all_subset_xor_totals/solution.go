package sum_of_all_subset_xor_totals

var answer int

func rec(nums []int, i int, xor int) {
	if i >= len(nums) {
		return
	}
	for j := i; j < len(nums); j++ {
		newXor := xor ^ nums[j]
		answer += newXor
		rec(nums, j+1, newXor)
	}
}

func subsetXORSum(nums []int) int {
	answer = 0
	rec(nums, 0, 0)
	return answer
}
