/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	l      []int
	min    int
	minList    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{l: []int{}, min: math.MaxInt32, minList: []int{}}
}

func (this *MinStack) Push(val int) {
	this.l = append(this.l, val)

	if val <= this.min {
		this.min = val
		this.minList = append(this.minList, val)
	}
}

func (this *MinStack) Pop() {
	v := this.l[len(this.l)-1]
	if v == this.min {
		this.minList = this.minList[0:len(this.minList)-1]

		if len(this.minList) == 0 {
			this.min = math.MaxInt32
		} else {
			this.min = this.minList[len(this.minList)-1]
		}
	}

	this.l = this.l[:len(this.l)-1]
}

func (this *MinStack) Top() int {
	return this.l[len(this.l)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

