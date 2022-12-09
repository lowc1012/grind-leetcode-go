/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	// nums [1, 2, 3, 4]

	result := make([]int, len(nums))

	prefix := 1
	for i, _ := range nums {
		result[i] = prefix
		prefix = prefix * nums[i]
	}
	// prefix [1, 1, 2, 6]

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = postfix * result[i]
		postfix = postfix * nums[i]
	}

	return result
}

// @lc code=end

