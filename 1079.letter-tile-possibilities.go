/*
 * @lc app=leetcode id=1079 lang=golang
 *
 * [1079] Letter Tile Possibilities
 */

// @lc code=start
func numTilePossibilities(tiles string) int {
    // Count frequency of each character
    charFreq := make(map[byte]int)
    for i := 0; i < len(tiles); i++ {
        charFreq[tiles[i]]++
    }
    
    return countSequences(charFreq)
}

func countSequences(charFreq map[byte]int) int {
    totalSequences := 0
    
    // Try using each available character
    for char, frequency := range charFreq {
        if frequency > 0 {
            // Use this character
            charFreq[char]--
            
            // Count: 1 (current char) + all sequences with remaining chars
            totalSequences += 1 + countSequences(charFreq)
            
            // Backtrack: restore the character
            charFreq[char]++
        }
    }
    
    return totalSequences
}

// @lc code=end

