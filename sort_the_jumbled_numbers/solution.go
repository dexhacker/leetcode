package sort_the_jumbled_numbers

import (
	"sort"
	"strconv"
)

var m []int

func mapNumber(num int) int {
	str := strconv.Itoa(num)
	runes := []rune(str)
	mappedNumber := ""
	for _, runeItem := range runes {
		numberItem := runeItem - '0'
		newNumberItem := m[numberItem]
		mappedNumber += strconv.Itoa(newNumberItem)
	}
	newNumber, err := strconv.Atoi(mappedNumber)
	if err != nil {
		return 0
	}
	return newNumber
}

func sortJumbled(mapping []int, nums []int) []int {
	m = mapping
	mp := make(map[int][]int)

	mappedNums := make([]int, 0)
	for _, num := range nums {
		mappedNum := mapNumber(num)
		if val, ok := mp[mappedNum]; ok {
			mp[mappedNum] = append(val, num)
		} else {
			mp[mappedNum] = []int{num}
			mappedNums = append(mappedNums, mappedNum)
		}
	}

	sort.Ints(mappedNums)

	answer := make([]int, 0)
	for _, mappedNum := range mappedNums {
		for _, num := range mp[mappedNum] {
			answer = append(answer, num)
		}
	}

	return answer
}
