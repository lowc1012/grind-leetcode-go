/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	// return the node as the root if there is only one node 
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// use two pointers to find the middle of the list
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow // save the previous node of slow
		slow = slow.Next // move slow one step
		fast = fast.Next.Next // move fast two steps
	}

	// slow is the middle node
	root := &TreeNode{
		Val: slow.Val,
	}

	// break the list into two halves
	// Use a prev pointer to keep track of the node
	// before slow. When slow reaches the middle,
	// set prev.Next to nil to break the list into two halves.
	if prev != nil {
		prev.Next = nil
	}

	// build the left and right subtrees
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}

// @lc code=end

