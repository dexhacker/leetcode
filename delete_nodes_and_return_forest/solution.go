package delete_nodes_and_return_forest

import "leetcode/structures"

var mp map[*structures.TreeNode]bool

func dfsDelete(item int, root *structures.TreeNode) bool {
	if root.Left != nil {
		if root.Left.Val == item {
			if root.Left.Left != nil {
				mp[root.Left.Left] = true
			}
			if root.Left.Right != nil {
				mp[root.Left.Right] = true
			}
			root.Left = nil
			return true
		}

		if dfsDelete(item, root.Left) {
			return true
		}
	}

	if root.Right != nil {
		if root.Right.Val == item {
			if root.Right.Left != nil {
				mp[root.Right.Left] = true
			}
			if root.Right.Right != nil {
				mp[root.Right.Right] = true
			}
			root.Right = nil
			return true
		}

		if dfsDelete(item, root.Right) {
			return true
		}
	}

	return false
}

func deleteItem(item int) {
	for subNode := range mp {
		if subNode.Val == item {
			delete(mp, subNode)
			if subNode.Left != nil {
				mp[subNode.Left] = true
			}
			if subNode.Right != nil {
				mp[subNode.Right] = true
			}
			return
		} else {
			if dfsDelete(item, subNode) {
				return
			}
		}
	}
}

func delNodes(root *structures.TreeNode, toDelete []int) []*structures.TreeNode {
	mp = make(map[*structures.TreeNode]bool)
	mp[root] = true

	for _, item := range toDelete {
		deleteItem(item)
	}

	answer := make([]*structures.TreeNode, 0)
	for k := range mp {
		answer = append(answer, k)
	}

	return answer
}
