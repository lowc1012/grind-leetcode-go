/*
 * @lc app=leetcode id=1297 lang=golang
 *
 * [1297] Maximum Number of Occurrences of a Substring
 */

// @lc code=start
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    freq := make(map[string]int)
    maxCount := 0

    // Only check substrings of minSize length
    for i := 0; i <= len(s)-minSize; i++ {
        substr := s[i:i+minSize]

        // Count unique characters in substring
        charCount := make(map[byte]int)
        for j := 0; j < len(substr); j++ {
            charCount[substr[j]]++
        }

        // Check if substring has at most maxLetters unique characters
        if len(charCount) <= maxLetters {
            freq[substr]++
            if freq[substr] > maxCount {
                maxCount = freq[substr]
            }
        }
    }

    return maxCount
}
// @lc code=end

