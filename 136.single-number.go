/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	a := 0
    for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}
// @lc code=end

