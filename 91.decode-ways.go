/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
    if s[0] == '0' {
		return 0
	}
	m := make([]int, len(s)+1)
	m[0], m[1] = 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
            m[i+1] += m[i]
        }
        v := 10 * int(s[i-1]-'0') + int(s[i]-'0')
        if s[i-1] != '0' && v < 27 {
            m[i+1] += m[i-1]
        }
	}
	return m[len(s)]
}
// @lc code=end

