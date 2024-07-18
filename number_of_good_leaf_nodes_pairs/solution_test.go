package number_of_good_leaf_nodes_pairs

import (
	"leetcode/structures"
	"testing"
)

func TestCountPairs(t *testing.T) {
	treeNode4 := &structures.TreeNode{Val: 4}
	treeNode5 := &structures.TreeNode{Val: 5}
	treeNode6 := &structures.TreeNode{Val: 6}
	treeNode7 := &structures.TreeNode{Val: 7}
	treeNode2 := &structures.TreeNode{Val: 2, Left: treeNode4, Right: treeNode5}
	treeNode3 := &structures.TreeNode{Val: 3, Left: treeNode6, Right: treeNode7}
	treeNode1 := &structures.TreeNode{Val: 1, Left: treeNode2, Right: treeNode3}
	result := countPairs(treeNode1, 4)
	if result != 6 {
		t.Errorf("result should be 6, got %d", result)
	}
}
