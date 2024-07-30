package minimum_deletions_to_make_string_balanced

import "testing"

func TestMinimumDeletions(t *testing.T) {
	tests := []struct {
		S        string
		Expected int
	}{
		{S: "aababbab", Expected: 2},
		{S: "bbaaaaabb", Expected: 2},
	}

	for _, test := range tests {
		result := minimumDeletions(test.S)
		if result != test.Expected {
			t.Errorf("minimum deletions failed, expected %d, got %d", test.Expected, result)
		}
	}
}
