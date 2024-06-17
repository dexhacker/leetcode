package sum_of_square_numbers

import (
	"fmt"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	tests := []struct {
		C        int
		Expected bool
	}{
		{C: 5, Expected: true},
		{C: 3, Expected: false},
		{C: 14, Expected: false},
		{C: 2, Expected: true},
		{C: 999999999, Expected: false},
		{C: 4, Expected: true},
		{C: 0, Expected: true},
	}

	for _, test := range tests {
		result := judgeSquareSum(test.C)
		if result != test.Expected {
			fmt.Printf("Result: %v, Expected: %v\n", result, test.Expected)
		}
	}
}
