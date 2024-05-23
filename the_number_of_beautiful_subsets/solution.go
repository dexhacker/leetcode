package the_number_of_beautiful_subsets

var answer int

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func rec(nums []int, k int, array []int, i int) {
	for j := i; j < len(nums); j++ {
		isValid := true
		newArray := make([]int, len(array))
		copy(newArray, array)
		for _, item := range array {
			diff := item - nums[j]
			if k == abs(diff) {
				isValid = false
				break
			}
		}

		if isValid {
			answer++
			newArray = append(newArray, nums[j])
			rec(nums, k, newArray, j+1)
		}
	}
}

func beautifulSubsets(nums []int, k int) int {
	answer = 0
	var array []int
	rec(nums, k, array, 0)
	return answer
}
