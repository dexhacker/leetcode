package maximum_score_words_formed_by_letters

func isValid(availableLetters []int, wordLetters []int) bool {
	for index := range availableLetters {
		if availableLetters[index] < wordLetters[index] {
			return false
		}
	}
	return true
}

func rec(words []string, i int, availableLetters []int, wordsScore []int, wordsLetters [][]int) int {

	maxScore := 0
	for j := i; j < len(words); j++ {
		if wordsScore[j] == 0 {
			continue
		}

		if isValid(availableLetters, wordsLetters[j]) {
			for index := range availableLetters {
				availableLetters[index] -= wordsLetters[j][index]
			}
			result := rec(words, j+1, availableLetters, wordsScore, wordsLetters) + wordsScore[j]
			if result > maxScore {
				maxScore = result
			}
			for index := range availableLetters {
				availableLetters[index] += wordsLetters[j][index]
			}
		}
	}
	return maxScore
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	availableLetters := make([]int, len(score))
	for _, letter := range letters {
		availableLetters[letter-'a']++
	}

	wordsScore := make([]int, len(words))
	var wordsLetters [][]int

	for index, word := range words {
		wordScore := 0
		wordLetters := make([]int, len(score))
		for _, letter := range word {
			wordLetters[letter-'a']++
			if wordLetters[letter-'a'] > availableLetters[letter-'a'] {
				wordScore = 0
				break
			}
			wordScore += score[letter-'a']
		}
		wordsScore[index] = wordScore
		wordsLetters = append(wordsLetters, wordLetters)
	}

	return rec(words, 0, availableLetters, wordsScore, wordsLetters)
}
