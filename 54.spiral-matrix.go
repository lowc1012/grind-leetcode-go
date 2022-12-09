/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	n, m := len(matrix), len(matrix[0])

	top, bottom, left, right := 0, n-1, 0, m-1
	for len(res) < n*m {

		for i := left; i <= right && len(res) < n*m; i++ {
			res = append(res, matrix[top][i])
		}

		for i := top + 1; i <= bottom && len(res) < n*m; i++ {
			res = append(res, matrix[i][right])
		}

		for i := right - 1; i >= left && len(res) < n*m; i-- {
			res = append(res, matrix[bottom][i])
		}

		for i := bottom - 1; i > top && len(res) < n*m; i-- {
			res = append(res, matrix[i][left])
		}

		top++
		right--
		left++
		bottom--

	}
	return res
}

// @lc code=end

