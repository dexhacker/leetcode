package find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance

import "testing"

func TestFindTheCity(t *testing.T) {
	tests := []struct {
		N                 int
		Edges             [][]int
		DistanceThreshold int
		Expected          int
	}{
		{
			N:                 4,
			Edges:             [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
			DistanceThreshold: 4,
			Expected:          3,
		},
		{
			N:                 5,
			Edges:             [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}},
			DistanceThreshold: 2,
			Expected:          0,
		},
	}
	for _, test := range tests {
		result := findTheCity(test.N, test.Edges, test.DistanceThreshold)
		if result != test.Expected {
			t.Errorf(
				"N = %d; Edges %v; Distance Threshold = %d; Expected %d",
				test.N,
				test.Edges,
				test.DistanceThreshold,
				result,
			)
		}
	}
}
