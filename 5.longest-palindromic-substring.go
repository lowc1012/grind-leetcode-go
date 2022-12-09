/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var start, end = 0, 0
	var maxLength int

	// O(n^2)
	for i := 0; i < len(s); i++ {

		// case 1: left == right (center剛好是其中一個字母)
		len1 := expandFromMiddle(i, i, s)

		// case 2: left != right (center在兩個字母中間)
		len2 := expandFromMiddle(i, i+1, s)

		maxLength = max(len1, len2)

		// 找到最長的 palindrome
		if maxLength > (end - start) {
			start = i - ((maxLength - 1) / 2)
			end = i + (maxLength / 2)
		}
	}

	return s[start:end+1]
}

// expandFromMiddle return the length of palindrome
func expandFromMiddle(left, right int, s string) int {
	if left > right {
		return 0
	}

	// 用 while 從 center 去找 plaindrome 的範圍, O(n)
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}

	// 當 s[left] 與 s[right] 不相同時 return
	// 這裡 -1 是因為左跟右要往內縮 => 長度 -2 
	// 所以 (right - left + 1) - 2 = right - left - 1
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end
