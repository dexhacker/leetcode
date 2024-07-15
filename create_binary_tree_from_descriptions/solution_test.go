package create_binary_tree_from_descriptions

import "testing"

func TestCreateBinaryTree(t *testing.T) {
	descriptions := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}
	result := createBinaryTree(descriptions)
	if result.Val != 50 {
		t.Errorf("result should be 50")
	}
	if result.Left.Val != 20 {
		t.Errorf("result should be 20")
	}
	if result.Right.Val != 80 {
		t.Errorf("result should be 80")
	}
	if result.Left.Left.Val != 15 {
		t.Errorf("result should be 15")
	}
	if result.Left.Right.Val != 17 {
		t.Errorf("result should be 17")
	}
	if result.Right.Left.Val != 19 {
		t.Errorf("result should be 19")
	}
}
