/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	var res int = 0
	
	// make a set
	set := make(map[string]struct{})

	// use two pointer to implement "sliding window"
	a := 0
	for i, xk := range s {
		_, ok := set[string(k)]
		for ok {
			delete(set, string(s[a]))
			a += 1
			_, ok = set[string(k)]
		}
		set[string(k)] = struct{}{}
		res = max(res, i - a + 1)
	} 
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

