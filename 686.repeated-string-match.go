/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 */

// @lc code=start
func repeatedStringMatch(a string, b string) int {
	// a = "abcd", b = "cdabcdab"
	lenA, lenB := len(a), len(b)

	kMin := (lenB + lenA - 1) / lenA
	for i := kMin; i <= kMin+1; i++ {
		repeat := strings.Repeat(a, i)
		if strings.Contains(repeat, b) {
			return i
		}
	}
	return -1
}

// @lc code=end

