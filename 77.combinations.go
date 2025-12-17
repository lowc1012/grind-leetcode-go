/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
    backtrack(&res, []int{}, 1, n, k)
    return res
}

func backtrack(res *[][]int, track []int, start, n, k int) {
    // end condition
    if k == 0 {
        t := make([]int, len(track))
        copy(t, track)
        *res = append(*res, t)
        return
    }

    // recursive exploration
    for i := start; i <= n; i++ {
        track = append(track, i)
        backtrack(res, track, i+1, n, k-1)
        track = track[:len(track)-1]
    }
}

// @lc code=end

