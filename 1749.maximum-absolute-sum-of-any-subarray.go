/*
 * @lc app=leetcode id=1749 lang=golang
 *
 * [1749] Maximum Absolute Sum of Any Subarray
 */

// @lc code=start
func maxAbsoluteSum(nums []int) int {
	// Kadane’s Algorithm
	n := len(nums)
	curMax := 0
	curMin := 0
	maxSum := curMax
	minSum := curMin

	for i := 0; i < n; i++ {
		curMax = max(curMax + nums[i], nums[i])
		maxSum = max(curMax, maxSum)

		curMin = min(curMin + nums[i], nums[i])
		minSum = min(curMin, minSum)
	}

	return max(abs(minSum), abs(maxSum))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

