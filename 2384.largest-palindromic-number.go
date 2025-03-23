/*
 * @lc app=leetcode id=2384 lang=golang
 *
 * [2384] Largest Palindromic Number
 */

// @lc code=start
import "strconv"

func reverse(s string) string {
	r := []rune(s)
	left, right := 0, len(r)-1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}

func largestPalindromic(num string) string {
    freqCount := make([]int, 10)
	for _, c := range num {
		freqCount[c-'0']++
	}

	var left []byte
	maxCenter := -1
	for i := 9; i >= 0; i-- {
		// Skip leading zero
        if len(left) == 0 && i == 0 {
            break
        }

		// build the left half
		pairs := freqCount[i]/2
		for j := 0; j < pairs; j++ {
			left = append(left, byte(i)+'0')
		}

		// find the center
		if freqCount[i]%2 == 1 && maxCenter == -1 {
			maxCenter = i
		}
	}

	// Handle case when left is empty (all zeros)
    if len(left) == 0 {
        if maxCenter != -1 {
            return strconv.Itoa(maxCenter)
        }
        return "0"
    }

	// if the maxCenter is not found, then the number is already a palindrome
	if maxCenter == -1 {
		return string(left) + reverse(string(left))
	}

	return string(left) + strconv.Itoa(maxCenter) + reverse(string(left))
}
// @lc code=end

