/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	count := 1
	top, bottom, left, right := 0, n-1, 0, n-1

	for count <= n*n {

		for i := left; i <= right; i++ {
			res[top][i] = count
			count++
		}

		for i := top + 1; i <= bottom; i++ {
			res[i][right] = count
			count++
		}

		for i := right - 1; i >= left; i-- {
			res[bottom][i] = count
			count++
		}

		for i := bottom - 1; i > top; i-- {
			res[i][left] = count
			count++
		}

		top++
		bottom--
		left++
		right--

	}

	return res
}

// @lc code=end

