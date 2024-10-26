/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// Kadane’s Algorithm
	currentSum := 0
	maxSum := nums[0]
	
	for i := 0; i < len(nums); i++ {
		currentSum = max(currentSum, 0)
		currentSum += nums[i]
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

// @lc code=end

