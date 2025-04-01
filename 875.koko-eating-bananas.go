/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start
import "math"

func minEatingSpeed(piles []int, h int) int {
    // find the max number of banana in a pile
    maxPile := 0
    for _, p := range piles {
        if p > maxPile {
            maxPile = p
        }
    }

    result := maxPile

    // use binary search to find the minimum speed to eat all bananas
    // time complexity: O(NlogM), N: num of piles, M: the max num of banana in a pile
    left, right := 1, maxPile
    for left <= right {
        k := left + (right-left)/2
    
        // calculate the total hours needed to eat all bananas at speed k
        hours := 0
        for _, p := range piles {
            hours += int(math.Ceil(float64(p) / float64(k)))
        }

        // adjust left or right to increase or decrease the speed
        if hours > h {
            left = k + 1
        } else {
            right = k - 1
            result = k
        }
    }
    return result
}
// @lc code=end

