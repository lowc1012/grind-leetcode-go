/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	var stack []rune
	for _, r := range s {
		if r == '[' || r == '{' || r == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}

			if r != ')' && stack[len(stack)-1] == '(' {
				return false
			}

			if r != '}' && stack[len(stack)-1] == '{' {
				return false
			}

			if r != ']' && stack[len(stack)-1] == '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}
    return false
}
// @lc code=end

