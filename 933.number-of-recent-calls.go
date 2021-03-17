/*
 * @lc app=leetcode id=933 lang=golang
 *
 * [933] Number of Recent Calls
 */

// @lc code=start
type RecentCounter struct {
    requestTime []int
	l int
}


func Constructor() RecentCounter {
    return RecentCounter{requestTime:[]int{},l:0}
}


func (this *RecentCounter) Ping(t int) int {
    this.requestTime = append(this.requestTime, t)
	var idx int
	this.l = len(this.requestTime)
	for i := 0; i < this.l ; i++ {
		if this.requestTime[i] >= t-3000 {
			idx = i
			break
		}
	}

	this.requestTime = this.requestTime[idx:this.l]
	this.l = this.l - idx
	return this.l
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

