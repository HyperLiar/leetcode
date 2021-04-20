/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	ver, hor, max := len(matrix[0]), len(matrix), 0
	dp := make([][]int, hor)
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, ver)
	}

	for i := 0; i < hor; i++ {
		for j := 0; j < ver; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
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

