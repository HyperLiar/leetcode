/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
	ans := 0

	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			i = max(m[s[j]], i)
		}

		ans = max(ans, j - i + 1)
		m[s[j]] = j + 1
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
// @lc code=end

