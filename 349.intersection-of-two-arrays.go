/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
	set := make(map[int]bool)

	for _, v := range nums1 {
		set[v] = true
	}

	for _, v := range nums2 {
		if set[v] == true {
			res = append(res, v)
			set[v] = false
		}
	}

	return res
}
// @lc code=end

