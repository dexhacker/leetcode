package maximum_score_words_formed_by_letters

import "testing"

func TestMaxScoreWords(t *testing.T) {
	tests := []struct {
		Words    []string
		Letters  []byte
		Score    []int
		Expected int
	}{
		{
			Words:    []string{"dog", "cat", "dad", "good"},
			Letters:  []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
			Score:    []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Expected: 23,
		},
		{
			Words:    []string{"xxxz", "ax", "bx", "cx"},
			Letters:  []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
			Score:    []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10},
			Expected: 27,
		},
		{
			Words:    []string{"leetcode"},
			Letters:  []byte{'l', 'e', 't', 'c', 'o', 'd'},
			Score:    []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			Expected: 0,
		},
	}

	for _, test := range tests {
		result := maxScoreWords(test.Words, test.Letters, test.Score)
		if result != test.Expected {
			t.Errorf("expected: %v, got: %v", test.Expected, result)
		}
	}
}
