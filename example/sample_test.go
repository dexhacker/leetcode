package example

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) = %d; expected %d", result, expected)
	}
}
