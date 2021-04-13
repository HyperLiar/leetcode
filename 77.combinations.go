/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	backTrack(&res, &cur, 1, n, k)

	return res
}

func backTrack(res *[][]int, cur *[]int, now, max, k int) {
	if len(*cur) == k {
		result := append([]int{}, *cur...)
		*res = append(*res, result)
		return
	}

	for i := now; i <= max; i++ {
		*cur = append(*cur, i)
		backTrack(res, cur, i+1, max, k)
		*cur = (*cur)[0:len(*cur)-1]
	}
}

// @lc code=end

