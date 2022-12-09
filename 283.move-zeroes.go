/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// the fast pointer is processing new elements
	slow, fast := 0, 0  
	for fast < len(nums) {

		// if the fast pointer finds a non-zero, swap the position with 
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}
// @lc code=end

