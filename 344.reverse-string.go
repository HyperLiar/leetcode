/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte)  {
	l := len(s)
	for i := 0; i < l / 2; i++ {
		s[i], s[l - i - 1] = s[l - i - 1], s[i]
	}
}
// @lc code=end

