/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    res := make([][]int, 0)

		sort.Ints(candidates)
		backtrack(&res, []int{}, candidates, target, 0)
		return res
}

func backtrack(res *[][]int, track []int, candidates []int, target, start int) {
	// end condition
	if target == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// remove duplicated combinations
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		
		// early termination to avoid TLE(Time Limit Exceeded)
		if candidates[i] > target {
			break
		}
		
		track = append(track, candidates[i])
		backtrack(res, track, candidates, target - candidates[i], i + 1)
		track = track[:len(track)-1]
	}
}

// @lc code=end

