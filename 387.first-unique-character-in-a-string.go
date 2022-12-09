/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	m := make(map[string]int)
	for _, j := range s {
		if _, ok := m[string(j)]; ok {
			m[string(j)] += 1
		} else {
			m[string(j)] = 1
		}
	}

	for i, v := range s {
		if m[string(v)] == 1 {
			return i
		}
	}
	
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

