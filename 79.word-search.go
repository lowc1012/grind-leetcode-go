/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	// base case
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	// explore the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, i, j, word, 0) {
				return true
			}

		}
	}

	return false
}

func backtrack(board [][]byte, i, j int, word string, index int) bool {
	// result condition
	if index == len(word) {
		return true
	}

	// check the boundary
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	// check the visited cell
	} else if board[i][j] == '#' {
		return false
	// check the current character
	} else if board[i][j] != word[index] {
		return false
	}

	// mark the visited cell 
	temp := board[i][j]
	board[i][j] = '#'

	// explore the neighbors
	if backtrack(board, i+1, j, word, index+1) ||
		backtrack(board, i-1, j, word, index+1) ||
		backtrack(board, i, j+1, word, index+1) ||
		backtrack(board, i, j-1, word, index+1) {
		return true
	}

	// restore the state
	board[i][j] = temp

	return false
}

// @lc code=end

