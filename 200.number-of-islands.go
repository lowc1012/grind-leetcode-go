/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
// time complexity: O(M * N)
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

// what if the matrix is too large?
// we can use a stack to simulate the recursive call to avoid deep recursion

// iterative DFS to mark the connected land cells as visited
func dfs(grid [][]byte, m, n int) {
	// use a stack to simulate the recursive call
	stack := make([][]int, 0)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// push the current land cell into the stack
	stack = append(stack, []int{m, n})

	// mark the current cell as visited
	for len(stack) > 0 {
		// pop the top element from the stack
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		row, col := cell[0], cell[1]
		if (row >= 0 && row < len(grid)) && (col >= 0 && col < len(grid[0])) && grid[row][col] == '1' {
			grid[row][col] = 'v' // mark as visited

			// push the neighbors into the stack
			for _, dir := range directions {
				stack = append(stack, []int{row + dir[0], col + dir[1]})
			}
		}
	}
}

// @lc code=end

