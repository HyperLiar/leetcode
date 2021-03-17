/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

// @lc code=start
type MyCircularQueue struct {
    itemList []int
	start int
	end int
	isWrite bool
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{itemList:make([]int, k), start:0, end: 0, isWrite: false}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
		return false
	}

	this.itemList[this.end] = value
	this.end = this.incr(this.end)
	this.isWrite = true

	return true
}

func (this *MyCircularQueue) incr(pos int) int {
	if pos == len(this.itemList) - 1 {
		return 0
	}

	return pos+1
}

func (this *MyCircularQueue) decr(pos int) int {
	if pos == 0 {
		return len(this.itemList) - 1
	}

	return pos-1
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
		return false
	}

	this.itemList[this.start] = 0
	this.start = this.incr(this.start)
	this.isWrite = false

	return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
		return -1
	}

	return this.itemList[this.start]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
		return -1
	}

	return this.itemList[this.decr(this.end)]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.start == this.end && !this.isWrite
}


func (this *MyCircularQueue) IsFull() bool {
    return this.start == this.end && this.isWrite
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

