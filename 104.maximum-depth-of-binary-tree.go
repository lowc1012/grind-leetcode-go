/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	maxL := maxDepth(root.Left)
	maxR := maxDepth(root.Right)
	
	if maxL > maxR {
		return maxL + 1
	} else {
		return maxR + 1
	}
}
// @lc code=end

