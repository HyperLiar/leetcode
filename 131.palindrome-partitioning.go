/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := make([][]string, 0)
	cur := []string{}

	dfs(dp, &res, s, &cur, 0)
	return res
}

func dfs(dp [][]bool, res *[][]string, s string, cur *[]string, i int) {
	if i >= len(s) {
		temp := []string{}
 		temp = append(temp, *cur...)
		*res = append(*res, temp)
	}

	for j := i; j < len(s); j++ {
		if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
			dp[i][j] = true
			*cur = append(*cur, s[i:j+1])
			dfs(dp, res, s, cur, j+1)
			*cur = (*cur)[0:len(*cur)-1]
		}
	}
}

// func dfs(dp [][]bool, res *[][]string, s string, cur *[]string, i int) {
// 	if i >= len(s) {
// 		temp := []string{}
// 		temp = append(temp, *cur...)
// 		*res = append(*res, temp)
// 	}

// 	for j := i; j < len(s); j++ {
// 		if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
// 			dp[i][j] = true
// 			*cur = append(*cur, s[i:j+1])
// 			dfs(dp, res, s, cur, j+1)
// 			*cur = (*cur)[0 : len(*cur)-1]
// 		}
// 	}
// }

// @lc code=end

