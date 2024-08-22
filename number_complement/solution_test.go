package number_complement

import "testing"

func TestFindComplement(t *testing.T) {
	tests := []struct {
		Num      int
		Expected int
	}{
		{Num: 5, Expected: 2},
		{Num: 1, Expected: 0},
	}

	for _, test := range tests {
		result := findComplement(test.Num)
		if result != test.Expected {
			t.Errorf("FindComplement(%d) = %d; expected %d", test.Num, result, test.Expected)
		}
	}
}
