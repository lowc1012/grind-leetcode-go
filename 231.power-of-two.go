/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	for i := n; i >= 2 && i%2 == 0; {
		i = i / 2
		if i == 1 {
			return true
		}

	}
	return false
}

// @lc code=end

