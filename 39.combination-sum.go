/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtrack(&result, []int{}, candidates, 0, target)
	return result
}

func backtrack(res *[][]int, track []int, candidates []int, start, target int) {
	if target < 0 {
		return
	}

	// desired condition
	if target == 0 {
		*res = append(*res, append([]int{}, track...))
		return
	}

	for i := start; i < len(candidates); i++ {
		track = append(track, candidates[i])
		backtrack(res, track, candidates, i, target-candidates[i])
		track = track[:len(track)-1]
	}
}

// @lc code=end

