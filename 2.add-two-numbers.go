/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  dummyNode := &ListNode{
		Val: 0,
	}

	cursor := dummyNode
	carry := 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10
		cursor.Next = &ListNode{
			Val: sum % 10,
		}
		cursor = cursor.Next
	}

	// Final carry handling
	if carry > 0 {
		cursor.Next = &ListNode{
			Val: carry,
		}
	}
	// Return the next node of dummyNode
	return dummyNode.Next
}
// @lc code=end

