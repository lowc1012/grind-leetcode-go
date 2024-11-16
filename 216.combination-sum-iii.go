/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backtrack(&res, []int{}, 0, k, n)
	return res
}

func backtrack(res *[][]int, track []int, start, k, target int) {

	// end condition
	if target == 0 && len(track) == k {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 1 + start; i <= 9; i++ {
		track = append(track, i)
		backtrack(res, track, i, k, target-i)
		track = track[:len(track)-1]
	}
}

// @lc code=end

