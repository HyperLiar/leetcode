/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	start, end := 0, 0
    for i, _ := range s {
		nowLen := maxPand(s, i)
		if nowLen > end - start {
			start, end = i - (nowLen - 1) / 2, i + nowLen / 2
		}
	}

	return s[start:end+1]
}

func maxPand(s string, a int) int {
	l1 := length(s, a, a)
	l2 := length(s, a, a+1)

	if l1 > l2 {
		return l1
	}
	return l2
}

func length(s string, a, b int) int {
	for a >= 0 && b < len(s) && s[a] == s[b] {
		a--
		b++
	}

	return b - a - 1
}
// @lc code=end

