package step_by_step_directions_from_a_binary_tree_node_to_another

import "leetcode/structures"

var parentMp map[*structures.TreeNode]*structures.TreeNode
var sourceValue int
var targetValue int
var source *structures.TreeNode
var target *structures.TreeNode
var answer string

func parentDfs(root *structures.TreeNode, parent *structures.TreeNode) {
	if root == nil {
		return
	}

	if parent != nil {
		parentMp[root] = parent
	}

	if root.Val == sourceValue {
		source = root
	}

	if root.Val == targetValue {
		target = root
	}

	parentDfs(root.Left, root)
	parentDfs(root.Right, root)
}

func dfs(start *structures.TreeNode, end *structures.TreeNode, parent *structures.TreeNode, letters []rune) {
	if start == nil {
		return
	}

	if start == end {
		answer = string(letters)
		return
	}

	size := len(letters)

	if start.Left != parent {
		letters = append(letters, 'L')
		dfs(start.Left, end, start, letters)
	}

	if start.Right != parent {
		letters = letters[:size]
		letters = append(letters, 'R')
		dfs(start.Right, end, start, letters)
	}

	if parentMp[start] != parent {
		letters = letters[:size]
		letters = append(letters, 'U')
		dfs(parentMp[start], end, start, letters)
	}
}

func getDirections(root *structures.TreeNode, startValue int, destValue int) string {
	parentMp = make(map[*structures.TreeNode]*structures.TreeNode)
	answer = ""
	sourceValue = startValue
	targetValue = destValue
	letters := make([]rune, 0)

	parentDfs(root, nil)
	dfs(source, target, nil, letters)
	return answer
}
