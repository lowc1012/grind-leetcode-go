/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, max, min *int) bool {
	if node == nil {
		return true
	}

	if (max != nil && node.Val >= *max) || (min != nil && node.Val <= *min) {
		return false
	}

	return validate(node.Left, &(node.Val), min) && validate(node.Right, max, &(node.Val))
}

// @lc code=end

