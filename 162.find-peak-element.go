/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	// binary search
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

