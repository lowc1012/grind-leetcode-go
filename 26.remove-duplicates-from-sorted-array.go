/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			// move slow pointer forward and replace the value
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// @lc code=end

