/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start

// what I know
// 1. each row/col are sorted in ascending left/top to right/bottom
// => binary search

func searchMatrix(matrix [][]int, target int) bool {

	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		} else {
			return true
		}
	}
	// TC: O(M+N), SC: O(1)
	return false
}

// @lc code=end

