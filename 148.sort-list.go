/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// merge sort, Time complexity: O(nlogn)
	if head == nil || head.Next == nil {
		return head
	}

	// step 1: split the list to two halves
	// Finding the middle element
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// slow point to the middle of list
		fast, slow, prev = fast.Next.Next, slow.Next, slow
	}

	// prev point to end of first left half
	prev.Next = nil

	// step 2: sort each halves
	h1, h2 := sortList(head), sortList(slow)

	// step3: merge h1 and h2
	return merge(h1, h2)
}

func merge(h1, h2 *ListNode) *ListNode {
	// Create a new linked list
	dummy := &ListNode{}
	p := dummy
	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			p.Next = h1
			h1 = h1.Next
		} else {
			p.Next = h2
			h2 = h2.Next
		}
		p = p.Next
	}
	if h1 != nil {
		p.Next = h1
	}

	if h2 != nil {
		p.Next = h2
	}

	return dummy.Next
}

// @lc code=end

