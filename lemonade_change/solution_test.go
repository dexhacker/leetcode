package lemonade_change

import "testing"

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		Bills    []int
		Expected bool
	}{
		{Bills: []int{5, 5, 5, 10, 20}, Expected: true},
		{Bills: []int{5, 5, 10, 10, 20}, Expected: false},
	}

	for _, test := range tests {
		result := lemonadeChange(test.Bills)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}
