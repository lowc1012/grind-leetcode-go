/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func backtrack(res *[][]string, track []string, start int, str string) {
	// return case
	if len(str) == start {
		tmp := make([]string, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(str); i++ {
		if isPalindrome(str, start, i) {
			track = append(track, str[start:i+1])
			backtrack(res, track, i+1, str)
			track = track[:len(track)-1]
		}
	}

}

func partition(s string) [][]string {
    // using backtracking, TC: O(N*2^N)
	result := make([][]string, 0)
	backtrack(&result, []string{}, 0, s)
	return result
}

// @lc code=end

