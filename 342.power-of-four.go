/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	if n & (n-1) != 0 {
		return false
	}

	return n & 0xAAAAAAAA == 0
}
// @lc code=end

