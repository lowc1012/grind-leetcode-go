/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
import "sort"

func frequencySort(s string) string {
	r := []rune(s)
	m := make(map[rune]int)
	for _, v := range r {
		m[v]++
	}

	sort.Slice(r, func(a, b int) bool {
		if m[r[a]] == m[r[b]] {
			return r[a] > r[b]
		}
		return m[r[a]] > m[r[b]]
	})
	return string(r)
}

// @lc code=end

