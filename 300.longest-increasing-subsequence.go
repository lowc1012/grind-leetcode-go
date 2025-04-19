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
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
            dp[i] = max(dp[i], dp[j]+1)
        }
        }
    }

    // Find the final result
    maxLength := 1
    for _, v := range dp {
        if v > maxLength {
            maxLength = v
        }
    }

    return maxLength
}

// @lc code=end

