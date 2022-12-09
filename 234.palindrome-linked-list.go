/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    nums := make([]int, 0)

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	j := len(nums)-1
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[j] {
			return false
		}
		j--
	}
	return true
}
// @lc code=end

