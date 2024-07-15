package create_binary_tree_from_descriptions

import (
	"leetcode/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *structures.TreeNode {
	mp := make(map[int]*structures.TreeNode)
	parentMap := make(map[*structures.TreeNode]*structures.TreeNode)

	var root *structures.TreeNode
	for _, item := range descriptions {
		val := item[0]
		valChild := item[1]
		isLeft := item[2] == 1
		var currentNode *structures.TreeNode
		var childNode *structures.TreeNode

		if node, ok := mp[val]; ok {
			currentNode = node
		} else {
			currentNode = &structures.TreeNode{Val: val}
			mp[val] = currentNode
		}

		if node, ok := mp[valChild]; ok {
			childNode = node
		} else {
			childNode = &structures.TreeNode{Val: valChild}
			mp[valChild] = childNode
		}

		if isLeft {
			currentNode.Left = childNode
		} else {
			currentNode.Right = childNode
		}
		parentMap[childNode] = currentNode

		if root == nil {
			root = currentNode
		}

		for parentMap[root] != nil {
			root = parentMap[root]
		}
	}

	return root
}
