/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	traversal(root, &result, 1)

	return result
}

func traversal(node *TreeNode, res *[][]int, level int) {
	if node == nil {
		return
	}

	if level > len(*res) {
		*res = append(*res, []int{})
	}

	(*res)[level-1] = append((*res)[level-1], node.Val)

	traversal(node.Left, res, level+1)
	traversal(node.Right, res, level+1)

}

// @lc code=end

