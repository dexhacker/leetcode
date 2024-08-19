package two_keys_keyboard

import "testing"

func TestMinSteps(t *testing.T) {
	tests := []struct {
		N        int
		Expected int
	}{
		{N: 3, Expected: 3},
		{N: 1, Expected: 0},
		{N: 999, Expected: 46},
		{N: 37, Expected: 37},
		{N: 111, Expected: 40},
	}

	for _, test := range tests {
		result := minSteps(test.N)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d", test.Expected, result)
		}
	}
}
