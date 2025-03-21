/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	
	result := 0
	for num := range m {
		// check if a start of a sequence
		if exist := m[num-1]; !exist {
			curr := num
			currStreak := 1

			for m[curr+1] {
				curr++
				currStreak++
			}

			result = max(result, currStreak)
		}
	}
	return result
}

// @lc code=end

