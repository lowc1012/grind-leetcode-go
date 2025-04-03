/*
 * @lc app=leetcode id=1380 lang=golang
 *
 * [1380] Lucky Numbers in a Matrix
 */

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	rowMins := make(map[int]bool)
	colMaxs := make(map[int]bool)

	for i := 0; i < m; i++ {
		minVal := matrix[i][0]
		for j := 0; j < n; j++ {
			if matrix[i][j] < minVal {
				minVal = matrix[i][j]
			}
		}
		rowMins[minVal] = true
	}

	for j := 0; j < n; j++ {
		maxVal := matrix[0][j]
		for i := 0; i < m; i++ {
			if matrix[i][j] > maxVal {
				maxVal = matrix[i][j]
			}
		}
		colMaxs[maxVal] = true
	}

    result := []int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rowMins[matrix[i][j]] &&
            colMaxs[matrix[i][j]] {
                result = append(result, matrix[i][j])
            }
        }
    }
    return result
}

// @lc code=end

