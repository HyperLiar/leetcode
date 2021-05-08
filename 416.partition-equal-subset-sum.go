/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}

	dp := make([]bool, sum/2+1)
	dp[0] = true

	for _, num := range nums {
		for j := sum / 2; j >= num; j-- {
			if dp[j-num] {
				dp[j] = dp[j-num]
			}
		}
	}

	return dp[sum/2]
}

// @lc code=end

