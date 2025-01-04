/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
		// if the input array is sorted? 
		// if the input array contains negative integers?

		// can not use sliding window because it's typically for non-negative numbers or fixed conditions
		// Instead, we can use a Prefix Sum and a hashmap to store the prefix sum and its frequency
		prefixSumMap := make(map[int]int)
		prefixSumMap[0] = 1

		currentSum, count := 0, 0
		for _, n := range nums {
				currentSum += n
				// if the currentSum - k exists in the map, then we can add the frequency of the currentSum - k to the count
				if val, exists := prefixSumMap[currentSum-k]; exists {
						count += val
				}
				prefixSumMap[currentSum]++
		}
		return count // Time: O(n), Space: O(n)
}
// @lc code=end

