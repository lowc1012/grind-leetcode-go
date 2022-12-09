/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	prv, curr := dummy, head
	for curr != nil && curr.Next != nil {
		prv.Next = curr.Next
		curr.Next = prv.Next.Next
		prv.Next.Next = curr

		prv = curr
		curr = curr.Next
	}

	return dummy.Next
}

// @lc code=end

