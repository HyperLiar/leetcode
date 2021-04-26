/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	return inDp(s)
}

func inDp(s string) int {
	res := 0
	dp := make([]int, len(s)+1)

	for i := 1; i <= len(s); i++ {
		j := i - 2 - dp[i-1]

		if s[i-1] == '(' || j < 0 || s[j] == ')' {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1] + 2 + dp[j]
			res = max(res, dp[i])
		}
	}
	return res
}

func twoPoint(s string) int {
	maxL := 0

	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxL = max(maxL, right*2)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxL = max(maxL, left*2)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxL
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// @lc code=end

