package relative_sort_array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	countNumbers := make(map[int]int)
	for _, item := range arr2 {
		countNumbers[item] = 0
	}

	notAppear := make([]int, 0)
	for _, item := range arr1 {
		if _, ok := countNumbers[item]; ok {
			countNumbers[item]++
		} else {
			notAppear = append(notAppear, item)
		}
	}

	sort.Ints(notAppear)

	answer := make([]int, 0)

	for _, item := range arr2 {
		for i := 0; i < countNumbers[item]; i++ {
			answer = append(answer, item)
		}
	}

	for _, item := range notAppear {
		answer = append(answer, item)
	}

	return answer
}
