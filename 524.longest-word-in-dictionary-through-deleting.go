/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
    result := ""
    for _, word := range dictionary {
        if isSubsequence(word, s) && (len(word) > len(result) || (len(word) == len(result) && word < result)) {
            result = word
        }
    }
    return result
}

func isSubsequence(word, s string) bool {
    i := 0
    for j := 0; j < len(s) && i < len(word); j++ {
        if s[j] == word[i] {
            i++
        }
    }
    return i == len(word)
}
// @lc code=end

