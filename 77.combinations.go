/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrack([]int{}, &res, 1, n, k)
	return res
}

func backtrack(curr []int, res *[][]int, start, n, k int) {
	if k == 0 {
		result := append([]int{}, curr...)
		*res = append(*res, result)
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

