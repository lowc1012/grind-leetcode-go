/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {

    m, n := len(grid), len(grid[0])

	// initialize DP table
	dp := make([][]int, m)
	for idx, _ := range dp {
		dp[idx] = make([]int, n)
	}

	// Initialize Starting Point
	dp[0][0] = grid[0][0]

	// fill first row
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// fill first column
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// fill the rest of table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
// @lc code=end

