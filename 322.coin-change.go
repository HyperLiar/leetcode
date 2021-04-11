/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount+1
	}

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i - coin] != amount+1 {
				dp[i] = min(dp[i], dp[i-coin] + 1)
			} 
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

