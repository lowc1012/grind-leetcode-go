/*
 * @lc app=leetcode id=1901 lang=golang
 *
 * [1901] Find a Peak Element II
 */

// @lc code=start
func getMaxRow(mat [][]int, col int) int {
    maxRow := 0
    for i := 1; i < len(mat); i++ {
        if mat[i][col] > mat[maxRow][col] {
            maxRow = i
        }
    }
    return maxRow
}

func findPeakGrid(mat [][]int) []int {
    _, n := len(mat), len(mat[0])
    left, right := 0, n-1
    for left <= right {
        midCol := left + (right-left)/2
        maxRow := getMaxRow(mat, midCol)

        // init left and right values
        leftVal, rightVal := -1, -1
        if midCol > 0 {
            leftVal = mat[maxRow][midCol-1]
        }
        if midCol < n-1 {
            rightVal = mat[maxRow][midCol+1]
        }

        // find the peak
        if mat[maxRow][midCol] > leftVal && mat[maxRow][midCol] > rightVal {
            return []int{maxRow, midCol}
        }

        // adjust left or right
        if mat[maxRow][midCol] < leftVal {
            right = midCol - 1
        } else {
            left = midCol + 1
        }
    }

    return []int{-1, -1}
}
// @lc code=end

