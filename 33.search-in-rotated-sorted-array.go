/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	// two pointers
	left, right := 0, len(nums)-1

	for left <= right {
		// using binary search because we must write an algorithm with O(log n) runtime complexity.
		mid := (left + right) / 2

		// Check if the target is at the mid index
        if nums[mid] == target {
            return mid
        } else if nums[left] <= nums[mid] {
            // Left part is sorted
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Right part is sorted
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
	}
	return -1
}
// @lc code=end

