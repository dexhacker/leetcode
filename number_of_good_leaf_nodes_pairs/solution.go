package number_of_good_leaf_nodes_pairs

import "leetcode/structures"

var parentMp map[*structures.TreeNode]*structures.TreeNode
var leafs map[*structures.TreeNode]bool
var answer int

func dfs(root *structures.TreeNode, parent *structures.TreeNode) {
	if root == nil {
		return
	}

	if parent != nil {
		parentMp[root] = parent
	}

	if root.Left == nil && root.Right == nil {
		leafs[root] = true
	}

	dfs(root.Left, root)
	dfs(root.Right, root)
}

func findPath(prev *structures.TreeNode, current *structures.TreeNode, distance int, isRight bool) {
	if current == nil {
		return
	}

	if leafs[current] {
		answer++
		return
	}

	if distance == 0 {
		return
	}

	if prev != parentMp[current] {
		findPath(current, parentMp[current], distance-1, isRight)
	}
	if prev != current.Left && isRight {
		findPath(current, current.Left, distance-1, isRight)
	}
	if prev != current.Right {
		findPath(current, current.Right, distance-1, true)
	}
}

func countPairs(root *structures.TreeNode, distance int) int {
	parentMp = make(map[*structures.TreeNode]*structures.TreeNode)
	leafs = make(map[*structures.TreeNode]bool)
	answer = 0

	dfs(root, nil)

	for leaf := range leafs {
		findPath(leaf, parentMp[leaf], distance-1, false)
	}

	return answer
}
