/*
 * @lc app=leetcode id=791 lang=golang
 *
 * [791] Custom Sort String
 */

// @lc code=start
import "sort"

func customSortString(order string, s string) string {
	// TC: O(nlogn), SC: O(len(order)+len(s))

	positionMap := make(map[rune]int)
	for k, v := range order {
		positionMap[v] = k
	}

	runeSlice := []rune(s)

	sort.Slice(runeSlice, func(a, b int) bool {
		x, exist := positionMap[runeSlice[a]]
		if !exist {
			x = a
		}

		y, exist := positionMap[runeSlice[b]]
		if !exist {
			y = b
		}
		return x < y
	})

	return string(runeSlice)
}

// @lc code=end

