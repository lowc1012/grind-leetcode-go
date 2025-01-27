/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1

	// O(n^2)
	for i := 0; i < len(s); i++ {
		// case 1: left == right (center剛好是其中一個字母)
		left, right := expandFromMiddle(i, i, s)
		if right-left+1 > maxLength {
			maxLength = right-left+1
			start = left
		}

		// case 2: left != right (center在兩個字母中間)
		left, right = expandFromMiddle(i, i+1, s)
		if right-left+1 > maxLength {
			maxLength = right-left+1
			start = left
		}
	}

	return s[start:start+maxLength]
}

// expandFromMiddle return the bounds of the palindrome
func expandFromMiddle(left, right int, s string) (int, int) {
	// 用 while 從 center 去找 plaindrome 的範圍, O(n)
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}

	return left + 1, right - 1
}
// @lc code=end
