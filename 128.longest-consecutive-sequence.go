/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
import "github.com/emirpasic/gods/sets/hashset"

func longestConsecutive(nums []int) int {
	s := hashset.New()
	for _, v := range nums {
		s.Add(v)
	}

	longestLen := 0
	for _, v := range s.Values() {
		if !s.Contains(v.(int) - 1) {
			current := v.(int)
			currentLen := 1
			for s.Contains(current + 1) {
				current += 1
				currentLen += 1
			}
			longestLen = max(longestLen, currentLen)
		}
	}

	return longestLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

