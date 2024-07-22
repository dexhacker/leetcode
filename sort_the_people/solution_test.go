package sort_the_people

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	tests := []struct {
		Names    []string
		Heights  []int
		Expected []string
	}{
		{
			Names:    []string{"Mary", "John", "Emma"},
			Heights:  []int{180, 165, 170},
			Expected: []string{"Mary", "Emma", "John"},
		},
		{
			Names:    []string{"Alice", "Bob", "Bob"},
			Heights:  []int{155, 185, 150},
			Expected: []string{"Bob", "Alice", "Bob"},
		},
	}

	for _, test := range tests {
		result := sortPeople(test.Names, test.Heights)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("expected: %v, got: %v", test.Expected, result)
		}
	}
}
