/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	white, blue := 0, 0
	counter := 0

	for _, n := range nums {
		if n == 0 {
			nums[counter] = 0
			counter++
		} else if n == 1 {
			white++
		} else {
			blue++
		}
	}

	for i := 0; i < white; i++ {
		nums[counter] = 1
		counter++
	}

	for i := 0; i < blue; i++ {
		nums[counter] = 2
		counter++
	}
}

// @lc code=end

