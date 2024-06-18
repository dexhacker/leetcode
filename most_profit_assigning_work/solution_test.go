package most_profit_assigning_work

import (
	"fmt"
	"testing"
)

func TestMaxProfitAssignment(t *testing.T) {
	tests := []struct {
		Difficulty []int
		Profit     []int
		Worker     []int
		Expected   int
	}{
		{
			Difficulty: []int{2, 4, 6, 8, 10},
			Profit:     []int{10, 20, 30, 40, 50},
			Worker:     []int{4, 5, 6, 7},
			Expected:   100,
		},
		{
			Difficulty: []int{85, 47, 57},
			Profit:     []int{24, 66, 99},
			Worker:     []int{40, 25, 25},
			Expected:   0,
		},
	}

	for _, test := range tests {
		result := maxProfitAssignment(test.Difficulty, test.Profit, test.Worker)
		if result != test.Expected {
			fmt.Printf("Expected %d, got %d\n", test.Expected, result)
		}
	}
}
