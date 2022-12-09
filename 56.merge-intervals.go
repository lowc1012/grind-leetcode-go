/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
import "sort"

func merge(intervals [][]int) [][]int {
	// sort approach, TC: O(nlogn), SC: O(n)
	res := [][]int{}

	// Sort with custom comparator
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // order by asc
	})

	// merge intervals
	start, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		// if no overlapping
		if interval[0] > end {
			// merge previos two nodes, and store them into result
			res = append(res, []int{start, end})
			// replace start and end with current node's
			start, end = interval[0], interval[1]
		} else {
			// compare and choose a new end
			end = max(interval[1], end)
		}
	}
	res = append(res, []int{start, end})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

