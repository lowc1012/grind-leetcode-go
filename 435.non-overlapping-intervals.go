/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start
import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	end, count := 0, 0
	

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		// if overlapping
		if curr[0] < intervals[end][1] {
			if curr[1] < intervals[end][1] {
				end = i
			}
			count++
		} else {
			end = i
		}
	}
	return count
}

// @lc code=end

