/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	dict := make(map[string]bool)

	for _, w := range wordDict {
		dict[w] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func dfs(s string, index int, words map[string]bool, cache map[int]bool) bool {
	if len(s) == index {
		return true
	}

	if v, ok := cache[index]; ok {
		return v
	}

	for i := index; i < len(s); i++ {
		str := s[index:i+1]
		if words[str] && dfs(s, i+1, words, cache) {
			return true
		}
	}

	cache[index] = false
	return false
}

func bfs(s string, words map[string]bool) bool {
	var queue []int
	visited := make([]bool, len(s))

	queue = append(queue, 0)

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if !visited[start] {
			for end := start + 1; end <= len(s); end++ {
				if words[s[start:end]] {
					queue = append(queue, end)
					if end == len(s) {
						return true
					}
				}
			}

			visited[start] = true
		}
	}

	return false
}
// @lc code=end

