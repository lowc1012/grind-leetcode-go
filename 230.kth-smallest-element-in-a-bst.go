/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	// 對 BST 進行 inorder traversal -> 排序

	valSlice := new([]int)

	inorderTraversal(root, valSlice)

	return (*valSlice)[k-1]
}

func inorderTraversal(root *TreeNode, valSlice *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, valSlice)
	*valSlice = append(*valSlice, root.Val)
	inorderTraversal(root.Right, valSlice)
}
// @lc code=end

