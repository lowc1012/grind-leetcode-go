/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var fast, slow *ListNode
	fast, slow = head, head
	
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 此時 fast.Next = target node 
	if fast == nil {
		head = head.Next
		return head
	}

	// 用 while 迴圈，移動 slow pointer
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
// @lc code=end

