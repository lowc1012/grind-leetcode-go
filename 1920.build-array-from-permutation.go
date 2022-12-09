/*
 * @lc app=leetcode id=1920 lang=golang
 *
 * [1920] Build Array from Permutation
 */

// @lc code=start
func buildArray(nums []int) []int {
	// copy origin input array
	//		   0 1 2 3 4 5
	// nums = [0,2,1,5,3,4]
	// ans[i] = nums[nums[i]]

	// n := len(nums)
	// temp := make([]int, n)
	// copy(temp, nums)

	// for i := 0; i < n; i++ {
	// 	nums[i] = temp[nums[i]]
	// }

	// TC: O(n), SC: O(n)
	// return nums

	// optimize the space complexity
	// basic math method
	n := len(nums)
	for i := 0; i < n; i++ {
		//  new value = orign value * n
		nums[i] = nums[i] * n
	}

	for i := 0; i < n; i++ {
		//  nums[i] = nums[origin value]/n + new value
		//          = nums[origin value]/n + orign value * n
		nums[i] += nums[nums[i]/n] / n
	}
	for i := 0; i < n; i++ {
		//	new value = nums[i] % n
		nums[i] %= n
	}
	// TC: O(n), SC: O(1)
	return nums
}

// @lc code=end

