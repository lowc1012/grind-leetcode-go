/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	if int(math.Abs(float64(rightHeight)-float64(leftHeight))) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

	return false
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(getHeight(node.Left), getHeight(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

