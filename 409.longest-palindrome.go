/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	m := make(map[rune]struct{})
	count := 0

	for _, r := range s {
		if _, exist := m[r]; exist {
			delete(m, r)
			count++
		} else {
			m[r] = struct{}{}
		}
	}

	if len(m) != 0 {
		return count*2 + 1
	}
	return count * 2
}

// @lc code=end

