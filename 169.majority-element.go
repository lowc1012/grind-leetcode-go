/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
// (method 1) use quick sort (sort.Ints) => TC: O(nlogn), SC: O(1)
// (method 2) use hash table => TC: O(n), SC: O(n)
// (method 3) Boyer-Moore Voting Algorithm => TC: O(n), SC: O(1)
func majorityElement(nums []int) int {
	// We can use Boyer-Moore without hash table to save memory space

	// initialize the result and the frequency of current majority
    result, majority := nums[0], 0

    // iterate each elements in nums
    for i := 0; i < len(nums); i++ {
        // select a new candidate from the current element
		// when the frequency is 0 
		if majority == 0 {
            result = nums[i]
        }
		// increase the majority if we encounter the same element
        if nums[i] == result {
            majority++
        } else {
			// If we encounter a different element, we decrease it
            majority--
        }
    }

    return result
}
