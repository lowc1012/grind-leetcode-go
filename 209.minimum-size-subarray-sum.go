/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// two pointers
	minLen := len(nums) + 1
	left, right, sum := 0, 0, 0
	for ; right < len(nums); right++ {
		sum = sum + nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum = sum - nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

