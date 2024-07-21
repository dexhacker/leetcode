package build_a_matrix_with_conditions

import (
	"reflect"
	"testing"
)

func TestBuildMatrix(t *testing.T) {
	tests := []struct {
		K             int
		RowConditions [][]int
		ColConditions [][]int
		Expected      [][]int
	}{
		{
			K:             3,
			RowConditions: [][]int{{1, 2}, {3, 2}},
			ColConditions: [][]int{{2, 1}, {3, 2}},
			Expected:      [][]int{{0, 0, 1}, {3, 0, 0}, {0, 2, 0}},
		},
		{
			K:             3,
			RowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 1}},
			ColConditions: [][]int{{2, 1}},
			Expected:      [][]int{},
		},
	}

	for _, test := range tests {
		result := buildMatrix(test.K, test.RowConditions, test.ColConditions)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}
