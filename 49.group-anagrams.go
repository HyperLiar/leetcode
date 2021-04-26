/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

	for str := range strs {
		if _, ok := m[str]
	}
}

func minAnagram(s string) string {
	m := s

	for i := 0; i < len(m); i++ {
		for j := i+1; j < len(m); j++ {
			if m[i] > m[j] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}

	return m
}
// @lc code=end

