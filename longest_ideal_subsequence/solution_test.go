package longest_ideal_subsequence

import "testing"

func TestLongestIdealString(t *testing.T) {
	tests := []struct {
		S        string
		K        int
		Expected int
	}{
		{"acfgbd", 2, 4},
		{"abcd", 3, 4},
		{"azaza", 25, 5},
		{"xzacd", 5, 3},
	}

	for _, test := range tests {
		result := longestIdealString(test.S, test.K)
		if result != test.Expected {
			t.Errorf("longestIdealString(%s, %v) = %v; expected %v", test.S, test.K, result, test.Expected)
		}
	}
}
