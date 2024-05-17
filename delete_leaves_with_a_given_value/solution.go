package delete_leaves_with_a_given_value

import "leetcode/structures"

func isLeave(root *structures.TreeNode, target int) bool {
	if root.Left != nil {
		leave := isLeave(root.Left, target)
		if leave && root.Left.Val == target {
			root.Left = nil
		}
	}

	if root.Right != nil {
		leave := isLeave(root.Right, target)
		if leave && root.Right.Val == target {
			root.Right = nil
		}
	}

	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

func removeLeafNodes(root *structures.TreeNode, target int) *structures.TreeNode {
	leave := isLeave(root, target)
	if leave && root.Val == target {
		return nil
	}
	return root
}
