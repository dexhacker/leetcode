package word_break_ii

import "strings"

var answer []string

func rec(s string, wordDict []string, i int, result []string) {
	if len(s) == i {
		strResult := strings.Join(result, " ")
		answer = append(answer, strResult)
		return
	}

	variants := make([]string, len(wordDict))
	copy(variants, wordDict)
	for j := i; j < len(s); j++ {
		newVariants := make([]string, 0)
		for _, variant := range variants {
			if variant[j-i] == s[j] {
				if len(variant)-1 == j-i {
					newResult := make([]string, len(result))
					copy(newResult, result)
					newResult = append(newResult, variant)
					rec(s, wordDict, j+1, newResult)
				} else {
					newVariants = append(newVariants, variant)
				}
			}
		}
		variants = newVariants
	}
}

func wordBreak(s string, wordDict []string) []string {
	answer = make([]string, 0)
	result := make([]string, 0)
	rec(s, wordDict, 0, result)
	return answer
}
