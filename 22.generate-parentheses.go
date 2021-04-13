/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
    res := make([]string, 0)
	start := ""

	backTrack(&res, start, n, 0, 0)
	return res
}

func backTrack(res *[]string, cur string, max, open, close int) {
	if len(cur) == max*2 {
		*res = append(*res, cur)
		return
	}

	if open < max {
		cur = cur + "("
		backTrack(res, cur, max, open+1, close)
		cur = cur[0:len(cur)-1]
	}

	if close < open {
		cur = cur + ")"
		backTrack(res, cur, max, open, close+1)
		cur = cur[0:len(cur)-1]
	}
}
// @lc code=end

