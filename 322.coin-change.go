/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// DP, tabulation
	table := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i >= v {
				table[i] = min(table[i], 1+table[i-v])
			}
		}

	}

	// if table[amount] == 0 {
	// 	return -1
	// }
	return table[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

