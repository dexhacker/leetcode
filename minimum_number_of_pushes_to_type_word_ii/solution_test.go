package minimum_number_of_pushes_to_type_word_ii

import "testing"

func TestMinimumPushes(t *testing.T) {
	tests := []struct {
		Word     string
		Expected int
	}{
		{Word: "abcde", Expected: 5},
		{Word: "xyzxyzxyzxyz", Expected: 12},
		{Word: "aabbccddeeffgghhiiiiii", Expected: 24},
	}

	for _, test := range tests {
		result := minimumPushes(test.Word)
		if result != test.Expected {
			t.Errorf("expected: %d, got: %d", test.Expected, result)
		}
	}
}
