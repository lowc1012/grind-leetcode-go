/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, m, n int) {
	if m >= 0 && n >= 0 && m < len(grid) && n < len(grid[0]) && grid[m][n] == '1' {
		grid[m][n] = '$'
		dfs(grid, m+1, n)
		dfs(grid, m-1, n)
		dfs(grid, m, n+1)
		dfs(grid, m, n-1)
	}
}

// @lc code=end

