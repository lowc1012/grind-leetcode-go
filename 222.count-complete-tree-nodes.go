/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	res := 0
	if root != nil {
		dfsTraversal(root, &res)
		res += 1
	}

	return res
}

func dfsTraversal(node *TreeNode, count *int) {
	if node.Right == nil && node.Left == nil {
		return
	} else {

		if node.Right != nil {
			*count += 1
			dfsTraversal(node.Right, count)
		}

		if node.Left != nil {
			*count += 1
			dfsTraversal(node.Left, count)
		}
	}
}

// @lc code=end

