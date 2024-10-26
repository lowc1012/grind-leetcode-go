/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
    maxProfit := 0
		minPrice := prices[0]

		for i := 1; i < len(prices); i++ {
			maxProfit = max(prices[i] - minPrice, maxProfit)
			minPrice = min(minPrice, prices[i])
		}

		return maxProfit
}
// @lc code=end

