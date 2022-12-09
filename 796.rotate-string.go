/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	// start from each char in s
	for start := 0; start >= 0 && start <= len(s); start++ {
		// restructure the given string
		str := s[start:len(s)] + s[0:start]
		if str == goal {
			return true
		}
	}

	return false
}

// @lc code=end

