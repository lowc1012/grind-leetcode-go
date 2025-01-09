/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
import "strings"

func frequencySort(s string) string {
	// use bucket sort
	freqMap := make(map[rune]int, 0)
	for _, ch := range s {
		freqMap[ch]++
	}

	// create a bucket to store the frequency of each character
	bucket := make([][]rune, len(s)+1)

	for ch, freq := range freqMap {
		bucket[freq] = append(bucket[freq], ch)
	}

	var res strings.Builder
	for i := len(bucket)-1; i > 0; i-- {
		for _, ch := range bucket[i] {
			res.WriteString(strings.Repeat(string(ch), i))
		}
	}

	return res.String()
}

// @lc code=end

