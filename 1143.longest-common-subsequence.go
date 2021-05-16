/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	
}
func byDp(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1)+1)

	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i, v1 := range text1 {
		for j, v2 := range text2 {
			if dp[i][j+1] > dp[i+1][j] {
				dp[i+1][j+1] = dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
			if v1 == v2 && dp[i][j]+1 > dp[i+1][j+1] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
		}
	}

	return dp[len(text1)][len(text2)]
}
// @lc code=end

