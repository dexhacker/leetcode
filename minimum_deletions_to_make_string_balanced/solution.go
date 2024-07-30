package minimum_deletions_to_make_string_balanced

func minimumDeletions(s string) int {
	aLetters := make([]int, len(s))
	bLetters := make([]int, len(s))

	countB := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			countB++
		}
		aLetters[i] = countB
	}

	countA := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'a' {
			countA++
		}
		bLetters[i] = countA
	}

	minAnswer := bLetters[0]
	for i := 0; i < len(s)-1; i++ {
		if minAnswer > aLetters[i]+bLetters[i+1] {
			minAnswer = aLetters[i] + bLetters[i+1]
		}
	}
	if aLetters[len(s)-1] < minAnswer {
		minAnswer = aLetters[len(s)-1]
	}
	return minAnswer
}
