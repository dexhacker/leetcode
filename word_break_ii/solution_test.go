package word_break_ii

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		S        string
		WordDict []string
		Expected []string
	}{
		{
			S:        "catsanddog",
			WordDict: []string{"cat", "cats", "and", "sand", "dog"},
			Expected: []string{"cat sand dog", "cats and dog"},
		},
		{
			S:        "pineapplepenapple",
			WordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			Expected: []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
		{
			S:        "catsandog",
			WordDict: []string{"cats", "dog", "sand", "and", "cat"},
			Expected: []string{},
		},
	}

	for _, test := range tests {
		result := wordBreak(test.S, test.WordDict)
		for i := range result {
			if result[i] != test.Expected[i] {
				t.Errorf("expected: %v, got: %v", test.Expected, result)
			}
		}
	}
}
