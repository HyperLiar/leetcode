/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start
type MyCircularDeque struct {
    list []int
	start int
	end int
	isWrite bool
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
		list: make([]int, k),
		start: 0,
		end: 0,
		isWrite: false,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
		return false
	}

	this.start = this.decr(this.start)
	this.list[this.start] = value
	this.isWrite = true
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
		return false
	}

	this.list[this.end] = value
	this.end = this.incr(this.end)
	this.isWrite = true
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
		return false
	}

	this.list[this.start] = 0
	this.start = this.incr(this.start)
	this.isWrite = false
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
		return false
	}

	this.end = this.decr(this.end)
	this.list[this.end] = 0
	this.isWrite = false
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
		return -1
	}

	return this.list[this.start]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
		return -1
	}

	return this.list[this.decr(this.end)]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.start == this.end && !this.isWrite
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.start == this.end && this.isWrite
}

func (this *MyCircularDeque) incr(val int) int {
	if val == len(this.list)-1 {
		return 0
	}

	return val+1
}

func (this *MyCircularDeque) decr(val int) int {
	if val == 0 {
		return len(this.list) - 1
	}

	return val - 1
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

