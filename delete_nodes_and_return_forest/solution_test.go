package delete_nodes_and_return_forest

import (
	"leetcode/structures"
	"testing"
)

func TestDelNodes(t *testing.T) {
	treeNode4 := &structures.TreeNode{Val: 4}
	treeNode5 := &structures.TreeNode{Val: 5}
	treeNode6 := &structures.TreeNode{Val: 6}
	treeNode7 := &structures.TreeNode{Val: 7}
	treeNode2 := &structures.TreeNode{Val: 2, Left: treeNode4, Right: treeNode5}
	treeNode3 := &structures.TreeNode{Val: 3, Left: treeNode6, Right: treeNode7}
	treeNode1 := &structures.TreeNode{Val: 1, Left: treeNode2, Right: treeNode3}
	result := delNodes(treeNode1, []int{3, 5})

	if result[0].Val != 1 {
		t.Errorf("result[0].Val should be 1, but got %d", result[0].Val)
	}

	if result[0].Left.Val != 2 {
		t.Errorf("result[0].Left.Val should be 2, but got %d", result[0].Left.Val)
	}

	if result[0].Left.Left.Val != 4 {
		t.Errorf("result[0].Left.Left.Val should be 4, but got %d", result[0].Left.Left.Val)
	}

	if result[1].Val != 6 {
		t.Errorf("result[1].Val should be 6, but got %d", result[1].Val)
	}

	if result[2].Val != 7 {
		t.Errorf("result[2].Val should be 7, but got %d", result[2].Val)
	}
}
