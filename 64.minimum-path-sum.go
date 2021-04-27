/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// @lc code=end

