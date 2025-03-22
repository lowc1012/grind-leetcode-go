/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
    // 1. sliding window - O(n-k)
    // left, right := 0, len(arr)-k
    // for left < right {
    //     if x - arr[left] > arr[left+k] - x {
    //         left++
    //     } else {
    //         right--
    //     }
    // }

    // return arr[left:left+k]

    // 2. use binary search to find left bound - O(log(n))
    left, right := 0, len(arr)-k
    for left < right {
        mid := left + (right-left)/2
        if x - arr[mid] > arr[mid+k] - x {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return arr[left:left+k]
}
// @lc code=end

