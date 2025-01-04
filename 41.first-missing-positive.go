/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// Place each number in its correct position
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// Swap the current number with the number at its correct position
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// Find the first missing positive
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// @lc code=end

