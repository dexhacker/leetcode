package binary_tree_postorder_traversal

import (
	"leetcode/structures"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	treeNode1 := &structures.TreeNode{Val: 1}
	treeNode2 := &structures.TreeNode{Val: 2}
	treeNode3 := &structures.TreeNode{Val: 3}
	treeNode1.Right = treeNode2
	treeNode2.Left = treeNode3
	expected := []int{3, 2, 1}
	result := postorderTraversal(treeNode1)
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], result[i])
		}
	}
}
