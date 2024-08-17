/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	// use a helper function to build BST
	// the time complexity is O(n) where n is the number of elements in the array
	return buildTree(nums, 0, len(nums)-1)
}

// buildTree builds a BST from the sorted array nums
func buildTree(nums []int, left, right int) *TreeNode {
	// the stop condition
	if left > right {
		return nil
	}

	// find the middle element as the root
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}

	// build left and right subtrees
	root.Left = buildTree(nums, left, mid-1)
	root.Right = buildTree(nums, mid+1, right)
	return root
}
// @lc code=end

