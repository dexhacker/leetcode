package evaluate_boolean_binary_tree

import (
	"leetcode/structures"
	"testing"
)

func TestEvaluateTree(t *testing.T) {
	test1TreeNode1 := &structures.TreeNode{Val: 2}
	test1TreeNode2 := &structures.TreeNode{Val: 1}
	test1TreeNode3 := &structures.TreeNode{Val: 3}
	test1TreeNode4 := &structures.TreeNode{Val: 0}
	test1TreeNode5 := &structures.TreeNode{Val: 1}
	test1TreeNode1.Left = test1TreeNode2
	test1TreeNode1.Right = test1TreeNode3
	test1TreeNode3.Left = test1TreeNode4
	test1TreeNode3.Right = test1TreeNode5

	test2TreeNode1 := &structures.TreeNode{Val: 0}

	tests := []struct {
		Root     *structures.TreeNode
		Expected bool
	}{
		{Root: test1TreeNode1, Expected: true},
		{Root: test2TreeNode1, Expected: false},
	}

	for _, test := range tests {
		result := evaluateTree(test.Root)
		if result != test.Expected {
			t.Errorf("test: %+v, expected: %t, got: %t", test.Root, test.Expected, result)
		}
	}
}
