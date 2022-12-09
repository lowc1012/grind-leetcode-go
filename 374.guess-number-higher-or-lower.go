/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	// binary search
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == -1 {
			right = mid - 1
		} else if res == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// @lc code=end

