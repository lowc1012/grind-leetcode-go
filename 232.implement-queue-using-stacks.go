/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	// if output stack is empty, transfering all elements
	// from input stack to output stack. => reverse the order
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			e := this.input[len(this.input)-1]

			this.output = append(this.output, e)
			this.input = this.input[:len(this.input)-1]
		}
	}

	// pop from the output stack and then return the peek of queue
	top := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]

	return top
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	// if output stack is empty, transfering all elements
	// from input stack to output stack. => reverse the order
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			e := this.input[len(this.input)-1]

			this.output = append(this.output, e)
			this.input = this.input[:len(this.input)-1]
		}
	}
	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.output) == 0 && len(this.input) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

