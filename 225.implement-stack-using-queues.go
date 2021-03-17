/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	in  *Queue
	out *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{in: newQueue(), out: newQueue()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.in.pushToback(x)
	for this.out.size() > 0 {
		this.in.pushToback(this.out.popFromFront())
	}

	this.in, this.out = this.out, this.in
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.out.popFromFront()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.out.peekFromFront()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.in.isEmpty() && this.out.isEmpty()
}

type Queue struct {
	l []int
}

func newQueue() *Queue {
	return &Queue{l: []int{}}
}

func (this *Queue) pushToback(x int) {
	this.l = append(this.l, x)
}

func (this *Queue) peekFromFront() int {
	return this.l[0]
}

func (this *Queue) popFromFront() (x int) {
	x = this.l[0]
	this.l = this.l[1:]
	return
}

func (this *Queue) size() int {
	return len(this.l)
}

func (this *Queue) isEmpty() bool {
	return this.size() == 0
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

