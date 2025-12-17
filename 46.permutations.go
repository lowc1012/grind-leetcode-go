/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start

// TC: O(N!)
func permute(nums []int) [][]int {
	// nums = [1, 2, 3]
	res := make([][]int, 0)
	backtrack(&res, nums, 0)
	return res
}

func backtrack(res *[][]int, track []int, index int) {
	if len(track) == index {
		t := make([]int, index)
		copy(t, track)
		*res = append(*res, t)
		return
	}

	for i := index; i < len(track); i++ {
		track[i], track[index] = track[index], track[i]
		backtrack(res, track, index+1)
		track[i], track[index] = track[index], track[i]
	}
}

// @lc code=end

