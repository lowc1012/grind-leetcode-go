/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := make([][]int, 0)
	backtrack(&res, []int{}, candidates, 0, 0, target)
	return res
}

func backtrack(res *[][]int, track []int, candidates []int, start, currSum, target int) {
	// end condition
	if currSum == target {
		s := make([]int, len(track))
		copy(s, track)
		*res = append(*res, s)
	}

	for i := start; i < len(candidates); i++ {
		if (currSum + candidates[i]) > target {
			break
		}
		track = append(track, candidates[i])
		currSum += candidates[i]
		backtrack(res, track, candidates, i, currSum, target)
		track = track[:len(track)-1]
		currSum -= candidates[i]
	}

}

// @lc code=end

