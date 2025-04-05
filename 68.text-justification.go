/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	currentLine := []string{}
	currentLength := 0

	for _, word := range words {
		// length of current words + length of next word + space slot
		if currentLength+len(word)+len(currentLine) > maxWidth {

			if len(currentLine) == 1 {
                firstWord := currentLine[0]
                result = append(result, firstWord + strings.Repeat(" ", maxWidth-len(firstWord)))
			} else {
				// calculate num of spaces
				numSpaces := maxWidth - currentLength
				spacesBetween := numSpaces / (len(currentLine) - 1)
				extraSpaces := numSpaces % (len(currentLine) - 1)

				var s strings.Builder
				for i := 0; i < len(currentLine)-1; i++ {
					s.WriteString(currentLine[i])
					s.WriteString(strings.Repeat(" ", spacesBetween))
					if i < extraSpaces {
						s.WriteString(" ")
					}
				}
				// add last word in current line
				s.WriteString(currentLine[len(currentLine)-1])
				result = append(result, s.String())
			}

			// create a new line
			currentLine = []string{}
			currentLength = 0
		}

		currentLine = append(currentLine, word)
		currentLength += len(word)
	}
	// left justification on last line
	lastLine := strings.Join(currentLine, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	result = append(result, lastLine)

	return result
}

// @lc code=end

