/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	// In order to remove the list "head"
	// use Sentinel Node as a pseudo-head
	sentinel := ListNode{
		Next: head,
	}

	var pred = &sentinel
	var cursor = sentinel.Next // head

	// O(n)
	for cursor != nil && cursor.Next != nil {
		p := cursor
		// find duplicated nodes
		for p != nil && p.Next != nil && p.Val == p.Next.Val {
			p = p.Next
		}

		if p != cursor {
			// skip sublist to delete
			pred.Next = p.Next
			cursor = p.Next
		} else {
			pred = cursor
			cursor = cursor.Next
		}

	}

	return sentinel.Next
}
// @lc code=end

