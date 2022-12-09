/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start

// TC: O(N!)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(&res, 0, nums)
	return res
}

func backtrack(res *[][]int, index int, nums []int) {

	if index == len(nums) {
		s := make([]int, len(nums))
		copy(s, nums)
		*res = append(*res, s)
	}

	// walk through nums, and swap two of elements in nums
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtrack(res, index+1, nums)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

// @lc code=end

