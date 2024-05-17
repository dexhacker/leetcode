package delete_leaves_with_a_given_value

import (
	"leetcode/structures"
	"testing"
)

func TestRemoveLeafNodes(t *testing.T) {
	test1TreeNode1 := &structures.TreeNode{Val: 1}
	test1TreeNode2 := &structures.TreeNode{Val: 2}
	test1TreeNode3 := &structures.TreeNode{Val: 3}
	test1TreeNode4 := &structures.TreeNode{Val: 2}
	test1TreeNode5 := &structures.TreeNode{Val: 2}
	test1TreeNode6 := &structures.TreeNode{Val: 4}
	test1TreeNode1.Left = test1TreeNode2
	test1TreeNode1.Right = test1TreeNode3
	test1TreeNode2.Left = test1TreeNode4
	test1TreeNode3.Left = test1TreeNode5
	test1TreeNode3.Right = test1TreeNode6

	result := removeLeafNodes(test1TreeNode1, 2)
	if result.Left != nil {
		t.Errorf("removeLeafNodes(result.Left) != nil; expected nil")
	}
	if result.Right.Left != nil {
		t.Errorf("removeLeafNodes(result.Right.Left) != nil; expected nil")
	}

	test2TreeNode1 := &structures.TreeNode{Val: 1}
	test2TreeNode2 := &structures.TreeNode{Val: 3}
	test2TreeNode3 := &structures.TreeNode{Val: 3}
	test2TreeNode4 := &structures.TreeNode{Val: 3}
	test2TreeNode5 := &structures.TreeNode{Val: 2}
	test2TreeNode1.Left = test2TreeNode2
	test2TreeNode1.Right = test2TreeNode3
	test2TreeNode2.Left = test2TreeNode4
	test2TreeNode2.Right = test2TreeNode5

	result = removeLeafNodes(test2TreeNode1, 3)
	if result.Right != nil {
		t.Errorf("removeLeafNodes(result.Right) != nil; expected nil")
	}
	if result.Left.Left != nil {
		t.Errorf("removeLeafNodes(result.Left.Left) != nil; expected nil")
	}

	test3TreeNode1 := &structures.TreeNode{Val: 1}
	test3TreeNode2 := &structures.TreeNode{Val: 2}
	test3TreeNode3 := &structures.TreeNode{Val: 2}
	test3TreeNode4 := &structures.TreeNode{Val: 2}
	test3TreeNode1.Left = test3TreeNode2
	test3TreeNode2.Left = test3TreeNode3
	test3TreeNode3.Left = test3TreeNode4

	result = removeLeafNodes(test3TreeNode1, 2)
	if result.Left != nil {
		t.Errorf("removeLeafNodes(result.Left) != nil; expected nil")
	}
}
