/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// use two pointers
	slow, fast := 0, 1

	// use a count to record the number of duplicates
	count := 1

	for fast < len(nums) {
		if nums[fast-1] == nums[fast] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}
	return slow + 1
}

// @lc code=end

