package longest_ideal_subsequence

func longestIdealString(s string, k int) int {
	letters := make([]int, 26)
	for _, c := range s {
		intLetter := int(c - 'a')
		i := intLetter - k
		if i < 0 {
			i = 0
		}

		finish := intLetter + k
		if finish > 25 {
			finish = 25
		}

		maximumLetterCount := 0
		for ; i <= finish; i++ {
			if letters[i] > maximumLetterCount {
				maximumLetterCount = letters[i]
			}
		}
		letters[intLetter] = maximumLetterCount + 1
	}

	maximum := 0
	for _, i := range letters {
		if maximum < i {
			maximum = i
		}
	}
	return maximum
}
