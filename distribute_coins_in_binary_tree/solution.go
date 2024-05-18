package distribute_coins_in_binary_tree

import "leetcode/structures"

var answer int

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func dfs(root *structures.TreeNode) int {
	coins := root.Val - 1

	if root.Left != nil {
		result := dfs(root.Left)
		answer += abs(result)
		coins += result
	}

	if root.Right != nil {
		result := dfs(root.Right)
		answer += abs(result)
		coins += result
	}

	return coins
}

func distributeCoins(root *structures.TreeNode) int {
	answer = 0
	dfs(root)
	return answer
}
