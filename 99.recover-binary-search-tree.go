/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	traversal(root, &first, &second, &prev)
	first.Val, second.Val = second.Val, first.Val
}

// 對 BST inorder, 輸出結果會是由小到大的數列
func traversal(node *TreeNode, first, second, prev **TreeNode) {
	if node == nil {
		return
	}

	traversal(node.Left, first, second, prev)

	// 若發現前一個 node 大於/等於 後面一個 node => swap
	if *prev != nil && (*prev).Val >= node.Val {
		if *first == nil {
			*first = *prev
		}
		*second = node
	}
	*prev = node

	traversal(node.Right, first, second, prev)
}

// @lc code=end

