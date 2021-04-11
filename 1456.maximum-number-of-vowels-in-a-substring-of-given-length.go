/*
 * @lc app=leetcode id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start
func maxVowels(s string, k int) int {
    m := make(map[byte]struct{})
	m['a'] = struct{}{}
	m['i'] = struct{}{}
	m['e'] = struct{}{}
	m['o'] = struct{}{}
	m['u'] = struct{}{}
    vowel := make([]int, 26)
    for v := range m {
        vowel[v-'a'] = 1
    }

	count, maxCount := 0, 0
	for i := 0; i < len(s); i++ {
		if vowel[s[i] - 'a'] == 1 {
			count++
		}
		if i >= k {
			if vowel[s[i-k]-'a'] == 1 {
				count--
			}
		}	
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
// @lc code=end

