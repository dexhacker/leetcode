package integer_to_english_words

import "testing"

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		Num      int
		Expected string
	}{
		{Num: 123, Expected: "One Hundred Twenty Three"},
		{Num: 12345, Expected: "Twelve Thousand Three Hundred Forty Five"},
		{Num: 1234567, Expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	}

	for _, test := range tests {
		result := numberToWords(test.Num)
		if result != test.Expected {
			t.Errorf("numberToWords(%d) = %s, want %s", test.Num, result, test.Expected)
		}
	}
}
