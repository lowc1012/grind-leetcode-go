/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, s)
	*s = append(*s, root.Val)
	traversal(root.Right, s)
}
// @lc code=end

