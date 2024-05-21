package subsets

import "slices"

var answer [][]int

func removeDuplicates(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func rec(nums []int, i int, array []int) {
	for j := i; j < len(nums); j++ {
		newArray := slices.Clone(array)
		newArray = append(newArray, nums[j])
		answer = append(answer, newArray)
		rec(nums, j+1, newArray)
	}
}

func subsets(nums []int) [][]int {
	nums = removeDuplicates(nums)
	answer = make([][]int, 0)
	var array []int
	answer = append(answer, array)
	rec(nums, 0, array)
	return answer
}
