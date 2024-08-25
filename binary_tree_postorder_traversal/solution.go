package binary_tree_postorder_traversal

import "leetcode/structures"

var answer []int

func dfs(root *structures.TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	dfs(root.Right)
	answer = append(answer, root.Val)
}

func postorderTraversal(root *structures.TreeNode) []int {
	answer = make([]int, 0)
	dfs(root)
	return answer
}
