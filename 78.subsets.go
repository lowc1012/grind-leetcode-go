/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start

// TC: O(N * 2^N), SC: O(N)
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack()
	return res
}

// 會產生 2^N 個結果
func backtrack(res *[][]int, track []int, index int, nums []int) {

	addToResult(res, curr)

	for i := index; i < len(nums); i++ {
		// select
		track = append(track, nums[i])

		// backtrack
		backtrack(res, track, i+1, nums)

		// deselect
		track = track[:len(track)-1]
	}

}

// O(N)
func addToResult(res *[][]int, curr []int) {
	s := make([]int, len(curr))
	copy(s, curr)
	*res = append(*res, s)
}

// @lc code=end

