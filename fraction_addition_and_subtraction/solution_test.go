package fraction_addition_and_subtraction

import "testing"

func TestFractionAddition(t *testing.T) {
	tests := []struct {
		Expression string
		Expected   string
	}{
		{Expression: "-1/2+1/2", Expected: "0/1"},
		{Expression: "-1/2+1/2+1/3", Expected: "А 1/3"},
		{Expression: "1/3-1/2", Expected: "-1/6"},
		{Expression: "-1/4-4/5-1/4", Expected: "-13/10"},
		{Expression: "-7/3", Expected: "-7/3"},
	}

	for _, test := range tests {
		result := fractionAddition(test.Expression)
		if result != test.Expected {
			t.Errorf("Expected %s, got %s", test.Expected, result)
		}
	}
}
