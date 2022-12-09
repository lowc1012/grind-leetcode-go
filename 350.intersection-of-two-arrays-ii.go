/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	// hash table
	set := make(map[int]int)
	for _, v := range nums1 {
		set[v] += 1
	}

	res := make([]int, 0)
	for _, v := range nums2 {
		if set[v] > 0 {
			res = append(res, v)
			set[v] -= 1
		}
	}

	return res
}
// @lc code=end

