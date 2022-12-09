/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
		return true
	}

	indexS, indexT := 0, 0
	for indexT < len(t) {
		if s[indexS] == t[indexT] {
			indexS++
			if indexS == len(s) {
				return true
			}
		}
		indexT++
	}
	return false
}
// @lc code=end

