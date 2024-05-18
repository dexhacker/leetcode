package distribute_coins_in_binary_tree

import (
	"leetcode/structures"
	"testing"
)

func TestDistributeCoins(t *testing.T) {
	test1Node1 := &structures.TreeNode{Val: 3}
	test1Node2 := &structures.TreeNode{Val: 0}
	test1Node3 := &structures.TreeNode{Val: 0}
	test1Node1.Left = test1Node2
	test1Node1.Right = test1Node3

	test2Node1 := &structures.TreeNode{Val: 0}
	test2Node2 := &structures.TreeNode{Val: 3}
	test2Node3 := &structures.TreeNode{Val: 0}
	test2Node1.Left = test2Node2
	test2Node1.Right = test2Node3

	tests := []struct {
		Root     *structures.TreeNode
		Expected int
	}{
		{Root: test1Node1, Expected: 2},
		{Root: test2Node1, Expected: 3},
	}

	for _, test := range tests {
		result := distributeCoins(test.Root)
		if result != test.Expected {
			t.Errorf("DistributeCoins(%v) = %d; want %d", test.Root, result, test.Expected)
		}
	}
}
