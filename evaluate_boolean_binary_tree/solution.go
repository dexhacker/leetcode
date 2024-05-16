package evaluate_boolean_binary_tree

import "leetcode/structures"

func evaluateTree(root *structures.TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	if root.Val == 2 {
		return left || right
	}

	return left && right
}
