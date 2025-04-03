/*
 * @lc app=leetcode id=1408 lang=golang
 *
 * [1408] String Matching in an Array
 */

// @lc code=start
func stringMatching(words []string) []string {
    result := make([]string, 0)
    visited := make(map[string]bool) // record the found substring

    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i == j {
                continue
            }

            if strings.Contains(words[j], words[i]) && !visited[words[i]] {
                result = append(result, words[i])
                visited[words[i]] = true
                break
            }
        }
    }
    return result
}
// @lc code=end

