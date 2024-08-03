package make_two_arrays_equal_by_reversing_subarrays

import "testing"

func TestCanBeEqual(t *testing.T) {
	tests := []struct {
		Target []int
		Arr    []int
		Expect bool
	}{
		{Target: []int{1, 2, 3, 4}, Arr: []int{2, 4, 1, 3}, Expect: true},
		{Target: []int{7}, Arr: []int{7}, Expect: true},
		{Target: []int{3, 7, 9}, Arr: []int{3, 7, 11}, Expect: false},
	}

	for _, test := range tests {
		result := canBeEqual(test.Target, test.Arr)
		if result != test.Expect {
			t.Errorf("%v != %v", result, test.Expect)
		}
	}
}
