/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	roadMap := make([][]int, n)
	for i := 0; i < n; i++ {
		roadMap[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		roadMap[0][j] = 1
	}

	for i := 0; i < n; i++ {
		roadMap[i][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			roadMap[i][j] = roadMap[i][j-1] + roadMap[i-1][j]
		}
	}

	return roadMap[n-1][m-1]
}

// @lc code=end

