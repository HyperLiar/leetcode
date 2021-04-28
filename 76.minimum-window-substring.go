/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	sL := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		sL[t[i]]++
	}

	left, cnt := 0, 0
	minLeft, minLen := -1, math.MaxInt32

	for i := 0; i < len(s); i++ {
		sL[s[i]]--
		if sL[s[i]] >= 0 {
			cnt++
		}

		for cnt == len(t) {
			if minLen > i-left+1 {
				minLen = i - left + 1
				minLeft = left
			}

			if left > len(s)-1 {
				break
			}
			sL[s[left]]++
			if sL[s[left]] > 0{
				cnt--
			}
			left++
		}
	}

	if minLeft == -1 {
		return ""
	}

	return s[minLeft : minLeft+minLen]
}

// @lc code=end

