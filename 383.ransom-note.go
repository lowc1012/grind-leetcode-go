/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	letterArray := make([]int, 26)

	runeRansomNote := []rune(ransomNote)
	runeMagazine := []rune(magazine)

	for i := 0; i < len(runeRansomNote); i++ {
		letterArray[runeRansomNote[i]-'a']++
	}

	for i := 0; i < len(runeMagazine); i++ {
		letterArray[runeMagazine[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if letterArray[i] > 0 {
			return false
		}
	}
	return true
}

// @lc code=end

