/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start
type heap struct {
	words []string
	m     map[string]int
	l     int
}

func (h *heap) less(i, j int) bool {
	wordI, wordJ := h.words[i], h.words[j]
	wordIC, wordJC := h.m[wordI], h.m[wordJ]

	if wordIC > wordJC {
		return false
	} else if wordIC < wordJC {
		return true
	} else {
		return wordI > wordJ
	}
}

func (h *heap) down(i, l int) {
	min := i
	left, right := i<<1+1, i<<2+1

	if left <= l && h.less(left, min) {
		min = left
	}

	if right <= l && h.less(right, min) {
		min = right
	}

	if i != min {
		h.words[i], h.words[min] = h.words[min], h.words[i]
		fmt.Println(h.words)
		h.down(min, l)
	}
}

func topKFrequent(words []string, k int) []string {
	h := &heap{m: make(map[string]int), words: make([]string, 0), l: k}
	for _, word := range words {
		h.m[word]++

	}

	for word, _ := range h.m {
		h.words = append(h.words, word)
	}

	lens := len(h.words) - 1

	for i := lens >> 1; i >= 0; i-- {
		h.down(i, lens)
	}

	for j := lens; j >= 1; j-- {
		lens--
		h.words[0], h.words[j] = h.words[j], h.words[0]
		h.down(0, lens)
	}

	fmt.Println(h.m, h.words)
	return h.words[0:k]
}

// @lc code=end

