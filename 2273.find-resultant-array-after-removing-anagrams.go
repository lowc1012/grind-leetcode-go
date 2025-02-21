/*
 * @lc app=leetcode id=2273 lang=golang
 *
 * [2273] Find Resultant Array After Removing Anagrams
 */

// @lc code=start
func removeAnagrams(words []string) []string {
	if len(words) < 2 {
		return words
	}

	result := []string{
		words[0],
	}
	for i := 1; i < len(words); i++ {
		// check if currenct word is an anagram of previous
		if !areAnagrams(words[i], words[i-1]) {
			result = append(result, words[i])
		}
	}

	return result
}

func areAnagrams(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	freqCount := [26]int{}
	for i := 0; i < len(w1); i++ {
		freqCount[w1[i]-'a']++
		freqCount[w2[i]-'a']--
	} 
	for _, c := range freqCount {
		if c != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

