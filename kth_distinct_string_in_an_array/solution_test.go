package kth_distinct_string_in_an_array

import "testing"

func TestKthDistinct(t *testing.T) {
	tests := []struct {
		Arr      []string
		K        int
		Expected string
	}{
		{Arr: []string{"d", "b", "c", "b", "c", "a"}, K: 2, Expected: "a"},
		{Arr: []string{"aaa", "aa", "a"}, K: 1, Expected: "aaa"},
		{Arr: []string{"a", "b", "a"}, K: 3, Expected: ""},
	}

	for _, test := range tests {
		result := kthDistinct(test.Arr, test.K)
		if result != test.Expected {
			t.Errorf("test: %v, expected: %v, got: %v", test.Arr, test.Expected, result)
		}
	}
}
