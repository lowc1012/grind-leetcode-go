/*
 * @lc app=leetcode id=1636 lang=golang
 *
 * [1636] Sort Array by Increasing Frequency
 */

// @lc code=start
import "sort"

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	sort.Slice(nums, func(a, b int) bool {
		if m[nums[a]] == m[nums[b]] {
			return nums[a] > nums[b]
		}
		return m[nums[a]] < m[nums[b]]
	})
	return nums
}

// @lc code=end

