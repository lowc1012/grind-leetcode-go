/*
 * @lc app=leetcode id=1721 lang=golang
 *
 * [1721] Swapping Nodes in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	ahead, behind := head, head
	for n := ahead; n != nil; n = n.Next {
		k--
		if k == 0 {
			ahead = n
		}

		if k < 0 {
			behind = behind.Next
		}
	}
	ahead.Val, behind.Val = behind.Val, ahead.Val
	return head
}

// @lc code=end

