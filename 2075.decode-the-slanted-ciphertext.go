/*
 * @lc app=leetcode id=2075 lang=golang
 *
 * [2075] Decode the Slanted Ciphertext
 */

// @lc code=start
func decodeCiphertext(encodedText string, rows int) string {
    // edge cases
    if rows == 1 {
        return encodedText
    }

    numCh := len(encodedText)
    if numCh == 0 {
        return ""
    }

    cols := numCh / rows
    res := make([]byte, 0, numCh)
    for startCol := 0; startCol < cols; startCol++ {
        for r, c := 0, startCol; r < rows && c < cols; r, c = r+1, c+1 {
            index := r * cols + c
            res = append(res, encodedText[index])
        }
    }

    // remove trailing spaces ' '
    i := len(res)-1
    for i >= 0 && res[i] == ' ' {
        i--
    }

    return string(res[:i+1])
}
// @lc code=end

