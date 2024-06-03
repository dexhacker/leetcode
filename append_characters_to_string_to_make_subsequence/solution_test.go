package append_characters_to_string_to_make_subsequence

import "testing"

func TestAppendCharacters(t *testing.T) {
	tests := []struct {
		S        string
		T        string
		Expected int
	}{
		{S: "coaching", T: "coding", Expected: 4},
		{S: "abcde", T: "a", Expected: 0},
		{S: "z", T: "abcde", Expected: 5},
	}

	for _, test := range tests {
		result := appendCharacters(test.S, test.T)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d", test.Expected, result)
		}
	}
}
