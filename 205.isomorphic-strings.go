/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// create a hashmap to hold the records of mappings
	// from s[i] to t[i]
	charMap := make(map[byte]byte)

	// create a hashset to check whether a char has been mapped to
	mappedSet := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		v, exist := charMap[s[i]]
		if exist {
			if v == t[i] {
				continue // has been mapped
			} else {
				return false
			}
		} else {
			if _, exist := mappedSet[t[i]]; exist {
				return false
			} else {
				charMap[s[i]] = t[i]
				mappedSet[t[i]] = struct{}{}
			}
		}
	}
	return true
}

// @lc code=end

