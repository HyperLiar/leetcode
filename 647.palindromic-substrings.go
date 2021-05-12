/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) int {
	n, res := len(s), 0

	dp := make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n)
	}
	
	for i := n-1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j - i <= 2 || dp[i + 1][j - 1]) {
				dp[i][j] = true
			}
			if dp[i][j] {
				res++
			}
		}
	}

	return res
}

func bruteForce(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		j := 0
		// odd
		for ; j <= i && j < len(s)-i; j++ {
			if s[i-j] != s[i+j] {
				break
			}
		}
		count += j

		fmt.Println(j, count)
		// even
		j = 0
		for ; j <= i && j+1 < len(s)-i; j++ {
			if s[i-j] != s[i+j+1] {
				break
			}
		}
		count += j
	}

	return count
}
// @lc code=end

