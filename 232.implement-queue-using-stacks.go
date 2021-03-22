/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
    in, out []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{in:[]int{}, out: []int{}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.in = append(this.in, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if len(this.out) == 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
		}
		this.in = this.in[:0]
	}

	x := this.out[len(this.out)-1]
	this.out = this.out[0:len(this.out)-1]
	return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.out) == 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
		}
		this.in = this.in[:0]
	}

	return this.out[len(this.out) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.in) == 0 && len(this.out) == 0
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

