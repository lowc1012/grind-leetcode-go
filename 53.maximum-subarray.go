/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Kadane’s Algorithm
	// find the maximum sum among all subarrays ending at that element. 
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum + nums[i], nums[i])
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

// @lc code=end

