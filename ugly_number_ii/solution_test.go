package ugly_number_ii

import "testing"

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		N        int
		Expected int
	}{
		{N: 10, Expected: 12},
		{N: 1, Expected: 1},
		{N: 11, Expected: 15},
	}

	for _, test := range tests {
		result := nthUglyNumber(test.N)
		if result != test.Expected {
			t.Errorf("result %d expected %d", result, test.Expected)
		}
	}
}
