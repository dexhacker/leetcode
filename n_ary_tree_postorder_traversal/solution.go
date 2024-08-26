package n_ary_tree_postorder_traversal

import "leetcode/structures"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var answer []int

func dfs(root *structures.Node) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		dfs(child)
	}
	answer = append(answer, root.Val)
}

func postorder(root *structures.Node) []int {
	answer = make([]int, 0)
	dfs(root)
	return answer
}
