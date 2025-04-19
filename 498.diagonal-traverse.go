/*
 * @lc app=leetcode id=498 lang=golang
 *
 * [498] Diagonal Traverse
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	rows, cols := len(mat), len(mat[0])
	if rows == 0 || cols == 0 {
		return []int{}
	}

	res := make([]int, 0, rows*cols)
	rightUp, leftDown := [2]int{-1, 1}, [2]int{1, -1}
	moveRightUp := true

	i, j := 0, 0
	for len(res) < rows*cols {
		res = append(res, mat[i][j])

		if moveRightUp {
			nextI := i + rightUp[0]
			nextJ := j + rightUp[1]

			if nextJ == cols {
				i++
				moveRightUp = false
			} else if nextI < 0 {
				j++
				moveRightUp = false
			} else {
				i, j = nextI, nextJ
			}

		} else {
			nextI := i + leftDown[0]
			nextJ := j + leftDown[1]

			if nextI == rows {
				j++
				moveRightUp = true
			} else if nextJ < 0 {
				i++
				moveRightUp = true
			} else {
				i, j = nextI, nextJ
			}
		}
	}
    return res
}

// @lc code=end

