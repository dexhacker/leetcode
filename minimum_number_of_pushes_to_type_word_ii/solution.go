package minimum_number_of_pushes_to_type_word_ii

import "sort"

func minimumPushes(word string) int {
	type LetterStruct struct {
		Letter rune
		Count  int
	}

	letterCount := make([]int, 26)
	runes := []rune(word)
	for _, letter := range runes {
		letterCount[letter-'a']++
	}

	letters := make([]*LetterStruct, 0)
	for letter, count := range letterCount {
		letters = append(letters, &LetterStruct{Letter: rune(letter), Count: count})
	}

	sort.Slice(letters, func(i, j int) bool {
		return letters[i].Count > letters[j].Count
	})

	answer := 0
	multiply := 0
	for i, letter := range letters {
		if i%8 == 0 {
			multiply++
		}
		answer += multiply * letter.Count
	}
	return answer
}
