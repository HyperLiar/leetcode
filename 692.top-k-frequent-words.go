/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	l := make(map[int][]string)
	res := make([]string, 0)

	max := 0
	for _, word := range words {
		m[word]++
		if m[word] > max {
			max = m[word]
		}
	}

	for word, freq := range m {
		l[freq] = append(l[freq], word)
	}

	for i := max; i >= 0; i-- {
		sort.Strings(l[i])
		res = append(res, l[i]...)
		if len(res) >= k {
			break
		}
	}

	return res[0:k]
}

// @lc code=end

