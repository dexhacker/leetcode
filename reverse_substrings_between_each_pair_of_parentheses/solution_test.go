package reverse_substrings_between_each_pair_of_parentheses

import "testing"

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		S        string
		Expected string
	}{
		{S: "(abcd)", Expected: "dcba"},
		{S: "(u(love)i)", Expected: "iloveu"},
		{S: "(ed(et(oc))el)", Expected: "leetcode"},
	}

	for _, test := range tests {
		result := reverseParentheses(test.S)
		if result != test.Expected {
			t.Errorf("expected: %s, got: %s", test.Expected, result)
		}
	}
}
