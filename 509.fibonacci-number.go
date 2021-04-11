/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
    if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	current, prev1, prev2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return current
}
// @lc code=end

