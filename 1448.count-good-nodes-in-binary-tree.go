/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := recursion(root, root.Val)
	return res
}

func recursion(node *TreeNode, max int) int {

	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= max {
		count = 1
		max = node.Val
	}

	return count + recursion(node.Left, max) + recursion(node.Right, max)
}

// @lc code=end

