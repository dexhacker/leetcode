package three_consecutive_odds

import (
	"fmt"
	"testing"
)

func TestThreeConsecutiveOdds(t *testing.T) {
	tests := []struct {
		Arr      []int
		Expected bool
	}{
		{Arr: []int{2, 6, 4, 1}, Expected: false},
		{Arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}, Expected: true},
	}

	for _, test := range tests {
		result := threeConsecutiveOdds(test.Arr)
		if result != test.Expected {
			fmt.Printf("FAIL - Expected: %v, Actual: %v\n", test.Expected, result)
		}
	}
}
