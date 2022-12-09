/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start

func lengthOfLIS(nums []int) int {
	return tabulation(nums, len(nums))
}

func tabulation(s []int, n int) int {
	// create a cache to store the LIS values
	// for ith element in input array
	// SC: O(n)
	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i]++
	}

	// Compute the LIS values in bottom-up manner
	// TC: O(log n*n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// if number increase && cache table need updated
			if s[i] > s[j] {
				cache[i] = max(cache[i], cache[j]+1)
			}
		}
	}

	max := 1
	for _, v := range cache {
		if max < v {
			max = v
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

