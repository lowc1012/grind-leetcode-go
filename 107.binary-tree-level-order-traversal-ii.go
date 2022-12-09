/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int

	traversal(root, &result, 0)

	reverse(result)

	return result
}

func traversal(node *TreeNode, result *[][]int, level int) {
	if node == nil {
		return
	}

	if level == len(*result) {
		(*result) = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], node.Val)

	traversal(node.Left, result, level+1)
	traversal(node.Right, result, level+1)
}

func reverse(s [][]int) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// @lc code=end

