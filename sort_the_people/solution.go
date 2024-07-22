package sort_the_people

import "sort"

func sortPeople(names []string, heights []int) []string {
	mp := make(map[int]string)
	answer := make([]string, 0)
	for i := range names {
		mp[heights[i]] = names[i]
	}

	sort.Ints(heights)
	for i := len(heights) - 1; i >= 0; i-- {
		answer = append(answer, mp[heights[i]])
	}

	return answer
}
