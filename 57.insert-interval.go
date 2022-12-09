/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	l, r := [][]int{}, [][]int{}
	a, b := newInterval[0], newInterval[1]
	for _, interval := range intervals {
		if interval[1] < a {
			l = append(l, interval)
		} else if interval[0] > b {
			r = append(r, interval)
		} else {
			a = min(a, interval[0])
			b = max(b, interval[1])
		}
	}
	return append(append(l, []int{a, b}), r...)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

