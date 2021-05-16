/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	if len(word2) == 0 {
		return len(word1)
	}
	if len(word1) == 0 {
		return len(word2)
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}

	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i, v1 := range word1 {
		for j, v2 := range word2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1], dp[i][j]) + 1
			}
		}
	}

	// max common subsequence
	return dp[len(word1)][len(word2)]
}

func min(i, j, k int) int {
	if i > j {
		if j > k {
			return k
		} else {
			return j
		}
	} else if i > k {
		return k
	} else {
		return i
	}
}

// @lc code=end

