/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int)  {
	// nums = []int{1,2,3,4,5,6,7}, k = 3

    n := len(nums)
	k = k % n // ensure k is within the bounds of n
	reverse(nums, 0, n-1) // nums = [7,6,5,4,3,2,1]
	reverse(nums, 0, k-1) // nums = [5,6,7,4,3,2,1]
	reverse(nums, k, n-1) // nums = [5,6,7,1,2,3,4]
}
// @lc code=end

