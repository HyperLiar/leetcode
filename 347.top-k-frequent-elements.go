/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	pairL := make([]Pair, len(m))
	i := 0
	for key, value := range m {
		pairL[i] = Pair{key, value}
		i++
	}

	sort.Slice(pairL, func(i, j int) bool {
		if pairL[i].v > pairL[j].v {
			return true
		} else {
			return false
		}
	})

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = pairL[i].k
	}

	return res
}

type Pair struct {
	k, v int
}

// @lc code=end

