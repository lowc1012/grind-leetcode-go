/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func partitionLabels(s string) []int {
	res := []int{}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}

	size, end := 0, 0
	for start := 0; start < len(s); start++ {
		size++
		if m[s[start]] > end {
			end = m[s[start]]
		}

		if start == end {
			res = append(res, size)
			// reset size
			size = 0
		}
	}
	return res
}

// @lc code=end

