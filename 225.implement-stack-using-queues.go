/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
    enqueue []int
	dequeue []int
}


func Constructor() MyStack {
    return MyStack{[]int{}, []int{}}
}


func (this *MyStack) Push(x int)  {
	this.enqueue = append(this.enqueue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.enqueue) - 1
    for i := 0; i < length; i++ {
		this.dequeue = append(this.dequeue, this.enqueue[0])
		this.enqueue = this.enqueue[1:]
	}

	top := this.enqueue[0]
	this.enqueue = this.dequeue
	this.dequeue = nil
	return top
}

/** Get the top element. */
func (this *MyStack) Top() int {
    peek := this.Pop()
	this.enqueue = append(this.enqueue, peek)
	return peek
}


func (this *MyStack) Empty() bool {
    if len(this.enqueue) == 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

