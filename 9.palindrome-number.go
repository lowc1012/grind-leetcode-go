/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	// Don't need to check negative numbers
	if x < 0 {
		return false
	}

	// The time complexity is O(log(n)), where n is the input number.
	// This is because we process each digit of the number once.
	return x == reverse(x)
}

// reverse reverses the digits of an integer.
func reverse(x int) int {
	reversed := 0
	for x != 0 {
		// Pop the last digit from x
		// for example, 123 % 10 = 3
		// and push it to the end of reversed
		reversed = reversed*10 + x%10

		// Remove the last digit from x
		x /= 10
	}
	return reversed
}
// @lc code=end

