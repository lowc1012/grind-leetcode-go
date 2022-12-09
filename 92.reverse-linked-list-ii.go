/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// head = [1,2,3,4,5], left = 2, right = 4
	// target group [2, 3, 4]
	start := head
	end := head
	var prev *ListNode = nil
	var next *ListNode = nil
	for i := 1; i < right; i++ {
		end = end.Next
		if i < left {
			prev = start
			start = start.Next
		}
	}
	// after this for-loop
	// prv start end
	//  |  |     |
	//[ 1, 2, 3, 4, 5]

	// store the next node of target group
	next = end.Next

	// reverse the list between start and end node (target group)
	reverse(start, end)

	// concatenate prev node and first node of "reversed" list
	if prev != nil {
		prev.Next = end
	}

	// concatenate last node of reversed list and
	// the next node of target group
	start.Next = next

	// start node is the first node
	// head will be the new start of the reversed list
	if left == 1 {
		head = end
	}
	return head
}

// reverse a list given the start and end node
func reverse(start, end *ListNode) {
	var prev *ListNode = nil
	var next *ListNode = nil
	curr := start
	for curr != end {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr.Next = prev
}

// @lc code=end

