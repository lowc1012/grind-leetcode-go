/*
 * @lc app=leetcode id=1286 lang=golang
 *
 * [1286] Iterator for Combination
 */

// @lc code=start
type CombinationIterator struct {
	combinations []string
	curr         int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combinations := []string{}

	var backtrack func([]byte, int)
	backtrack = func(track []byte, start int) {
		if len(track) == combinationLength {
			combinations = append(combinations, string(track))
			return
		}

		for i := start; i < len(characters); i++ {
			track = append(track, characters[i])
			backtrack(track, i+1)
			track = track[:len(track)-1]
		}
	}

	backtrack([]byte{}, 0)
	return CombinationIterator{combinations, 0}
}

func (this *CombinationIterator) Next() string {
	result := this.combinations[this.curr]
	this.curr++
	return result
}

func (this *CombinationIterator) HasNext() bool {
	return this.curr < len(this.combinations)
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

