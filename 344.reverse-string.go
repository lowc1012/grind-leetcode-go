/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {
	end := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[end], s[i] = s[i], s[end]
		end--
	}
}

// @lc code=end

