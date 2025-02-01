/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, n := range nums {
		switch n {
		case 0:
			red++
		case 1:
			white++
		default:
			blue++
		}
	}

	curr := 0
	for i := 0; i < red; i++ {
		nums[curr] = 0
		curr++
	}
	for i := 0; i < white; i++ {
		nums[curr] = 1
		curr++
	}
	for i := 0; i < blue; i++ {
		nums[curr] = 2
		curr++
	}
}

// @lc code=end

