package find_valid_matrix_given_row_and_column_sums

import (
	"reflect"
	"testing"
)

func TestRestoreMatrix(t *testing.T) {
	tests := []struct {
		RowSum   []int
		ColSum   []int
		Expected [][]int
	}{
		{RowSum: []int{3, 8}, ColSum: []int{4, 7}, Expected: [][]int{{3, 0}, {1, 7}}},
		{RowSum: []int{5, 7, 10}, ColSum: []int{8, 6, 8}, Expected: [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}}},
	}

	for _, test := range tests {
		result := restoreMatrix(test.RowSum, test.ColSum)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Expected: %v, Actual: %v", test.Expected, result)
		}
	}
}
