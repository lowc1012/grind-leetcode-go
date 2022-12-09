/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	left, pivot := 0, 0
	right := len(nums)-1

	for left <= right {
		pivot = left + (right - left / 2)

		if target == nums[pivot] {
			return pivot
		}

		if target < nums[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}
// @lc code=end

