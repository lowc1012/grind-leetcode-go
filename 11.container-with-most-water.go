/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	maxContainer := 0
	left, right := 0, len(height)-1
	for left < right {
		h := min(height[left], height[right])
		c := (right - left) * h
		if c > maxContainer {
			maxContainer = c
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxContainer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

