/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	if lenHaystack < lenNeedle {
		return -1
	}

	if lenNeedle == 0 {
		return 0
	}

	return strings.Index(haystack, needle)

}
// @lc code=end

