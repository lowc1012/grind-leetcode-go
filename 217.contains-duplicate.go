/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {

	set := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, exist := set[nums[i]]; exist {
			return true
		}

		set[nums[i]] = struct{}{}
	}
	return false
}

// @lc code=end

