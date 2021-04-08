/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(j * max(dp[i-j], i-j),dp[i])
		}
	}

	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// @lc code=end

