/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// TC: O(n)
	if root == nil {
		return nil
	}

	connect(root.Left)
	connect(root.Right)
	merge(root.Left, root.Right)
	return root
}

func merge(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	merge(left.Right, right.Left)
}

// @lc code=end

