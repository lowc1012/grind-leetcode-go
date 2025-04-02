/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
    maxAllowed := math.MaxInt32
    isNegative := false
    if x < 0 {
        isNegative = true
        x = -x
        maxAllowed = math.MaxInt32 + 1
    }

    result := 0
    for x > 0 {
        pop := x % 10
        x /= 10

        if result > maxAllowed/10 || (result == maxAllowed/10 && pop > maxAllowed%10) {
            return 0
        }
        result = pop + result*10
    }

    if isNegative {
        result = -result
    }
    return result
}
// @lc code=end
