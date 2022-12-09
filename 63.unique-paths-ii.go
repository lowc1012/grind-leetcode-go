/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	table := make([][]int, m)
	for i, _ := range table {
		table[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 0 {
					table[i][j] = 1
				} else {
					return 0
				}
			} else if i == 0 && obstacleGrid[i][j] == 0 {
				table[i][j] = table[i][j-1]
			} else if j == 0 && obstacleGrid[i][j] == 0 {
				table[i][j] = table[i-1][j]
			} else {
				if obstacleGrid[i][j] == 0 {
					table[i][j] = table[i-1][j] + table[i][j-1]
				}
			}
		}
	}
	// TC: O(mn), SC: O(mn)
	return table[m-1][n-1]
}

// @lc code=end

