/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start

func lengthOfLIS(nums []int) int {
	// Dynamic Programming
	if len(nums) == 0 {
		return 0
	}

	// Init DP state
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // LIS is at least one
	}

	// Recurrence Relation
	// For each element at index i (from 1 to n-1), compare it with all elements before it (from 0 to i-1).
	// If nums[j] < nums[i] (where j ranges from 0 to i-1), it means the element at index i can be a part of a longer subsequence.
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// increasing
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
            // find the result
			if dp[i] > maxLength {
				maxLength = dp[i]
			}
		}
	}

	return maxLength
}

// @lc code=end

