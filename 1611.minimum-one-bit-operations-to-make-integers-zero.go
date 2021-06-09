/*
 * @lc app=leetcode id=1611 lang=golang
 *
 * [1611] Minimum One Bit Operations to Make Integers Zero
 */

// @lc code=start
func minimumOneBitOperations(n int) int {
	ans := 0
	for n > 0 {
		ans ^= n
		n >>= 1
	}
	return ans
}

// @lc code=end

