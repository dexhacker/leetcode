package lucky_numbers_in_a_matrix

import (
	"reflect"
	"testing"
)

func TestLuckyNumbers(t *testing.T) {
	tests := []struct {
		Matrix [][]int
		Expect []int
	}{
		{Matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, Expect: []int{15}},
		{Matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, Expect: []int{12}},
		{Matrix: [][]int{{7, 8}, {1, 2}}, Expect: []int{7}},
		{Matrix: [][]int{{3, 7, 8}, {9, 21, 13}, {20, 16, 17}}, Expect: []int{}},
	}

	for _, test := range tests {
		result := luckyNumbers(test.Matrix)
		if !reflect.DeepEqual(result, test.Expect) {
			t.Errorf("luckyNumbers(%v) = %v, want %v", test.Matrix, result, test.Expect)
		}
	}
}
