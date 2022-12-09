/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) int {
	// TC: O(n^2)
	count := 0
	for i := 0; i < len(s); i++ {
		// case 1
		count += helper(s, i, i)

		// case 2
		count += helper(s, i, i+1)
	}
	return count
}

func helper(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

// @lc code=end

