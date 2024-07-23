package sort_array_by_increasing_frequency

import "sort"

func frequencySort(nums []int) []int {
	counting := make(map[int]int)
	for _, num := range nums {
		if val, ok := counting[num]; ok {
			counting[num] = val + 1
		} else {
			counting[num] = 1
		}
	}

	frequencyMp := make(map[int][]int)
	for num, frequency := range counting {
		if val, ok := frequencyMp[frequency]; ok {
			frequencyMp[frequency] = append(val, num)
		} else {
			frequencyMp[frequency] = []int{num}
		}
	}

	frequencies := make([]int, 0)
	for frequency := range frequencyMp {
		frequencies = append(frequencies, frequency)
	}

	sort.Ints(frequencies)

	answer := make([]int, 0)
	for _, frequency := range frequencies {
		numbers := frequencyMp[frequency]
		sort.Ints(numbers)
		for i := len(numbers) - 1; i >= 0; i-- {
			for j := 0; j < frequency; j++ {
				answer = append(answer, numbers[i])
			}
		}
	}

	return answer
}
