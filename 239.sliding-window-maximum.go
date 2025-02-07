/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, 0)
	if len(nums) == 0 || k == 0 {
		return result
	}

	// use a doubly-ended queue to store the index
	// of elements that are inside of window
	deque := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(deque) > 0 && deque[0] < i+1-k {
			deque = deque[1:] // remove the front of deque
		}

		// remove elements from the back of deque
		// which are smaller than the current element
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i+1 >= k {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
// @lc code=end

