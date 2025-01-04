/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
  res := make([][]int, 0)
	backtrack(&res, nums, 0)
	return res
}

func backtrack(res *[][]int, track []int, index int) {

	if len(track) == index {
		t := make([]int, len(track))
		copy(t, track)
		*res = append(*res, t)
		return
	}

	// use hashmap to ignore duplicated element iteration with fixed `index`
	m := make(map[int]bool)

	for i := index; i < len(track); i++ {
		if m[track[i]] {
			continue
		} else {
			m[track[i]] = true
		}

		//  swap elements
		track[i], track[index] = track[index], track[i]
		// Recursively generate permutations for the next index
		backtrack(res, track, index+1)
		// restore the state
		track[i], track[index] = track[index], track[i]
	}
}
// @lc code=end

