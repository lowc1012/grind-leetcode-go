/*
 * @lc app=leetcode id=2486 lang=golang
 *
 * [2486] Append Characters to String to Make Subsequence
 */

// @lc code=start
func appendCharacters(s string, t string) int {
    // iterate through s and t, O(len(s) + len(t))
    a, b := 0, 0
    for a < len(s) && b < len(t) {
        if s[a] == t[b] {
            b++
        }
        a++
    }
    return len(t) - b // remaining chars in t need to be appended
}
// @lc code=end

