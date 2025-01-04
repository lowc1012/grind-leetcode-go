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

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	return validate(root, math.Inf(-1), math.Inf(1))
}

func validate(node *TreeNode, min, max float64) bool {
	if node == nil {
		return true
	}

	if (float64(node.Val) >= max) || (float64(node.Val) <= min) {
		return false
	}

	return validate(node.Left, min, float64(node.Val)) && validate(node.Right, float64(node.Val), max)
}

// @lc code=end

