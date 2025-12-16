/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	bracketMap := map[rune]rune{
        '}': '{',
        ']': '[',
        ')': '(',
    }

    stack := make([]rune, 0)
    for _, r := range s {
        switch r {
            case '}', ']', ')':
                if len(stack) == 0 || stack[len(stack)-1] != bracketMap[r] {
                    return false
                }
                stack = stack[:len(stack)-1]
            default:
                stack = append(stack, r)
        }
    }
    return len(stack) == 0
}
// @lc code=end

