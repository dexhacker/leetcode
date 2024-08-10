package regions_cut_by_slashes

import "testing"

func TestRegionsBySlashes(t *testing.T) {
	tests := []struct {
		Grid     []string
		Expected int
	}{
		{Grid: []string{" /", "/ "}, Expected: 2},
		{Grid: []string{" /", "  "}, Expected: 1},
		{Grid: []string{"/\\", "\\/"}, Expected: 5},
	}

	for _, test := range tests {
		result := regionsBySlashes(test.Grid)
		if result != test.Expected {
			t.Errorf("regionsBySlashes(%q) = %d; expected %d", test.Grid, result, test.Expected)
		}
	}
}
