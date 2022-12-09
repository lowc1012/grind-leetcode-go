/*
 * @lc app=leetcode id=1464 lang=golang
 *
 * [1464] Maximum Product of Two Elements in an Array
 */

// @lc code=start
func maxProduct(nums []int) int {
	max, preMax := 0, 0

	for _, v := range nums {
		if v > max {
			preMax = max
			max = v
		} else if v > preMax {
			preMax = v
		}
	}

	return (max - 1) * (preMax - 1)
}

// @lc code=end

