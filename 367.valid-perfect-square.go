/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	left := 0
	right := num

	for left <= right {
		mid := left + (right-left) / 2
		if mid * mid == num {
			return true
		}
	
		if mid * mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
    
	return false
}
// @lc code=end

