/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])

	// Init DP state
	// The value dp[i][j] will represent the side length of the largest square composed
	// entirely of '1's whose bottom-right corner is at the original matrix cell matrix[i-1][j-1].
	dp := make([][]int, rows+1) // for edge case
	for i := range dp {
		dp[i] = make([]int, cols+1) // for edge case
	}

	// Recurrence Relation
	// When i=1 or j=1, if matrix[i-1][j-1] is '1', the min function involving the initial zeros will
	// evaluate to 0, correctly resulting in dp[i][j] = 0 + 1 = 1.
	maxSide := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			// If matrix[i-1][j-1] is '1', then a square can end here. The maximum possible side
			// length of this square is limited by the squares ending at its top (dp[i-1][j]),
			// left (dp[i][j-1]), and top-left (dp[i-1][j-1]) neighbors.
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				// Update the maximum side length found so far
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

// @lc code=end

