package number_of_senior_citizens

import "testing"

func TestCountSeniors(t *testing.T) {
	tests := []struct {
		Details  []string
		Expected int
	}{
		{Details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}, Expected: 2},
		{Details: []string{"1313579440F2036", "2921522980M5644"}, Expected: 0},
		{Details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010", "9273438290F6012"}, Expected: 2},
	}

	for _, test := range tests {
		result := countSeniors(test.Details)
		if result != test.Expected {
			t.Errorf("got %d, want %d", result, test.Expected)
		}
	}
}
