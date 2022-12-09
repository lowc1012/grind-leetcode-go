/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	j := len(s)-1
	for i := 0; i < len(s); i++ {
		if string(s[i]) != string(s[j]) {
			return false
		}
		j--
	}
    return true
}
// @lc code=end

