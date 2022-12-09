/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
    tail1, tail2, finished := m-1, n-1, m+n-1
	for tail1 >= 0 && tail2 >= 0 {
		if nums1[tail1] > nums2[tail2] {
			nums1[finished] = nums1[tail1]
			tail1 -= 1
		} else {
			nums1[finished] = nums2[tail2]
			tail2 -= 1
		}
		finished -= 1
	}

	for tail2 >= 0 {
		nums1[finished] = nums2[tail2]
		finished -= 1
		tail2 -= 1
	}
}
// @lc code=end

