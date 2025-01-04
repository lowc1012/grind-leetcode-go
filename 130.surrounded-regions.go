/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte) {
	// Mark all 'O' cells connected to the boundary with a temporary marker '$'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	// Convert all remaining 'O' cells to 'X' and revert '$' back to 'O'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '$' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, m, n int) {
	stack := make([][]int, 0)
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	stack = append(stack, []int{m, n})
	for len(stack) > 0 {
		// pop the top element of stack
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		row, col := cell[0], cell[1]
		if row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) && board[row][col] == 'O' {
			// mark the cell as '$'
			board[row][col] = '$'
			for _, d := range directions {
				stack = append(stack, []int{row + d[0], col + d[1]})
			}
		}
	}
}

// @lc code=end
