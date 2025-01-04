/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	track := make([]int, 0, k)
	backtrack(track, &res, 1, n, k)
	return res
}

func backtrack(curr []int, res *[][]int, start, n, k int) {
	if k == 0 {
		track := make([]int, len(curr))
		copy(track, curr)
		*res = append(*res, track)
		return
	}

	for i := start; i <= n; i++ {
		curr = append(curr, i)
		backtrack(curr, res, i+1, n, k-1)
		curr = curr[:len(curr)-1]
	}
	return
}

// @lc code=end

