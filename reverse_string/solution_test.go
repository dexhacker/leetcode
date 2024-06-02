package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		S        []byte
		Expected []byte
	}{
		{S: []byte("hello"), Expected: []byte("olleh")},
		{S: []byte("Hannah"), Expected: []byte("hannaH")},
	}

	for _, test := range tests {
		reverseString(test.S)
		for i := range test.S {
			if test.S[i] != test.Expected[i] {
				t.Error("Expected", test.Expected, "got", test.S[i])
			}
		}
	}
}
