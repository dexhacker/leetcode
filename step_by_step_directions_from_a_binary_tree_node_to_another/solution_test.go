package step_by_step_directions_from_a_binary_tree_node_to_another

import (
	"leetcode/structures"
	"testing"
)

func TestGetDirections(t *testing.T) {
	treeNode3 := &structures.TreeNode{Val: 3}
	treeNode6 := &structures.TreeNode{Val: 6}
	treeNode4 := &structures.TreeNode{Val: 4}
	treeNode1 := &structures.TreeNode{Val: 1, Left: treeNode3}
	treeNode2 := &structures.TreeNode{Val: 2, Left: treeNode6, Right: treeNode4}
	treeNode5 := &structures.TreeNode{Val: 5, Left: treeNode1, Right: treeNode2}
	result := getDirections(treeNode5, 3, 6)
	if result != "UURL" {
		t.Errorf("result should be 'UURL'")
	}
}
