/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */
// @lc code=start
func topKFrequent(words []string, k int) []string {
	// return bucketSort(words, k)
	freqArray := make(map[string]int)

	for _, word := range words {
		freqArray[word]++
	}

	h := newHeap(k)
	for word, freq := range freqArray {
		w := WordAndFreq{word: word, freq: freq}
		h.insert(&w)
	}

	res := make([]string, k)

	for i := k - 1; i >= 0; i-- {
		res[i] = h.l[1].word
		h.l[1], h.l[h.size] = h.l[h.size], nil
		h.size--
		h.sink(1)
	}

	return res
}

type WordAndFreq struct {
	word string
	freq int
}

func (w *WordAndFreq) less(w1 *WordAndFreq) bool {
	if w.freq > w1.freq {
		return true
	} else if w.freq == w1.freq {
		return w.word < w1.word
	}

	return false
}

type heap struct {
	l    []*WordAndFreq
	k    int
	size int
}

func newHeap(k int) heap {
	return heap{l: make([]*WordAndFreq, k+1), k: k}
}

func (h *heap) insert(wf *WordAndFreq) {
	// heap is full
	// the top of heap is less than wf
	if h.size == h.k {
		if !h.l[1].less(wf) {
			h.remove()
		} else {
			return
		}
	}
	
	// add the new item
	h.l[h.size+1] = wf
	 // update the size instance variable
	h.size++
	h.swim(h.size)

}

func (h *heap) remove() {
	// put a big item to top
	h.l[1] = h.l[h.size]
	h.l[h.size] = nil
	 // update the size instance variable
	h.size--
	// rebuild heap
	h.sink(1)
}

func (h *heap) swim(i int) {
	for i > 1 && h.l[i/2].less(h.l[i]) {
		h.l[i/2], h.l[i] = h.l[i], h.l[i/2]
		i /= 2
	}
}

func (h *heap) sink(i int) {
	for i*2 <= h.size {
		j := i * 2

		if j < h.size && h.l[j].less(h.l[j+1]) {
			j++
		}

		if h.l[j].less(h.l[i]) {
			break
		}

		h.l[j], h.l[i] = h.l[i], h.l[j]
		i = j
	}
}

func bucketSort(words []string, k int) []string {
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

