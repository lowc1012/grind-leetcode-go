/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    fresh := 0
    queue := make([][]int, 0)

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, []int{i, j})
            } else if grid[i][j] == 1 {
                fresh++
            }
        }
    }

    if fresh == 0 {
        return 0
    }

    dir := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    result := 0
    for len(queue) > 0 && fresh > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            curr := queue[0]
            queue = queue[1:]
            for _, d := range dir {
                x, y := curr[0]+d[0], curr[1]+d[1]
                if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == 1 {
                    grid[x][y] = 2
                    fresh--
                    queue = append(queue, []int{x, y})
                }
            }
        }
        result++
    }
    if fresh > 0 {
        return -1
    }
    return result
}
// @lc code=end

