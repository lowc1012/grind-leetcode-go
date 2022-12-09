/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {

	// record the (k - node.Val)
	m := make(map[int]struct{})
	res := traversal(root, k, m)
	return res
}

func traversal(node *TreeNode, k int, m map[int]struct{}) bool {
	if node == nil {
		return false
	}

	if _, exist := m[node.Val]; exist {
		return true
	} else {
		m[k-node.Val] = struct{}{}
	}

	leftCheck := traversal(node.Left, k, m)
	rightCheck := traversal(node.Right, k, m)
	return rightCheck || leftCheck
}

// @lc code=end

