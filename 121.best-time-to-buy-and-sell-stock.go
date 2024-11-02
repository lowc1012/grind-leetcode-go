/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 1. brute-force
	// using nested loops to iterate over all possible subarrays
	// outer loop marks the start point, inner loop marks the end point of subarray
	// to find the max profix (end - start)

	// 2. kadane's algo
	maxP := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		maxP = max(prices[i]-minPrice, maxP)
		minPrice = min(minPrice, prices[i])
	}

	return maxP
}

// @lc code=end

