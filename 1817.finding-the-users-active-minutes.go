/*
 * @lc app=leetcode id=1817 lang=golang
 *
 * [1817] Finding the Users Active Minutes
 */

// @lc code=start
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	m := make(map[int]map[int]struct{})
	for _, log := range logs {
		if m[log[0]] == nil {
			m[log[0]] = make(map[int]struct{})
			m[log[0]][log[1]] = struct{}{}
		} else {
			m[log[0]][log[1]] = struct{}{}
		}
	}

	result := make([]int, k)
	for _, mapEntry := range m {
		// mapEntry is hashset acturally which stores the minutes
		if len(mapEntry) > 0 {
			result[len(mapEntry)-1] += 1
		}

	}
	return result
}

// @lc code=end

