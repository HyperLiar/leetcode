/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	return search(haystack, needle)
}

func search(word, pattern string) int {
	i, j := 0, 0
	next := makeNext(pattern)

	for i < len(word) && j < len(pattern) {
		if j == -1 || word[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(pattern) {
		return i - j
	}

	return -1
}

func makeNext(pattern string) []int {
	next := make([]int, len(pattern)+1)
	next[0] = -1
	i, j := 0, -1

	for i < len(pattern) {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}

// @lc code=end

