/*
 * @lc app=leetcode id=1486 lang=golang
 *
 * [1486] XOR Operation in an Array
 */

// @lc code=start
func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start+2 * i
	}
	return res
}

// @lc code=end

