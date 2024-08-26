package n_ary_tree_postorder_traversal

import (
	"leetcode/structures"
	"testing"
)

func TestPostorder(t *testing.T) {
	node1 := &structures.Node{Val: 1}
	node3 := &structures.Node{Val: 3}
	node2 := &structures.Node{Val: 2}
	node4 := &structures.Node{Val: 4}
	node5 := &structures.Node{Val: 5}
	node6 := &structures.Node{Val: 6}
	node3.Children = append(node3.Children, node5, node6)
	node1.Children = append(node1.Children, node3, node2, node4)
	expected := []int{5, 6, 3, 2, 4, 1}
	result := postorder(node1)
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("postorder(%d) failed. Expected: %d, got: %d", i, expected[i], result[i])
		}
	}
}
