/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// @lc code=end

